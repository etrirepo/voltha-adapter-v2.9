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

// Dot1AgDefaultMdLevelClassID is the 16-bit ID for the OMCI
// Managed entity Dot1ag default MD level
const Dot1AgDefaultMdLevelClassID = ClassID(301) // 0x012d

var dot1agdefaultmdlevelBME *ManagedEntityDefinition

// Dot1AgDefaultMdLevel (Class ID: #301 / 0x012d)
//	The collection of the functionality called a maintenance half-function (MHF) is not explicitly
//	modelled as a ME by either [IEEE 802.1ag] or the OMCI. The ONU automatically creates MHFs
//	according to parameters specified in a dot1ag MD or a dot1ag MA ME; the dot1ag default MD level
//	ME catches the corner cases not covered by other MEs, specifically VLANs not included by any
//	defined MA.
//
//	The dot1ag default MD level comprises a configurable table, each entry of which specifies
//	default MHF functionality for some set of VLANs. Once a set of VLANs is defined, operations to
//	different table entries or to dot1ag MAs that conflict with the set membership should be denied.
//	In addition, catch-all attributes are defined to specify MHF functionality when there is no
//	match to either a table entry or an MA.
//
//	Relationships
//		An ONU that supports [IEEE 802.1ag] automatically creates one instance of this ME for each MAC
//		bridge or IEEE 802.1p mapper, depending on the ONU's provisioning model. It should not create an
//		instance for an IEEE 802.1p mapper that is associated with a MAC bridge.
//
//	Attributes
//		Managed Entity Id
//			This attribute uniquely identifies an instance of this ME. Through an identical ID, this ME is
//			implicitly linked to an instance of the MAC bridge service profile ME or an IEEE 802.1p mapper
//			ME. It is expected that an ONU will implement CFM on bridges or on IEEE-802.1p mappers, but not
//			both, depending on its provisioning model. For precision, the reference is disambiguated by the
//			value of the layer 2 type pointer attribute. (R) (mandatory) (2-bytes)
//
//		Layer 2 Type
//			This attribute specifies whether the dot1ag default MD level ME is associated with a MAC bridge
//			service profile (value 0) or an IEEE 802.1p mapper (value-1). (R) (mandatory) (1-byte)
//
//		Catchall Level
//			This attribute ranges from 0..7 and specifies the MD level of MHFs created when no specific
//			match is found. (R,-W) (mandatory) (1-byte)
//
//		Catchall Mhf Creation
//			This attribute determines whether, when no more specific match is found, the bridge creates an
//			MHF or not. This attribute is an enumeration with the following values:
//
//			1	None. The bridge does not create any MHFs. This is the default value.
//
//			2	Default. The bridge can create MHFs on this VID on any port through which the VID can pass.
//
//			3	Explicit. The bridge can create MHFs on this VID on any port through which the VID can pass,
//			but only if an MEP exists at some lower maintenance level.
//
//			(R,-W) (mandatory) (1-byte)
//
//		Catchall Sender Id Permission
//			This attribute determines the contents of the sender ID TLV included in CFM messages transmitted
//			by MPs when no more specific match is found. This attribute is identical to that defined in the
//			description of the dot1ag MD ME (i.e., excluding code point 5, defer). (R,-W) (mandatory)
//			(1-byte)
//
//		Default Md Level Table
//			Each entry is a vector of fields, indexed by primary VLAN ID.
//
//			Primary VLAN ID (2-bytes)
//
//			Table control: This field controls the meaning of a set operation. The 1-byte size of this field
//			is included in get/get-next operations, but its value is undefined under get-next and should be
//			ignored by the OLT. (1-byte)
//
//			1	Add record to table; overwrite existing record, if any.
//
//			2	Delete record from table.
//
//			3	Clear all entries from table. This action may affect service and should be used judiciously.
//
//			Other values are reserved.
//
//			Status: This Boolean field indicates whether this table entry is in effect (true) or whether
//			(false) it has been overridden by the existence of an MA for the same VID and MD level as this
//			table's entry, and on which an up MEP is defined. This attribute is read-only. Space should be
//			allocated for it during set operations, but the value is not used. (1-byte)
//
//			Level: This field ranges from 0..7 and specifies the MD level of MHFs under the control of this
//			instance of the dot1ag default MD level. The additional value 0xFF instructs the bridge to use
//			the value in the catch-all level attribute. (1-byte)
//
//			MHF creation: This attribute determines whether the bridge creates an MHF or not, under
//			circumstances defined in clause 22.2.3 of [IEEE 802.1ag]. This attribute is an enumeration with
//			the following values. (1-byte)
//
//			1	None. No MHFs are created on this bridge for this MA.
//
//			2	Default. The bridge can create MHFs on this VID on any port through which the VID can pass.
//
//			3	Explicit. The bridge can create MHFs on this VID on any port through which the VID can pass,
//			but only if an MEP exists at some lower maintenance level.
//
//			4	Defer. This value causes the ONU to use the setting of the catch-all MHF creation attribute.
//			This is recommended to be the default value.
//
//			Sender ID permission: This attribute determines the contents of the sender ID TLV included in
//			CFM messages transmitted by MPs controlled by this MA. (1-byte)
//
//			1	None: the sender ID TLV is not to be sent, default.
//
//			2	Chassis: the chassis ID length, chassis ID subtype, and chassis ID fields of the sender ID TLV
//			are to be sent, but not the management address fields.
//
//			3	Manage: the management address fields of the sender ID TLV are to be sent, but the chassis ID
//			length is to be transmitted with a 0 value, and the chassis ID subtype, and chassis ID fields
//			are not to be sent.
//
//			4	ChassisManage: all chassis ID and management address fields are to be sent.
//
//			5	Defer: the contents of the sender ID TLV is determined by the catch-all sender ID permission
//			attribute.
//
//			Associated VLANs list: This field comprises a list of up to 11 additional VLAN IDs associated
//			with the primary VLAN, 2-bytes each. Unused placeholders, possibly including the entire field,
//			are set to 0. (22-bytes)
//
//			(R,-W) (mandatory) (29-bytes * N entries)
//
type Dot1AgDefaultMdLevel struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	dot1agdefaultmdlevelBME = &ManagedEntityDefinition{
		Name:    "Dot1AgDefaultMdLevel",
		ClassID: 301,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
			Set,
			SetTable,
		),
		AllowedAttributeMask: 0xf800,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: ByteField("Layer2Type", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2: ByteField("CatchallLevel", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, Write), false, false, false, 2),
			3: ByteField("CatchallMhfCreation", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read, Write), false, false, false, 3),
			4: ByteField("CatchallSenderIdPermission", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, Write), false, false, false, 4),
			5: TableField("DefaultMdLevelTable", TableAttributeType, 0x0800, TableInfo{nil, 29}, mapset.NewSetWith(Read, Write), false, false, false, 5),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewDot1AgDefaultMdLevel (class ID 301) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewDot1AgDefaultMdLevel(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*dot1agdefaultmdlevelBME, params...)
}
