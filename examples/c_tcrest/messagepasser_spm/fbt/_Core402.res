<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="_Core402" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2017-00-01" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="gen" Type="Gen" x="175" y="262.5" />
  <FB Name="p50_2" Type="Pass50" x="2012.5" y="262.5" />
  <FB Name="print" Type="Print" x="2012.5" y="1837.5" />
  <FB Name="p50_1" Type="Pass50" x="962.5" y="262.5" />
  <FB Name="p50_3" Type="Pass50" x="3062.5" y="262.5" />
  <FB Name="p50_5" Type="Pass50" x="1312.5" y="1006.25" />
  <FB Name="p50_4" Type="Pass50" x="262.5" y="1006.25" />
  <FB Name="p50_6" Type="Pass50" x="2362.5" y="1006.25" />
  <FB Name="p50_7" Type="Pass50" x="3412.5" y="1006.25" />
  <FB Name="p50_8" Type="Pass50" x="656.25" y="1837.5" />
  <EventConnections><Connection Source="gen.CountChanged" Destination="p50_1.CountChanged" />
<Connection Source="p50_2.OutCountChanged" Destination="p50_3.CountChanged" />
<Connection Source="p50_1.OutCountChanged" Destination="p50_2.CountChanged" />
<Connection Source="p50_3.OutCountChanged" Destination="p50_4.CountChanged" />
<Connection Source="p50_5.OutCountChanged" Destination="p50_6.CountChanged" />
<Connection Source="p50_4.OutCountChanged" Destination="p50_5.CountChanged" />
<Connection Source="p50_6.OutCountChanged" Destination="p50_7.CountChanged" />
<Connection Source="p50_7.OutCountChanged" Destination="p50_8.CountChanged" />
<Connection Source="p50_8.OutCountChanged" Destination="print.CountChanged" /></EventConnections>
  <DataConnections><Connection Source="gen.Count" Destination="p50_1.Count" />
<Connection Source="p50_2.OutCount" Destination="p50_3.Count" />
<Connection Source="p50_1.OutCount" Destination="p50_2.Count" />
<Connection Source="p50_3.OutCount" Destination="p50_4.Count" />
<Connection Source="p50_5.OutCount" Destination="p50_6.Count" />
<Connection Source="p50_4.OutCount" Destination="p50_5.Count" />
<Connection Source="p50_6.OutCount" Destination="p50_7.Count" />
<Connection Source="p50_7.OutCount" Destination="p50_8.Count" />
<Connection Source="p50_8.OutCount" Destination="print.Count" /></DataConnections>
</FBNetwork>
</ResourceType>