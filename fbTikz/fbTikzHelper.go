package main

import (
	"fmt"

	"github.com/PRETgroup/goFB/iec61499"
)

//FBTikzDynamicHelper overleads the iec61499.FB type with extra render functions
type FBTikzDynamicHelper struct {
	iec61499.FB
	Characteristics FBTikzCharacteristics
}

//FBTikzStaticHelper overleads the FBTikzDynamicHelper type by saving the info
type FBTikzStaticHelper struct {
	InstanceName string
	Col          int
	Row          int
	FBTikzDynamicHelper
	Points FBTikzPoints
}

//FBTikzPoint represents a point in Tikz-space (i.e. X,Y)
type FBTikzPoint struct {
	X float64
	Y float64
}

//AddX adds x to a given point and returns it as a new point
func (p FBTikzPoint) AddX(x float64) FBTikzPoint {
	return FBTikzPoint{X: p.X + x, Y: p.Y}
}

//AddY adds y to a given point and returns it as a new point
func (p FBTikzPoint) AddY(y float64) FBTikzPoint {
	return FBTikzPoint{X: p.X, Y: p.Y + y}
}

//Add adds x,y to a given point and returns it as a new point
func (p FBTikzPoint) Add(x, y float64) FBTikzPoint {
	return FBTikzPoint{X: p.X + x, Y: p.Y + y}
}

//Sub subtracts x,y from a given point and returns it as a new point
func (p FBTikzPoint) Sub(x, y float64) FBTikzPoint {
	return FBTikzPoint{X: p.X - x, Y: p.Y - y}
}

func (p FBTikzPoint) String() string {
	return fmt.Sprintf("(%.2f,%.2f)", p.X, p.Y)
}

//NonZero is to be used when checking if a point has been set to be 0,0
func (p FBTikzPoint) NonZero() bool {
	return p.X != 0 || p.Y != 0
}

//Render constants
const (
	NeckInset       float64 = 1
	NeckHeight      float64 = 1
	TextSpacing     float64 = 1
	TextOffset      float64 = 1
	MinIOLineLength float64 = 1
	MinBlockWidth   float64 = 7

	FirstLinkAssociationOffset float64 = 0.2
	LinkAssocationSpacing      float64 = 0.2
	LinkAssociationCircleDia   float64 = 0.075
)

/*
                                  width
                |--------------------|

---		origin: +--------------------+   ---
 |				|                    |    |
 |				| a               c  |   --- textoffset, to calculate (origin + offsety) or (origin + textoffset + textspacing * portoffset)
 |				|                    |    |
 |				| b               d  |   --- textspacing
 |		 		|                    |
---eventheight:	+----+          +----+
 | 			 		 |          |
---	neckheight:	+----+          +----+
 |				|                    |
 |				|                 e  |   --- to calculate, (origin + offsety) or (origin + eventheight + neckheight + textoffset + textspacing * portoffset)
 |				|                    |
 |				|                 f  |
 |				|                    |
---	dataheight:	+--------------------+

*/

//FBTikzCharacteristics is used to store information about a given FB that is independent of its origin location
type FBTikzCharacteristics struct {
	Width                    float64 // == width of the FB
	EventsHeight             float64 // == height of the events area
	NeckHeight               float64 // == height of the fb "neck"
	DataHeight               float64 // == height of the data area
	TextSpacing              float64 // == spacing between i/o labels
	TextOffset               float64 // == offset from top of area to first text label
	PortLineLength           float64 // == length of the i/o lines
	LinkAssociationCircleDia float64 // == diameter of link association circles

}

//FBTikzPoints is used to store information about a given FB that depends on the location
type FBTikzPoints struct {
	Origin             FBTikzPoint //x = xlocation, y = ylocation
	InstanceNameAnchor FBTikzPoint
	NameAnchor         FBTikzPoint
	IOInfo             map[string]FBTikzIOInfo
	Border             FBTikzBorderBox
}

//FBTikzIOInfo is used for port info, location, data-variable association
type FBTikzIOInfo struct {
	//OffsetY    float64 //the difference between the block origin and the ypos for this position
	Anchor       FBTikzPoint //where this port is located (left- or right-alignedness is determined by input- or output-ness)
	PortAnchor   FBTikzPoint //where the end of the port is located
	LinkAnchor   FBTikzPoint //if zerod, not linked, if !zerod, this is the position to render the link circle
	LinkLineDest FBTikzPoint //if zerod, don't draw (this is the event), if !zerod, draw a line from LinkAnchor to this position
}

//NewFBTikzDynamicHelper will convert a fb to the dynamic helper type
func NewFBTikzDynamicHelper(fb iec61499.FB) FBTikzDynamicHelper {
	help := FBTikzDynamicHelper{
		FB: fb,
	}
	characteristics := FBTikzCharacteristics{
		TextSpacing:              TextSpacing,
		TextOffset:               TextOffset,
		NeckHeight:               NeckHeight,
		LinkAssociationCircleDia: LinkAssociationCircleDia,
	}
	//todo: calculate width
	characteristics.Width = 10

	//todo: claculate port line width
	characteristics.PortLineLength = 1

	//calculate events height
	eventLength := help.GetEventsSize() + 1
	characteristics.EventsHeight = float64(eventLength) * (characteristics.TextSpacing)

	//calculate data height
	dataLength := help.GetVarsSize() + 1
	characteristics.DataHeight = float64(dataLength) * (characteristics.TextSpacing)
	help.Characteristics = characteristics
	return help
}

//NewFBTikzStaticHelper will convert an FB to a FBTikzHelper and calculate all
//necessary render information
func NewFBTikzStaticHelper(fb iec61499.FB, origin FBTikzPoint, name string) FBTikzStaticHelper {
	help := FBTikzStaticHelper{
		InstanceName:        name,
		FBTikzDynamicHelper: NewFBTikzDynamicHelper(fb),
	}

	help.Points = help.FBTikzDynamicHelper.CalcPoints(origin)
	return help
}

//ToStatic is a helpful wrapper function for CalcPoints,
//and converts a FBTikzDynamicHelper to a FBTikzStaticHelper
func (f FBTikzDynamicHelper) ToStatic(origin FBTikzPoint, name string, col int, row int) FBTikzStaticHelper {
	points := f.CalcPoints(origin)
	return FBTikzStaticHelper{
		InstanceName:        name,
		Col:                 col,
		Row:                 row,
		FBTikzDynamicHelper: f,
		Points:              points,
	}
}

//CalcPoints will calculate all locations of all points of interest inside an FBTikzDynamicHelper
//with reference to a provided origin (which is the top-left of the image)
func (f FBTikzDynamicHelper) CalcPoints(origin FBTikzPoint) FBTikzPoints {
	points := FBTikzPoints{
		Origin: origin.AddX(1),
		IOInfo: make(map[string]FBTikzIOInfo),
	}

	//calculate port locations, association link circle locations, and port text locations
	assocPos := 0

	for i, port := range f.InterfaceList.EventInputs {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.AddY(f.Characteristics.TextOffset + float64(i)*f.Characteristics.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(-f.Characteristics.PortLineLength)
		if len(port.With) > 0 {
			info.LinkAnchor = info.Anchor.AddX(-FirstLinkAssociationOffset - LinkAssocationSpacing*float64(assocPos))
			assocPos++
		}
		points.IOInfo[port.Name] = info
		for _, w := range port.With { //pre-save the linkAnchor destination for the associated vars
			points.IOInfo[w.Var] = FBTikzIOInfo{LinkLineDest: info.LinkAnchor}
		}
	}

	assocPos = 0

	for i, port := range f.InterfaceList.EventOutputs {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.Add(f.Characteristics.Width, f.Characteristics.TextOffset+float64(i)*f.Characteristics.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(f.Characteristics.PortLineLength)
		if len(port.With) > 0 {
			info.LinkAnchor = info.Anchor.AddX(FirstLinkAssociationOffset + LinkAssocationSpacing*float64(assocPos))
			assocPos++
		}
		points.IOInfo[port.Name] = info
		for _, w := range port.With { //pre-save the linkAnchor destination for the associated vars
			points.IOInfo[w.Var] = FBTikzIOInfo{LinkLineDest: info.LinkAnchor}
		}
	}

	for i, port := range f.InterfaceList.InputVars {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.AddY(f.Characteristics.EventsHeight + f.Characteristics.NeckHeight + f.Characteristics.TextOffset + float64(i)*f.Characteristics.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(-f.Characteristics.PortLineLength)
		info.LinkLineDest = points.IOInfo[port.Name].LinkLineDest
		if info.LinkLineDest.NonZero() {
			info.LinkAnchor.X = info.LinkLineDest.X
			info.LinkAnchor.Y = info.PortAnchor.Y
		}
		points.IOInfo[port.Name] = info
	}

	for i, port := range f.InterfaceList.OutputVars {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.Add(f.Characteristics.Width, f.Characteristics.EventsHeight+f.Characteristics.NeckHeight+f.Characteristics.TextOffset+float64(i)*f.Characteristics.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(f.Characteristics.PortLineLength)
		info.LinkLineDest = points.IOInfo[port.Name].LinkLineDest
		if info.LinkLineDest.NonZero() {
			info.LinkAnchor.X = info.LinkLineDest.X
			info.LinkAnchor.Y = info.PortAnchor.Y
		}
		points.IOInfo[port.Name] = info
	}

	points.InstanceNameAnchor = points.Origin.AddX(f.Characteristics.Width / 2)
	points.NameAnchor = points.InstanceNameAnchor.AddY(f.Characteristics.EventsHeight + f.Characteristics.NeckHeight + f.Characteristics.DataHeight)

	//calculate border
	points.Border = FBTikzBorderBox{}

	points.Border.EventsTopLeft = points.Origin
	points.Border.EventsTopRight = points.Origin.AddX(f.Characteristics.Width)
	points.Border.EventsBottomLeft = points.Border.EventsTopLeft.AddY(f.Characteristics.EventsHeight)
	points.Border.EventsBottomRight = points.Border.EventsTopRight.AddY(f.Characteristics.EventsHeight)

	points.Border.NeckTopLeft = points.Border.EventsBottomLeft.AddX(NeckInset)
	points.Border.NeckTopRight = points.Border.EventsBottomRight.AddX(-NeckInset)
	points.Border.NeckBottomLeft = points.Border.NeckTopLeft.AddY(f.Characteristics.NeckHeight)
	points.Border.NeckBottomRight = points.Border.NeckTopRight.AddY(f.Characteristics.NeckHeight)

	points.Border.DataTopLeft = points.Border.EventsBottomLeft.AddY(f.Characteristics.NeckHeight)
	points.Border.DataTopRight = points.Border.EventsBottomRight.AddY(f.Characteristics.NeckHeight)
	points.Border.DataBottomLeft = points.Border.DataTopLeft.AddY(f.Characteristics.DataHeight)
	points.Border.DataBottomRight = points.Border.DataTopRight.AddY(f.Characteristics.DataHeight)

	return points
}

//FBTikzBorderBox contains the list of points that make up the border of the structure
type FBTikzBorderBox struct {
	EventsTopLeft     FBTikzPoint
	EventsTopRight    FBTikzPoint
	EventsBottomLeft  FBTikzPoint
	EventsBottomRight FBTikzPoint
	NeckTopLeft       FBTikzPoint
	NeckTopRight      FBTikzPoint
	NeckBottomLeft    FBTikzPoint
	NeckBottomRight   FBTikzPoint
	DataTopLeft       FBTikzPoint
	DataTopRight      FBTikzPoint
	DataBottomLeft    FBTikzPoint
	DataBottomRight   FBTikzPoint
}

//GetVarsSize returns the height of the vars area that needs to be drawen
func (f FBTikzDynamicHelper) GetVarsSize() int {
	if len(f.InterfaceList.InputVars) > len(f.InterfaceList.OutputVars) {
		return len(f.InterfaceList.InputVars)
	}
	return len(f.InterfaceList.OutputVars)
}

//GetEventsSize returns the height of the vars area that needs to be drawen
func (f FBTikzDynamicHelper) GetEventsSize() int {
	if len(f.InterfaceList.EventInputs) > len(f.InterfaceList.EventOutputs) {
		return len(f.InterfaceList.EventInputs)
	}
	return len(f.InterfaceList.EventOutputs)
}

//FBTikzStaticConnectionBuilder is used to create FBTikzStaticConnections
//so that they don't overlap
type FBTikzStaticConnectionBuilder struct {
	GlobalOrigin FBTikzPoint //will probably be 0,0: represents the top-left of the top-left-most block
	GlobalHeight float64     //represents the distance to the bottom of the lowest block

	BottomHorizWireCount int
	TopHorizWireCount    int
	Columns              []FBTikzStaticConnectionColumn

	Connections []FBTikzStaticConnection
}

func NewFBTikzStaticConnectionBuilder(origin FBTikzPoint, height float64, columns []FBTikzStaticConnectionColumn) FBTikzStaticConnectionBuilder {
	return FBTikzStaticConnectionBuilder{
		TopHorizWireCount:    2, //set the first index for topHorizWires to be "2", meaning they're a bit higher (and look a bit nicer)
		BottomHorizWireCount: 0, //set the first index for bottomHorizWires to be "0", meaning they're a bit higher (and look a bit nicer)
		GlobalOrigin:         origin,
		GlobalHeight:         height,
		Columns:              columns,
	}
}

//FBTikzStaticConnectionColumn is used to store information about columns within the
//FBTikzStaticConnectionBuilder, it allows for renderers to find the column origin
//as well as store the current offset for vertical wires
type FBTikzStaticConnectionColumn struct {
	Origin            FBTikzPoint
	IncomingVertCount int
	OutgoingVertCount int
}

//FBTikzStaticConnection holds the data needed to draw a connection line in
//a FB network
//line points are as follows
//SourceAnchor, IntermediatePoints[0], ... IntermediatePoints[n], DestAnchor
type FBTikzStaticConnection struct {
	SourceText         string //if the connection is to parent, then this will indicate that by being not ""
	SourceAnchor       FBTikzPoint
	DestText           string //if the connection is to parent, then this will indicate that by being not ""
	DestAnchor         FBTikzPoint
	IntermediatePoints []FBTikzPoint //any points that are needed along the way
}

//AddNormalFBTikzStaticConnection will create a new static connection between two normal anchors
//it will endeavor to prevent overlaps on the blocks
//Rules:
//1. if a wire goes from column x to column y=x+1
//it will travel just between these columns, left to right
//it will increment the columnVertWireCounts[x] value and take that position
//2. if a wire goes from column x to column y > x + 1
//it will travel "up and over"
//it will increment both the columnVertWireCounts[x] value and the columnVertWireCounts[y-1] value
//3. if a wire goes from column x to column y <= x
//it will travel "down and under"
//it will increment both the columnVertWireCounts[x] value and the columnVetWireCounts[y] value
func (b *FBTikzStaticConnectionBuilder) AddNormalFBTikzStaticConnection(sourceAnchor FBTikzPoint, sourceBlockCol int, destAnchor FBTikzPoint, destBlockCol int) {

	//ensure we have enough vertWireCounts, fill with zero if not
	for sourceBlockCol > len(b.Columns) {
		panic("sourceCol size exceeds number of columns")
	}
	for destBlockCol > len(b.Columns) {
		panic("destCol size exceeds number of columns")
	}

	WireSpacing := TextOffset

	link := FBTikzStaticConnection{
		SourceAnchor: sourceAnchor,
		DestAnchor:   destAnchor,
	}

	if destBlockCol == sourceBlockCol+1 {
		//case 1, make some intermediate links
		changeAnchor1 := sourceAnchor.AddX(WireSpacing * float64(b.Columns[sourceBlockCol].OutgoingVertCount))
		changeAnchor2 := destAnchor
		changeAnchor2.X = changeAnchor1.X
		b.Columns[sourceBlockCol].OutgoingVertCount++
		link.IntermediatePoints = []FBTikzPoint{changeAnchor1, changeAnchor2}

	} else if destBlockCol > sourceBlockCol+1 {
		//case 2, make some intermediate links for "up and over"
		//travel right to turn up location
		changeAnchor1 := sourceAnchor.AddX(WireSpacing * float64(b.Columns[sourceBlockCol].OutgoingVertCount))
		b.Columns[sourceBlockCol].OutgoingVertCount++

		//travel up to turn right location
		changeAnchor2 := changeAnchor1
		changeAnchor2.Y = b.GlobalOrigin.Y - WireSpacing*float64(b.TopHorizWireCount)
		b.TopHorizWireCount++

		//travel right to turn down location
		changeAnchor3 := changeAnchor2
		changeAnchor3.X = b.Columns[destBlockCol].Origin.X - float64(b.Columns[destBlockCol].IncomingVertCount)*WireSpacing
		b.Columns[destBlockCol].IncomingVertCount++

		//travel down to turn right location
		changeAnchor4 := changeAnchor3
		changeAnchor4.Y = destAnchor.Y

		link.IntermediatePoints = []FBTikzPoint{changeAnchor1, changeAnchor2, changeAnchor3, changeAnchor4}

	} else {
		//case 3, make some intermediate links for "down and under"
		fmt.Printf("SourceCol:%v,DestCol:%v\n", sourceBlockCol, destBlockCol)

		//travel right to turn down location
		changeAnchor1 := sourceAnchor.AddX(WireSpacing * float64(b.Columns[sourceBlockCol].OutgoingVertCount))
		b.Columns[sourceBlockCol].OutgoingVertCount++

		//travel down to turn left location
		changeAnchor2 := changeAnchor1
		changeAnchor2.Y = b.GlobalOrigin.Y + b.GlobalHeight + WireSpacing*float64(b.BottomHorizWireCount)
		b.BottomHorizWireCount++

		//travel left to turn up location
		changeAnchor3 := changeAnchor2
		changeAnchor3.X = b.Columns[destBlockCol].Origin.X - float64(b.Columns[destBlockCol].IncomingVertCount)*WireSpacing
		b.Columns[destBlockCol].IncomingVertCount++

		//travel up to turn left location
		changeAnchor4 := changeAnchor3
		changeAnchor4.Y = destAnchor.Y

		link.IntermediatePoints = []FBTikzPoint{changeAnchor1, changeAnchor2, changeAnchor3, changeAnchor4}

	}

	b.Connections = append(b.Connections, link)

}
