// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package deprecatedzconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ACEMatch struct {
	Type  string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ACEMatch) Reset()                    { *m = ACEMatch{} }
func (m *ACEMatch) String() string            { return proto.CompactTextString(m) }
func (*ACEMatch) ProtoMessage()               {}
func (*ACEMatch) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ACEMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ACEMatch) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ACEAction struct {
	Drop                bool   `protobuf:"varint,1,opt,name=drop" json:"drop,omitempty"`
	Limit               bool   `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Limitrate           uint32 `protobuf:"varint,3,opt,name=limitrate" json:"limitrate,omitempty"`
	Limitunit           string `protobuf:"bytes,4,opt,name=limitunit" json:"limitunit,omitempty"`
	Limitburst          uint32 `protobuf:"varint,5,opt,name=limitburst" json:"limitburst,omitempty"`
	Allocate            bool   `protobuf:"varint,6,opt,name=allocate" json:"allocate,omitempty"`
	Exportprivate       bool   `protobuf:"varint,7,opt,name=exportprivate" json:"exportprivate,omitempty"`
	Allocationprefix    []byte `protobuf:"bytes,8,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32 `protobuf:"varint,9,opt,name=allocationprefixlen" json:"allocationprefixlen,omitempty"`
}

func (m *ACEAction) Reset()                    { *m = ACEAction{} }
func (m *ACEAction) String() string            { return proto.CompactTextString(m) }
func (*ACEAction) ProtoMessage()               {}
func (*ACEAction) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ACEAction) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *ACEAction) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *ACEAction) GetLimitrate() uint32 {
	if m != nil {
		return m.Limitrate
	}
	return 0
}

func (m *ACEAction) GetLimitunit() string {
	if m != nil {
		return m.Limitunit
	}
	return ""
}

func (m *ACEAction) GetLimitburst() uint32 {
	if m != nil {
		return m.Limitburst
	}
	return 0
}

func (m *ACEAction) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *ACEAction) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *ACEAction) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *ACEAction) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

type ACE struct {
	Matches []*ACEMatch  `protobuf:"bytes,1,rep,name=matches" json:"matches,omitempty"`
	Actions []*ACEAction `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *ACE) Reset()                    { *m = ACE{} }
func (m *ACE) String() string            { return proto.CompactTextString(m) }
func (*ACE) ProtoMessage()               {}
func (*ACE) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ACE) GetMatches() []*ACEMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *ACE) GetActions() []*ACEAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

type NameToEid struct {
	Hostname string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Eids     []string `protobuf:"bytes,2,rep,name=eids" json:"eids,omitempty"`
}

func (m *NameToEid) Reset()                    { *m = NameToEid{} }
func (m *NameToEid) String() string            { return proto.CompactTextString(m) }
func (*NameToEid) ProtoMessage()               {}
func (*NameToEid) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *NameToEid) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *NameToEid) GetEids() []string {
	if m != nil {
		return m.Eids
	}
	return nil
}

type AdditionInfoDevice struct {
	Underlayip string `protobuf:"bytes,1,opt,name=underlayip" json:"underlayip,omitempty"`
	Hostname   string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	City       string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	Region     string `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	Country    string `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	Loc        string `protobuf:"bytes,6,opt,name=loc" json:"loc,omitempty"`
	Org        string `protobuf:"bytes,7,opt,name=org" json:"org,omitempty"`
}

func (m *AdditionInfoDevice) Reset()                    { *m = AdditionInfoDevice{} }
func (m *AdditionInfoDevice) String() string            { return proto.CompactTextString(m) }
func (*AdditionInfoDevice) ProtoMessage()               {}
func (*AdditionInfoDevice) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *AdditionInfoDevice) GetUnderlayip() string {
	if m != nil {
		return m.Underlayip
	}
	return ""
}

func (m *AdditionInfoDevice) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *AdditionInfoDevice) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *AdditionInfoDevice) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AdditionInfoDevice) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AdditionInfoDevice) GetLoc() string {
	if m != nil {
		return m.Loc
	}
	return ""
}

func (m *AdditionInfoDevice) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

type OverlayNetwork struct {
	Iid           uint32              `protobuf:"varint,1,opt,name=iid" json:"iid,omitempty"`
	Ip            string              `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Lispsignature string              `protobuf:"bytes,3,opt,name=lispsignature" json:"lispsignature,omitempty"`
	Acls          []*ACE              `protobuf:"bytes,4,rep,name=acls" json:"acls,omitempty"`
	Nmtoeid       []*NameToEid        `protobuf:"bytes,5,rep,name=nmtoeid" json:"nmtoeid,omitempty"`
	Addinfodev    *AdditionInfoDevice `protobuf:"bytes,6,opt,name=addinfodev" json:"addinfodev,omitempty"`
}

func (m *OverlayNetwork) Reset()                    { *m = OverlayNetwork{} }
func (m *OverlayNetwork) String() string            { return proto.CompactTextString(m) }
func (*OverlayNetwork) ProtoMessage()               {}
func (*OverlayNetwork) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *OverlayNetwork) GetIid() uint32 {
	if m != nil {
		return m.Iid
	}
	return 0
}

func (m *OverlayNetwork) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *OverlayNetwork) GetLispsignature() string {
	if m != nil {
		return m.Lispsignature
	}
	return ""
}

func (m *OverlayNetwork) GetAcls() []*ACE {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *OverlayNetwork) GetNmtoeid() []*NameToEid {
	if m != nil {
		return m.Nmtoeid
	}
	return nil
}

func (m *OverlayNetwork) GetAddinfodev() *AdditionInfoDevice {
	if m != nil {
		return m.Addinfodev
	}
	return nil
}

type EIDAllocation struct {
	Allocate            bool   `protobuf:"varint,1,opt,name=allocate" json:"allocate,omitempty"`
	Exportprivate       bool   `protobuf:"varint,2,opt,name=exportprivate" json:"exportprivate,omitempty"`
	Allocationprefix    []byte `protobuf:"bytes,3,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32 `protobuf:"varint,4,opt,name=allocationprefixlen" json:"allocationprefixlen,omitempty"`
}

func (m *EIDAllocation) Reset()                    { *m = EIDAllocation{} }
func (m *EIDAllocation) String() string            { return proto.CompactTextString(m) }
func (*EIDAllocation) ProtoMessage()               {}
func (*EIDAllocation) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *EIDAllocation) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *EIDAllocation) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *EIDAllocation) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *EIDAllocation) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

type EIDConfigDetails struct {
	Iid           uint32         `protobuf:"varint,1,opt,name=iid" json:"iid,omitempty"`
	Eidalloc      *EIDAllocation `protobuf:"bytes,2,opt,name=eidalloc" json:"eidalloc,omitempty"`
	Eid           string         `protobuf:"bytes,3,opt,name=eid" json:"eid,omitempty"`
	Lispsignature string         `protobuf:"bytes,4,opt,name=lispsignature" json:"lispsignature,omitempty"`
	Pemcert       []byte         `protobuf:"bytes,5,opt,name=pemcert,proto3" json:"pemcert,omitempty"`
	Pemprivatekey []byte         `protobuf:"bytes,6,opt,name=pemprivatekey,proto3" json:"pemprivatekey,omitempty"`
}

func (m *EIDConfigDetails) Reset()                    { *m = EIDConfigDetails{} }
func (m *EIDConfigDetails) String() string            { return proto.CompactTextString(m) }
func (*EIDConfigDetails) ProtoMessage()               {}
func (*EIDConfigDetails) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *EIDConfigDetails) GetIid() uint32 {
	if m != nil {
		return m.Iid
	}
	return 0
}

func (m *EIDConfigDetails) GetEidalloc() *EIDAllocation {
	if m != nil {
		return m.Eidalloc
	}
	return nil
}

func (m *EIDConfigDetails) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func (m *EIDConfigDetails) GetLispsignature() string {
	if m != nil {
		return m.Lispsignature
	}
	return ""
}

func (m *EIDConfigDetails) GetPemcert() []byte {
	if m != nil {
		return m.Pemcert
	}
	return nil
}

func (m *EIDConfigDetails) GetPemprivatekey() []byte {
	if m != nil {
		return m.Pemprivatekey
	}
	return nil
}

type EIDOverlayConfig struct {
	Eidcfgdetails *EIDConfigDetails `protobuf:"bytes,1,opt,name=eidcfgdetails" json:"eidcfgdetails,omitempty"`
	Acls          []*ACE            `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	Ntoeid        []*NameToEid      `protobuf:"bytes,3,rep,name=ntoeid" json:"ntoeid,omitempty"`
}

func (m *EIDOverlayConfig) Reset()                    { *m = EIDOverlayConfig{} }
func (m *EIDOverlayConfig) String() string            { return proto.CompactTextString(m) }
func (*EIDOverlayConfig) ProtoMessage()               {}
func (*EIDOverlayConfig) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *EIDOverlayConfig) GetEidcfgdetails() *EIDConfigDetails {
	if m != nil {
		return m.Eidcfgdetails
	}
	return nil
}

func (m *EIDOverlayConfig) GetAcls() []*ACE {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *EIDOverlayConfig) GetNtoeid() []*NameToEid {
	if m != nil {
		return m.Ntoeid
	}
	return nil
}

type UnderlayNetwork struct {
	Acls []*ACE `protobuf:"bytes,1,rep,name=acls" json:"acls,omitempty"`
}

func (m *UnderlayNetwork) Reset()                    { *m = UnderlayNetwork{} }
func (m *UnderlayNetwork) String() string            { return proto.CompactTextString(m) }
func (*UnderlayNetwork) ProtoMessage()               {}
func (*UnderlayNetwork) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *UnderlayNetwork) GetAcls() []*ACE {
	if m != nil {
		return m.Acls
	}
	return nil
}

func init() {
	proto.RegisterType((*ACEMatch)(nil), "ACEMatch")
	proto.RegisterType((*ACEAction)(nil), "ACEAction")
	proto.RegisterType((*ACE)(nil), "ACE")
	proto.RegisterType((*NameToEid)(nil), "NameToEid")
	proto.RegisterType((*AdditionInfoDevice)(nil), "AdditionInfoDevice")
	proto.RegisterType((*OverlayNetwork)(nil), "OverlayNetwork")
	proto.RegisterType((*EIDAllocation)(nil), "EIDAllocation")
	proto.RegisterType((*EIDConfigDetails)(nil), "EIDConfigDetails")
	proto.RegisterType((*EIDOverlayConfig)(nil), "EIDOverlayConfig")
	proto.RegisterType((*UnderlayNetwork)(nil), "UnderlayNetwork")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x96, 0x9d, 0x34, 0x89, 0x4f, 0x9a, 0xde, 0xde, 0xe9, 0xd5, 0x95, 0x75, 0x75, 0x05, 0x91,
	0xa9, 0x44, 0x54, 0x24, 0x17, 0xa5, 0x48, 0x2c, 0x58, 0x85, 0x24, 0x8b, 0x22, 0x51, 0xd0, 0x08,
	0x36, 0xec, 0x5c, 0xcf, 0x24, 0x1d, 0xd5, 0x9e, 0xb1, 0xc6, 0xe3, 0xd0, 0xf4, 0x0d, 0x78, 0x1a,
	0xc4, 0x96, 0x27, 0x60, 0xc3, 0x3b, 0xa1, 0x39, 0xfe, 0x69, 0xd3, 0x14, 0x01, 0xbb, 0x73, 0xbe,
	0xf3, 0x33, 0xe7, 0x9c, 0xef, 0xb3, 0x61, 0x20, 0xb9, 0xf9, 0xa8, 0xf4, 0x65, 0x98, 0x69, 0x65,
	0x54, 0xf0, 0x0c, 0x7a, 0x93, 0xe9, 0xfc, 0x75, 0x64, 0xe2, 0x0b, 0x42, 0xa0, 0x6d, 0xd6, 0x19,
	0xf7, 0x9d, 0xa1, 0x33, 0xf2, 0x28, 0xda, 0xe4, 0x1f, 0xd8, 0x59, 0x45, 0x49, 0xc1, 0x7d, 0x17,
	0xc1, 0xd2, 0x09, 0xbe, 0xb8, 0xe0, 0x4d, 0xa6, 0xf3, 0x49, 0x6c, 0x84, 0x92, 0xb6, 0x8e, 0x69,
	0x95, 0x61, 0x5d, 0x8f, 0xa2, 0x6d, 0xeb, 0x12, 0x91, 0x0a, 0x83, 0x75, 0x3d, 0x5a, 0x3a, 0xe4,
	0x7f, 0xf0, 0xd0, 0xd0, 0x91, 0xe1, 0x7e, 0x6b, 0xe8, 0x8c, 0x06, 0xf4, 0x06, 0x68, 0xa2, 0x85,
	0x14, 0xc6, 0x6f, 0xe3, 0x7b, 0x37, 0x00, 0x79, 0x00, 0x80, 0xce, 0x79, 0xa1, 0x73, 0xe3, 0xef,
	0x60, 0xf1, 0x2d, 0x84, 0xfc, 0x07, 0xbd, 0x28, 0x49, 0x54, 0x6c, 0x5b, 0x77, 0xf0, 0xd1, 0xc6,
	0x27, 0x87, 0x30, 0xe0, 0x57, 0x99, 0xd2, 0x26, 0xd3, 0x62, 0x65, 0x13, 0xba, 0x98, 0xb0, 0x09,
	0x92, 0x23, 0xd8, 0xaf, 0x2a, 0x84, 0x92, 0x99, 0xe6, 0x0b, 0x71, 0xe5, 0xf7, 0x86, 0xce, 0x68,
	0x97, 0x6e, 0xe1, 0xe4, 0x29, 0x1c, 0xdc, 0xc5, 0x12, 0x2e, 0x7d, 0x0f, 0xc7, 0xba, 0x2f, 0x14,
	0xbc, 0x85, 0xd6, 0x64, 0x3a, 0x27, 0x8f, 0xa0, 0x9b, 0xda, 0x6b, 0xf3, 0xdc, 0x77, 0x86, 0xad,
	0x51, 0x7f, 0xec, 0x85, 0x35, 0x01, 0xb4, 0x8e, 0x90, 0x43, 0xe8, 0x46, 0x78, 0xdb, 0xdc, 0x77,
	0x31, 0x09, 0xc2, 0xe6, 0xdc, 0xb4, 0x0e, 0x05, 0x2f, 0xc0, 0x3b, 0x8b, 0x52, 0xfe, 0x4e, 0xcd,
	0x05, 0xb3, 0xeb, 0x5f, 0xa8, 0xdc, 0xc8, 0x28, 0xad, 0x09, 0x6c, 0x7c, 0x4b, 0x10, 0x17, 0xac,
	0xec, 0xe5, 0x51, 0xb4, 0x83, 0xaf, 0x0e, 0x90, 0x09, 0x63, 0xc2, 0xb6, 0x3a, 0x95, 0x0b, 0x35,
	0xe3, 0x2b, 0x11, 0x73, 0x7b, 0xe5, 0x42, 0x32, 0xae, 0x93, 0x68, 0x2d, 0xb2, 0xaa, 0xd1, 0x2d,
	0x64, 0xe3, 0x19, 0x77, 0xfb, 0x99, 0x58, 0x98, 0x35, 0x12, 0xeb, 0x51, 0xb4, 0xc9, 0xbf, 0xd0,
	0xd1, 0x7c, 0x29, 0x94, 0xac, 0x08, 0xad, 0x3c, 0xe2, 0x43, 0x37, 0x56, 0x85, 0x34, 0x7a, 0x8d,
	0x54, 0x7a, 0xb4, 0x76, 0xc9, 0x3e, 0xb4, 0x12, 0x15, 0x23, 0x85, 0x1e, 0xb5, 0xa6, 0x45, 0x94,
	0x5e, 0x22, 0x67, 0x1e, 0xb5, 0x66, 0xf0, 0xdd, 0x81, 0xbd, 0x37, 0x2b, 0x9c, 0xe9, 0xac, 0x94,
	0xb3, 0x4d, 0x12, 0x82, 0xe1, 0xc4, 0x03, 0x6a, 0x4d, 0xb2, 0x07, 0xae, 0xc8, 0xaa, 0x21, 0x5d,
	0x91, 0x59, 0x11, 0x24, 0x22, 0xcf, 0x72, 0xb1, 0x94, 0x91, 0x29, 0x34, 0xaf, 0xe6, 0xdc, 0x04,
	0x89, 0x0f, 0xed, 0x28, 0x4e, 0x72, 0xbf, 0x8d, 0x77, 0x6f, 0xdb, 0xbb, 0x53, 0x44, 0x2c, 0x29,
	0x32, 0x35, 0x8a, 0x0b, 0xe6, 0xef, 0x54, 0xa4, 0x34, 0xe7, 0xa7, 0x75, 0x88, 0x9c, 0x00, 0x44,
	0x8c, 0x09, 0xb9, 0x50, 0x8c, 0xaf, 0x70, 0x8b, 0xfe, 0xf8, 0x20, 0xdc, 0xbe, 0x34, 0xbd, 0x95,
	0x16, 0x7c, 0x76, 0x60, 0x30, 0x3f, 0x9d, 0x4d, 0x1a, 0xd9, 0x6c, 0xa8, 0xd9, 0xf9, 0x95, 0x9a,
	0xdd, 0xdf, 0x55, 0x73, 0xeb, 0xcf, 0xd4, 0xdc, 0xfe, 0xb9, 0x9a, 0xbf, 0x39, 0xb0, 0x3f, 0x3f,
	0x9d, 0x4d, 0x95, 0x5c, 0x88, 0xe5, 0x8c, 0x9b, 0x48, 0x24, 0xf9, 0x3d, 0x1c, 0x1c, 0x41, 0x8f,
	0x0b, 0x86, 0x0d, 0x70, 0xca, 0xfe, 0x78, 0x2f, 0xdc, 0x58, 0x94, 0x36, 0x71, 0x5b, 0x6d, 0x6f,
	0x5b, 0xb2, 0x62, 0xcd, 0x6d, 0xc6, 0xda, 0xf7, 0x33, 0xd6, 0xcd, 0x78, 0x1a, 0x73, 0x5d, 0xfe,
	0x15, 0x76, 0x69, 0xed, 0xda, 0xfa, 0x8c, 0xa7, 0xd5, 0x41, 0x2e, 0xf9, 0x1a, 0xe9, 0xd8, 0xa5,
	0x9b, 0x60, 0xf0, 0xa9, 0x5c, 0xa5, 0xd2, 0x53, 0xb9, 0x11, 0x79, 0x0e, 0x03, 0x2e, 0x58, 0xbc,
	0x58, 0xb2, 0x72, 0x37, 0x5c, 0xaa, 0x3f, 0xfe, 0x3b, 0xbc, 0xbb, 0x34, 0xdd, 0xcc, 0x6b, 0xf4,
	0xe3, 0x6e, 0xe9, 0x27, 0x80, 0x8e, 0x2c, 0xe5, 0xd3, 0xda, 0x92, 0x4f, 0x15, 0x09, 0x9e, 0xc0,
	0x5f, 0xef, 0xab, 0x8f, 0xad, 0x16, 0x76, 0xdd, 0xd0, 0xb9, 0xdb, 0xf0, 0xe5, 0x2b, 0x78, 0x18,
	0xab, 0x34, 0xbc, 0xe6, 0x8c, 0xb3, 0x28, 0x8c, 0x13, 0x55, 0xb0, 0xb0, 0xc8, 0xb9, 0xb6, 0xf2,
	0x2a, 0x7f, 0xef, 0x1f, 0x1e, 0x2f, 0x85, 0xb9, 0x28, 0xce, 0xc3, 0x58, 0xa5, 0xc7, 0x65, 0xde,
	0x71, 0x94, 0x89, 0x63, 0xc6, 0x33, 0xcd, 0xad, 0x96, 0xd8, 0x75, 0x8c, 0xcb, 0x9c, 0x77, 0x30,
	0xff, 0xe4, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x93, 0x15, 0xa5, 0x1f, 0x06, 0x00, 0x00,
}
