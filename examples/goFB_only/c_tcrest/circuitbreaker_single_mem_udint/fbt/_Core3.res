<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core3" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-16" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="tx" Type="ArgoTx" x="2975" y="1225">
    <Parameter Name="ChanId" Value="3" />
  </FB>
  <FB Name="messageHandler" Type="SawmillMessageHandler" x="1520.3125" y="1170.3125" />
  <FB Name="sawmill" Type="SawmillModule" x="514.0625" y="1039.0625" />
  <EventConnections><Connection Source="tx.SuccessChanged" Destination="messageHandler.TxSuccessChanged" />
<Connection Source="messageHandler.TxDataPresent" Destination="tx.DataPresent" />
<Connection Source="sawmill.MessageChange" Destination="messageHandler.MessageChanged" /></EventConnections>
  <DataConnections><Connection Source="tx.Success" Destination="messageHandler.TxSuccess" />
<Connection Source="messageHandler.TxData" Destination="tx.Data" />
<Connection Source="sawmill.Message" Destination="messageHandler.Message" /></DataConnections>
</FBNetwork>
</ResourceType>