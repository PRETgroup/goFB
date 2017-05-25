<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="p" Type="plant" x="1093.75" y="1356.25" />
  <FB Name="c" Type="controller" x="2056.25" y="1356.25" />
  <EventConnections><Connection Source="p.update" Destination="c.update" />
<Connection Source="c.S2" Destination="p.S2" />
<Connection Source="c.S3" Destination="p.S3" /></EventConnections>
  <DataConnections><Connection Source="p.x" Destination="c.x" />
<Connection Source="p.y" Destination="c.y" /></DataConnections>
</FBNetwork>
</ResourceType>