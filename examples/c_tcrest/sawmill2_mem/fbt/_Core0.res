<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core0" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="saw1rx" Type="ArgoRx" x="568.75" y="1006.25">
    <Parameter Name="ChanId" Value="1" />
  </FB>
  <FB Name="saw2rx" Type="ArgoRx" x="568.75" y="1575">
    <Parameter Name="ChanId" Value="2" />
  </FB>
  <FB Name="saw3rx" Type="ArgoRx" x="568.75" y="2143.75">
    <Parameter Name="ChanId" Value="3" />
  </FB>
  <FB Name="statusprint" Type="PrintStatus" x="2231.25" y="1268.75" />
  <EventConnections><Connection Source="saw1rx.DataPresent" Destination="statusprint.StatusUpdate" />
<Connection Source="saw2rx.DataPresent" Destination="statusprint.StatusUpdate" />
<Connection Source="saw3rx.DataPresent" Destination="statusprint.StatusUpdate" /></EventConnections>
  <DataConnections><Connection Source="saw1rx.Data" Destination="statusprint.Saw1Status" />
<Connection Source="saw2rx.Data" Destination="statusprint.Saw2Status" />
<Connection Source="saw3rx.Data" Destination="statusprint.Saw3Status" /></DataConnections>
</FBNetwork>
</ResourceType>