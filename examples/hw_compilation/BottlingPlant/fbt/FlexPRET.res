<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<FBType Name="FlexPRET" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="UOA" Version="0.1" Author="hpea485" Date="2016-00-28" />
<CompilerInfo header="" classdef="">
</CompilerInfo>
<InterfaceList>
<EventInputs>
  <Event Name="InjectorArmFinishMovement" Comment="">
  </Event>
  <Event Name="EmergencyStopChanged" Comment="">
    <With Var="EmergencyStop" />
  </Event>
  <Event Name="CanisterPressureChanged" Comment="">
    <With Var="CanisterPressure" />
  </Event>
  <Event Name="FillContentsAvailableChanged" Comment="">
    <With Var="FillContentsAvailable" />
  </Event>
  <Event Name="LasersChanged" Comment="">
    <With Var="DoorSiteLaser" />
    <With Var="InjectSiteLaser" />
    <With Var="RejectSiteLaser" />
    <With Var="RejectBinLaser" />
    <With Var="AcceptBinLaser" />
  </Event>
  <Event Name="DoorOverride" Comment="">
  </Event>
  <Event Name="VacuumTimerElapsed" Comment="">
  </Event>
</EventInputs>
<EventOutputs>
  <Event Name="DoorReleaseCanister" Comment="">
  </Event>
  <Event Name="ConveyorChanged" Comment="">
    <With Var="ConveyorSpeed" />
  </Event>
  <Event Name="InjectorPositionChanged" Comment="">
    <With Var="InjectorPosition" />
  </Event>
  <Event Name="InjectorControlsChanged" Comment="">
    <With Var="InjectorContentsValveOpen" />
    <With Var="InjectorVacuumRun" />
    <With Var="InjectorPressurePumpRun" />
  </Event>
  <Event Name="FillContentsChanged" Comment="">
    <With Var="FillContents" />
  </Event>
  <Event Name="StartVacuumTimer" Comment="">
  </Event>
  <Event Name="GoRejectArm" Comment="">
  </Event>
  <Event Name="CanisterCountChanged" Comment="">
    <With Var="CanisterCount" />
  </Event>
  <Event Name="InjectDone" Comment="">
  </Event>
</EventOutputs>
<InputVars>
  <VarDeclaration Name="EmergencyStop" Type="BOOL" Comment="" />
  <VarDeclaration Name="CanisterPressure" Type="BYTE" Comment="" />
  <VarDeclaration Name="FillContentsAvailable" Type="BYTE" Comment="" />
  <VarDeclaration Name="DoorSiteLaser" Type="BOOL" Comment="" />
  <VarDeclaration Name="InjectSiteLaser" Type="BOOL" Comment="" />
  <VarDeclaration Name="RejectSiteLaser" Type="BOOL" Comment="" />
  <VarDeclaration Name="RejectBinLaser" Type="BOOL" Comment="" />
  <VarDeclaration Name="AcceptBinLaser" Type="BOOL" Comment="" />
</InputVars>
<OutputVars>
  <VarDeclaration Name="ConveyorSpeed" Type="BYTE" Comment="" />
  <VarDeclaration Name="InjectorPosition" Type="BYTE" Comment="" />
  <VarDeclaration Name="InjectorContentsValveOpen" Type="BOOL" Comment="" />
  <VarDeclaration Name="InjectorVacuumRun" Type="BOOL" Comment="" />
  <VarDeclaration Name="InjectorPressurePumpRun" Type="BOOL" Comment="" />
  <VarDeclaration Name="FillContents" Type="BOOL" Comment="" />
  <VarDeclaration Name="CanisterCount" Type="BYTE" Comment="" />
</OutputVars>
</InterfaceList>
<FBNetwork>
  <FB Name="CCounter" Type="CanisterCounter" x="3543.75" y="1662.5" />
  <FB Name="Door" Type="DoorController" x="3543.75" y="787.5" />
  <FB Name="Injector" Type="InjectorController" x="3237.5" y="2362.5" />
  <FB Name="Conveyor" Type="ConveyorController" x="5613.56461636749" y="2815.82283526615" />
  <FB Name="RejectArm" Type="RejectArmController" x="5687.5" y="1793.75" />
  <EventConnections><Connection Source="InjectorArmFinishMovement" Destination="Injector.InjectorArmFinishedMovement" />
<Connection Source="EmergencyStopChanged" Destination="Door.EmergencyStopChanged" />
<Connection Source="EmergencyStopChanged" Destination="Injector.EmergencyStopChanged" />
<Connection Source="EmergencyStopChanged" Destination="Conveyor.EmergencyStopChanged" />
<Connection Source="CanisterPressureChanged" Destination="Injector.CanisterPressureChanged" />
<Connection Source="FillContentsAvailableChanged" Destination="Injector.FillContentsAvailableChanged" />
<Connection Source="LasersChanged" Destination="CCounter.LasersChanged" />
<Connection Source="LasersChanged" Destination="RejectArm.LasersChanged" />
<Connection Source="LasersChanged" Destination="Conveyor.LasersChanged" />
<Connection Source="DoorOverride" Destination="Door.ReleaseDoorOverride" />
<Connection Source="VacuumTimerElapsed" Destination="Injector.VacuumTimerElapsed" />
<Connection Source="CCounter.CanisterCountChanged" Destination="CanisterCountChanged" />
<Connection Source="Door.DoorReleaseCanister" Destination="DoorReleaseCanister" />
<Connection Source="Injector.InjectDone" Destination="Door.BottlingDone" />
<Connection Source="Injector.InjectDone" Destination="Conveyor.InjectDone" />
<Connection Source="Injector.InjectorPositionChanged" Destination="InjectorPositionChanged" />
<Connection Source="Injector.InjectorControlsChanged" Destination="InjectorControlsChanged" />
<Connection Source="Injector.RejectCanister" Destination="RejectArm.RejectCanister" />
<Connection Source="Injector.FillContentsChanged" Destination="FillContentsChanged" />
<Connection Source="Injector.StartVacuumTimer" Destination="StartVacuumTimer" />
<Connection Source="Conveyor.ConveyorChanged" Destination="ConveyorChanged" />
<Connection Source="Conveyor.ConveyorStoppedForInject" Destination="Injector.ConveyorStoppedForInject" />
<Connection Source="RejectArm.GoRejectArm" Destination="GoRejectArm" /></EventConnections>
  <DataConnections><Connection Source="EmergencyStop" Destination="Door.EmergencyStop" />
<Connection Source="EmergencyStop" Destination="Injector.EmergencyStop" />
<Connection Source="EmergencyStop" Destination="Conveyor.EmergencyStop" />
<Connection Source="CanisterPressure" Destination="Injector.CanisterPressure" />
<Connection Source="FillContentsAvailable" Destination="Injector.FillContentsAvailable" />
<Connection Source="DoorSiteLaser" Destination="CCounter.DoorSiteLaser" />
<Connection Source="InjectSiteLaser" Destination="Conveyor.InjectSiteLaser" />
<Connection Source="RejectSiteLaser" Destination="RejectArm.RejectSiteLaser" />
<Connection Source="RejectBinLaser" Destination="CCounter.RejectBinLaser" />
<Connection Source="AcceptBinLaser" Destination="CCounter.AcceptBinLaser" />
<Connection Source="CCounter.CanisterCount" Destination="CanisterCount" />
<Connection Source="Injector.InjectorPosition" Destination="InjectorPosition" />
<Connection Source="Injector.InjectorContentsValveOpen" Destination="InjectorContentsValveOpen" />
<Connection Source="Injector.InjectorVacuumRun" Destination="InjectorVacuumRun" />
<Connection Source="Injector.InjectorPressurePumpRun" Destination="InjectorPressurePumpRun" />
<Connection Source="Injector.FillContents" Destination="FillContents" />
<Connection Source="Conveyor.ConveyorSpeed" Destination="ConveyorSpeed" /></DataConnections>
</FBNetwork>
</FBType>