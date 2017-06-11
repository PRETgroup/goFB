<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="o" Type="oven" x="1225" y="2012.5" />
  <FB Name="c" Type="conveyor" x="656.25" y="1050" />
  <FB Name="contr" Type="controller" x="2256.77083333333" y="586.979166666667" />
  <EventConnections><Connection Source="o.xChange" Destination="contr.update" />
<Connection Source="c.xdchange" Destination="contr.update" />
<Connection Source="contr.Conveyor_On" Destination="c.on" />
<Connection Source="contr.Conveyor_Off" Destination="c.off" />
<Connection Source="contr.Oven_Start" Destination="o.start" />
<Connection Source="contr.Oven_Remove" Destination="o.remove" />
<Connection Source="contr.Oven_Done" Destination="o.done" /></EventConnections>
  <DataConnections><Connection Source="o.x" Destination="contr.Oven_Vo" />
<Connection Source="c.d" Destination="contr.Conveyor_D" /></DataConnections>
</FBNetwork>
</ResourceType>