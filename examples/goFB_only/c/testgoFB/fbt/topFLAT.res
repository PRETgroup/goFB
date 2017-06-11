<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="topFLAT" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="pf1" Type="passforward" x="262.5" y="568.75">
    <Parameter Name="printf_id" Value="1" />
  </FB>
  <FB Name="pf2" Type="passforward" x="1312.5" y="568.75">
    <Parameter Name="printf_id" Value="2" />
  </FB>
  <FB Name="pf3" Type="passforward" x="2362.5" y="568.75">
    <Parameter Name="printf_id" Value="3" />
  </FB>
  <FB Name="pf4" Type="passforward" x="3412.5" y="568.75">
    <Parameter Name="printf_id" Value="4" />
  </FB>
  <EventConnections><Connection Source="pf1.DataOutChanged" Destination="pf2.DataInChanged" />
<Connection Source="pf2.DataOutChanged" Destination="pf3.DataInChanged" />
<Connection Source="pf3.DataOutChanged" Destination="pf4.DataInChanged" />
<Connection Source="pf4.DataOutChanged" Destination="pf1.DataInChanged" /></EventConnections>
  <DataConnections><Connection Source="pf1.DataOut" Destination="pf2.DataIn" />
<Connection Source="pf2.DataOut" Destination="pf3.DataIn" />
<Connection Source="pf3.DataOut" Destination="pf4.DataIn" />
<Connection Source="pf4.DataOut" Destination="pf1.DataIn" /></DataConnections>
</FBNetwork>
</ResourceType>