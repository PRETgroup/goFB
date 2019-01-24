package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/PRETgroup/goFB/iec61499"
)

//FBTikz is the structure we use to hold everything
type FBTikz struct {
	Blocks []iec61499.FB
}

//OutputFile is used when returning the converted data from the iec61499
type OutputFile struct {
	Name      string
	Extension string
	Contents  []byte
}

//FindBlockForName will return the block and true if it can find a block in an FBTikz blocks
//that matches the provided name
func (f *FBTikz) FindBlockForName(name string) (iec61499.FB, bool) {
	for i := 0; i < len(f.Blocks); i++ {
		if f.Blocks[i].Name == name {
			return f.Blocks[i], true
		}
	}
	return iec61499.FB{}, false
}

//AddBlock should be called for each block in the network
func (f *FBTikz) AddBlock(iec61499bytes []byte) error {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	f.Blocks = append(f.Blocks, FB)

	return nil
}

//ConvertAll will compile all the FBs to Tikz files
func (f *FBTikz) ConvertAll() ([]OutputFile, error) {
	finishedConversions := make([]OutputFile, 0, len(f.Blocks))

	for i := 0; i < len(f.Blocks); i++ {

		output := &bytes.Buffer{}
		origin := FBTikzPoint{0.0, 0.0}
		fbh := NewFBTikzStaticHelper(f.Blocks[i], origin, "")
		if err := tikzTemplate.ExecuteTemplate(output, "tikzBlockIO", fbh); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + f.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: f.Blocks[i].Name + "-tikz", Extension: "tex", Contents: output.Bytes()})
	}

	return finishedConversions, nil
}

//ConvertInternal will render the internals of a single FB to a Tikz file
func (f *FBTikz) ConvertInternal(name string) ([]OutputFile, error) {
	//create dynamic helpers for all blocks and map them relative to their names
	blocks := make(map[string]FBTikzDynamicHelper)
	for _, block := range f.Blocks {
		blocks[block.Name] = NewFBTikzDynamicHelper(block)
	}

	//make sure top block is found
	top, ok := blocks[name]
	if !ok {
		return nil, errors.New("couldn't find block with name '" + name + "'")
	}
	if top.CompositeFB == nil {
		return nil, errors.New("can only draw internals of composite FBs for now")
	}

	//create instance slice for rendering
	instances := make([]FBTikzStaticHelper, len(top.CompositeFB.FBs))

	//define global origin
	origin := FBTikzPoint{0.0, 0.0}

	//for each instance in the network, create in the instance slice
	for i, fbRef := range top.CompositeFB.FBs {
		instDef, ok := blocks[fbRef.Type]
		if !ok {
			return nil, errors.New("couldn't find block reference for name '" + fbRef.Name + "'")
		}
		fbRefX, err := strconv.ParseFloat(fbRef.X, 64)
		if err != nil {
			return nil, errors.New("Problem parsing X position in block ref name'" + fbRef.Name + "': " + err.Error())
		}
		fbRefY, err := strconv.ParseFloat(fbRef.Y, 64)
		if err != nil {
			return nil, errors.New("Problem parsing Y position in block ref name'" + fbRef.Name + "': " + err.Error())
		}
		//todo: determine this correctly
		xGap := 14.0
		yGap := 20.0
		instOrig := origin.Add(xGap*fbRefX, yGap*fbRefY) //line up blocks with a 7 point gap
		instance := instDef.ToStatic(instOrig, fbRef.Name)
		instances[i] = instance
	}

	output := &bytes.Buffer{}
	if err := tikzTemplate.ExecuteTemplate(output, "tikzBlockInternalNetwork", instances); err != nil {
		return nil, errors.New("Couldn't format template (fb) of" + name + ": " + err.Error())
	}

	return []OutputFile{
		OutputFile{Name: name + "-network-tikz", Extension: "tex", Contents: output.Bytes()},
	}, nil
}
