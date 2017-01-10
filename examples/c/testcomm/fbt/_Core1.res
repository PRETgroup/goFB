<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core1" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-09" />
<CompilerInfo header="" classdef="">
</CompilerInfo>
<VarDeclaration Name="TxChanId" Type="UDINT" Comment="" />
<FBNetwork>
  <FB Name="tx" Type="ArgoTx" x="2012.5" y="962.5" />
  <FB Name="prod" Type="Producer" x="918.75" y="962.5" />
  <EventConnections><Connection Source="tx.SuccessChanged" Destination="prod.TxSuccessChanged" />
<Connection Source="prod.DataPresent" Destination="tx.DataPresent" /></EventConnections>
  <DataConnections><Connection Source="TxChanId" Destination="tx.ChanId" />
<Connection Source="tx.Success" Destination="prod.TxSuccess" />
<Connection Source="prod.Data" Destination="tx.Data" /></DataConnections>
</FBNetwork>
</ResourceType>