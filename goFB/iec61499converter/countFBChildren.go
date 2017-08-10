package iec61499converter

import (
	"errors"

	"github.com/kiwih/goFB/iec61499"
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
		childFBType := findBlockDefinitionForType(fbs, childFBRef.Type)
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
		childFBType := findBlockDefinitionForType(fbs, childFBRef.Type)
		if childFBType == nil {
			return nil, errors.New("Couldn't find instance type")
		}

		childInstanceID := myInstanceID + instanceOffset
		chi, err := FBToInstanceGraph(childFBType, fbs, instanceName+"->"+childFBRef.Name, childInstanceID, myInstanceID)
		if err != nil {
			return nil, err
		}
		nodes[0].ChildNodeIDs = append(nodes[0].ChildNodeIDs, childInstanceID)
		nodes = append(nodes, chi...)
		instanceOffset += (childFBType.NumChildren + 1)

	}

	return nodes, nil
}
