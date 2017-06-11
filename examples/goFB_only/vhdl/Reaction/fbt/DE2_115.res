<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="DE2_115" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2016-00-31" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="gpio" Type="BFB_GPIO" x="2424.47916666667" y="1192.1875" />
  <FB Name="reaction" Type="BFB_Reaction" x="2318.75" y="1968.75" />
  <EventConnections><Connection Source="gpio.rx_rd" Destination="reaction.rx_change" />
<Connection Source="reaction.tx_change" Destination="gpio.tx_rd" /></EventConnections>
  <DataConnections><Connection Source="gpio.rx_data" Destination="reaction.rx" />
<Connection Source="reaction.tx" Destination="gpio.tx_data" /></DataConnections>
</FBNetwork>
</ResourceType>