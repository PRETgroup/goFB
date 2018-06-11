package eca

import (
	"errors"
	"strings"

	"github.com/PRETgroup/goFB/iec61499"
)

//ComputeFBChildrenCounts counts and stores all FB children in all FBtypes across the network
func ComputeFBChildrenCounts(fbs []iec61499.FB) error {
	for i := 0; i < len(fbs); i++ {
		fbs[i].NumChildren = -1 //mark uncounted
	}

	for i := 0; i < len(fbs); i++ {
		_, err := GetFBChildrenCounts(&fbs[i], fbs)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetFBChildrenCounts recursively counts the number of fbChildren a fbtype has
func GetFBChildrenCounts(fb *iec61499.FB, fbs []iec61499.FB) (int, error) {
	if fb.NumChildren != -1 {
		return fb.NumChildren, nil
	}
	count := 0
	//define the unprocessed children
	children := make([]iec61499.FBReference, 0)

	if fb.CompositeFB != nil {
		children = append(children, fb.CompositeFB.FBs...)
	}

	if fb.Resources != nil {
		children = append(children, fb.Resources...)
	}

	for _, childFBRef := range children {
		childFBType := iec61499.FindBlockDefinitionForType(fbs, childFBRef.Type)
		if childFBType == nil {
			return 0, errors.New("Couldn't find instance type")
		}

		nChi, err := GetFBChildrenCounts(childFBType, fbs)
		if err != nil {
			return 0, err
		}
		count += nChi + 1
	}

	fb.NumChildren = count
	return count, nil
}

//An InstanceNode is used with the event MoC to store the unique identifiers for all fbs in the instantiated network
type InstanceNode struct {
	InstanceID   int
	ParentID     int //if ParentID == InstanceID then it has no parent, i.e. is the top
	InstanceName string
	FBType       string
	ChildNodeIDs []int
}

//FBToInstanceGraph will construct the []InstanceNode for a given FB network
func FBToInstanceGraph(fb *iec61499.FB, fbs []iec61499.FB, instanceName string, myInstanceID int, parentID int) ([]InstanceNode, error) {
	nodes := make([]InstanceNode, 0)

	me := InstanceNode{
		InstanceID:   myInstanceID,
		ParentID:     parentID,
		InstanceName: instanceName,
		FBType:       fb.Name,
		ChildNodeIDs: make([]int, 0),
	}

	nodes = append(nodes, me)

	//define the unprocessed children
	children := make([]iec61499.FBReference, 0)

	if fb.CompositeFB != nil {
		children = append(children, fb.CompositeFB.FBs...)
	}

	if fb.Resources != nil {
		children = append(children, fb.Resources...)
	}

	instanceOffset := 1

	for _, childFBRef := range children {
		childFBType := iec61499.FindBlockDefinitionForType(fbs, childFBRef.Type)
		if childFBType == nil {
			return nil, errors.New("Couldn't find instance type")
		}

		childInstanceID := myInstanceID + instanceOffset
		chi, err := FBToInstanceGraph(childFBType, fbs, childFBRef.Name, childInstanceID, myInstanceID)
		if err != nil {
			return nil, err
		}
		nodes[0].ChildNodeIDs = append(nodes[0].ChildNodeIDs, childInstanceID)
		nodes = append(nodes, chi...)
		instanceOffset += (childFBType.NumChildren + 1)

	}

	return nodes, nil
}

//InstanceConnection is used to help us find connections globally between BFBs
type InstanceConnection struct {
	InstanceID int    //ID of an instance in an InstanceGraph
	PortName   string //name of port on instance
}

//FindDestinations : Given an instance and an output port, find the instances where that port go
func FindDestinations(fromInstanceID int, fromName string, instG []InstanceNode, fbs []iec61499.FB) []InstanceConnection {
	myInst := instG[fromInstanceID]
	me := iec61499.FindBlockDefinitionForType(fbs, myInst.FBType)
	if me == nil {
		return nil
	}

	unresolvedConns := make([]InstanceConnection, 0)  //we'll queue mystery connections as we find them
	destinationConns := make([]InstanceConnection, 0) //final connections loaded in here

	unresolvedConns = append(unresolvedConns, InstanceConnection{InstanceID: fromInstanceID, PortName: fromName})

	for len(unresolvedConns) > 0 {
		//pop the first unresolved connection
		var unresolvedConn InstanceConnection
		unresolvedConn, unresolvedConns = unresolvedConns[0], unresolvedConns[1:]
		examineInst := instG[unresolvedConn.InstanceID]
		examineFB := iec61499.FindBlockDefinitionForType(fbs, examineInst.FBType)
		if examineFB == nil {
			//whoops, something bad has happened, we can't resolve the instance ID
			return nil
		}
		if examineFB.CompositeFB == nil {
			//if this is an input to this block, we're finished here
			found := false
			for _, eventIn := range examineFB.EventInputs {
				if eventIn.Name == unresolvedConn.PortName {
					found = true
					break
				}
			}
			if found == false {
				for _, dataIn := range examineFB.InputVars {
					if dataIn.Name == unresolvedConn.PortName {
						found = true
						break
					}
				}
			}
			if found == true {
				destinationConns = append(destinationConns,
					InstanceConnection{
						InstanceID: unresolvedConn.InstanceID,
						PortName:   unresolvedConn.PortName,
					})
				continue
			}

			//if this is an output from this block, it is also unresolved, and needs to be fixed at a higher level
			//we can't do much with this at this level
			unresolvedConns = append(unresolvedConns,
				InstanceConnection{
					InstanceID: examineInst.ParentID,
					PortName:   examineInst.InstanceName + "." + unresolvedConn.PortName,
				})
			continue
		}
		//we are a composite FB now, and we're examining a source port
		//(either an input to this block, or an output of a child block)
	inner: //we need an inner loop because one source might have multiple destinations
		for _, conn := range append(examineFB.CompositeFB.DataConnections, examineFB.CompositeFB.EventConnections...) {
			if conn.Source == unresolvedConn.PortName {
				//the destination of this connection might be a resolution to this
				//we'll check later
				if strings.Contains(conn.Destination, ".") {
					//this goes to a child block
					destParts := strings.Split(conn.Destination, ".")
					childInstName := destParts[0]
					childPortName := destParts[1]

					for _, childInstID := range examineInst.ChildNodeIDs {
						if instG[childInstID].InstanceName == childInstName {
							unresolvedConns = append(unresolvedConns, InstanceConnection{
								InstanceID: childInstID,
								PortName:   childPortName,
							})
							continue inner
						}
					}
					//if we're still here, something has derped, and we can't find a matching instance name
					return nil
				}
				//okay, the destination is an output port of this CFB
				//parent will need to resolve this
				unresolvedConns = append(unresolvedConns, InstanceConnection{
					InstanceID: examineInst.ParentID,
					PortName:   examineInst.InstanceName + "." + conn.Destination,
				})
				continue inner
			}
		}
	}

	//job's done
	return destinationConns
}

//FindSources : Given an instance and an input port, find the instances where that port comes from
func FindSources(toInstanceID int, toName string, instG []InstanceNode, fbs []iec61499.FB) []InstanceConnection {
	myInst := instG[toInstanceID]
	me := iec61499.FindBlockDefinitionForType(fbs, myInst.FBType)
	if me == nil {
		return nil
	}

	unresolvedConns := make([]InstanceConnection, 0) //we'll queue mystery connections as we find them
	sourceConns := make([]InstanceConnection, 0)     //final connections loaded in here

	unresolvedConns = append(unresolvedConns, InstanceConnection{InstanceID: toInstanceID, PortName: toName})

	for len(unresolvedConns) > 0 {
		//pop the first unresolved connection
		var unresolvedConn InstanceConnection
		unresolvedConn, unresolvedConns = unresolvedConns[0], unresolvedConns[1:]
		examineInst := instG[unresolvedConn.InstanceID]
		examineFB := iec61499.FindBlockDefinitionForType(fbs, examineInst.FBType)
		if examineFB == nil {
			//whoops, something bad has happened, we can't resolve the instance ID
			return nil
		}
		if examineFB.CompositeFB == nil {
			//if this is an output from this block, we're finished here
			found := false
			for _, eventIn := range examineFB.EventOutputs {
				if eventIn.Name == unresolvedConn.PortName {
					found = true
					break
				}
			}
			if found == false {
				for _, dataIn := range examineFB.OutputVars {
					if dataIn.Name == unresolvedConn.PortName {
						found = true
						break
					}
				}
			}
			if found == true {
				sourceConns = append(sourceConns,
					InstanceConnection{
						InstanceID: unresolvedConn.InstanceID,
						PortName:   unresolvedConn.PortName,
					})
				continue
			}

			//if this is an input to this block, it is also unresolved, and needs to be fixed at a higher level
			//we can't do much with this at this level
			unresolvedConns = append(unresolvedConns,
				InstanceConnection{
					InstanceID: examineInst.ParentID,
					PortName:   examineInst.InstanceName + "." + unresolvedConn.PortName,
				})
			continue
		}
		//we are a composite FB now, and we're examining a destination port
		//(either an output from this block, or an input of a child block)
	inner: //we need an inner loop because one source might have multiple destinations
		for _, conn := range append(examineFB.CompositeFB.DataConnections, examineFB.CompositeFB.EventConnections...) {
			if conn.Destination == unresolvedConn.PortName {
				//the source of this connection might be a resolution to this
				//we'll check later
				if strings.Contains(conn.Source, ".") {
					//this goes to a child block
					sourceParts := strings.Split(conn.Source, ".")
					childInstName := sourceParts[0]
					childPortName := sourceParts[1]

					for _, childInstID := range examineInst.ChildNodeIDs {
						if instG[childInstID].InstanceName == childInstName {
							unresolvedConns = append(unresolvedConns, InstanceConnection{
								InstanceID: childInstID,
								PortName:   childPortName,
							})
							continue inner
						}
					}
					//if we're still here, something has derped, and we can't find a matching instance name
					return nil
				}
				//okay, the source is an input port of this CFB
				//parent will need to resolve this
				unresolvedConns = append(unresolvedConns, InstanceConnection{
					InstanceID: examineInst.ParentID,
					PortName:   examineInst.InstanceName + "." + conn.Source,
				})
				continue inner
			}
		}
	}

	//job's done
	return sourceConns
}
