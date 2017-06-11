<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core022" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-06" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="g" Type="Gen" x="743.75" y="787.5" />
  <FB Name="p1" Type="Pass10" x="1268.75" y="787.5" />
  <FB Name="p2" Type="Pass10" x="2275" y="787.5" />
  <FB Name="pr" Type="Print" x="3281.25" y="787.5" />
  <EventConnections><Connection Source="g.CountChanged" Destination="p1.CountChanged" />
<Connection Source="p1.OutCountChanged" Destination="p2.CountChanged" />
<Connection Source="p2.OutCountChanged" Destination="pr.CountChanged" /></EventConnections>
  <DataConnections><Connection Source="g.Count" Destination="p1.Count" />
<Connection Source="p1.OutCount" Destination="p2.Count" />
<Connection Source="p2.OutCount" Destination="pr.Count" /></DataConnections>
</FBNetwork>
</ResourceType>