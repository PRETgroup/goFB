<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="p" Type="plant" x="1925" y="1312.5" />
  <FB Name="c" Type="controller" x="3062.5" y="1312.5" />
  <EventConnections><Connection Source="p.update" Destination="c.update" />
<Connection Source="c.add1" Destination="p.add1" />
<Connection Source="c.add2" Destination="p.add2" />
<Connection Source="c.remove1" Destination="p.remove1" />
<Connection Source="c.remove2" Destination="p.remove2" /></EventConnections>
  <DataConnections><Connection Source="p.x" Destination="c.x" /></DataConnections>
</FBNetwork>
</ResourceType>