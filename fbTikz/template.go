package main

import (
	"strings"
	"text/template"
)

const tikzTemplateStr = `\documentclass{standalone}
\usepackage{tikz}

\begin{document}
\begin{tikzpicture}[x=5mm,y=-5mm]
{{$width := 7}}
\draw (0,0) -- ({{$width}},0); %top bar

%top sides and event I/O
{{$varLen := .GetEventsSize -}}
{{range $i, $names := .GetMatchedEvents}}
\draw (0,{{$i}}) -- (0,{{add $i 1}}) {{if $names.Input}}node [anchor=west] { {{textsafe $names.Input}} } -- (-1,{{add $i 1}}) {{end}}; %vert line and label and link line if necessary
\draw (({{$width}},{{$i}}) -- (({{$width}},{{add $i 1}}) {{if $names.Output}}node [anchor=east] { {{textsafe $names.Output}} } -- ({{add $width 1}},{{add $i 1}}) {{end}}; 
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
{{range $i, $names := .GetMatchedVars}}
\draw (0,{{add $baseVars $i}}) -- (0,{{add $i (add $baseVars 1)}}) {{if $names.Input}}node [anchor=west,yshift=-0.5] { {{textsafe $names.Input}} } -- (-1,{{add $i (add $baseVars 1)}}) {{end}}; 
\draw ({{$width}},{{add $baseVars $i}}) -- ({{$width}},{{add $i (add $baseVars 1)}}) {{if $names.Output}}node [anchor=east,yshift=-0.5] { {{textsafe $names.Output}} } -- ({{add $width 1}},{{add $i (add $baseVars 1)}}) {{end}}; 
{{end}}

%bottom container
\draw (0,{{add $baseVars $eventLen}}) -- (0,{{add $baseVars (add $eventLen 1)}}); %left
\draw ({{$width}},{{add $baseVars $eventLen}}) -- ({{$width}},{{add $baseVars (add $eventLen 1)}}); %right
\draw (0,{{add $baseVars (add $eventLen 1)}}) -- ({{$width}},{{add $baseVars (add $eventLen 1)}}); %bottom bar

\end{tikzpicture}
\end{document}
`

var tikzTemplateFuncMap = template.FuncMap{
	"add":      add,
	"sub":      sub,
	"textsafe": textsafe,
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func textsafe(s string) string {
	return strings.Replace(s, "_", "\\_", -1)
}

var tikzTemplate = template.Must(template.New("").Funcs(tikzTemplateFuncMap).Parse(tikzTemplateStr))
