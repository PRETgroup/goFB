<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="topFLAT" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="pf1" Type="passforward" x="1356.25" y="1181.25" />
  <FB Name="pf2" Type="passforward" x="3208.33333333333" y="1199.47916666667" />
  <EventConnections><Connection Source="pf1.DataOutChanged" Destination="pf2.DataInChanged" />
<Connection Source="pf2.DataOutChanged" Destination="pf1.DataInChanged" /></EventConnections>
  <DataConnections><Connection Source="pf1.DataOut" Destination="pf2.DataIn" />
<Connection Source="pf2.DataOut" Destination="pf1.DataIn" /></DataConnections>
</FBNetwork>
</ResourceType>