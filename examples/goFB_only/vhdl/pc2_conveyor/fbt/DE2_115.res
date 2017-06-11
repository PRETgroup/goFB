<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="DE2_115" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2016-00-08" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="conveyor" Type="Conveyor_SIFB" x="2406.25" y="612.5" />
  <FB Name="globals" Type="Globals_SIFB" x="1225" y="1618.75" />
  <FB Name="boxdropper" Type="BoxDropper_SIFB" x="3018.75" y="1925" />
  <EventConnections><Connection Source="globals.global_run_changed" Destination="conveyor.conveyor_run_changed" />
<Connection Source="globals.global_run_changed" Destination="boxdropper.box_dropper_run_changed" /></EventConnections>
  <DataConnections><Connection Source="globals.global_run" Destination="conveyor.conveyor_run" />
<Connection Source="globals.global_run_infinite" Destination="boxdropper.box_dropper_run" /></DataConnections>
</FBNetwork>
</ResourceType>