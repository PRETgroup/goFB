<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_CBCore3" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2019-00-29" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="timer" Type="SifbTimer" x="339.0625" y="586.979166666667" />
  <FB Name="tx" Type="ArgoTx" x="3937.5" y="831.25">
    <Parameter Name="ChanId" Value="1" />
  </FB>
  <FB Name="hmi" Type="SifbManagementControls" x="262.5" y="1268.75" />
  <FB Name="led" Type="SifbIntLed" x="3018.75" y="1881.25" />
  <FB Name="cb" Type="CfbBreakerController" x="1400" y="743.75" />
  <FB Name="amm" Type="SifbAmmeter" x="350" y="2012.5" />
  <FB Name="msgh" Type="SawmillMessageHandler" x="2800" y="831.25" />
  <EventConnections><Connection Source="timer.Tick" Destination="cb.tick" />
<Connection Source="tx.SuccessChanged" Destination="msgh.TxSuccessChanged" />
<Connection Source="hmi.i_set_change" Destination="cb.i_set_change" />
<Connection Source="hmi.brk" Destination="cb.brk" />
<Connection Source="hmi.rst" Destination="cb.rst" />
<Connection Source="cb.b_change" Destination="msgh.MessageChanged" />
<Connection Source="cb.b_change" Destination="led.i_change" />
<Connection Source="amm.i_measured" Destination="cb.i_measured" />
<Connection Source="msgh.TxDataPresent" Destination="tx.DataPresent" /></EventConnections>
  <DataConnections><Connection Source="tx.Success" Destination="msgh.TxSuccess" />
<Connection Source="hmi.i_set" Destination="cb.i_set" />
<Connection Source="cb.b" Destination="msgh.Message" />
<Connection Source="cb.b" Destination="led.i" />
<Connection Source="amm.i" Destination="cb.i" />
<Connection Source="msgh.TxData" Destination="tx.Data" /></DataConnections>
</FBNetwork>
</ResourceType>