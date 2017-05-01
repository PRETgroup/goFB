<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_WaterHeaterPlant" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="waterHeaterPlantODE1" Type="WaterHeaterPlantODE" x="1837.5" y="1050">
    <Parameter Name="K" Value="0.075" />
    <Parameter Name="H" Value="150" />
  </FB>
  <FB Name="t_tx1" Type="interResourceTxLReal" x="2800" y="1093.75">
    <Parameter Name="Channel" Value="2" />
  </FB>
  <FB Name="heat_rx1" Type="interResourceRxBool" x="918.75" y="1093.75">
    <Parameter Name="Channel" Value="0" />
  </FB>
  <FB Name="StartGen" Type="StartGen" x="743.75" y="481.25">
    <Parameter Name="SetDeltaTime" Value="0.5" />
  </FB>
  <FB Name="waterHeaterPlandODE2" Type="WaterHeaterPlantODE" x="1837.5" y="2012.5">
    <Parameter Name="K" Value="0.075" />
    <Parameter Name="H" Value="150" />
  </FB>
  <FB Name="t_tx2" Type="interResourceTxLReal" x="2800" y="2143.75">
    <Parameter Name="Channel" Value="3" />
  </FB>
  <FB Name="heat_rx2" Type="interResourceRxBool" x="875" y="2231.25">
    <Parameter Name="Channel" Value="1" />
  </FB>
  <EventConnections><Connection Source="waterHeaterPlantODE1.Ychange" Destination="t_tx1.Tx" />
<Connection Source="heat_rx1.Rx" Destination="waterHeaterPlantODE1.HeatChange" />
<Connection Source="StartGen.Tick" Destination="waterHeaterPlantODE1.Tick" />
<Connection Source="StartGen.Tick" Destination="waterHeaterPlandODE2.Tick" />
<Connection Source="waterHeaterPlandODE2.Ychange" Destination="t_tx2.Tx" />
<Connection Source="heat_rx2.Rx" Destination="waterHeaterPlandODE2.HeatChange" /></EventConnections>
  <DataConnections><Connection Source="waterHeaterPlantODE1.Y" Destination="t_tx1.Data" />
<Connection Source="heat_rx1.Data" Destination="waterHeaterPlantODE1.Heat" />
<Connection Source="StartGen.DeltaTime" Destination="waterHeaterPlantODE1.DeltaTime" />
<Connection Source="StartGen.DeltaTime" Destination="waterHeaterPlandODE2.DeltaTime" />
<Connection Source="waterHeaterPlandODE2.Y" Destination="t_tx2.Data" />
<Connection Source="heat_rx2.Data" Destination="waterHeaterPlandODE2.Heat" /></DataConnections>
</FBNetwork>
</ResourceType>