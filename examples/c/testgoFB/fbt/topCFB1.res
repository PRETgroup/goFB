<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="topCFB1" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="cf1" Type="container_one" x="350" y="875">
    <Parameter Name="printf_id" Value="1" />
  </FB>
  <FB Name="cf2" Type="container_one" x="1400" y="875">
    <Parameter Name="printf_id" Value="2" />
  </FB>
  <FB Name="cf3" Type="container_one" x="2450" y="875">
    <Parameter Name="printf_id" Value="3" />
  </FB>
  <FB Name="cf4" Type="container_one" x="3500" y="875">
    <Parameter Name="printf_id" Value="4" />
  </FB>
  <EventConnections><Connection Source="cf1.DataOutChanged" Destination="cf2.DataInChanged" />
<Connection Source="cf2.DataOutChanged" Destination="cf3.DataInChanged" />
<Connection Source="cf3.DataOutChanged" Destination="cf4.DataInChanged" />
<Connection Source="cf4.DataOutChanged" Destination="cf1.DataInChanged" /></EventConnections>
  <DataConnections><Connection Source="cf1.DataOut" Destination="cf2.DataIn" />
<Connection Source="cf2.DataOut" Destination="cf3.DataIn" />
<Connection Source="cf3.DataOut" Destination="cf4.DataIn" />
<Connection Source="cf4.DataOut" Destination="cf1.DataIn" /></DataConnections>
</FBNetwork>
</ResourceType>