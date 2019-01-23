package main

import (
	"strings"
	"text/template"
)

const tikzTemplateStr = `\documentclass{standalone}
\usepackage[rgb]{xcolor}
\usepackage{tikz}

\begin{document}{{$fb := .}}
\begin{tikzpicture}[x=5mm,y=-5mm]
\definecolor{eventWire}{HTML}{6C8EBF}
\definecolor{dataWire}{HTML}{B85450}
{{$border := $fb.Points.Border}}
\draw 	{{$border.EventsTopLeft}} -- 
		{{$border.EventsTopRight}} --
		{{$border.EventsBottomRight}} --
		{{$border.NeckTopRight}} --
		{{$border.NeckBottomRight}} --
		{{$border.DataTopRight}} --
		{{$border.DataBottomRight}} --
		{{$border.DataBottomLeft}} --
		{{$border.DataTopLeft}} --
		{{$border.NeckBottomLeft}} --
		{{$border.NeckTopLeft}} --
		{{$border.EventsBottomLeft}} --
		cycle;

{{range $i, $port := $fb.InterfaceList.EventInputs}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=west] { {{textsafe $port.Name}} };
{{end}}

{{range $i, $port := $fb.InterfaceList.EventOutputs}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=east] { {{textsafe $port.Name}} };
{{end}}

{{range $i, $port := $fb.InterfaceList.InputVars}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=west] { {{textsafe $port.Name}} };
{{end}}

{{range $i, $port := $fb.InterfaceList.OutputVars}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=east] { {{textsafe $port.Name}} };
{{end}}


\end{tikzpicture}
\end{document}
`

var tikzTemplateFuncMap = template.FuncMap{
	"add":      add,
	"sub":      sub,
	"addf":     addf,
	"subf":     subf,
	"mulf":     mulf,
	"intf":     intf,
	"textsafe": textsafe,
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func addf(x, y float64) float64 {
	return x + y
}

func subf(x, y float64) float64 {
	return x - y
}

func mulf(x, y float64) float64 {
	return x * y
}

func intf(x int) float64 {
	return float64(x)
}

func textsafe(s string) string {
	return strings.Replace(s, "_", "\\_", -1)
}

var tikzTemplate = template.Must(template.New("").Funcs(tikzTemplateFuncMap).Parse(tikzTemplateStr))

/*
{{$width := 7}}{{$io := .GetTikzIO}}
\draw (0,0) -- ({{$width}},0); %top bar

%top sides and event I/O
{{$varLen := .GetEventsSize -}}
{{range $i, $names := $io.Events}}
\draw (0,{{$i}}) -- (0,{{add $i 1}}) {{if $names.Input}}node [anchor=west] { {{textsafe $names.Input}} } {{end}}; %vert line and label
{{if $names.Input}}
	\draw [eventWire] (0,{{add $i 1}}) -- (-1,{{add $i 1}}); %link line
	{{if $names.InputAssoc.PosX}}
		\draw ({{subf -0.2 (mulf 0.2 (intf $names.InputAssoc.PosX))}},{{add $i 1}}) circle (0.3mm); %association circle
	{{end}}
{{end}}
\draw (({{$width}},{{$i}}) -- (({{$width}},{{add $i 1}}) {{if $names.Output}}node [anchor=east] { {{textsafe $names.Output}} } {{end}}; %vert line and label
{{if $names.Output}}
	\draw [eventWire] (({{$width}},{{add $i 1}}) -- ({{add $width 1}},{{add $i 1}}); %link line
	{{if $names.OutputAssoc.PosX}}
		\draw ({{addf (addf (intf $width) 0.2) (mulf 0.2 (intf $names.OutputAssoc.PosX))}},{{add $i 1}}) circle (0.3mm); %association circle
	{{end}}
{{end}}
{{end}}

%left indent
\draw (0,{{$varLen}}) -- (0,{{add $varLen 1}});
\draw (0,{{add $varLen 1}}) -- (1,{{add $varLen 1}});
\draw (1,{{add $varLen 1}}) -- (1,{{add $varLen 2}});
\draw (1,{{add $varLen 2}}) -- (0,{{add $varLen 2}});

%right indent
\draw (({{$width}},{{$varLen}}) -- ({{$width}},{{add $varLen 1}});
\draw (({{$width}},{{add $varLen 1}}) -- ({{sub $width 1}},{{add $varLen 1}});
\draw (({{sub $width 1}},{{add $varLen 1}}) -- ({{sub $width 1}},{{add $varLen 2}});
\draw (({{sub $width 1}},{{add $varLen 2}}) -- ({{$width}},{{add $varLen 2}});

%bottom sides and bottom I/O
{{$baseVars := add $varLen 2}}{{$eventLen := .GetVarsSize -}}
{{range $i, $names := $io.Data}}
\draw (0,{{add $baseVars $i}}) -- (0,{{add $i (add $baseVars 1)}}) {{if $names.Input}}node [anchor=west,yshift=-0.5] { {{textsafe $names.Input}} } {{end}}; %vert line and label
{{if $names.Input}}
\draw [dataWire] (0,{{add $i (add $baseVars 1)}}) -- (-1,{{add $i (add $baseVars 1)}}); %link line
	{{if $names.InputAssoc.PosX}}
		\draw ({{subf -0.2 (mulf 0.2 (intf $names.InputAssoc.PosX))}},{{add $i (add $baseVars 1)}}) circle (0.3mm); %association circle
	{{end}}
	\draw [dataWire,dashed] ({{subf -0.2 (mulf 0.2 (intf $names.InputAssoc.PosX))}},{{add $i (add $baseVars 1)}})
		-- ({{subf -0.2 (mulf 0.2 (intf $names.InputAssoc.PosX))}},{{add $names.InputAssoc.PosEvent 1}}); %association line
{{end}}
\draw ({{$width}},{{add $baseVars $i}}) -- ({{$width}},{{add $i (add $baseVars 1)}}) {{if $names.Output}}node [anchor=east,yshift=-0.5] { {{textsafe $names.Output}} } {{end}};
{{if $names.Output}}
\draw [dataWire] ({{$width}},{{add $i (add $baseVars 1)}}) -- ({{add $width 1}},{{add $i (add $baseVars 1)}}); %link line
	{{if $names.OutputAssoc.PosX}}
	\draw ({{addf (addf (intf $width) 0.2) (mulf 0.2 (intf $names.OutputAssoc.PosX))}},{{add $i (add $baseVars 1)}}) circle (0.3mm); %association circle
	{{end}}
	\draw [dataWire,dashed] ({{addf (addf (intf $width) 0.2) (mulf 0.2 (intf $names.OutputAssoc.PosX))}},{{add $i (add $baseVars 1)}})
		-- ({{addf (addf (intf $width) 0.2) (mulf 0.2 (intf $names.OutputAssoc.PosX))}},{{add $names.InputAssoc.PosEvent 1}}); %association line

{{end}}
{{end}}

%bottom container
\draw (0,{{add $baseVars $eventLen}}) -- (0,{{add $baseVars (add $eventLen 1)}}); %left
\draw ({{$width}},{{add $baseVars $eventLen}}) -- ({{$width}},{{add $baseVars (add $eventLen 1)}}); %right
\draw (0,{{add $baseVars (add $eventLen 1)}}) -- ({{$width}},{{add $baseVars (add $eventLen 1)}}); %bottom bar


*/
