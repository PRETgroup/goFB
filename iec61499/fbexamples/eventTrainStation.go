package fbexamples

//EventTrainStationFBT is a simple example system designed to be used with an event-driven MoC
//The FBT files are provided here for usage in examples and automated testing
var EventTrainStationFBT = []string{`<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
	<FBType Name="Enforcer" Comment="">
		<Identification Standard="61499-2"></Identification>
		<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
		<CompilerInfo header="" classdef=""></CompilerInfo>
		<InterfaceList>
			<EventInputs>
				<Event Name="SysReady_env_tx" Comment=""></Event>
				<Event Name="RiChange_env_tx" Comment=""></Event>
				<Event Name="RnChange_env_tx" Comment=""></Event>
				<Event Name="RsChange_env_tx" Comment=""></Event>
				<Event Name="DwiChange_env_tx" Comment=""></Event>
				<Event Name="DwoChange_env_tx" Comment=""></Event>
				<Event Name="DwnChange_env_tx" Comment=""></Event>
				<Event Name="DwsChange_env_tx" Comment=""></Event>
				<Event Name="DnChange_env_tx" Comment=""></Event>
				<Event Name="DsChange_env_tx" Comment=""></Event>
				<Event Name="SoChange_env_tx" Comment=""></Event>
				<Event Name="SChange_ctrl_tx" Comment=""></Event>
				<Event Name="WChange_ctrl_tx" Comment=""></Event>
			</EventInputs>
			<EventOutputs>
				<Event Name="SysReady_ctrl_rx" Comment=""></Event>
				<Event Name="RiChange_ctrl_rx" Comment=""></Event>
				<Event Name="RnChange_ctrl_rx" Comment=""></Event>
				<Event Name="RsChange_ctrl_rx" Comment=""></Event>
				<Event Name="DwiChange_ctrl_rx" Comment=""></Event>
				<Event Name="DwoChange_ctrl_rx" Comment=""></Event>
				<Event Name="DwnChange_ctrl_rx" Comment=""></Event>
				<Event Name="DwsChange_ctrl_rx" Comment=""></Event>
				<Event Name="DnChange_ctrl_rx" Comment=""></Event>
				<Event Name="DsChange_ctrl_rx" Comment=""></Event>
				<Event Name="SoChange_ctrl_rx" Comment=""></Event>
				<Event Name="SChange_env_rx" Comment=""></Event>
				<Event Name="WChange_env_rx" Comment=""></Event>
			</EventOutputs>
			<InputVars>
				<VarDeclaration Name="RiReq_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="RnReq_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="RsReq_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwiPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwoPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwnPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwsPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DnPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DsPrs_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SoGrn_env_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SiGrn_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SnGrn_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SsGrn_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WiDvrg_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WoDvrg_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WnDvrg_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WsDvrg_ctrl_tx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			</InputVars>
			<OutputVars>
				<VarDeclaration Name="RiReq_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="RnReq_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="RsReq_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwiPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwoPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwnPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DwsPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DnPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="DsPrs_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SoGrn_ctrl_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SiGrn_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SnGrn_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="SsGrn_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WiDvrg_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WoDvrg_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WnDvrg_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
				<VarDeclaration Name="WsDvrg_env_rx" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			</OutputVars>
		</InterfaceList>
		<FBNetwork>
			<FB Name="enfRx" Type="EnvRxEnforcer" x="" y=""></FB>
			<EventConnections>
				<Connection Source="SoChange_env_tx" Destination="enfRx.SoChange" dx1=""></Connection>
				<Connection Source="SChange_ctrl_tx" Destination="enfRx.SChange_in" dx1=""></Connection>
				<Connection Source="WChange_ctrl_tx" Destination="enfRx.WChange_in" dx1=""></Connection>
				<Connection Source="DwiChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="DwoChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="DwnChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="DwsChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="DnChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="DsChange_env_tx" Destination="enfRx.DChange_in" dx1=""></Connection>
				<Connection Source="RiChange_env_tx" Destination="RiChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="RnChange_env_tx" Destination="RnChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="RsChange_env_tx" Destination="RsChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="SysReady_env_tx" Destination="SysReady_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwiChange_env_tx" Destination="DwiChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwoChange_env_tx" Destination="DwoChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwnChange_env_tx" Destination="DwnChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwsChange_env_tx" Destination="DwsChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="DnChange_env_tx" Destination="DnChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="DsChange_env_tx" Destination="DsChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="SoChange_env_tx" Destination="SoChange_ctrl_rx" dx1=""></Connection>
				<Connection Source="enfRx.SChange_out" Destination="SChange_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.WChange_out" Destination="WChange_env_rx" dx1=""></Connection>
			</EventConnections>
			<DataConnections>
				<Connection Source="SoGrn_env_tx" Destination="enfRx.SoGrn" dx1=""></Connection>
				<Connection Source="SiGrn_ctrl_tx" Destination="enfRx.SiGrn_in" dx1=""></Connection>
				<Connection Source="SnGrn_ctrl_tx" Destination="enfRx.SnGrn_in" dx1=""></Connection>
				<Connection Source="SsGrn_ctrl_tx" Destination="enfRx.SsGrn_in" dx1=""></Connection>
				<Connection Source="WiDvrg_ctrl_tx" Destination="enfRx.WiDvrg_in" dx1=""></Connection>
				<Connection Source="WoDvrg_ctrl_tx" Destination="enfRx.WoDvrg_in" dx1=""></Connection>
				<Connection Source="WnDvrg_ctrl_tx" Destination="enfRx.WnDvrg_in" dx1=""></Connection>
				<Connection Source="WsDvrg_ctrl_tx" Destination="enfRx.WsDvrg_in" dx1=""></Connection>
				<Connection Source="DwiPrs_env_tx" Destination="enfRx.DwiPrs_in" dx1=""></Connection>
				<Connection Source="DwoPrs_env_tx" Destination="enfRx.DwoPrs_in" dx1=""></Connection>
				<Connection Source="DwnPrs_env_tx" Destination="enfRx.DwnPrs_in" dx1=""></Connection>
				<Connection Source="DwsPrs_env_tx" Destination="enfRx.DwsPrs_in" dx1=""></Connection>
				<Connection Source="DnPrs_env_tx" Destination="enfRx.DnPrs_in" dx1=""></Connection>
				<Connection Source="DsPrs_env_tx" Destination="enfRx.DsPrs_in" dx1=""></Connection>
				<Connection Source="enfRx.SiGrn_out" Destination="SiGrn_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.SnGrn_out" Destination="SnGrn_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.SsGrn_out" Destination="SsGrn_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.WiDvrg_out" Destination="WiDvrg_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.WoDvrg_out" Destination="WoDvrg_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.WnDvrg_out" Destination="WnDvrg_env_rx" dx1=""></Connection>
				<Connection Source="enfRx.WsDvrg_out" Destination="WsDvrg_env_rx" dx1=""></Connection>
				<Connection Source="RiReq_env_tx" Destination="RiReq_ctrl_rx" dx1=""></Connection>
				<Connection Source="RnReq_env_tx" Destination="RnReq_ctrl_rx" dx1=""></Connection>
				<Connection Source="RsReq_env_tx" Destination="RsReq_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwiPrs_env_tx" Destination="DwiPrs_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwoPrs_env_tx" Destination="DwoPrs_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwnPrs_env_tx" Destination="DwnPrs_ctrl_rx" dx1=""></Connection>
				<Connection Source="DwsPrs_env_tx" Destination="DwsPrs_ctrl_rx" dx1=""></Connection>
				<Connection Source="DnPrs_env_tx" Destination="DnPrs_ctrl_rx" dx1=""></Connection>
				<Connection Source="DsPrs_env_tx" Destination="DsPrs_ctrl_rx" dx1=""></Connection>
			</DataConnections>
		</FBNetwork>
	</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="Environment" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs>
			<Event Name="SChange" Comment=""></Event>
			<Event Name="WChange" Comment=""></Event>
		</EventInputs>
		<EventOutputs>
			<Event Name="SysReady" Comment=""></Event>
			<Event Name="RiChange" Comment=""></Event>
			<Event Name="RnChange" Comment=""></Event>
			<Event Name="RsChange" Comment=""></Event>
			<Event Name="DwiChange" Comment=""></Event>
			<Event Name="DwoChange" Comment=""></Event>
			<Event Name="DwnChange" Comment=""></Event>
			<Event Name="DwsChange" Comment=""></Event>
			<Event Name="DnChange" Comment=""></Event>
			<Event Name="DsChange" Comment=""></Event>
			<Event Name="SoChange" Comment=""></Event>
		</EventOutputs>
		<InputVars>
			<VarDeclaration Name="SiGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InputVars>
		<OutputVars>
			<VarDeclaration Name="RiReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RnReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RsReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SoGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</OutputVars>
	</InterfaceList>
	<FBNetwork>
		<FB Name="etx" Type="EnvTxSystemRx" x="" y=""></FB>
		<FB Name="erx" Type="EnvRxSystemTx" x="" y=""></FB>
		<EventConnections>
			<Connection Source="etx.RiChange" Destination="RiChange" dx1=""></Connection>
			<Connection Source="etx.RnChange" Destination="RnChange" dx1=""></Connection>
			<Connection Source="etx.RsChange" Destination="RsChange" dx1=""></Connection>
			<Connection Source="etx.DwiChange" Destination="DwiChange" dx1=""></Connection>
			<Connection Source="etx.DwoChange" Destination="DwoChange" dx1=""></Connection>
			<Connection Source="etx.DwnChange" Destination="DwnChange" dx1=""></Connection>
			<Connection Source="etx.DwsChange" Destination="DwsChange" dx1=""></Connection>
			<Connection Source="etx.DnChange" Destination="DnChange" dx1=""></Connection>
			<Connection Source="etx.DsChange" Destination="DsChange" dx1=""></Connection>
			<Connection Source="etx.SoChange" Destination="SoChange" dx1=""></Connection>
			<Connection Source="SChange" Destination="erx.SChange" dx1=""></Connection>
			<Connection Source="etx.SoChange" Destination="erx.SChange" dx1=""></Connection>
			<Connection Source="WChange" Destination="erx.WChange" dx1=""></Connection>
			<Connection Source="etx.DwiChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.DwoChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.DwnChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.DwsChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.DnChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.DsChange" Destination="erx.DChangeD" dx1=""></Connection>
			<Connection Source="etx.SysReady" Destination="SysReady" dx1=""></Connection>
		</EventConnections>
		<DataConnections>
			<Connection Source="etx.RiReq" Destination="RiReq" dx1=""></Connection>
			<Connection Source="etx.RnReq" Destination="RnReq" dx1=""></Connection>
			<Connection Source="etx.RsReq" Destination="RsReq" dx1=""></Connection>
			<Connection Source="etx.DwiPrs" Destination="DwiPrs" dx1=""></Connection>
			<Connection Source="etx.DwoPrs" Destination="DwoPrs" dx1=""></Connection>
			<Connection Source="etx.DwnPrs" Destination="DwnPrs" dx1=""></Connection>
			<Connection Source="etx.DwsPrs" Destination="DwsPrs" dx1=""></Connection>
			<Connection Source="etx.DnPrs" Destination="DnPrs" dx1=""></Connection>
			<Connection Source="etx.DsPrs" Destination="DsPrs" dx1=""></Connection>
			<Connection Source="etx.SoGrn" Destination="SoGrn" dx1=""></Connection>
			<Connection Source="SiGrn" Destination="erx.SiGrn" dx1=""></Connection>
			<Connection Source="etx.SoGrn" Destination="erx.SoGrn" dx1=""></Connection>
			<Connection Source="SnGrn" Destination="erx.SnGrn" dx1=""></Connection>
			<Connection Source="SsGrn" Destination="erx.SsGrn" dx1=""></Connection>
			<Connection Source="WiDvrg" Destination="erx.WiDvrg" dx1=""></Connection>
			<Connection Source="WoDvrg" Destination="erx.WoDvrg" dx1=""></Connection>
			<Connection Source="WnDvrg" Destination="erx.WnDvrg" dx1=""></Connection>
			<Connection Source="WsDvrg" Destination="erx.WsDvrg" dx1=""></Connection>
			<Connection Source="etx.DwiPrs" Destination="erx.DwiPrsD" dx1=""></Connection>
			<Connection Source="etx.DwoPrs" Destination="erx.DwoPrsD" dx1=""></Connection>
			<Connection Source="etx.DwnPrs" Destination="erx.DwnPrsD" dx1=""></Connection>
			<Connection Source="etx.DwsPrs" Destination="erx.DwsPrsD" dx1=""></Connection>
			<Connection Source="etx.DnPrs" Destination="erx.DnPrsD" dx1=""></Connection>
			<Connection Source="etx.DsPrs" Destination="erx.DsPrsD" dx1=""></Connection>
		</DataConnections>
	</FBNetwork>
</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="EnvRxEnforcer" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs>
			<Event Name="SoChange" Comment="">
				<With Var="SoGrn"></With>
			</Event>
			<Event Name="SChange_in" Comment="">
				<With Var="SiGrn_in"></With>
				<With Var="SnGrn_in"></With>
				<With Var="SsGrn_in"></With>
			</Event>
			<Event Name="WChange_in" Comment="">
				<With Var="WiDvrg_in"></With>
				<With Var="WoDvrg_in"></With>
				<With Var="WnDvrg_in"></With>
				<With Var="WsDvrg_in"></With>
			</Event>
			<Event Name="DChange_in" Comment="">
				<With Var="DwiPrs_in"></With>
				<With Var="DwoPrs_in"></With>
				<With Var="DwnPrs_in"></With>
				<With Var="DwsPrs_in"></With>
				<With Var="DnPrs_in"></With>
				<With Var="DsPrs_in"></With>
			</Event>
		</EventInputs>
		<EventOutputs>
			<Event Name="EnforcerAbortedAction" Comment=""></Event>
			<Event Name="SChange_out" Comment="">
				<With Var="SiGrn_out"></With>
				<With Var="SnGrn_out"></With>
				<With Var="SsGrn_out"></With>
			</Event>
			<Event Name="WChange_out" Comment="">
				<With Var="WiDvrg_out"></With>
				<With Var="WoDvrg_out"></With>
				<With Var="WnDvrg_out"></With>
				<With Var="WsDvrg_out"></With>
			</Event>
		</EventOutputs>
		<InputVars>
			<VarDeclaration Name="SoGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SiGrn_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrs_in" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InputVars>
		<OutputVars>
			<VarDeclaration Name="SiGrn_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg_out" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</OutputVars>
	</InterfaceList>
	<BasicFB>
		<InternalVars>
			<VarDeclaration Name="SiGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InternalVars>
		<ECC>
			<ECState Name="idle" Comment="" x="" y="">
				<ECAction Algorithm="idle_alg0"></ECAction>
			</ECState>
			<ECState Name="copyVals" Comment="" x="" y="">
				<ECAction Algorithm="copyInVals"></ECAction>
			</ECState>
			<ECState Name="startCheck" Comment="" x="" y=""></ECState>
			<ECState Name="singleGreenSiSnSs" Comment="" x="" y=""></ECState>
			<ECState Name="singleGreenSiSnSs_correction" Comment="" x="" y="">
				<ECAction Output="EnforcerAbortedAction"></ECAction>
				<ECAction Algorithm="singleGreenSiSnSs_correction_alg0"></ECAction>
			</ECState>
			<ECState Name="SoRequiredForSnSs" Comment="" x="" y="">
				<ECAction Algorithm="SoRequiredForSnSs_alg0"></ECAction>
			</ECState>
			<ECState Name="SoRequiredForSnSs_correction" Comment="" x="" y="">
				<ECAction Output="EnforcerAbortedAction"></ECAction>
				<ECAction Algorithm="SoRequiredForSnSs_correction_alg0"></ECAction>
				<ECAction Algorithm="SoRequiredForSnSs_correction_alg1"></ECAction>
			</ECState>
			<ECState Name="WiWsMatch" Comment="" x="" y=""></ECState>
			<ECState Name="WiWsMatch_correction" Comment="" x="" y="">
				<ECAction Algorithm="WiWsMatch_correction_alg0"></ECAction>
			</ECState>
			<ECState Name="WiWnMatch" Comment="" x="" y=""></ECState>
			<ECState Name="WiWnMatch_correction" Comment="" x="" y="">
				<ECAction Algorithm="WiWnMatch_correction_alg0"></ECAction>
			</ECState>
			<ECState Name="SsWMatch" Comment="" x="" y=""></ECState>
			<ECState Name="SsWMatch_correction" Comment="" x="" y="">
				<ECAction Algorithm="SsWMatch_correction_alg0"></ECAction>
			</ECState>
			<ECState Name="SnWMatch" Comment="" x="" y=""></ECState>
			<ECState Name="SnWMatch_correction" Comment="" x="" y="">
				<ECAction Algorithm="SnWMatch_correction_alg0"></ECAction>
			</ECState>
			<ECState Name="finished" Comment="" x="" y="">
				<ECAction Output="WChange_out"></ECAction>
				<ECAction Output="SChange_out"></ECAction>
				<ECAction Algorithm="copyOutVals"></ECAction>
			</ECState>
			<ECTransition Source="idle" Destination="copyVals" Condition="( SChange_in || WChange_in )" x="" y=""></ECTransition>
			<ECTransition Source="copyVals" Destination="startCheck" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="startCheck" Destination="singleGreenSiSnSs" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="singleGreenSiSnSs" Destination="singleGreenSiSnSs_correction" Condition="( SiGrn + SnGrn + SsGrn ) &gt; 1" x="" y=""></ECTransition>
			<ECTransition Source="singleGreenSiSnSs" Destination="SoRequiredForSnSs" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="singleGreenSiSnSs_correction" Destination="SoRequiredForSnSs" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SoRequiredForSnSs" Destination="SoRequiredForSnSs_correction" Condition="SoGrn == 0 &amp;&amp; ( SnGrn == 1 || SsGrn == 1 )" x="" y=""></ECTransition>
			<ECTransition Source="SoRequiredForSnSs" Destination="WiWsMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SoRequiredForSnSs_correction" Destination="WiWsMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="WiWsMatch" Destination="WiWsMatch_correction" Condition="SiGrn == 1 &amp;&amp; WiDvrg == 1 &amp;&amp; WsDvrg != 1" x="" y=""></ECTransition>
			<ECTransition Source="WiWsMatch" Destination="WiWnMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="WiWsMatch_correction" Destination="WiWnMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="WiWnMatch" Destination="WiWnMatch_correction" Condition="SiGrn == 1 &amp;&amp; WiDvrg == 0 &amp;&amp; WnDvrg != 0" x="" y=""></ECTransition>
			<ECTransition Source="WiWnMatch" Destination="SsWMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="WiWnMatch_correction" Destination="SsWMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SsWMatch" Destination="SsWMatch_correction" Condition="SsGrn == 1 &amp;&amp; ( WsDvrg != 0 || WoDvrg != 0 )" x="" y=""></ECTransition>
			<ECTransition Source="SsWMatch" Destination="SnWMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SsWMatch_correction" Destination="SnWMatch" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SnWMatch" Destination="SnWMatch_correction" Condition="SnGrn == 1 &amp;&amp; ( WnDvrg != 1 || WoDvrg != 1 )" x="" y=""></ECTransition>
			<ECTransition Source="SnWMatch" Destination="finished" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="SnWMatch_correction" Destination="finished" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="finished" Destination="idle" Condition="true" x="" y=""></ECTransition>
		</ECC>
		<Algorithm Name="idle_alg0" Comment="">
			<Other Language="C" Text="printf(&#34;idle\r\n&#34;);"></Other>
		</Algorithm>
		<Algorithm Name="singleGreenSiSnSs_correction_alg0" Comment="">
			<Other Language="C" Text="me-&gt;SiGrn = me-&gt;SiGrn_out; me-&gt;SnGrn = me-&gt;SnGrn_out; me-&gt;SsGrn = me-&gt;SsGrn_out;"></Other>
		</Algorithm>
		<Algorithm Name="SoRequiredForSnSs_alg0" Comment="">
			<Other Language="C" Text="printf(&#34;Enforcer: SoRequiredForSnSs\r\n&#34;);"></Other>
		</Algorithm>
		<Algorithm Name="SoRequiredForSnSs_correction_alg0" Comment="">
			<Other Language="C" Text="printf(&#34;Enforcer: SoRequiredForSnSs_correction\r\n&#34;);"></Other>
		</Algorithm>
		<Algorithm Name="SoRequiredForSnSs_correction_alg1" Comment="">
			<Other Language="C" Text="me-&gt;SnGrn = 0; me-&gt;SsGrn = 0;"></Other>
		</Algorithm>
		<Algorithm Name="WiWsMatch_correction_alg0" Comment="">
			<Other Language="C" Text="me-&gt;WsDvrg = 1;"></Other>
		</Algorithm>
		<Algorithm Name="WiWnMatch_correction_alg0" Comment="">
			<Other Language="C" Text="me-&gt;WnDvrg = 0;"></Other>
		</Algorithm>
		<Algorithm Name="SsWMatch_correction_alg0" Comment="">
			<Other Language="C" Text="me-&gt;WsDvrg = 0; me-&gt;WoDvrg = 0;"></Other>
		</Algorithm>
		<Algorithm Name="SnWMatch_correction_alg0" Comment="">
			<Other Language="C" Text="me-&gt;WnDvrg = 1; me-&gt;WoDvrg = 1;"></Other>
		</Algorithm>
		<Algorithm Name="copyInVals" Comment="">
			<Other Language="C" Text="&#xA;&#x9;&#x9;//printf(&#34;copying in vals\r\n&#34;);&#xA;&#x9;&#x9;me-&gt;SiGrn = me-&gt;SiGrn_in;&#xA;&#x9;&#x9;me-&gt;SnGrn = me-&gt;SnGrn_in;&#xA;&#x9;&#x9;me-&gt;SsGrn = me-&gt;SsGrn_in;&#xA;&#xA;&#x9;&#x9;me-&gt;WiDvrg = me-&gt;WiDvrg_in;&#xA;&#x9;&#x9;me-&gt;WoDvrg = me-&gt;WoDvrg_in;&#xA;&#x9;&#x9;me-&gt;WnDvrg = me-&gt;WnDvrg_in;&#xA;&#x9;&#x9;me-&gt;WsDvrg = me-&gt;WsDvrg_in;&#xA;&#x9;"></Other>
		</Algorithm>
		<Algorithm Name="copyOutVals" Comment="">
			<Other Language="C" Text="&#xA;&#x9;&#x9;//printf(&#34;copying out vals\r\n&#34;);&#xA;&#x9;&#x9;me-&gt;SiGrn_out = me-&gt;SiGrn;&#xA;&#x9;&#x9;me-&gt;SnGrn_out = me-&gt;SnGrn;&#xA;&#x9;&#x9;me-&gt;SsGrn_out = me-&gt;SsGrn;&#xA;&#xA;&#x9;&#x9;me-&gt;WiDvrg_out = me-&gt;WiDvrg;&#xA;&#x9;&#x9;me-&gt;WoDvrg_out = me-&gt;WoDvrg;&#xA;&#x9;&#x9;me-&gt;WnDvrg_out = me-&gt;WnDvrg;&#xA;&#x9;&#x9;me-&gt;WsDvrg_out = me-&gt;WsDvrg;&#xA;&#x9;"></Other>
		</Algorithm>
	</BasicFB>
</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="EnvRxSystemTx" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs>
			<Event Name="SChange" Comment="">
				<With Var="SiGrn"></With>
				<With Var="SoGrn"></With>
				<With Var="SnGrn"></With>
				<With Var="SsGrn"></With>
			</Event>
			<Event Name="WChange" Comment="">
				<With Var="WiDvrg"></With>
				<With Var="WoDvrg"></With>
				<With Var="WnDvrg"></With>
				<With Var="WsDvrg"></With>
			</Event>
			<Event Name="DChangeD" Comment="">
				<With Var="DwiPrsD"></With>
				<With Var="DwoPrsD"></With>
				<With Var="DwnPrsD"></With>
				<With Var="DwsPrsD"></With>
				<With Var="DnPrsD"></With>
				<With Var="DsPrsD"></With>
			</Event>
		</EventInputs>
		<EventOutputs></EventOutputs>
		<InputVars>
			<VarDeclaration Name="SiGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SoGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrsD" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InputVars>
		<OutputVars></OutputVars>
	</InterfaceList>
	<Service>
		<Autogenerate Language="C" ArbitraryText="&#xA;#include &lt;stdio.h&gt;&#xA;#include &lt;stdlib.h&gt;&#xA;" InStructText="" PreInitText="" InitText="" RunText="&#xA;//if any change occurs&#xA;if(me-&gt;inputEvents.event.SChange || &#xA;        me-&gt;inputEvents.event.WChange || &#xA;        me-&gt;inputEvents.event.DChangeD) {&#xA;&#xA;    printf(&#34;\r\n\r\n\r\n&#34;&#xA;    &#34;         Si%s     Dwi%s Wi%s  Wn%s Dwn%s                Dn%s       \r\n&#34;&#xA;    &#34;in -&gt; ---%s----------[%s]----=%s----%s=----[%s]---%s--------------[%s]-----------X\r\n&#34;&#xA;    &#34;                              \\  /             Sn%s                          \r\n&#34;&#xA;    &#34;                               \\/                                             \r\n&#34;&#xA;    &#34;                               /\\                                             \r\n&#34;&#xA;    &#34;                              /  \\                                             \r\n&#34;&#xA;    &#34;out&lt;- ---%s----------[%s]----=%s----%s=----[%s]---%s--------------[%s]-----------X\r\n&#34;&#xA;    &#34;         So%s     Dwo%s Wo%s  Ws%s Dws%s  Ss%s         Ds%s          \r\n&#34;,&#xA;    &#xA;    me-&gt;SiGrn ? TF_GRN &#34;(g)&#34; T_NRM : TF_RED &#34;(r)&#34; T_NRM,&#xA;    me-&gt;DwiPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;,&#xA;    me-&gt;WiDvrg ? TF_YEL &#34;(d)&#34; T_NRM : &#34;(s)&#34;,&#xA;    me-&gt;WnDvrg ? TF_YEL &#34;(d)&#34; T_NRM : &#34;(s)&#34;,&#xA;    me-&gt;DwnPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;,&#xA;    me-&gt;DnPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;,&#xA;    &#xA;    me-&gt;SiGrn ? TB_GRN &#34;&gt;&#34; T_NRM : TB_RED &#34;&gt;&#34; T_NRM,&#xA;    me-&gt;DwiPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;    me-&gt;WiDvrg ? TF_YEL &#34;\\&#34; T_NRM : TF_YEL &#34;-&#34; T_NRM,&#xA;    me-&gt;WnDvrg ? TF_YEL &#34;/&#34; T_NRM : TF_YEL &#34;-&#34; T_NRM,&#xA;    me-&gt;DwnPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;    me-&gt;SnGrn ? TB_GRN &#34;&lt;&#34; T_NRM : TB_RED &#34;&lt;&#34; T_NRM,&#xA;    me-&gt;DnPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;    &#xA;    me-&gt;SnGrn ? TF_GRN &#34;(g)&#34; T_NRM : TF_RED &#34;(r)&#34; T_NRM,&#xA;    &#xA;    me-&gt;SoGrn ? TB_GRN &#34;&lt;&#34; T_NRM : TB_RED &#34;&lt;&#34; T_NRM,&#xA;    me-&gt;DwoPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;    me-&gt;WoDvrg ? TF_YEL &#34;/&#34; T_NRM : TF_YEL &#34;-&#34; T_NRM,&#xA;    me-&gt;WsDvrg ? TF_YEL &#34;\\&#34; T_NRM : TF_YEL &#34;-&#34; T_NRM,&#xA;    me-&gt;DwsPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;    me-&gt;SsGrn ? TB_GRN &#34;&lt;&#34; T_NRM : TB_RED &#34;&lt;&#34; T_NRM,&#xA;    me-&gt;DsPrsD ? TF_CYN &#34;HH&#34; T_NRM : &#34;  &#34;,&#xA;&#xA;    me-&gt;SoGrn ? TF_GRN &#34;(g)&#34; T_NRM : TF_RED &#34;(r)&#34; T_NRM,&#xA;    me-&gt;DwoPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;,&#xA;    me-&gt;WoDvrg ? TF_YEL &#34;(d)&#34; T_NRM : &#34;(s)&#34;,&#xA;    me-&gt;WsDvrg ? TF_YEL &#34;(d)&#34; T_NRM : &#34;(s)&#34;,&#xA;    me-&gt;DwsPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;,&#xA;    me-&gt;SsGrn ? TF_GRN &#34;(g)&#34; T_NRM : TF_RED &#34;(r)&#34; T_NRM,&#xA;    me-&gt;DsPrsD ? TF_CYN &#34;(p)&#34; T_NRM : &#34;( )&#34;);&#xA;}&#xA;&#xA;&#xA;" ShutdownText=""></Autogenerate>
	</Service>
</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="EnvTxSystemRx" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs></EventInputs>
		<EventOutputs>
			<Event Name="RiChange" Comment="">
				<With Var="RiReq"></With>
			</Event>
			<Event Name="RnChange" Comment="">
				<With Var="RnReq"></With>
			</Event>
			<Event Name="RsChange" Comment="">
				<With Var="RsReq"></With>
			</Event>
			<Event Name="DwiChange" Comment="">
				<With Var="DwiPrs"></With>
			</Event>
			<Event Name="DwoChange" Comment="">
				<With Var="DwoPrs"></With>
			</Event>
			<Event Name="DwnChange" Comment="">
				<With Var="DwnPrs"></With>
			</Event>
			<Event Name="DwsChange" Comment="">
				<With Var="DwsPrs"></With>
			</Event>
			<Event Name="DnChange" Comment="">
				<With Var="DnPrs"></With>
			</Event>
			<Event Name="DsChange" Comment="">
				<With Var="DsPrs"></With>
			</Event>
			<Event Name="SoChange" Comment="">
				<With Var="SoGrn"></With>
			</Event>
			<Event Name="SysReady" Comment=""></Event>
		</EventOutputs>
		<InputVars></InputVars>
		<OutputVars>
			<VarDeclaration Name="RiReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RnReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RsReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SoGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</OutputVars>
	</InterfaceList>
	<Service>
		<Autogenerate Language="C" ArbitraryText="&#xA;#include &lt;stdio.h&gt;&#xA;#include &lt;stdlib.h&gt;&#xA;" InStructText="&#xA;int last_RiReq, last_RnReq, last_RsReq, last_DwiPrs, last_DwoPrs, last_DwnPrs, last_DwsPrs, last_DnPrs, last_DsPrs, last_SoGrn;&#xA;int initialised;&#xA;" PreInitText="" InitText="" RunText="&#xA;FILE *file;&#xA;file = fopen(&#34;commands.txt&#34;, &#34;r&#34;);&#xA;if (file) {&#xA;    int c;&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DwiPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DwiPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DwiChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 3);&#xA;        //}&#xA;        me-&gt;DwiPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DwiPrs event and set to %i\r\n&#34;, me-&gt;DwiPrs);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DwoPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DwoPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DwoChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 4);&#xA;        //}&#xA;        me-&gt;DwoPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DwoPrs event and set to %i\r\n&#34;, me-&gt;DwoPrs);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DwnPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DwnPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DwnChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 5);&#xA;        //}&#xA;        me-&gt;DwnPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DwnPrs event and set to %i\r\n&#34;, me-&gt;DwnPrs);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DwsPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DwsPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DwsChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 6);&#xA;        //}&#xA;        me-&gt;DwsPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DwsPrs event and set to %i\r\n&#34;, me-&gt;DwsPrs);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DnPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DnPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DnChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 7);&#xA;        //}&#xA;        me-&gt;DnPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DnPrs event and set to %i\r\n&#34;, me-&gt;DnPrs);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_DsPrs &amp;&amp; (c == &#39;-&#39; || c == &#39;p&#39;)) {&#xA;        me-&gt;last_DsPrs = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.DsChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 8);&#xA;        //}&#xA;        me-&gt;DsPrs = (c == &#39;p&#39;);&#xA;        printf(&#34;EnvTxSystemRx: DsPrs event and set to %i\r\n&#34;, me-&gt;DsPrs);&#xA;    }    &#xA;&#xA;    //get rid of space&#xA;    c = getc(file);&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_RiReq &amp;&amp; (c == &#39;-&#39; || c == &#39;r&#39;)) {&#xA;        me-&gt;last_RiReq = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.RiChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 0);&#xA;        //}&#xA;        me-&gt;RiReq = (c == &#39;r&#39;);&#xA;        printf(&#34;EnvTxSystemRx: RiReq event and set to %i\r\n&#34;, me-&gt;RiReq);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_RnReq &amp;&amp; (c == &#39;-&#39; || c == &#39;r&#39;)) {&#xA;        me-&gt;last_RnReq = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.RnChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 1);&#xA;        //}&#xA;        me-&gt;RnReq = (c == &#39;r&#39;);&#xA;        printf(&#34;EnvTxSystemRx: RnReq event and set to %i\r\n&#34;, me-&gt;RnReq);&#xA;    }&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_RsReq &amp;&amp; (c == &#39;-&#39; || c == &#39;r&#39;)) {&#xA;        me-&gt;last_RsReq = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.RsChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 2);&#xA;        //}&#xA;        me-&gt;RsReq = (c == &#39;r&#39;);&#xA;        printf(&#34;EnvTxSystemRx: RsReq event and set to %i\r\n&#34;, me-&gt;RsReq);&#xA;    }&#xA;&#xA;    //get rid of space&#xA;    c = getc(file);&#xA;&#xA;    c = getc(file);&#xA;    if(c != me-&gt;last_SoGrn &amp;&amp; (c == &#39;r&#39; || c == &#39;g&#39;)) {&#xA;        me-&gt;last_SoGrn = c;&#xA;        //if(me-&gt;initialised) {&#xA;            me-&gt;outputEvents.event.SoChange = 1;&#xA;            PushEvent(me-&gt;myInstanceID, 9);&#xA;        //}&#xA;        me-&gt;SoGrn = (c == &#39;g&#39;);&#xA;        printf(&#34;EnvTxSystemRx: SoChange event and set to %i\r\n&#34;, me-&gt;SoGrn);&#xA;    }&#xA;&#xA;    if(me-&gt;initialised == 0) {&#xA;        me-&gt;initialised = 1;&#xA;        me-&gt;outputEvents.event.SysReady = 1;&#xA;        PushEvent(me-&gt;myInstanceID, 10);&#xA;    }&#xA;    &#xA;    &#xA;    fclose(file);&#xA;} else {&#xA;    printf(&#34;EnvTxSystemRx: Couldn&#39;t open commands.txt\n&#34;);&#xA;    while(1);&#xA;}&#xA;" ShutdownText=""></Autogenerate>
	</Service>
</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="Top" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs></EventInputs>
		<EventOutputs></EventOutputs>
		<InputVars></InputVars>
		<OutputVars></OutputVars>
	</InterfaceList>
	<FBNetwork>
		<FB Name="ctrl" Type="TrainCtrl" x="" y=""></FB>
		<FB Name="env" Type="Environment" x="" y=""></FB>
		<FB Name="enf" Type="Enforcer" x="" y=""></FB>
		<EventConnections>
			<Connection Source="enf.RiChange_ctrl_rx" Destination="ctrl.RiChange" dx1=""></Connection>
			<Connection Source="enf.RnChange_ctrl_rx" Destination="ctrl.RnChange" dx1=""></Connection>
			<Connection Source="enf.RsChange_ctrl_rx" Destination="ctrl.RsChange" dx1=""></Connection>
			<Connection Source="enf.DwiChange_ctrl_rx" Destination="ctrl.DwiChange" dx1=""></Connection>
			<Connection Source="enf.DwoChange_ctrl_rx" Destination="ctrl.DwoChange" dx1=""></Connection>
			<Connection Source="enf.DwnChange_ctrl_rx" Destination="ctrl.DwnChange" dx1=""></Connection>
			<Connection Source="enf.DwsChange_ctrl_rx" Destination="ctrl.DwsChange" dx1=""></Connection>
			<Connection Source="enf.DnChange_ctrl_rx" Destination="ctrl.DnChange" dx1=""></Connection>
			<Connection Source="enf.DsChange_ctrl_rx" Destination="ctrl.DsChange" dx1=""></Connection>
			<Connection Source="enf.SoChange_ctrl_rx" Destination="ctrl.SoChange" dx1=""></Connection>
			<Connection Source="enf.SysReady_ctrl_rx" Destination="ctrl.SysReady" dx1=""></Connection>
			<Connection Source="env.SysReady" Destination="enf.SysReady_env_tx" dx1=""></Connection>
			<Connection Source="env.SoChange" Destination="enf.SoChange_env_tx" dx1=""></Connection>
			<Connection Source="ctrl.SChange" Destination="enf.SChange_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.WChange" Destination="enf.WChange_ctrl_tx" dx1=""></Connection>
			<Connection Source="env.DwiChange" Destination="enf.DwiChange_env_tx" dx1=""></Connection>
			<Connection Source="env.DwoChange" Destination="enf.DwoChange_env_tx" dx1=""></Connection>
			<Connection Source="env.DwnChange" Destination="enf.DwnChange_env_tx" dx1=""></Connection>
			<Connection Source="env.DwsChange" Destination="enf.DwsChange_env_tx" dx1=""></Connection>
			<Connection Source="env.DnChange" Destination="enf.DnChange_env_tx" dx1=""></Connection>
			<Connection Source="env.DsChange" Destination="enf.DsChange_env_tx" dx1=""></Connection>
			<Connection Source="env.RiChange" Destination="enf.RiChange_env_tx" dx1=""></Connection>
			<Connection Source="env.RnChange" Destination="enf.RnChange_env_tx" dx1=""></Connection>
			<Connection Source="env.RsChange" Destination="enf.RsChange_env_tx" dx1=""></Connection>
			<Connection Source="enf.SChange_env_rx" Destination="env.SChange" dx1=""></Connection>
			<Connection Source="enf.WChange_env_rx" Destination="env.WChange" dx1=""></Connection>
		</EventConnections>
		<DataConnections>
			<Connection Source="enf.RiReq_ctrl_rx" Destination="ctrl.RiReq" dx1=""></Connection>
			<Connection Source="enf.RnReq_ctrl_rx" Destination="ctrl.RnReq" dx1=""></Connection>
			<Connection Source="enf.RsReq_ctrl_rx" Destination="ctrl.RsReq" dx1=""></Connection>
			<Connection Source="enf.SoGrn_ctrl_rx" Destination="ctrl.SoGrn" dx1=""></Connection>
			<Connection Source="enf.DwiPrs_ctrl_rx" Destination="ctrl.DwiPrs" dx1=""></Connection>
			<Connection Source="enf.DwoPrs_ctrl_rx" Destination="ctrl.DwoPrs" dx1=""></Connection>
			<Connection Source="enf.DwnPrs_ctrl_rx" Destination="ctrl.DwnPrs" dx1=""></Connection>
			<Connection Source="enf.DwsPrs_ctrl_rx" Destination="ctrl.DwsPrs" dx1=""></Connection>
			<Connection Source="enf.DnPrs_ctrl_rx" Destination="ctrl.DnPrs" dx1=""></Connection>
			<Connection Source="enf.DsPrs_ctrl_rx" Destination="ctrl.DsPrs" dx1=""></Connection>
			<Connection Source="env.SoGrn" Destination="enf.SoGrn_env_tx" dx1=""></Connection>
			<Connection Source="ctrl.SiGrn" Destination="enf.SiGrn_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.SnGrn" Destination="enf.SnGrn_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.SsGrn" Destination="enf.SsGrn_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.WiDvrg" Destination="enf.WiDvrg_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.WoDvrg" Destination="enf.WoDvrg_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.WnDvrg" Destination="enf.WnDvrg_ctrl_tx" dx1=""></Connection>
			<Connection Source="ctrl.WsDvrg" Destination="enf.WsDvrg_ctrl_tx" dx1=""></Connection>
			<Connection Source="env.DwiPrs" Destination="enf.DwiPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.DwoPrs" Destination="enf.DwoPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.DwnPrs" Destination="enf.DwnPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.DwsPrs" Destination="enf.DwsPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.DnPrs" Destination="enf.DnPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.DsPrs" Destination="enf.DsPrs_env_tx" dx1=""></Connection>
			<Connection Source="env.RiReq" Destination="enf.RiReq_env_tx" dx1=""></Connection>
			<Connection Source="env.RnReq" Destination="enf.RnReq_env_tx" dx1=""></Connection>
			<Connection Source="env.RsReq" Destination="enf.RsReq_env_tx" dx1=""></Connection>
			<Connection Source="enf.SiGrn_env_rx" Destination="env.SiGrn" dx1=""></Connection>
			<Connection Source="enf.SnGrn_env_rx" Destination="env.SnGrn" dx1=""></Connection>
			<Connection Source="enf.SsGrn_env_rx" Destination="env.SsGrn" dx1=""></Connection>
			<Connection Source="enf.WiDvrg_env_rx" Destination="env.WiDvrg" dx1=""></Connection>
			<Connection Source="enf.WoDvrg_env_rx" Destination="env.WoDvrg" dx1=""></Connection>
			<Connection Source="enf.WnDvrg_env_rx" Destination="env.WnDvrg" dx1=""></Connection>
			<Connection Source="enf.WsDvrg_env_rx" Destination="env.WsDvrg" dx1=""></Connection>
		</DataConnections>
	</FBNetwork>
</FBType>`,
	`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="TrainCtrl" Comment="">
	<Identification Standard="61499-2"></Identification>
	<VersionInfo Organization="" Version="" Author="" Date=""></VersionInfo>
	<CompilerInfo header="" classdef=""></CompilerInfo>
	<InterfaceList>
		<EventInputs>
			<Event Name="SysReady" Comment=""></Event>
			<Event Name="RiChange" Comment="">
				<With Var="RiReq"></With>
			</Event>
			<Event Name="RnChange" Comment="">
				<With Var="RnReq"></With>
			</Event>
			<Event Name="RsChange" Comment="">
				<With Var="RsReq"></With>
			</Event>
			<Event Name="DwiChange" Comment="">
				<With Var="DwiPrs"></With>
			</Event>
			<Event Name="DwoChange" Comment="">
				<With Var="DwoPrs"></With>
			</Event>
			<Event Name="DwnChange" Comment="">
				<With Var="DwnPrs"></With>
			</Event>
			<Event Name="DwsChange" Comment="">
				<With Var="DwsPrs"></With>
			</Event>
			<Event Name="DnChange" Comment="">
				<With Var="DnPrs"></With>
			</Event>
			<Event Name="DsChange" Comment="">
				<With Var="DsPrs"></With>
			</Event>
			<Event Name="abort" Comment=""></Event>
		</EventInputs>
		<EventOutputs>
			<Event Name="SChange" Comment="">
				<With Var="SiGrn"></With>
				<With Var="SnGrn"></With>
				<With Var="SsGrn"></With>
			</Event>
			<Event Name="WChange" Comment="">
				<With Var="WiDvrg"></With>
				<With Var="WoDvrg"></With>
				<With Var="WnDvrg"></With>
				<With Var="WsDvrg"></With>
			</Event>
		</EventOutputs>
		<InputVars>
			<VarDeclaration Name="RiReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RnReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="RsReq" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwiPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwoPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DwsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DnPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="DsPrs" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InputVars>
		<OutputVars>
			<VarDeclaration Name="SiGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SnGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="SsGrn" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WiDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WoDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WnDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="WsDvrg" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</OutputVars>
	</InterfaceList>
	<BasicFB>
		<InternalVars>
			<VarDeclaration Name="busyN" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
			<VarDeclaration Name="busyS" Type="BOOL" InitialValue="false" Comment=""></VarDeclaration>
		</InternalVars>
		<ECC>
			<ECState Name="init" Comment="" x="" y=""></ECState>
			<ECState Name="idle" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="ClrSignals"></ECAction>
			</ECState>
			<ECState Name="n_allow_train_exit_0" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNExit"></ECAction>
			</ECState>
			<ECState Name="n_allow_train_exit_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNExitHalf"></ECAction>
			</ECState>
			<ECState Name="s_allow_train_exit_0" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSExit"></ECAction>
			</ECState>
			<ECState Name="s_allow_train_exit_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSExitHalf"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_0" Comment="" x="" y=""></ECState>
			<ECState Name="i_allow_train_entrance_s" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetSEntrance"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_s_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetEntranceHalf"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_n" Comment="" x="" y="">
				<ECAction Output="WChange"></ECAction>
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetNEntrance"></ECAction>
			</ECState>
			<ECState Name="i_allow_train_entrance_n_passed_first_signal" Comment="" x="" y="">
				<ECAction Output="SChange"></ECAction>
				<ECAction Algorithm="SetEntranceHalf"></ECAction>
			</ECState>
			<ECTransition Source="init" Destination="idle" Condition="SysReady" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="n_allow_train_exit_0" Condition="RnChange &amp;&amp; RnReq == true" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="s_allow_train_exit_0" Condition="RsChange &amp;&amp; RsReq == true" x="" y=""></ECTransition>
			<ECTransition Source="idle" Destination="i_allow_train_entrance_0" Condition="RiChange &amp;&amp; RiReq == true" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_0" Destination="n_allow_train_exit_passed_first_signal" Condition="DwnChange &amp;&amp; DwnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_0" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_passed_first_signal" Destination="idle" Condition="DwoChange &amp;&amp; DwoPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="n_allow_train_exit_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_0" Destination="s_allow_train_exit_passed_first_signal" Condition="DwsChange &amp;&amp; DwsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_0" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_passed_first_signal" Destination="idle" Condition="DwoChange &amp;&amp; DwoPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="s_allow_train_exit_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="i_allow_train_entrance_s" Condition="busyS == false &amp;&amp; DsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="i_allow_train_entrance_n" Condition="busyN == false &amp;&amp; DnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_0" Destination="idle" Condition="true" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s" Destination="i_allow_train_entrance_s_passed_first_signal" Condition="DwiChange &amp;&amp; DwiPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_s_passed_first_signal" Destination="idle" Condition="DwsChange &amp;&amp; DwsPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n" Destination="i_allow_train_entrance_n_passed_first_signal" Condition="DwiChange &amp;&amp; DwiPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n" Destination="idle" Condition="abort" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n_passed_first_signal" Destination="idle" Condition="DwnChange &amp;&amp; DwnPrs == false" x="" y=""></ECTransition>
			<ECTransition Source="i_allow_train_entrance_n_passed_first_signal" Destination="idle" Condition="abort" x="" y=""></ECTransition>
		</ECC>
		<Algorithm Name="ClrSignals" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: ClrSignals\n&#34;);&#xA;        me-&gt;SiGrn = false;&#xA;        me-&gt;SnGrn = false;&#xA;        me-&gt;SsGrn = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNEntrance" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNEntrance\r\n&#34;);&#xA;        me-&gt;SiGrn = true;&#xA;        me-&gt;WiDvrg = false;&#xA;        me-&gt;WnDvrg = false;&#xA;        me-&gt;busyN = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSEntrance" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSEntrance\r\n&#34;);&#xA;        me-&gt;SiGrn = true;&#xA;        me-&gt;WiDvrg = true;&#xA;        me-&gt;WsDvrg = false; //this is an error!&#xA;        me-&gt;busyS = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetEntranceHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetEntranceHalf\n&#34;);&#xA;        me-&gt;SiGrn = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNExit" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNExit\n&#34;);&#xA;        me-&gt;SnGrn = true;&#xA;        me-&gt;WnDvrg = true;&#xA;        me-&gt;WoDvrg = true;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetNExitHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetNExitHalf\n&#34;);&#xA;        me-&gt;SnGrn = false;&#xA;        me-&gt;busyN = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSExit" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSExit\n&#34;);&#xA;        me-&gt;SsGrn = true;&#xA;        me-&gt;WsDvrg = false;&#xA;        me-&gt;WoDvrg = false;&#xA;    "></Other>
		</Algorithm>
		<Algorithm Name="SetSExitHalf" Comment="">
			<Other Language="C" Text="&#xA;        printf(&#34;TrainCtrl: SetSExitHalf\n&#34;);&#xA;        me-&gt;SsGrn = false;&#xA;        me-&gt;busyS = false;&#xA;    "></Other>
		</Algorithm>
	</BasicFB>
</FBType>`}
