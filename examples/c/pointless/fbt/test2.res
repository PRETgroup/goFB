<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >
<ResourceType Name="test2" Comment="" >
<Identification Standard="61499-2" />
<VersionInfo Organization="HAMMONDSDESKTOP" Version="0.1" Author="Hammond" Date="2016-00-15" />
<CompilerInfo header="" classdef="">
</CompilerInfo>
<VarDeclaration Name="ac1_default_var" Type="INT" InitialValue="[0, 1, 2, 3]" ArraySize="4" Comment="" />
<VarDeclaration Name="ac2_default_var" Type="INT" InitialValue="[3, 2, 1, 0]" ArraySize="4" Comment="" />
<FBNetwork>
  <FB Name="ac1" Type="ArrayCopier" x="2975" y="1925" />
  <FB Name="ac2" Type="ArrayCopier" x="1312.5" y="1925" />
  <EventConnections><Connection Source="ac1.out" Destination="ac2.in" />
<Connection Source="ac1.set_default_out" Destination="ac1.set_default_in" />
<Connection Source="ac2.out" Destination="ac1.in" />
<Connection Source="ac2.set_default_out" Destination="ac2.set_default_in" /></EventConnections>
  <DataConnections><Connection Source="ac1_default_var" Destination="ac1.default_var" />
<Connection Source="ac2_default_var" Destination="ac2.default_var" />
<Connection Source="ac1.out_var" Destination="ac2.in_var" />
<Connection Source="ac2.out_var" Destination="ac1.in_var" /></DataConnections>
</FBNetwork>
</ResourceType>