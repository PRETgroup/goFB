<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-02" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="pid" Type="PID" x="2143.75" y="437.5">
    <Parameter Name="PGain" Value="0.4" />
    <Parameter Name="IGain" Value="0.05" />
    <Parameter Name="DGain" Value="0.5" />
    <Parameter Name="WindupGuard" Value="1.0" />
  </FB>
  <FB Name="manager" Type="Manager" x="787.5" y="525">
    <Parameter Name="TicksPerVal" Value="100" />
  </FB>
  <FB Name="plant" Type="Plant" x="2187.5" y="1662.5" />
  <EventConnections><Connection Source="pid.ControlChanged" Destination="plant.ControlChange" />
<Connection Source="manager.Zero" Destination="pid.Zero" />
<Connection Source="manager.Zero" Destination="plant.Zero" />
<Connection Source="manager.Tick" Destination="pid.Tick" />
<Connection Source="plant.ValueChange" Destination="manager.ActualValueChanged" />
<Connection Source="plant.ValueChange" Destination="pid.ActualValueChanged" /></EventConnections>
  <DataConnections><Connection Source="pid.Control" Destination="plant.Control" />
<Connection Source="manager.DesiredValue" Destination="pid.DesiredValue" />
<Connection Source="plant.Value" Destination="pid.ActualValue" />
<Connection Source="plant.Value" Destination="manager.ActualValue" /></DataConnections>
</FBNetwork>
</ResourceType>