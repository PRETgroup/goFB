package iec61499converter

import (
	"errors"

	"github.com/kiwih/goFB/iec61499"
)

var currentInstanceID int

//An InstanceGraph is used with the event MoC to store the unique identifiers for all fbs in the instantiated network
type InstanceGraph struct {
	InstanceID   int
	InstanceName string
	FBType       string
	ChildNodes   []InstanceGraph
}

func nextInstanceID() int {
	c := currentInstanceID
	currentInstanceID++
	return c
}

//FBToInstanceGraph will construct the InstanceGraph for a given FB network
func FBToInstanceGraph(fb *iec61499.FB, fbs []iec61499.FB, instanceName string) (InstanceGraph, error) {
	instG := InstanceGraph{
		InstanceID:   nextInstanceID(),
		InstanceName: instanceName,
		FBType:       fb.Name,
		ChildNodes:   make([]InstanceGraph, 0),
	}

	//define the unprocessed children
	children := make([]iec61499.FBReference, 0)

	if fb.CompositeFB != nil {
		children = append(children, fb.CompositeFB.FBs...)
	}

	if fb.Resources != nil {
		children = append(children, fb.Resources...)
	}

	for _, childFBRef := range children {
		childFBType := findBlockDefinitionForType(fbs, childFBRef.Type)
		if childFBType == nil {
			return InstanceGraph{}, errors.New("Couldn't find instance type")
		}

		chi, err := FBToInstanceGraph(childFBType, fbs, childFBRef.Name)
		if err != nil {
			return InstanceGraph{}, err
		}
		instG.ChildNodes = append(instG.ChildNodes, chi)

	}

	return instG, nil
}
