<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_CoreSingle" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-22" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="saw1" Type="SawmillModule" x="1050" y="481.25" />
  <FB Name="saw2" Type="SawmillModule" x="1050" y="1050" />
  <FB Name="saw3" Type="SawmillModule" x="1050" y="1618.75" />
  <FB Name="statusprint" Type="PrintStatus" x="2493.75" y="1093.75" />
  <EventConnections><Connection Source="saw1.MessageChange" Destination="statusprint.StatusUpdate" />
<Connection Source="saw2.MessageChange" Destination="statusprint.StatusUpdate" />
<Connection Source="saw3.MessageChange" Destination="statusprint.StatusUpdate" /></EventConnections>
  <DataConnections><Connection Source="saw1.Message" Destination="statusprint.Saw1Status" />
<Connection Source="saw2.Message" Destination="statusprint.Saw2Status" />
<Connection Source="saw3.Message" Destination="statusprint.Saw3Status" /></DataConnections>
</FBNetwork>
</ResourceType>