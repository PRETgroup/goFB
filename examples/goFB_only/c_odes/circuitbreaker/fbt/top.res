<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="top" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-25" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="load" Type="load" x="1225" y="2012.5" />
  <FB Name="contr" Type="controller" x="2256.77083333333" y="586.979166666667" />
  <EventConnections>
    <Connection Source="contr.off" Destination="load.off" />
    <Connection Source="contr.on" Destination="load.on" />
    <Connection Source="contr.fault" Destination="load.fault" />
  </EventConnections>
  <DataConnections></DataConnections>
</FBNetwork>
</ResourceType>