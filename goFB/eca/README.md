# goFB/eca
goFB/eca, or goFBs _Event Chain Analysis_ tool, is for analysing function blocks running under the Event-driven Model of Computation (MoC).

### Primary Function

* Scan FB networks running under Event-Driven MoC to detect infinite loops and derive worst-case invokation chains.

### Example Output

For the railway_station_v2 event MoC project:
```
	Longest Trace analysis:
	[  0]: (o)(EnvTxSystemRx)Top.env.etx.RiChange
	[  1]: (i)(TrainCtrl)Top.ctrl.RiChange
	[  2]: (o)(TrainCtrl)Top.ctrl.WChange
	[  3]: (o)(TrainCtrl)Top.ctrl.SChange
	[  4]: (i)(EnvRxEnforcer)Top.enf.enfRx.WChange_in
	[  5]: (i)(EnvRxEnforcer)Top.enf.enfRx.SChange_in
	[  6]: (o)(EnvRxEnforcer)Top.enf.enfRx.EnforcerAbortedAction
	[  7]: (o)(EnvRxEnforcer)Top.enf.enfRx.EnforcerAbortedAction
	[  8]: (o)(EnvRxEnforcer)Top.enf.enfRx.WChange_out
	[  9]: (o)(EnvRxEnforcer)Top.enf.enfRx.SChange_out
	[ 10]: (o)(EnvRxEnforcer)Top.enf.enfRx.EnforcerAbortedAction
	[ 11]: (o)(EnvRxEnforcer)Top.enf.enfRx.EnforcerAbortedAction
	[ 12]: (o)(EnvRxEnforcer)Top.enf.enfRx.WChange_out
	[ 13]: (o)(EnvRxEnforcer)Top.enf.enfRx.SChange_out
	[ 14]: (i)(EnvRxSystemTx)Top.env.erx.WChange
	[ 15]: (i)(EnvRxSystemTx)Top.env.erx.SChange
	[ 16]: (i)(EnvRxSystemTx)Top.env.erx.WChange
	[ 17]: (i)(EnvRxSystemTx)Top.env.erx.SChange
	REQUIRED EVENT RING BUFFER SIZE: 8 elements
	REQUIRED EVENT LINEAR BUFFER SIZE: 11 elements
```
This indicates:
* The worst case trace is when `RiChange` (the _Request Input_ to the train station) occurs, which is outputted from `env.etx`.
* The enforcer is executed twice, once for `WChange` (_Switch Change_) from the controller, and once for `SChange` (_Signal Change_), which is to be expected.
* The worst-case is when the enforcer corrects two things each time (this is not possible, but only data variable analysis would discover this).
* Then, the environment `env.erx` receives the events to trigger the physical hardware.

### Methodology

1. First, derive the _instance graph_ for the network, which encapsulates all the instances of each type of FB.
2. Then, derive, for each INPUT EVENT to a FB, compute the set of possible OUTPUT EVENTS.
	* This is done by checking states which depend upon those inputs, and executing them to their penultimate state according to event semantics.
	* It will output a trace of transitions and output events.
3. Now list all possible _event sources_, i.e. all events that are emitted from SIFBs.
4. Using the results of (1), (2), and (3), derive the _invokation trace set_, which is a list of all possible traces between all instances of FBs in the network, starting from the SIFB event sources.
    * This is done by taking an event source,
	* Finding its destinations, 
	* And using the possible OUTPUT EVENTS set from those BFBs it invokes,
	* Adding the possible events to the trace (branch/duplicating the trace as necessary).
5. Finally, we can convert the data into easy-to-understand format, and perform analysis on the _invokation trace set_.