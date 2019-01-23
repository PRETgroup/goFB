package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"

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
		fbh := NewFBTikzHelper(f.Blocks[i], origin)
		if err := tikzTemplate.ExecuteTemplate(output, "", fbh); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + f.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: f.Blocks[i].Name + "-tikz", Extension: "tex", Contents: output.Bytes()})
	}

	return finishedConversions, nil
}

//FBTikzHelper overleads the iec61499.FB type with extra helper functions
type FBTikzHelper struct {
	iec61499.FB
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

func (p FBTikzPoint) String() string {
	return fmt.Sprintf("(%.2f,%.2f)", p.X, p.Y)
}

//Render constants
const (
	NeckInset       float64 = 1
	NeckHeight      float64 = 1
	TextSpacing     float64 = 1
	TextOffset      float64 = 0.5
	MinIOLineLength float64 = 1
	MinBlockWidth   float64 = 7
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
	Origin          FBTikzPoint //x = xlocation, y = ylocation
	Width           float64     // == width of the FB
	EventsHeight    float64     // == height of the events area
	NeckHeight      float64     // == height of the fb "neck"
	DataHeight      float64     // == height of the data area
	TextSpacing     float64     // == spacing between i/o labels
	TextOffset      float64     // == offset from top of area to first text label
	MinIOLineLength float64     // == minimum length of the i/o lines
	EventsInfo      map[string]FBTikzIOInfo
	Border          FBTikzBorderBox
}

//FBTikzIOInfo is used for port info, location, data-variable association
type FBTikzIOInfo struct {
	OffsetY    float64 //the difference between the block origin and the ypos for this position
	LinkX      float64 //if 0, not linked, if !0, this is the position to render the link in
	LinkEventY float64 //if 0, this is the event, and no action required, if !0, this is the height to draw the vertical association link to
}

//NewFBTikzHelper will convert an FB to a FBTikzHelper and calculate all
//necessary render information
func NewFBTikzHelper(fb iec61499.FB, origin FBTikzPoint) FBTikzHelper {
	help := FBTikzHelper{
		FB: fb,
		Points: FBTikzPoints{
			Origin:      origin,
			TextSpacing: TextSpacing,
			TextOffset:  TextOffset,
			NeckHeight:  NeckHeight,
			EventsInfo:  make(map[string]FBTikzIOInfo),
		},
	}

	//calculate width
	help.Points.Width = 7

	//calculate events height
	eventLength := help.GetEventsSize() + 1
	help.Points.EventsHeight = float64(eventLength) * (help.Points.TextSpacing)

	//calculate data height
	dataLength := help.GetVarsSize() + 1
	help.Points.DataHeight = float64(dataLength) * (help.Points.TextSpacing)

	for i, port := range help.FB.InterfaceList.EventInputs {
		info := FBTikzIOInfo{}
		info.OffsetY = help.Points.TextOffset + float64(i)*help.Points.TextSpacing
		help.Points.EventsInfo[port.Name] = info
	}

	for i, port := range help.FB.InterfaceList.EventOutputs {
		info := FBTikzIOInfo{}
		info.OffsetY = help.Points.TextOffset + float64(i)*help.Points.TextSpacing
		help.Points.EventsInfo[port.Name] = info
	}

	//calculate border
	help.Points.Border = FBTikzBorderBox{}

	help.Points.Border.EventsTopLeft = help.Points.Origin
	help.Points.Border.EventsTopRight = help.Points.Origin.AddX(help.Points.Width)
	help.Points.Border.EventsBottomLeft = help.Points.Border.EventsTopLeft.AddY(help.Points.EventsHeight)
	help.Points.Border.EventsBottomRight = help.Points.Border.EventsTopRight.AddY(help.Points.EventsHeight)

	help.Points.Border.NeckTopLeft = help.Points.Border.EventsBottomLeft.AddX(NeckInset)
	help.Points.Border.NeckTopRight = help.Points.Border.EventsBottomRight.AddX(-NeckInset)
	help.Points.Border.NeckBottomLeft = help.Points.Border.NeckTopLeft.AddY(help.Points.NeckHeight)
	help.Points.Border.NeckBottomRight = help.Points.Border.NeckTopRight.AddY(help.Points.NeckHeight)

	help.Points.Border.DataTopLeft = help.Points.Border.EventsBottomLeft.AddY(help.Points.NeckHeight)
	help.Points.Border.DataTopRight = help.Points.Border.EventsBottomRight.AddY(help.Points.NeckHeight)
	help.Points.Border.DataBottomLeft = help.Points.Border.DataTopLeft.AddY(help.Points.DataHeight)
	help.Points.Border.DataBottomRight = help.Points.Border.DataTopRight.AddY(help.Points.DataHeight)

	return help
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
func (f FBTikzHelper) GetVarsSize() int {
	if len(f.InterfaceList.InputVars) > len(f.InterfaceList.OutputVars) {
		return len(f.InterfaceList.InputVars)
	}
	return len(f.InterfaceList.OutputVars)
}

//GetEventsSize returns the height of the vars area that needs to be drawen
func (f FBTikzHelper) GetEventsSize() int {
	if len(f.InterfaceList.EventInputs) > len(f.InterfaceList.EventOutputs) {
		return len(f.InterfaceList.EventInputs)
	}
	return len(f.InterfaceList.EventOutputs)
}

//FBTikzIO will hold the events and data slices for rendering purposes
type FBTikzIO struct {
	Events []FBTikzIONames
	Data   []FBTikzIONames
}

//FBTikzIONames is used for rendering matched IO on a horizontal level
//You can use the InputAssocPos to link between events and data, by using the same index
//and input/output-ness
type FBTikzIONames struct {
	Input       string
	InputAssoc  FBTikzIOAssocPos
	Output      string
	OutputAssoc FBTikzIOAssocPos
}

//FBTikzIOAssocPos is used for data-variable association
type FBTikzIOAssocPos struct {
	PosX     int //if 0, not linked
	PosEvent int //if 0, this is the event
}

//GetTikzIO will return the IO of the block in a helpful and easy-to-render structre
func (f FBTikzHelper) GetTikzIO() FBTikzIO {
	IO := FBTikzIO{
		Events: make([]FBTikzIONames, f.GetEventsSize()),
		Data:   make([]FBTikzIONames, f.GetVarsSize()),
	}

	//sort out the event area for a FB
	inputEventsPos := 0
	outputEventsPos := 0

	inputAssocPos := make(map[string]FBTikzIOAssocPos)
	outputAssocPos := make(map[string]FBTikzIOAssocPos)

	for i := 0; i < len(IO.Events); i++ {
		if i < len(f.InterfaceList.EventInputs) {
			IO.Events[i].Input = f.InterfaceList.EventInputs[i].Name
			if len(f.InterfaceList.EventInputs[i].With) > 0 {
				inputEventsPos++
				IO.Events[i].InputAssoc = FBTikzIOAssocPos{PosX: inputEventsPos, PosEvent: 0}
				for _, with := range f.InterfaceList.EventInputs[i].With {
					inputAssocPos[with.Var] = FBTikzIOAssocPos{PosX: inputEventsPos, PosEvent: i}
				}
			}
		}
		if i < len(f.InterfaceList.EventOutputs) {
			IO.Events[i].Output = f.InterfaceList.EventOutputs[i].Name
			if len(f.InterfaceList.EventOutputs[i].With) > 0 {
				outputEventsPos++
				IO.Events[i].OutputAssoc = FBTikzIOAssocPos{PosX: outputEventsPos, PosEvent: 0}
				for _, with := range f.InterfaceList.EventOutputs[i].With {
					outputAssocPos[with.Var] = FBTikzIOAssocPos{PosX: outputEventsPos, PosEvent: i}
				}
			}
		}
	}

	//fmt.Printf("%+v\r\n", IO.Events)

	//var names
	for i := 0; i < len(IO.Data); i++ {
		if i < len(f.InterfaceList.InputVars) {
			IO.Data[i].Input = f.InterfaceList.InputVars[i].Name
			if pos, ok := inputAssocPos[f.InterfaceList.InputVars[i].Name]; ok {
				IO.Data[i].InputAssoc = pos
			}
		}
		if i < len(f.InterfaceList.OutputVars) {
			IO.Data[i].Output = f.InterfaceList.OutputVars[i].Name
			if pos, ok := outputAssocPos[f.InterfaceList.OutputVars[i].Name]; ok {
				IO.Data[i].OutputAssoc = pos
			}
		}
	}

	return IO
}
