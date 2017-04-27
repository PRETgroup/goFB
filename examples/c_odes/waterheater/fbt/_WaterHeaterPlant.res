<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_WaterHeaterPlant" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="waterHeaterPlantODE" Type="WaterHeaterPlantODE" x="1793.75" y="1071.875">
    <Parameter Name="K" Value="0.075" />
    <Parameter Name="H" Value="150" />
  </FB>
  <FB Name="t_tx" Type="interResourceTxLReal" x="3193.75" y="1093.75">
    <Parameter Name="Channel" Value="1" />
  </FB>
  <FB Name="heat_rx" Type="interResourceRxBool" x="612.5" y="1050">
    <Parameter Name="Channel" Value="0" />
  </FB>
  <FB Name="tickgen" Type="TickGen" x="743.75" y="218.75">
    <Parameter Name="SetDeltaTime" Value="2.5" />
  </FB>
  <EventConnections><Connection Source="waterHeaterPlantODE.Ychange" Destination="t_tx.Tx" />
<Connection Source="heat_rx.Rx" Destination="waterHeaterPlantODE.HeatChange" />
<Connection Source="tickgen.Tick" Destination="waterHeaterPlantODE.Tick" /></EventConnections>
  <DataConnections><Connection Source="waterHeaterPlantODE.Y" Destination="t_tx.Data" />
<Connection Source="heat_rx.Data" Destination="waterHeaterPlantODE.Heat" />
<Connection Source="tickgen.DeltaTime" Destination="waterHeaterPlantODE.DeltaTime" /></DataConnections>
</FBNetwork>
</ResourceType>