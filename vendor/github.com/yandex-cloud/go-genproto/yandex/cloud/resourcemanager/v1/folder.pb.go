// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/resourcemanager/v1/folder.proto

package resourcemanager // import "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status of the folder.
type Folder_Status int32

const (
	// Unspecified.
	Folder_STATUS_UNSPECIFIED Folder_Status = 0
	// The folder is active.
	Folder_ACTIVE Folder_Status = 1
	// The folder is being deleted.
	Folder_DELETING Folder_Status = 2
)

var Folder_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "ACTIVE",
	2: "DELETING",
}
var Folder_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"ACTIVE":             1,
	"DELETING":           2,
}

func (x Folder_Status) String() string {
	return proto.EnumName(Folder_Status_name, int32(x))
}
func (Folder_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_folder_2f4391f4930a4699, []int{0, 0}
}

// A Folder resource. For more information, see [Folder](/resource-manager/concepts/resources-hierarchy#folder).
type Folder struct {
	// ID of the folder.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the cloud that the folder belongs to.
	CloudId string `protobuf:"bytes,2,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the folder.
	// The name is unique within the cloud. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the folder. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the folder.
	Status               Folder_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.resourcemanager.v1.Folder_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Folder) Reset()         { *m = Folder{} }
func (m *Folder) String() string { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()    {}
func (*Folder) Descriptor() ([]byte, []int) {
	return fileDescriptor_folder_2f4391f4930a4699, []int{0}
}
func (m *Folder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Folder.Unmarshal(m, b)
}
func (m *Folder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Folder.Marshal(b, m, deterministic)
}
func (dst *Folder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Folder.Merge(dst, src)
}
func (m *Folder) XXX_Size() int {
	return xxx_messageInfo_Folder.Size(m)
}
func (m *Folder) XXX_DiscardUnknown() {
	xxx_messageInfo_Folder.DiscardUnknown(m)
}

var xxx_messageInfo_Folder proto.InternalMessageInfo

func (m *Folder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Folder) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *Folder) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Folder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Folder) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Folder) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Folder) GetStatus() Folder_Status {
	if m != nil {
		return m.Status
	}
	return Folder_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Folder)(nil), "yandex.cloud.resourcemanager.v1.Folder")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.resourcemanager.v1.Folder.LabelsEntry")
	proto.RegisterEnum("yandex.cloud.resourcemanager.v1.Folder_Status", Folder_Status_name, Folder_Status_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/resourcemanager/v1/folder.proto", fileDescriptor_folder_2f4391f4930a4699)
}

var fileDescriptor_folder_2f4391f4930a4699 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x86, 0x0b, 0x24, 0x24, 0x19, 0xaa, 0x08, 0xad, 0xaa, 0x8a, 0xfa, 0x12, 0x94, 0x13, 0x87,
	0x66, 0x57, 0x71, 0x2e, 0x4d, 0x7a, 0x72, 0x13, 0x5c, 0xa1, 0x46, 0x51, 0x05, 0xb8, 0x87, 0x5e,
	0xac, 0x85, 0x5d, 0x53, 0x54, 0x60, 0xad, 0x65, 0xb1, 0xea, 0xb7, 0xea, 0x23, 0x56, 0xdd, 0xc5,
	0x92, 0xe5, 0x8b, 0x73, 0x9b, 0x19, 0xfe, 0xff, 0x1b, 0x66, 0x76, 0xe0, 0xe3, 0x96, 0x76, 0x8c,
	0xff, 0x21, 0x65, 0x23, 0x06, 0x46, 0x24, 0xef, 0xc5, 0x20, 0x4b, 0xde, 0xd2, 0x8e, 0x56, 0x5c,
	0x92, 0xcd, 0x2d, 0x59, 0x89, 0x86, 0x71, 0x89, 0xd7, 0x52, 0x28, 0x81, 0xae, 0x8c, 0x1a, 0x6b,
	0x35, 0x3e, 0x50, 0xe3, 0xcd, 0xed, 0xe4, 0xaa, 0x12, 0xa2, 0x6a, 0x38, 0xd1, 0xf2, 0x62, 0x58,
	0x11, 0x55, 0xb7, 0xbc, 0x57, 0xb4, 0x5d, 0x1b, 0xc2, 0xf5, 0x5f, 0x07, 0xdc, 0xb9, 0x46, 0xa2,
	0x4b, 0xb0, 0x6b, 0x16, 0x58, 0xa1, 0x15, 0x5d, 0xa4, 0x76, 0xcd, 0xd0, 0x07, 0x38, 0xd7, 0xdc,
	0x65, 0xcd, 0x02, 0x5b, 0x57, 0xcf, 0x74, 0x9e, 0x30, 0x74, 0x0f, 0x50, 0x4a, 0x4e, 0x15, 0x67,
	0x4b, 0xaa, 0x02, 0x27, 0xb4, 0x22, 0x6f, 0x3a, 0xc1, 0xa6, 0x17, 0xde, 0xf5, 0xc2, 0xf9, 0xae,
	0x57, 0x7a, 0x31, 0xaa, 0x67, 0x0a, 0x21, 0x38, 0xe9, 0x68, 0xcb, 0x83, 0x13, 0x4d, 0xd4, 0x31,
	0x0a, 0xc1, 0x63, 0xbc, 0x2f, 0x65, 0xbd, 0x56, 0xb5, 0xe8, 0x82, 0x53, 0xfd, 0x69, 0xbf, 0x84,
	0xbe, 0x81, 0xdb, 0xd0, 0x82, 0x37, 0x7d, 0xe0, 0x86, 0x4e, 0xe4, 0x4d, 0xef, 0xf0, 0x91, 0xc9,
	0xb1, 0x19, 0x0a, 0x3f, 0x6b, 0x57, 0xdc, 0x29, 0xb9, 0x4d, 0x47, 0x04, 0x9a, 0x83, 0xdb, 0x2b,
	0xaa, 0x86, 0x3e, 0x38, 0x0b, 0xad, 0xe8, 0x72, 0x8a, 0x5f, 0x0b, 0xcb, 0xb4, 0x2b, 0x1d, 0xdd,
	0x93, 0x7b, 0xf0, 0xf6, 0xf0, 0xc8, 0x07, 0xe7, 0x37, 0xdf, 0x8e, 0x0b, 0xfc, 0x1f, 0xa2, 0x77,
	0x70, 0xba, 0xa1, 0xcd, 0xc0, 0xc7, 0xf5, 0x99, 0xe4, 0xc1, 0xfe, 0x64, 0x5d, 0x3f, 0x80, 0x6b,
	0x60, 0xe8, 0x3d, 0xa0, 0x2c, 0x9f, 0xe5, 0x8b, 0x6c, 0xb9, 0x78, 0xc9, 0xbe, 0xc7, 0x8f, 0xc9,
	0x3c, 0x89, 0x9f, 0xfc, 0x37, 0x08, 0xc0, 0x9d, 0x3d, 0xe6, 0xc9, 0x8f, 0xd8, 0xb7, 0xd0, 0x5b,
	0x38, 0x7f, 0x8a, 0x9f, 0xe3, 0x3c, 0x79, 0xf9, 0xea, 0xdb, 0x5f, 0x16, 0x3f, 0xb3, 0xaa, 0x56,
	0xbf, 0x86, 0x02, 0x97, 0xa2, 0x25, 0xe6, 0xd7, 0x6f, 0xcc, 0xbd, 0x54, 0xe2, 0xa6, 0xe2, 0x9d,
	0x7e, 0x00, 0x72, 0xe4, 0x90, 0x3e, 0x1f, 0x94, 0x0a, 0x57, 0xdb, 0xee, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x91, 0xd1, 0x66, 0x6f, 0x82, 0x02, 0x00, 0x00,
}