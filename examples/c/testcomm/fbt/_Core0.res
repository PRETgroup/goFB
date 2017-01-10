<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core0" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>
<VarDeclaration Name="RxChanId" Type="UDINT" Comment="" />
<FBNetwork>
  <FB Name="rx" Type="ArgoRx" x="1662.5" y="1225" />
  <FB Name="print" Type="PrintInt" x="2581.25" y="1225" />
  <EventConnections><Connection Source="rx.DataPresent" Destination="print.DataPresent" /></EventConnections>
  <DataConnections><Connection Source="RxChanId" Destination="rx.ChanId" />
<Connection Source="rx.Data" Destination="print.Data" /></DataConnections>
</FBNetwork>
</ResourceType>