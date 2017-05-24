<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="p" Type="plant" x="481.25" y="1159.375" />
  <FB Name="c" Type="controller" x="2318.75" y="1268.75" />
  <EventConnections><Connection Source="p.update" Destination="c.update" />
<Connection Source="c.TURNON" Destination="p.TURNON" />
<Connection Source="c.TURNOFF" Destination="p.TURNOFF" /></EventConnections>
  <DataConnections><Connection Source="p.y" Destination="c.y" />
<Connection Source="p.x" Destination="c.x" /></DataConnections>
</FBNetwork>
</ResourceType>