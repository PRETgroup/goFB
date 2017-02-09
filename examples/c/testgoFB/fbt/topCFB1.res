<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="topCFB1" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="cf1" Type="container" x="1494.79166666667" y="1261.45833333333" />
  <FB Name="cf2" Type="container" x="3150" y="1268.75" />
  <EventConnections><Connection Source="cf1.DataOutChanged" Destination="cf2.DataInChanged" />
<Connection Source="cf2.DataOutChanged" Destination="cf1.DataInChanged" /></EventConnections>
  <DataConnections><Connection Source="cf1.DataOut" Destination="cf2.DataIn" />
<Connection Source="cf2.DataOut" Destination="cf1.DataIn" /></DataConnections>
</FBNetwork>
</ResourceType>