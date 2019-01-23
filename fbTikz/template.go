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
	\draw {{$portInfo.Anchor}} node[anchor=west] { {{textsafe $port.Name}} }; %port text
	\draw [eventWire] {{$portInfo.Anchor}} -- {{$portInfo.PortAnchor}}; %port line
	{{if $portInfo.LinkAnchor.NonZero}}\draw {{$portInfo.LinkAnchor}} circle ({{$fb.Points.LinkAssociationCircleDia}}); %association circle{{end}}
{{end}}

{{range $i, $port := $fb.InterfaceList.EventOutputs}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=east] { {{textsafe $port.Name}} };
	\draw [eventWire] {{$portInfo.Anchor}} -- {{$portInfo.PortAnchor}};
	{{if $portInfo.LinkAnchor.NonZero}}\draw {{$portInfo.LinkAnchor}} circle ({{$fb.Points.LinkAssociationCircleDia}}); %association circle{{end}}
{{end}}

{{range $i, $port := $fb.InterfaceList.InputVars}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=west] { {{textsafe $port.Name}} };
	\draw [dataWire] {{$portInfo.Anchor}} -- {{$portInfo.PortAnchor}};
	{{if $portInfo.LinkAnchor.NonZero}}
		\draw {{$portInfo.LinkAnchor}} circle ({{$fb.Points.LinkAssociationCircleDia}}); %association circle
		\draw [eventWire,dashed] {{$portInfo.LinkAnchor}} -- {{$portInfo.LinkLineDest}};
	{{end}}
{{end}}

{{range $i, $port := $fb.InterfaceList.OutputVars}}
	{{$portInfo := (index $fb.Points.EventsInfo $port.Name)}}
	\draw {{$portInfo.Anchor}} node[anchor=east] { {{textsafe $port.Name}} };
	\draw [dataWire] {{$portInfo.Anchor}} -- {{$portInfo.PortAnchor}};
	{{if $portInfo.LinkAnchor.NonZero}}
		\draw {{$portInfo.LinkAnchor}} circle ({{$fb.Points.LinkAssociationCircleDia}}); %association circle
		\draw [eventWire,dashed] {{$portInfo.LinkAnchor}} -- {{$portInfo.LinkLineDest}};
	{{end}}
{{end}}


\end{tikzpicture}
\end{document}
`

var tikzTemplateFuncMap = template.FuncMap{
	"textsafe": textsafe,
}

func textsafe(s string) string {
	return strings.Replace(s, "_", "\\_", -1)
}

var tikzTemplate = template.Must(template.New("").Funcs(tikzTemplateFuncMap).Parse(tikzTemplateStr))
