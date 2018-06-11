# goFB/eca
goFB/eca, or goFBs _Event Chain Analysis_ tool, is for analysing function blocks running under the Event-driven Model of Computation (MoC).

### Primary Function

* Scan FB networks running under Event-Driven MoC to detect infinite loops and derive worst-case invokation chains.

### Methodology

1. For each INPUT EVENT to a FB, compute the set of possible OUTPUT EVENTS.
	* This is done by checking states which depend upon those inputs, and executing them to their penultimate state according to event semantics.
	* It will output a trace of transitions and output events.