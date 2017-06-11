<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core102" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-01" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="gen" Type="Gen" x="262.5" y="1006.25" />
  <FB Name="print" Type="Print" x="3762.5" y="1006.25" />
  <FB Name="p50_1" Type="Pass50" x="1050" y="1006.25" />
  <FB Name="p50_2" Type="Pass50" x="2318.75" y="1006.25" />
  <EventConnections><Connection Source="gen.CountChanged" Destination="p50_1.CountChanged" />
<Connection Source="p50_1.OutCountChanged" Destination="p50_2.CountChanged" />
<Connection Source="p50_2.OutCountChanged" Destination="print.CountChanged" /></EventConnections>
  <DataConnections><Connection Source="gen.Count" Destination="p50_1.Count" />
<Connection Source="p50_1.OutCount" Destination="p50_2.Count" />
<Connection Source="p50_2.OutCount" Destination="print.Count" /></DataConnections>
</FBNetwork>
</ResourceType>