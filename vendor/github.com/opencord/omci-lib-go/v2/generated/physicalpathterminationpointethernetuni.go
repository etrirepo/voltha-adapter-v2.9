/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 * Copyright 2020-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// PhysicalPathTerminationPointEthernetUniClassID is the 16-bit ID for the OMCI
// Managed entity Physical path termination point Ethernet UNI
const PhysicalPathTerminationPointEthernetUniClassID = ClassID(11) // 0x000b

var physicalpathterminationpointethernetuniBME *ManagedEntityDefinition

// PhysicalPathTerminationPointEthernetUni (Class ID: #11 / 0x000b)
//	This ME represents the point at an Ethernet UNI where the physical path terminates and Ethernet
//	physical level functions are performed.
//
//	The ONU automatically creates an instance of this ME per port:
//
//	o	when the ONU has Ethernet ports built into its factory configuration;
//
//	o	when a cardholder is provisioned to expect a circuit pack of the Ethernet type;
//
//	o	when a cardholder provisioned for plug-and-play is equipped with a circuit pack of the
//	Ethernet type. Note that the installation of a plug-and-play card may indicate the presence of
//	Ethernet ports via equipment ID as well as its type, and indeed may cause the ONU to instantiate
//	a port-mapping package that specifies Ethernet ports.
//
//	The ONU automatically deletes instances of this ME when a cardholder is neither provisioned to
//	expect an Ethernet circuit pack, nor is it equipped with an Ethernet circuit pack.
//
//	Relationships
//		An instance of this ME is associated with each instance of a pre-provisioned or real Ethernet
//		port.
//
//	Attributes
//		Managed Entity Id
//			This attribute uniquely identifies each instance of this ME. This 2 byte number indicates the
//			physical position of the UNI. The first byte is the slot ID (defined in clause 9.1.5). The
//			second byte is the port ID, with the range 1..255. (R) (mandatory) (2-bytes)
//
//		Expected Type
//			This attribute supports pre-provisioning. It is coded as follows:
//
//			0	Autosense
//
//			1 to 254	One of the values from Table-9.1.5-1 that is compatible with an Ethernet circuit pack
//
//			Upon ME instantiation, the ONU sets this attribute to 0. (R,-W) (mandatory) (1-byte)
//
//		Sensed Type
//			When a circuit pack is present, this attribute represents its type as one of the values from
//			Table-9.1.5-1. If the value of the expected type is not 0, then the value of the sensed type
//			should be the same as the value of the expected type. Upon ME instantiation, the ONU sets this
//			attribute to 0. See also the note in the following AVC table.
//
//			(R) (mandatory if the ONU supports circuit packs with configurable interface types, e.g., 10/100
//			BASE-T card) (1-byte)
//
//		Auto Detection Configuration
//			This attribute sets the following Ethernet port configuration.
//
//			Upon ME instantiation, the ONU sets this attribute to 0. (R,-W) (mandatory for interfaces with
//			autodetection options) (1-byte)
//
//		Ethernet Loopback Configuration
//			This attribute sets the following Ethernet loopback configuration.
//
//			0	No loopback
//
//			3	Loop 3, loopback of downstream traffic after PHY transceiver. Loop-3 is depicted in Figure
//			9.5.1-1.
//
//			Note that normal bridge behaviour may defeat the loopback signal unless broadcast MAC addresses
//			are used. Although it does not reach the physical interface, [IEEE 802.1ag] loopback is
//			preferred.
//
//			Upon ME instantiation, the ONU sets this attribute to 0. (R,-W) (mandatory) (1-byte)
//
//		Administrative State
//			This attribute locks (1) and unlocks (0) the functions performed by this ME. Administrative
//			state is further described in clause A.1.6. (R,-W) (mandatory) (1-byte)
//
//		Operational State
//			This attribute indicates whether the ME is capable of performing its function. Valid values are
//			enabled (0) and disabled (1). (R) (optional) (1-byte)
//
//		Configuration Ind
//			This attribute indicates the configuration status of the Ethernet UNI.
//
//			0x01	10BASE-T full duplex
//
//			0x02	100BASE-T full duplex
//
//			0x03	Gigabit Ethernet full duplex
//
//			0x04	10Gb/s Ethernet full duplex
//
//			0x05	2.5Gb/s Ethernet full duplex
//
//			0x06	5Gb/s Ethernet full duplex
//
//			0x07	25Gb/s Ethernet full duplex
//
//			0x08	40Gb/s Ethernet full duplex
//
//			0x11	10BASE-T half duplex
//
//			0x12	100BASE-T half duplex
//
//			0x13	Gigabit Ethernet half duplex
//
//			The value 0 indicates that the configuration status is unknown (e.g., Ethernet link is not
//			established or the circuit pack is not yet installed). Upon ME instantiation, the ONU sets this
//			attribute to 0. (R) (mandatory) (1-byte)
//
//		Max Frame Size
//			This attribute denotes the maximum frame size allowed across this interface. Upon ME
//			instantiation, the ONU sets the attribute to 1518. (R,-W) (mandatory for G-PON, optional for
//			ITU-T G.986 systems) (2 bytes)
//
//		Dte Or Dce Ind
//			This attribute specifies the following Ethernet interface wiring.
//
//			0	DCE or MDI-X (default).
//
//			1	DTE or MDI.
//
//			2	Automatic selection
//
//			(R,-W) (mandatory) (1-byte)
//
//		Pause Time
//			This attribute allows the PPTP to ask the subscriber terminal to temporarily suspend sending
//			data. Units are in pause quanta (1 pause quantum is 512 bit times of the particular
//			implementation). Values: 0..0xFFFF. Upon ME instantiation, the ONU sets this attribute to 0.
//			(R,-W) (optional) (2-bytes)
//
//		Bridged Or Ip Ind
//			This attribute specifies whether the Ethernet interface is bridged or derived from an IP router
//			function.
//
//			0	Bridged
//
//			1	IP router
//
//			2	Depends on the parent circuit pack. 2 means that the circuit pack's bridged or IP ind
//			attribute is either 0 or 1.
//
//			Upon ME instantiation, the ONU sets this attribute to 2. (R,-W) (optional) (1-byte)
//
//		Arc
//			See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Arc Interval
//			See clause A.1.4.3. (R,-W) (optional) (1-byte)
//
//		Pppoe Filter
//			This attribute controls filtering of PPPoE packets on this Ethernet port. The value 0 allows
//			packets of all types. The value 1 discards everything but PPPoE packets. The default value is 0.
//			(R,-W) (optional) (1-byte)
//
//		Power Control
//			This attribute controls whether power is provided to an external equipment over the Ethernet
//			PPTP. The value 1 enables power over the Ethernet port. The default value 0 disables power feed.
//			(R,-W) (optional) (1-byte)
//
//			NOTE - This attribute is the equivalent of the acPSEAdminControl variable defined in clause
//			30.9.1.2.1 of [IEEE 802.3]. Other variables related to PoE appear in the PoE control ME.
//
type PhysicalPathTerminationPointEthernetUni struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	physicalpathterminationpointethernetuniBME = &ManagedEntityDefinition{
		Name:    "PhysicalPathTerminationPointEthernetUni",
		ClassID: 11,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1:  ByteField("ExpectedType", EnumerationAttributeType, 0x8000, 0, mapset.NewSetWith(Read, Write), false, false, false, 1),
			2:  ByteField("SensedType", EnumerationAttributeType, 0x4000, 0, mapset.NewSetWith(Read), true, false, false, 2),
			3:  ByteField("AutoDetectionConfiguration", EnumerationAttributeType, 0x2000, 0, mapset.NewSetWith(Read, Write), false, false, false, 3),
			4:  ByteField("EthernetLoopbackConfiguration", EnumerationAttributeType, 0x1000, 0, mapset.NewSetWith(Read, Write), false, false, false, 4),
			5:  ByteField("AdministrativeState", EnumerationAttributeType, 0x0800, 0, mapset.NewSetWith(Read, Write), false, false, false, 5),
			6:  ByteField("OperationalState", EnumerationAttributeType, 0x0400, 0, mapset.NewSetWith(Read), true, true, false, 6),
			7:  ByteField("ConfigurationInd", EnumerationAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, false, false, 7),
			8:  Uint16Field("MaxFrameSize", UnsignedIntegerAttributeType, 0x0100, 1518, mapset.NewSetWith(Read, Write), false, false, false, 8),
			9:  ByteField("DteOrDceInd", EnumerationAttributeType, 0x0080, 0, mapset.NewSetWith(Read, Write), false, false, false, 9),
			10: Uint16Field("PauseTime", UnsignedIntegerAttributeType, 0x0040, 0, mapset.NewSetWith(Read, Write), false, true, false, 10),
			11: ByteField("BridgedOrIpInd", EnumerationAttributeType, 0x0020, 0, mapset.NewSetWith(Read, Write), false, true, false, 11),
			12: ByteField("Arc", EnumerationAttributeType, 0x0010, 0, mapset.NewSetWith(Read, Write), true, true, false, 12),
			13: ByteField("ArcInterval", UnsignedIntegerAttributeType, 0x0008, 0, mapset.NewSetWith(Read, Write), false, true, false, 13),
			14: ByteField("PppoeFilter", EnumerationAttributeType, 0x0004, 0, mapset.NewSetWith(Read, Write), false, true, false, 14),
			15: ByteField("PowerControl", EnumerationAttributeType, 0x0002, 0, mapset.NewSetWith(Read, Write), false, true, false, 15),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
		Alarms: AlarmMap{
			0: "LAN-LOS",
		},
	}
}

// NewPhysicalPathTerminationPointEthernetUni (class ID 11) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewPhysicalPathTerminationPointEthernetUni(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*physicalpathterminationpointethernetuniBME, params...)
}
