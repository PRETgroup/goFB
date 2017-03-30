<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-29" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="g" Type="TickGen" x="1487.5" y="1181.25" />
  <FB Name="ha" Type="TrainHA" x="2100" y="1181.25">
    <Parameter Name="Vs" Value="1.0" />
    <Parameter Name="Vf" Value="2.0" />
  </FB>
  <FB Name="p" Type="Output" x="2843.75" y="1181.25" />
  <EventConnections><Connection Source="g.tick" Destination="ha.tick" />
<Connection Source="ha.update" Destination="p.output" /></EventConnections>
  <DataConnections><Connection Source="g.delta" Destination="ha.delta" />
<Connection Source="ha.x" Destination="p.var" /></DataConnections>
</FBNetwork>
</ResourceType>