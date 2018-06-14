package eca

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/PRETgroup/goFB/iec61499"
)

//An EventTrace is a set of instantaneous EventTraceSteps
type EventTrace []EventTraceStep

//An EventTraceStep shows a possible instantaneous step in a given trace, based on the ECStateName, and saves the output events
type EventTraceStep struct {
	InboundTransition iec61499.ECTransition
	OutputEvents      []string
}

//GetECStateName returns the given state's name in this EventTraceStep
func (ets EventTraceStep) GetECStateName() string {
	return ets.InboundTransition.Destination
}

//ChainOutputs are a slice of event outputs
type ChainOutputs []string

//DeriveBFBEventTraceSet takes an IEC61499 FB and creates a set of event output traces for that single FB
//e.g. for input "A", trace "s1->s2->s3" emits "B", "C", and trace "s1->s4_s3" also emits "B", "C"
func DeriveBFBEventTraceSet(fb iec61499.FB) (map[string][]EventTrace, error) {
	//1. For each INPUT EVENT to a FB, compute the set of possible OUTPUT EVENTS.
	//* This is done by checking states which depend upon those inputs, and executing them to their penultimate state according to event semantics.

	if fb.BasicFB == nil {
		return nil, errors.New("Can only derive Eventchain set for BFBs")
	}

	chains := make(map[string][]EventTrace, 0)

	//for each input event
	for _, evI := range fb.InterfaceList.EventInputs {
		chains[evI.Name] = make([]EventTrace, 0)
		// chain := EventChain{
		// 	InputName: evI.Name,
		// }
		//fmt.Printf("Input: %s\n", evI.Name)

		//find all transitions that depend upon that event
		startTrs := make([]iec61499.ECTransition, 0)
		for _, tr := range fb.BasicFB.Transitions {
			if strings.Contains(tr.Condition, evI.Name) {
				startTrs = append(startTrs, tr)
			}
		}

		//for each starting transition, follow them through to all penultimate states
		for _, sTr := range startTrs {
			searchTraces := make([]EventTrace, 1)
			//outpEs := make([]string, 0) //output Events for this transition start
			searchTraces[0] = make([]EventTraceStep, 1)
			searchTraces[0][0] = EventTraceStep{
				InboundTransition: sTr,
			}

			for i := 0; i < len(searchTraces); i++ {
				var searchTr *iec61499.ECTransition
				searchTr = &searchTraces[i][len(searchTraces[i])-1].InboundTransition
				//find the destination state for this transition
				for _, s := range fb.BasicFB.States {
					if s.Name == searchTr.Destination {
						//fmt.Printf("(state '%s')", s.Name)
						//append all output events
						for _, ecA := range s.ECActions {
							if ecA.Output != "" {
								searchTraces[i][len(searchTraces[i])-1].OutputEvents = append(searchTraces[i][len(searchTraces[i])-1].OutputEvents, ecA.Output)
							}
						}

						//now check outgoing transitions from this state
						//foreach transition with this state as the source
						firstFound := false
						for _, tr := range fb.BasicFB.Transitions {
							if tr.Source == s.Name {
								//add them to the search if they feature no events (ie if they are instantaneous)
								instantaneous := true
								for _, evIc := range fb.InterfaceList.EventInputs {
									if strings.Contains(tr.Condition, evIc.Name) {
										instantaneous = false
										break
									}
								}
								if instantaneous == true {
									//check destination of transition (make sure we've not been at this state before, else we have an instantaneous self-loop / an infinite loop)
									firstI := 0
									if firstFound {
										firstI = 1
									}
									for _, traceStep := range searchTraces[i+firstI] {
										if traceStep.GetECStateName() == tr.Destination {
											//we've been here before!
											return nil, fmt.Errorf("error in Block '%s', instataneous self-loop to state '%s'", fb.Name, tr.Destination)
										}
									}

									if firstFound == false {
										//this is the first finding, so add it to this searchTrace (a linear addition to the trace)
										searchTraces[i] = append(searchTraces[i], EventTraceStep{
											InboundTransition: tr,
										})
										i--               //we subtract one from i here to indicate that this searchTrace isn't yet complete (so it will be rechecked on next loop iteration)
										firstFound = true //we can only do this linear addition once, any other discovery must make a new searchTrace
									} else {
										//otherwise, continue on the next searchTrace (this is a branching addition to the trace)
										tempSearch := make(EventTrace, len(searchTraces[i+1])-1)
										copy(tempSearch, searchTraces[i+1][:len(searchTraces[i+1])-1])
										tempSearch = append(tempSearch, EventTraceStep{
											InboundTransition: tr,
										})
										searchTraces = append(searchTraces, tempSearch)
									}

								}
							}
						}
						break //there's only one state with this name, so now we've found it, stop searching
					}
				}
			}
			chains[evI.Name] = append(chains[evI.Name], searchTraces...)
		}
		//chains = append(chains, chain)
	}

	return chains, nil
}

//DeriveBFBEventChainSet takes an IEC61499 FB and creates a set of event outputs for that single FB, removing duplicates
//e.g. for input "A", you will emit "B", "C"
func DeriveBFBEventChainSet(fb iec61499.FB) (map[string][]ChainOutputs, error) {
	traces, err := DeriveBFBEventTraceSet(fb)
	if err != nil {
		return nil, err
	}
	chains := make(map[string][]ChainOutputs)
	for input, traces := range traces {
		chainOutputSet := make([]ChainOutputs, 0)
		for _, trace := range traces {
			outputs := make(ChainOutputs, 0)
			for _, traceStep := range trace {
				outputs = append(outputs, traceStep.OutputEvents...)
			}

			//check to make sure that no existing chainOutput with this order of output events exists
			match := false
			for i := 0; i < len(chainOutputSet); i++ {
				if len(chainOutputSet[i]) != len(outputs) {
					continue
				}
				for j := 0; j < len(chainOutputSet[i]); j++ {
					if chainOutputSet[i][j] != outputs[j] {
						continue
					}
				}
				//still here? exact match
				match = true
				break
			}
			if match == false {
				chainOutputSet = append(chainOutputSet, outputs)
			}
		}
		chains[input] = chainOutputSet
	}
	return chains, nil
}

//DeriveAllBFBEventChainSets will call DeriveBFBEventChainSet on all BFBs in a set and store it in a helpful map
func DeriveAllBFBEventChainSets(instG []InstanceNode, fbs []iec61499.FB) (map[int]map[string][]ChainOutputs, error) {
	allChains := make(map[int]map[string][]ChainOutputs)
	for instID, inst := range instG {
		instFBT := iec61499.FindBlockDefinitionForType(fbs, inst.FBType)
		if instFBT == nil {
			return nil, errors.New("Bad FB set")
		}
		if instFBT.BasicFB != nil {
			dat, err := DeriveBFBEventChainSet(*instFBT)
			if err != nil {
				return nil, err
			}
			allChains[instID] = dat
		}
	}
	return allChains, nil
}

//ListSIFBEventSources will scan a set of blocks and an instance graph and will return a set of InstanceConnections which
//reflect possible event sources (ie output events on SIFBs)
func ListSIFBEventSources(instG []InstanceNode, fbs []iec61499.FB) ([]InstanceConnection, error) {
	conns := make([]InstanceConnection, 0)
	for instID, inst := range instG {
		instFBT := iec61499.FindBlockDefinitionForType(fbs, inst.FBType)
		if instFBT == nil {
			return nil, errors.New("Bad FB set")
		}
		if instFBT.IsSIFB() {
			for _, outpE := range instFBT.EventOutputs {
				conns = append(conns, InstanceConnection{
					InstanceID: instID,
					PortName:   outpE.Name,
				})
			}
		}
	}
	return conns, nil
}

//An InvokationTrace is used to store a given trace of events in an execution queue,
//as well as the current "position" of the queue (for processing)
//once the processing is complete, position == len(queue)
type InvokationTrace struct {
	Queue    []InstanceConnection
	position int
}

//at each "tick", we have an event queue, and a current event
/*
tick 0:
	add A
tick 1:
	toProcess: A
	add B, C			OR			add Q
tick 2:
	toProcess: B, C		OR			toProcess: Q
	add D							add null
tick 3:
	toProcess: C, D		OR			DONE
	...
*/

//DeriveInstanceInvocationTraceSet will list all possible InstanceInvocationTraces for a given input event
func DeriveInstanceInvocationTraceSet(source InstanceConnection, instG []InstanceNode, fbs []iec61499.FB, allEventChains map[int]map[string][]ChainOutputs) ([]InvokationTrace, error) {
	//determine type of source, is it an input or is it an output?
	//

	// //check if it is input

	traces := make([]InvokationTrace, 0)

	traces = append(traces, InvokationTrace{
		Queue:    []InstanceConnection{source},
		position: 0,
	})

	count := 0

	//foreach destination, match it in the BFB chain set
	for i := 0; i < len(traces); {
		//grab the source node (the trace position)
		if traces[i].position >= len(traces[i].Queue) {
			i++
			continue //we're done with this trace
		}
		source := traces[i].Queue[traces[i].position]

		//classify this node, is it a destination or is it a source?
		sourceFBType := instG[source.InstanceID].FBType
		sourceFBT := iec61499.FindBlockDefinitionForType(fbs, sourceFBType)
		if sourceFBT == nil {
			return nil, errors.New("Bad FB set (#1)")
		}
		sourceIsInput := false
		if sourceFBT.HasEventNamed(true, source.PortName) {
			sourceIsInput = true
		}

		//if the source is an output, then we need to find the destinations and add them to the trace
		if sourceIsInput == false {
			destinations := FindDestinations(source.InstanceID, source.PortName, instG, fbs)

			//add invoked ports to the trace
			//for _, destination := range destinations {
			traces[i].Queue = append(traces[i].Queue, destinations...)
			//}

		} else { //sourceIsInput = true

			//TODO: ensure that we haven't had a self-loop
			//TODO: this can be done roughly at the moment, by just making sure we don't invoke the same block twice
			destination := source
			//otherwise, we need to grab the possible expansions from the destinations

			//all destinations in traces should be penultimate connections (ie BFBs or SIFBs, not CFBs)
			//if they are BFBs, they could continue spawning events in the trace
			destType := instG[destination.InstanceID].FBType
			instFBT := iec61499.FindBlockDefinitionForType(fbs, destType)
			if instFBT == nil {
				return nil, errors.New("Bad FB set")
			}
			if instFBT.BasicFB != nil {
				//load the possible output events for this bfb given this input invokation
				//Then, queue the destinations those events go to
				//And, make sure you're tracing everything at the same time

				bfbEventChains := allEventChains[destination.InstanceID][destination.PortName]
				//for each possible output chain based on the inputted event
				for bfbEventChainI, bfbEventChain := range bfbEventChains {
					if bfbEventChainI == len(bfbEventChains)-1 {
						//linear trace
						for _, bfbOutputEventPort := range bfbEventChain { //for each possible outputted port
							//source: bfbOutputEventPort
							//instance: destination.InstanceID
							outSource := InstanceConnection{
								InstanceID: destination.InstanceID,
								PortName:   bfbOutputEventPort,
							}

							traces[i].Queue = append(traces[i].Queue, outSource)
						}

					} else {
						//branching trace
						tempQueue := make([]InstanceConnection, len(traces[i].Queue))
						copy(tempQueue, traces[i].Queue)
						tempTrace := InvokationTrace{
							Queue:    tempQueue,
							position: traces[i].position + 1,
						}

						for _, bfbOutputEventPort := range bfbEventChain { //for each possible outputted port
							//source: bfbOutputEventPort
							//instance: destination.InstanceID
							outSource := InstanceConnection{
								InstanceID: destination.InstanceID,
								PortName:   bfbOutputEventPort,
							}

							tempTrace.Queue = append(tempTrace.Queue, outSource)
						}

						traces = append(traces, tempTrace)
					}
				}

			}
		}
		//increment the position index
		traces[i].position++

		count++
		if count > 50000 {
			bytes, _ := json.MarshalIndent(traces, "", "\t")
			fmt.Printf("\n%s\n", bytes)
			fmt.Printf("Aborted")
			return nil, errors.New("Aborted due to state space explosion")
		}

	}
	fmt.Printf("Finished after %d iterations\n", count)

	return traces, nil
}
