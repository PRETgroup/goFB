<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="topMANY" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-14" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="basic" Type="container_two_basic" x="918.75" y="1137.5">
    <Parameter Name="printf_id1" Value="1" />
    <Parameter Name="printf_id2" Value="2" />
  </FB>
  <FB Name="mixed" Type="container_two_mixed" x="2318.75" y="1137.5">
    <Parameter Name="printf_id1" Value="3" />
    <Parameter Name="printf_id2" Value="4" />
  </FB>
  <EventConnections><Connection Source="basic.DataOutChanged" Destination="mixed.DataInChanged" />
<Connection Source="mixed.DataOutChanged" Destination="basic.DataInChanged" /></EventConnections>
  <DataConnections><Connection Source="basic.DataOut" Destination="mixed.DataIn" />
<Connection Source="mixed.DataOut" Destination="basic.DataIn" /></DataConnections>
</FBNetwork>
</ResourceType>