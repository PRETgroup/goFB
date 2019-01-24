package main

import (
	"fmt"

	"github.com/PRETgroup/goFB/iec61499"
)

//FBTikzDynamicHelper overleads the iec61499.FB type with extra render functions
type FBTikzDynamicHelper iec61499.FB

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

//FBTikzPoints is used to store information about where everything is for a given FB
type FBTikzPoints struct {
	Origin                   FBTikzPoint //x = xlocation, y = ylocation
	Width                    float64     // == width of the FB
	EventsHeight             float64     // == height of the events area
	NeckHeight               float64     // == height of the fb "neck"
	DataHeight               float64     // == height of the data area
	TextSpacing              float64     // == spacing between i/o labels
	TextOffset               float64     // == offset from top of area to first text label
	PortLineLength           float64     // == length of the i/o lines
	LinkAssociationCircleDia float64     // == diameter of link association circles
	InstanceNameAnchor       FBTikzPoint
	NameAnchor               FBTikzPoint
	IOInfo                   map[string]FBTikzIOInfo
	Border                   FBTikzBorderBox
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
	return FBTikzDynamicHelper(fb)
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
		Origin:                   origin.AddX(1),
		TextSpacing:              TextSpacing,
		TextOffset:               TextOffset,
		NeckHeight:               NeckHeight,
		LinkAssociationCircleDia: LinkAssociationCircleDia,
		IOInfo:                   make(map[string]FBTikzIOInfo),
	}
	//todo: calculate width
	points.Width = 10

	//todo: claculate port line width
	points.PortLineLength = 1

	//calculate events height
	eventLength := f.GetEventsSize() + 1
	points.EventsHeight = float64(eventLength) * (points.TextSpacing)

	//calculate data height
	dataLength := f.GetVarsSize() + 1
	points.DataHeight = float64(dataLength) * (points.TextSpacing)

	//calculate port locations, association link circle locations, and port text locations
	assocPos := 0

	for i, port := range f.InterfaceList.EventInputs {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.AddY(points.TextOffset + float64(i)*points.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(-points.PortLineLength)
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
		info.Anchor = points.Origin.Add(points.Width, points.TextOffset+float64(i)*points.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(points.PortLineLength)
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
		info.Anchor = points.Origin.AddY(points.EventsHeight + points.NeckHeight + points.TextOffset + float64(i)*points.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(-points.PortLineLength)
		info.LinkLineDest = points.IOInfo[port.Name].LinkLineDest
		if info.LinkLineDest.NonZero() {
			info.LinkAnchor.X = info.LinkLineDest.X
			info.LinkAnchor.Y = info.PortAnchor.Y
		}
		points.IOInfo[port.Name] = info
	}

	for i, port := range f.InterfaceList.OutputVars {
		info := FBTikzIOInfo{}
		info.Anchor = points.Origin.Add(points.Width, points.EventsHeight+points.NeckHeight+points.TextOffset+float64(i)*points.TextSpacing)
		info.PortAnchor = info.Anchor.AddX(points.PortLineLength)
		info.LinkLineDest = points.IOInfo[port.Name].LinkLineDest
		if info.LinkLineDest.NonZero() {
			info.LinkAnchor.X = info.LinkLineDest.X
			info.LinkAnchor.Y = info.PortAnchor.Y
		}
		points.IOInfo[port.Name] = info
	}

	points.InstanceNameAnchor = points.Origin.AddX(points.Width / 2)
	points.NameAnchor = points.InstanceNameAnchor.AddY(points.EventsHeight + points.NeckHeight + points.DataHeight)

	//calculate border
	points.Border = FBTikzBorderBox{}

	points.Border.EventsTopLeft = points.Origin
	points.Border.EventsTopRight = points.Origin.AddX(points.Width)
	points.Border.EventsBottomLeft = points.Border.EventsTopLeft.AddY(points.EventsHeight)
	points.Border.EventsBottomRight = points.Border.EventsTopRight.AddY(points.EventsHeight)

	points.Border.NeckTopLeft = points.Border.EventsBottomLeft.AddX(NeckInset)
	points.Border.NeckTopRight = points.Border.EventsBottomRight.AddX(-NeckInset)
	points.Border.NeckBottomLeft = points.Border.NeckTopLeft.AddY(points.NeckHeight)
	points.Border.NeckBottomRight = points.Border.NeckTopRight.AddY(points.NeckHeight)

	points.Border.DataTopLeft = points.Border.EventsBottomLeft.AddY(points.NeckHeight)
	points.Border.DataTopRight = points.Border.EventsBottomRight.AddY(points.NeckHeight)
	points.Border.DataBottomLeft = points.Border.DataTopLeft.AddY(points.DataHeight)
	points.Border.DataBottomRight = points.Border.DataTopRight.AddY(points.DataHeight)

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

func NewFBTikzStaticConnectionBuilder(origin FBTikzPoint, columns []FBTikzStaticConnectionColumn) FBTikzStaticConnectionBuilder {
	return FBTikzStaticConnectionBuilder{
		GlobalOrigin: origin,
		Columns:      columns,
	}
}

//FBTikzStaticConnectionColumn is used to store information about columns within the
//FBTikzStaticConnectionBuilder, it allows for renderers to find the column origin
//as well as store the current offset for vertical wires
type FBTikzStaticConnectionColumn struct {
	Origin        FBTikzPoint
	VertWireCount int
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

	fmt.Printf("SourceCol:%v,DestCol:%v\n", sourceBlockCol, destBlockCol)

	if destBlockCol == sourceBlockCol+1 {
		//case 1, make some intermediate links
		changeAnchor1 := sourceAnchor.AddX(WireSpacing * float64(b.Columns[sourceBlockCol*3].VertWireCount+1))
		changeAnchor2 := destAnchor
		changeAnchor2.X = changeAnchor1.X
		b.Columns[sourceBlockCol*3].VertWireCount++
		link.IntermediatePoints = []FBTikzPoint{changeAnchor1, changeAnchor2}
		b.Connections = append(b.Connections, link)
	} else if destBlockCol > sourceBlockCol+1 {
		//case 2, make some intermediate links for "up and over"
		// changeAnchor1 := sourceAnchor.AddX(WireSpacing * float64(b.ColumnVertWireCounts[sourceCol]))
		// b.ColumnVertWireCounts[sourceCol]++

		// changeAnchor2 := b.GlobalOrigin.AddY(WireSpacing * float64(b.TopHorizWireCount))
		// b.TopHorizWireCount++
		// changeAnchor2.X = changeAnchor1.X

		// changeAnchor4 :=
		// changeAnchor2 := destAnchor
		// changeAnchor2.X = changeAnchor1.X
		// b.ColumnVertWireCounts[sourceCol]++
		// link.IntermediatePoints = []FBTikzPoint{changeAnchor1, changeAnchor2}

		//panic("ahahaha")
	} else {
		//case 3, make some intermediate links for "down and under"
		//panic("nonononon")
	}

}
