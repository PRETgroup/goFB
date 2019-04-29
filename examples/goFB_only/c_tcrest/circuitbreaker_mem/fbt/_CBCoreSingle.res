<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_CBCoreSingle" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2019-00-29" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="amm1" Type="SifbAmmeter" x="1268.75" y="1706.25" />
  <FB Name="timer1" Type="SifbTimer" x="831.25" y="306.25" />
  <FB Name="cb1" Type="CfbBreakerController" x="2318.75" y="787.5" />
  <FB Name="hm1" Type="SifbManagementControls" x="568.75" y="918.75" />
  <FB Name="led1" Type="SifbIntLed" x="3237.5" y="218.75" />
  <FB Name="hm3" Type="SifbManagementControls" x="3543.75" y="1268.75" />
  <FB Name="led3" Type="SifbIntLed" x="5993.75" y="656.25" />
  <FB Name="amm3" Type="SifbAmmeter" x="3850" y="2362.5" />
  <FB Name="cb3" Type="CfbBreakerController" x="4812.5" y="568.75" />
  <FB Name="timer3" Type="SifbTimer" x="3937.5" y="700" />
  <FB Name="hm2" Type="SifbManagementControls" x="525" y="3325" />
  <FB Name="led2" Type="SifbIntLed" x="2756.25" y="3193.75" />
  <FB Name="amm2" Type="SifbAmmeter" x="525" y="4287.5" />
  <FB Name="cb2" Type="CfbBreakerController" x="1531.25" y="3281.25" />
  <FB Name="timer2" Type="SifbTimer" x="393.75" y="2406.25" />
  <FB Name="print" Type="SifbCBPrintStatus" x="4940.46731788081" y="3465.53543622391" />
  <EventConnections><Connection Source="amm1.i_measured" Destination="cb1.i_measured" />
<Connection Source="timer1.Tick" Destination="cb1.tick" />
<Connection Source="cb1.b_change" Destination="led1.i_change" />
<Connection Source="cb1.b_change" Destination="print.StatusUpdate" />
<Connection Source="hm1.i_set_change" Destination="cb1.i_set_change" />
<Connection Source="hm1.brk" Destination="cb1.brk" />
<Connection Source="hm1.rst" Destination="cb1.rst" />
<Connection Source="hm3.i_set_change" Destination="cb3.i_set_change" />
<Connection Source="hm3.brk" Destination="cb3.brk" />
<Connection Source="hm3.rst" Destination="cb3.rst" />
<Connection Source="amm3.i_measured" Destination="cb3.i_measured" />
<Connection Source="cb3.b_change" Destination="led3.i_change" />
<Connection Source="cb3.b_change" Destination="print.StatusUpdate" />
<Connection Source="timer3.Tick" Destination="cb3.tick" />
<Connection Source="hm2.i_set_change" Destination="cb2.i_set_change" />
<Connection Source="hm2.brk" Destination="cb2.brk" />
<Connection Source="hm2.rst" Destination="cb2.rst" />
<Connection Source="amm2.i_measured" Destination="cb2.i_measured" />
<Connection Source="cb2.b_change" Destination="led2.i_change" />
<Connection Source="cb2.b_change" Destination="print.StatusUpdate" />
<Connection Source="timer2.Tick" Destination="cb2.tick" /></EventConnections>
  <DataConnections><Connection Source="amm1.i" Destination="cb1.i" />
<Connection Source="cb1.b" Destination="led1.i" />
<Connection Source="cb1.b" Destination="print.St1" />
<Connection Source="hm1.i_set" Destination="cb1.i_set" />
<Connection Source="hm3.i_set" Destination="cb3.i_set" />
<Connection Source="amm3.i" Destination="cb3.i" />
<Connection Source="cb3.b" Destination="led3.i" />
<Connection Source="cb3.b" Destination="print.St3" />
<Connection Source="hm2.i_set" Destination="cb2.i_set" />
<Connection Source="amm2.i" Destination="cb2.i" />
<Connection Source="cb2.b" Destination="led2.i" />
<Connection Source="cb2.b" Destination="print.St2" /></DataConnections>
</FBNetwork>
</ResourceType>