<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_WaterHeaterPlantAndController" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-01" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="p1" Type="WaterHeaterPlantODE" x="2501.04166666667" y="831.25">
    <Parameter Name="K" Value="0.075" />
    <Parameter Name="H" Value="150" />
  </FB>
  <FB Name="p2" Type="WaterHeaterPlantODE" x="2537.5" y="1750">
    <Parameter Name="K" Value="0.075" />
    <Parameter Name="H" Value="150" />
  </FB>
  <FB Name="c" Type="Controller" x="1006.25" y="1531.25" />
  <FB Name="gen" Type="StartGen" x="962.5" y="656.25">
    <Parameter Name="SetDeltaTime" Value="0.5" />
  </FB>
  <EventConnections><Connection Source="p1.Ychange" Destination="c.TChange1" />
<Connection Source="p2.Ychange" Destination="c.TChange2" />
<Connection Source="c.HeatChange" Destination="p1.HeatChange" />
<Connection Source="c.HeatChange" Destination="p2.HeatChange" />
<Connection Source="gen.Start" Destination="p1.Start" />
<Connection Source="gen.Start" Destination="p2.Start" /></EventConnections>
  <DataConnections><Connection Source="p1.Y" Destination="c.T1" />
<Connection Source="p2.Y" Destination="c.T2" />
<Connection Source="c.Heat1" Destination="p1.Heat" />
<Connection Source="c.Heat2" Destination="p2.Heat" />
<Connection Source="gen.DeltaTime" Destination="p1.DeltaTime" />
<Connection Source="gen.DeltaTime" Destination="p2.DeltaTime" /></DataConnections>
</FBNetwork>
</ResourceType>