<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_pizza" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-12" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="conv" Type="conveyorHFB" x="1137.5" y="568.75">
    <Parameter Name="DeltaTime" Value="0.5" />
    <Parameter Name="M" Value="10" />
  </FB>
  <FB Name="cont" Type="controller" x="2275" y="525" />
  <EventConnections><Connection Source="conv.XChange" Destination="cont.XChange" />
<Connection Source="conv.DChange" Destination="cont.DChange" />
<Connection Source="cont.Start" Destination="conv.Start" />
<Connection Source="cont.Start" Destination="conv.MChange" />
<Connection Source="cont.Off" Destination="conv.Off" />
<Connection Source="cont.On" Destination="conv.On" /></EventConnections>
  <DataConnections><Connection Source="conv.X" Destination="cont.X" />
<Connection Source="conv.D" Destination="cont.D" /></DataConnections>
</FBNetwork>
</ResourceType>