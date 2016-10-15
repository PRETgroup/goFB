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
  <FB Name="Conveyor" Type="ConveyorController" x="5613.56461636749" y="2815.82283526615" />
  <FB Name="RejectArm" Type="RejectArmController" x="5687.5" y="1793.75" />
  <FB Name="Pumps" Type="InjectorPumpsController" x="3412.5" y="3631.25" />
  <FB Name="Motor" Type="InjectorMotorController" x="3412.5" y="2625" />
  <EventConnections><Connection Source="IO.InjectorArmFinishMovement" Destination="Motor.InjectorArmFinishedMovement" />
<Connection Source="IO.EmergencyStopChanged" Destination="Door.EmergencyStopChanged" />
<Connection Source="IO.EmergencyStopChanged" Destination="Conveyor.EmergencyStopChanged" />
<Connection Source="IO.EmergencyStopChanged" Destination="Motor.EmergencyStopChanged" />
<Connection Source="IO.EmergencyStopChanged" Destination="Pumps.EmergencyStopChanged" />
<Connection Source="IO.CanisterPressureChanged" Destination="Pumps.CanisterPressureChanged" />
<Connection Source="IO.FillContentsAvailableChanged" Destination="Pumps.FillContentsAvailableChanged" />
<Connection Source="IO.LasersChanged" Destination="CCounter.LasersChanged" />
<Connection Source="IO.LasersChanged" Destination="RejectArm.LasersChanged" />
<Connection Source="IO.LasersChanged" Destination="Conveyor.LasersChanged" />
<Connection Source="IO.DoorOverride" Destination="Door.ReleaseDoorOverride" />
<Connection Source="IO.VacuumTimerElapsed" Destination="Pumps.VacuumTimerElapsed" />
<Connection Source="CCounter.CanisterCountChanged" Destination="IO.CanisterCountChanged" />
<Connection Source="Door.DoorReleaseCanister" Destination="IO.DoorReleaseCanister" />
<Connection Source="Conveyor.ConveyorChanged" Destination="IO.ConveyorChanged" />
<Connection Source="Conveyor.ConveyorStoppedForInject" Destination="Motor.ConveyorStoppedForInject" />
<Connection Source="RejectArm.GoRejectArm" Destination="IO.GoRejectArm" />
<Connection Source="Pumps.PumpFinished" Destination="Motor.PumpFinished" />
<Connection Source="Pumps.RejectCanister" Destination="RejectArm.RejectCanister" />
<Connection Source="Pumps.InjectorControlsChanged" Destination="IO.InjectorControlsChanged" />
<Connection Source="Pumps.FillContentsChanged" Destination="IO.FillContentsChanged" />
<Connection Source="Pumps.StartVacuumTimer" Destination="IO.StartVacuumTimer" />
<Connection Source="Motor.StartPump" Destination="Pumps.StartPump" />
<Connection Source="Motor.InjectDone" Destination="Door.BottlingDone" />
<Connection Source="Motor.InjectDone" Destination="Conveyor.InjectDone" />
<Connection Source="Motor.InjectDone" Destination="IO.InjectDone" />
<Connection Source="Motor.InjectorPositionChanged" Destination="IO.InjectorPositionChanged" /></EventConnections>
  <DataConnections><Connection Source="IO.EmergencyStop" Destination="Door.EmergencyStop" />
<Connection Source="IO.EmergencyStop" Destination="Conveyor.EmergencyStop" />
<Connection Source="IO.EmergencyStop" Destination="Motor.EmergencyStop" />
<Connection Source="IO.EmergencyStop" Destination="Pumps.EmergencyStop" />
<Connection Source="IO.CanisterPressure" Destination="Pumps.CanisterPressure" />
<Connection Source="IO.FillContentsAvailable" Destination="Pumps.FillContentsAvailable" />
<Connection Source="IO.DoorSiteLaser" Destination="CCounter.DoorSiteLaser" />
<Connection Source="IO.InjectSiteLaser" Destination="Conveyor.InjectSiteLaser" />
<Connection Source="IO.RejectSiteLaser" Destination="RejectArm.RejectSiteLaser" />
<Connection Source="IO.RejectBinLaser" Destination="CCounter.RejectBinLaser" />
<Connection Source="IO.AcceptBinLaser" Destination="CCounter.AcceptBinLaser" />
<Connection Source="CCounter.CanisterCount" Destination="IO.CanisterCount" />
<Connection Source="Conveyor.ConveyorSpeed" Destination="IO.ConveyorSpeed" />
<Connection Source="Pumps.InjectorContentsValveOpen" Destination="IO.InjectorContentsValveOpen" />
<Connection Source="Pumps.InjectorVacuumRun" Destination="IO.InjectorVacuumRun" />
<Connection Source="Pumps.InjectorPressurePumpRun" Destination="IO.InjectorPressurePumpRun" />
<Connection Source="Pumps.FillContents" Destination="IO.FillContents" />
<Connection Source="Motor.InjectorPosition" Destination="IO.InjectorPosition" /></DataConnections>
</FBNetwork>
</ResourceType>