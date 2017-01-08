<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core1" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>
<VarDeclaration Name="Core1Input" Type="INT" Comment="" />
<FBNetwork>
  <FB Name="tx" Type="ArgoTx" x="2012.5" y="962.5" />
  <FB Name="prod" Type="Producer" x="1006.25" y="962.5" />
  <EventConnections><Connection Source="tx.BusyChanged" Destination="prod.TxStatusChanged" />
<Connection Source="prod.DataPresent" Destination="tx.DataPresent" /></EventConnections>
  <DataConnections><Connection Source="tx.Busy" Destination="prod.TxBusy" />
<Connection Source="prod.Data" Destination="tx.Data" /></DataConnections>
</FBNetwork>
</ResourceType>