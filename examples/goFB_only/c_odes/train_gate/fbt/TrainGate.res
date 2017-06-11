<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="TrainGate" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-23" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="p" Type="Plant" x="1192.1875" y="1261.45833333333" />
  <FB Name="c" Type="Controller" x="2366.14583333333" y="1469.27083333333" />
  <EventConnections><Connection Source="p.update_t" Destination="c.update" />
<Connection Source="c.UP" Destination="p.UP" />
<Connection Source="c.DOWN" Destination="p.DOWN" /></EventConnections>
  <DataConnections><Connection Source="p.y" Destination="c.y1" />
<Connection Source="p.x" Destination="c.x" /></DataConnections>
</FBNetwork>
</ResourceType>