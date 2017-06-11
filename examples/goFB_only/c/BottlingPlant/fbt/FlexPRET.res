<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="FlexPRET" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2016-00-28" />
<CompilerInfo header="" classdef="">
</CompilerInfo>

<FBNetwork>
  <FB Name="IO" Type="IOManager" x="831.25" y="962.5" />
  <FB Name="CCounter" Type="CanisterCounter" x="3543.75" y="1662.5" />
  <FB Name="Door" Type="DoorController" x="3543.75" y="787.5" />
  <FB Name="Injector" Type="InjectorController" x="3237.5" y="2362.5" />
  <FB Name="Conveyor" Type="ConveyorController" x="5613.56461636749" y="2815.82283526615" />
  <FB Name="RejectArm" Type="RejectArmController" x="5687.5" y="1793.75" />
  <EventConnections><Connection Source="IO.InjectorArmFinishMovement" Destination="Injector.InjectorArmFinishedMovement" />
<Connection Source="IO.EmergencyStopChanged" Destination="Door.EmergencyStopChanged" />
<Connection Source="IO.EmergencyStopChanged" Destination="Injector.EmergencyStopChanged" />
<Connection Source="IO.EmergencyStopChanged" Destination="Conveyor.EmergencyStopChanged" />
<Connection Source="IO.CanisterPressureChanged" Destination="Injector.CanisterPressureChanged" />
<Connection Source="IO.FillContentsAvailableChanged" Destination="Injector.FillContentsAvailableChanged" />
<Connection Source="IO.LasersChanged" Destination="CCounter.LasersChanged" />
<Connection Source="IO.LasersChanged" Destination="RejectArm.LasersChanged" />
<Connection Source="IO.LasersChanged" Destination="Conveyor.LasersChanged" />
<Connection Source="IO.DoorOverride" Destination="Door.ReleaseDoorOverride" />
<Connection Source="IO.VacuumTimerElapsed" Destination="Injector.VacuumTimerElapsed" />
<Connection Source="CCounter.CanisterCountChanged" Destination="IO.CanisterCountChanged" />
<Connection Source="Door.DoorReleaseCanister" Destination="IO.DoorReleaseCanister" />
<Connection Source="Injector.InjectDone" Destination="Door.BottlingDone" />
<Connection Source="Injector.InjectDone" Destination="Conveyor.InjectDone" />
<Connection Source="Injector.InjectorPositionChanged" Destination="IO.InjectorPositionChanged" />
<Connection Source="Injector.InjectorControlsChanged" Destination="IO.InjectorControlsChanged" />
<Connection Source="Injector.RejectCanister" Destination="RejectArm.RejectCanister" />
<Connection Source="Injector.FillContentsChanged" Destination="IO.FillContentsChanged" />
<Connection Source="Injector.StartVacuumTimer" Destination="IO.StartVacuumTimer" />
<Connection Source="Conveyor.ConveyorChanged" Destination="IO.ConveyorChanged" />
<Connection Source="Conveyor.ConveyorStoppedForInject" Destination="Injector.ConveyorStoppedForInject" />
<Connection Source="RejectArm.GoRejectArm" Destination="IO.GoRejectArm" /></EventConnections>
  <DataConnections><Connection Source="IO.EmergencyStop" Destination="Door.EmergencyStop" />
<Connection Source="IO.EmergencyStop" Destination="Injector.EmergencyStop" />
<Connection Source="IO.EmergencyStop" Destination="Conveyor.EmergencyStop" />
<Connection Source="IO.CanisterPressure" Destination="Injector.CanisterPressure" />
<Connection Source="IO.FillContentsAvailable" Destination="Injector.FillContentsAvailable" />
<Connection Source="IO.DoorSiteLaser" Destination="CCounter.DoorSiteLaser" />
<Connection Source="IO.InjectSiteLaser" Destination="Conveyor.InjectSiteLaser" />
<Connection Source="IO.RejectSiteLaser" Destination="RejectArm.RejectSiteLaser" />
<Connection Source="IO.RejectBinLaser" Destination="CCounter.RejectBinLaser" />
<Connection Source="IO.AcceptBinLaser" Destination="CCounter.AcceptBinLaser" />
<Connection Source="CCounter.CanisterCount" Destination="IO.CanisterCount" />
<Connection Source="Injector.InjectorPosition" Destination="IO.InjectorPosition" />
<Connection Source="Injector.InjectorContentsValveOpen" Destination="IO.InjectorContentsValveOpen" />
<Connection Source="Injector.InjectorVacuumRun" Destination="IO.InjectorVacuumRun" />
<Connection Source="Injector.InjectorPressurePumpRun" Destination="IO.InjectorPressurePumpRun" />
<Connection Source="Injector.FillContents" Destination="IO.FillContents" />
<Connection Source="Conveyor.ConveyorSpeed" Destination="IO.ConveyorSpeed" /></DataConnections>
</FBNetwork>
</ResourceType>