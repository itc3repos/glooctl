// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtualmesh.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// A Virtual Mesh is a container for a set of Virtual Services that will be used to generate a single proxy config
// to be applied to one or more Envoy nodes. The Virtual Mesh is best understood as an in-mesh application's localized view
// of the rest of the mesh.
// Each domains for each Virtual Services contained in a Virtual Mesh cannot appear more than once, or the Virtual Mesh
// will be invalid.
type VirtualMesh struct {
	// Name of the virtual mesh. Envoy nodes will be assigned a config corresponding with virtual mesh they are assigned.
	// Envoy instances must specify the virtual mesh they belong to when they register to Gloo.
	//
	// Currently this is done by specifying the name of the virtual mesh as a prefix to the Envoy's Node ID
	// which can be specified with the `--service-node` flag, or in the Envoy instance's bootstrap config.
	//
	// Names must be unique and follow the following syntax rules:
	// One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the list of names of the virtual services this vmesh includes.
	VirtualServices []string `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services,omitempty"`
	// Status indicates the validation status of the virtual mesh resource.
	// Status is read-only by clients, and set by gloo during validation
	Status *Status `protobuf:"bytes,6,opt,name=status" json:"status,omitempty" testdiff:"ignore"`
	// Metadata contains the resource metadata for the virtual mesh
	Metadata *Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *VirtualMesh) Reset()                    { *m = VirtualMesh{} }
func (m *VirtualMesh) String() string            { return proto.CompactTextString(m) }
func (*VirtualMesh) ProtoMessage()               {}
func (*VirtualMesh) Descriptor() ([]byte, []int) { return fileDescriptorVirtualmesh, []int{0} }

func (m *VirtualMesh) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMesh) GetVirtualServices() []string {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *VirtualMesh) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMesh) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMesh)(nil), "v1.VirtualMesh")
}
func (this *VirtualMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMesh)
	if !ok {
		that2, ok := that.(VirtualMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if this.VirtualServices[i] != that1.VirtualServices[i] {
			return false
		}
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("virtualmesh.proto", fileDescriptorVirtualmesh) }

var fileDescriptorVirtualmesh = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xe5, 0x52, 0x05, 0xea, 0x56, 0x40, 0x2d, 0x90, 0xa2, 0x0a, 0x95, 0x28, 0x53, 0x58,
	0x12, 0x15, 0x26, 0x18, 0xbb, 0x77, 0x49, 0x25, 0x56, 0xe4, 0x36, 0x7f, 0x5c, 0x4b, 0x4d, 0x8c,
	0xec, 0xdf, 0x79, 0x26, 0x56, 0x5e, 0x88, 0x81, 0x47, 0xe0, 0x09, 0x50, 0x9c, 0x1f, 0x86, 0x6e,
	0x77, 0xdf, 0xf9, 0xe4, 0xfb, 0xf9, 0xbc, 0xd3, 0x16, 0xbd, 0x3c, 0x36, 0xe0, 0x0e, 0xf9, 0xbb,
	0x35, 0x68, 0xc4, 0xa8, 0x5b, 0x2d, 0xee, 0x94, 0x31, 0xea, 0x08, 0x45, 0x20, 0x3b, 0x5f, 0x17,
	0x0e, 0xad, 0xdf, 0xe3, 0xf0, 0x62, 0xb1, 0x3c, 0x4d, 0x2b, 0x6f, 0x25, 0x6a, 0xd3, 0x52, 0x7e,
	0xa3, 0x8c, 0x32, 0x41, 0x16, 0xbd, 0x22, 0x3a, 0x73, 0x28, 0xd1, 0x3b, 0x72, 0x97, 0x0d, 0xa0,
	0xac, 0x24, 0xca, 0xc1, 0xa7, 0x9f, 0x8c, 0x4f, 0x5f, 0x87, 0x2d, 0x1b, 0x70, 0x07, 0x21, 0xf8,
	0xb8, 0x95, 0x0d, 0xc4, 0x2c, 0x61, 0xd9, 0xa4, 0x0c, 0x5a, 0x3c, 0xf0, 0x6b, 0x9a, 0xfb, 0xe6,
	0xc0, 0x76, 0x7a, 0x0f, 0x2e, 0x1e, 0x25, 0x67, 0xd9, 0xa4, 0xbc, 0x22, 0xbe, 0x25, 0x2c, 0x9e,
	0x79, 0x34, 0x7c, 0x17, 0x47, 0x09, 0xcb, 0xa6, 0x8f, 0x3c, 0xef, 0x56, 0xf9, 0x36, 0x90, 0xf5,
	0xed, 0xcf, 0xd7, 0xfd, 0x1c, 0xc1, 0x61, 0xa5, 0xeb, 0xfa, 0x25, 0xd5, 0xaa, 0x35, 0x16, 0xd2,
	0x92, 0x0a, 0x22, 0xe3, 0x17, 0x7f, 0xdb, 0xe2, 0xf3, 0x50, 0x9e, 0xf5, 0xe5, 0x0d, 0xb1, 0xf2,
	0x3f, 0x5d, 0x8f, 0x3f, 0xbe, 0x97, 0x6c, 0x17, 0x85, 0x03, 0x9e, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x16, 0xf9, 0x29, 0xbe, 0x4b, 0x01, 0x00, 0x00,
}
