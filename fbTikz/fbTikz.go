package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"sort"
	"strconv"
	"strings"

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
	instances := make(map[string]FBTikzStaticHelper)

	//define global origin
	origin := FBTikzPoint{0.0, 0.0}

	//determine the number of columns (3 columns for each "x" in fbRef)
	//and determine the total height
	//column 1 = incoming vertical wires
	//column 2 = the blocks
	//column 3 = outgoing vertical wires

	//capture all unique column labels for X position, and increment a counter each time it's
	//captured
	type CountAndIndex struct {
		FBNames          []string
		Index            int //Index position of column (0, 1, 2, etc)
		NumIncomingVerts int //number of slots required for incoming vertical wires (either down or up)
		NumOutgoingVerts int //number of slots required for outgoing vertical wires (either down or up)
	}
	colLabelCounts := make(map[float64]CountAndIndex)
	for _, fbRef := range top.CompositeFB.FBs {
		fbRefX, err := strconv.ParseFloat(fbRef.X, 64)
		if err != nil {
			return nil, errors.New("Problem parsing X position in block ref name'" + fbRef.Name + "': " + err.Error())
		}
		cai := colLabelCounts[fbRefX]
		cai.FBNames = append(cai.FBNames, fbRef.Name)
		colLabelCounts[fbRefX] = cai
	}
	//then sort those unique labels in order
	var colLabels []float64
	for key := range colLabelCounts {
		colLabels = append(colLabels, key)
	}
	sort.Float64s(colLabels) //using colLabels, we can now work out which index a given column is in
	//and put them back in the map
	for i, key := range colLabels {
		cai := colLabelCounts[key]
		cai.Index = i
		colLabelCounts[key] = cai
	}

	//find the maximum height for a given label
	maxColHeight := float64(0)
	for _, key := range colLabels {
		cai := colLabelCounts[key]
		colHeight := float64(0)
	out3:
		for _, sname := range cai.FBNames {
			for _, ref := range top.CompositeFB.FBs {
				if ref.Name == sname {
					def := blocks[ref.Type]
					colHeight += def.Characteristics.DataHeight + def.Characteristics.NeckHeight + def.Characteristics.DataHeight + 2
					break out3
				}
			}
		}
		if colHeight > maxColHeight {
			maxColHeight = colHeight
		}
	}

	//there are len(colLabels)*3 columns
	//to determine their widths, we need to know how many incoming and outgoing wires there
	//will be for each block in the column
	//.. for each column, for each block, classify incoming wires and count, classify outgoing wires and count
	//.. if a wire connects any two blocks in the network it counts as a vertical wire

	//.. for each connection, find the source and destination links, and increment the appropriate counters
	for _, fbConn := range top.CompositeFB.EventConnections {
		sourceParts := strings.Split(fbConn.Source, ".")
		destParts := strings.Split(fbConn.Destination, ".")

		if len(sourceParts) == 2 && len(destParts) == 2 {
			foundSource := false
			foundDest := false

		out:
			for _, key := range colLabels {
				cai := colLabelCounts[key]

				for _, name := range cai.FBNames {
					if name == sourceParts[0] {
						cai.NumOutgoingVerts++
						colLabelCounts[key] = cai
						foundSource = true
					}
					if name == destParts[0] {
						cai.NumIncomingVerts++
						colLabelCounts[key] = cai
						foundDest = true
					}
					if foundSource && foundDest {
						break out
					}

				}
			}
		}
	}

	//now we have a nice structure with all the column information, so let's make the columns that we need
	columns := make([]FBTikzStaticConnectionColumn, 0, len(colLabels)*3)
	ColumnSpacing := TextSpacing

	//this mechanism to uses columns of blocks,
	// which have spacing between them to ensure that they have enough room for vertical wires
	nextOrigin := origin
	for _, key := range colLabels {
		cai := colLabelCounts[key]

		nextOrigin = nextOrigin.AddX(float64(cai.NumIncomingVerts+1) * ColumnSpacing)
		columnBlock := FBTikzStaticConnectionColumn{
			Origin:                      nextOrigin,
			IncomingFromTopVertCount:    1, //starting offsets for incoming and outgoing verts
			IncomingFromBottomVertCount: cai.NumIncomingVerts,

			OutgoingToTopVertCount:    1, //starting offsets for incoming and outgoing verts
			OutgoingToBottomVertCount: cai.NumOutgoingVerts,
		}
		nextOrigin = columnBlock.Origin.AddX(float64(12) * ColumnSpacing)

		nextOrigin = nextOrigin.AddX(float64(cai.NumOutgoingVerts+1) * ColumnSpacing)

		columns = append(columns, columnBlock)
	}

	//for each instance in the network, create in the instance slice
	for _, fbRef := range top.CompositeFB.FBs {
		instDef, ok := blocks[fbRef.Type]
		if !ok {
			return nil, errors.New("couldn't find block reference for name '" + fbRef.Name + "'")
		}

		//find the column that they're in
		refColIndex := -1

	out2:
		for _, key := range colLabels {
			cai := colLabelCounts[key]
			for _, name := range cai.FBNames {
				if name == fbRef.Name {
					refColIndex = cai.Index
					break out2
				}
			}
		}
		if refColIndex == -1 {
			return nil, errors.New("Something went wrong\r\n")
		}
		refCol := columns[refColIndex]

		fbRefY, err := strconv.ParseFloat(fbRef.Y, 64)
		if err != nil {
			return nil, errors.New("Problem parsing Y position in block ref name'" + fbRef.Name + "': " + err.Error())
		}

		yGap := 20.0 //todo: fix
		instOrig := refCol.Origin.AddY(yGap * fbRefY)

		instance := instDef.ToStatic(instOrig, fbRef.Name, refColIndex, int(fbRefY)) //todo: the "int" casts here are egregariously bad
		instances[fbRef.Name] = instance
	}

	//create connections

	//1. create links that go "up and over"
	// --- (from TOP to BOTTOM of column)

	//2. create links that go "down and under"
	// --- (from BOTTOM to TOP of column)

	//3. cr??

	staticLinksFactory := NewFBTikzStaticConnectionBuilder(origin, maxColHeight, columns)
	eventConnections := make([]FBTikzStaticConnection, 0, len(top.CompositeFB.EventConnections))
	for _, fbConn := range top.CompositeFB.EventConnections {
		if strings.Contains(fbConn.Source, ".") && strings.Contains(fbConn.Destination, ".") {
			sourceParts := strings.Split(fbConn.Source, ".")
			destParts := strings.Split(fbConn.Destination, ".")
			sourceInstance, ok := instances[sourceParts[0]]
			if !ok {
				return nil, errors.New("couldn't find block reference for name '" + sourceParts[0] + "'")
			}
			destInstance, ok := instances[destParts[0]]
			if !ok {
				return nil, errors.New("couldn't find block reference for name '" + destParts[0] + "'")
			}

			sourcePort, ok := sourceInstance.Points.IOInfo[sourceParts[1]]
			if !ok {
				return nil, errors.New("couldn't find block port for name '" + sourceParts[1] + "'")
			}
			destPort, ok := destInstance.Points.IOInfo[destParts[1]]
			if !ok {
				return nil, errors.New("couldn't find block port for name '" + destParts[1] + "'")
			}

			staticLinksFactory.AddNormalFBTikzStaticConnection(sourcePort.PortAnchor, sourceInstance.Col, destPort.PortAnchor, destInstance.Col)
		}
	}
	eventConnections = append(eventConnections, staticLinksFactory.Connections...)

	type TikzBlockInternalNetwork struct {
		Instances   map[string]FBTikzStaticHelper
		Connections []FBTikzStaticConnection
	}

	render := TikzBlockInternalNetwork{
		Instances:   instances,
		Connections: eventConnections,
	}

	output := &bytes.Buffer{}
	if err := tikzTemplate.ExecuteTemplate(output, "tikzBlockInternalNetwork", render); err != nil {
		return nil, errors.New("Couldn't format template (fb) of" + name + ": " + err.Error())
	}

	return []OutputFile{
		OutputFile{Name: name + "-network-tikz", Extension: "tex", Contents: output.Bytes()},
	}, nil
}
