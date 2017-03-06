<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="FakePacemakerTop" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-06" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="pmaker" Type="VVI_Pacemaker" x="1881.25" y="1181.25" />
  <EventConnections><Connection Source="pmaker.VPace" Destination="pmaker.VPulse" />
<Connection Source="pmaker.VRP_Start_Timer" Destination="pmaker.VRP_Timer_Timeout" />
<Connection Source="pmaker.LRI_Timer_Start" Destination="pmaker.LRI_Timer_Timeout" /></EventConnections>
</FBNetwork>
</ResourceType>