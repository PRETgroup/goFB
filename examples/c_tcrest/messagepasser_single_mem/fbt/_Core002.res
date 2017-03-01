<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core002" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-01" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="print" Type="Print" x="3368.75" y="1181.25" />
  <FB Name="gen" Type="Gen" x="1618.75" y="1181.25" />
  <FB Name="Reference" Type="Pass" x="2275" y="525" />
  <EventConnections><Connection Source="gen.CountChanged" Destination="Reference.CountChanged" />
<Connection Source="Reference.OutCountChanged" Destination="print.CountChanged" /></EventConnections>
  <DataConnections><Connection Source="gen.Count" Destination="Reference.Count" />
<Connection Source="Reference.OutCount" Destination="print.Count" /></DataConnections>
</FBNetwork>
</ResourceType>