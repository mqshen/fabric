// Code generated by protoc-gen-go.
// source: peer/fabric_message.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message_Type int32

const (
	// Undefined exists to prevent invalid message construction.
	Message_UNDEFINED Message_Type = 0
	// Handshake messages.
	Message_DISCOVERY Message_Type = 1
	// Sent to catch up with existing peers.
	Message_SYNC Message_Type = 2
)

var Message_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "DISCOVERY",
	2: "SYNC",
}
var Message_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"DISCOVERY": 1,
	"SYNC":      2,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (Message_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

// A Message encapsulates a payload of the indicated type in this message.
type Message struct {
	// Type of this message.
	Type Message_Type `protobuf:"varint,1,opt,name=type,enum=protos.Message_Type" json:"type,omitempty"`
	// Version indicates message protocol version.
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// The payload in this message. The way it should be unmarshaled
	// depends in the type field
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func init() {
	proto.RegisterType((*Message)(nil), "protos.Message")
	proto.RegisterEnum("protos.Message_Type", Message_Type_name, Message_Type_value)
}

func init() { proto.RegisterFile("peer/fabric_message.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0x4d, 0x2d,
	0xd2, 0x4f, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0x8e, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xd3, 0x19, 0xb9, 0xd8, 0x7d,
	0x21, 0x32, 0x42, 0x1a, 0x5c, 0x2c, 0x25, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c,
	0x46, 0x22, 0x10, 0x95, 0xc5, 0x7a, 0x50, 0x69, 0xbd, 0x90, 0xca, 0x82, 0xd4, 0x20, 0xb0, 0x0a,
	0x21, 0x09, 0x2e, 0xf6, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x18, 0x17, 0x24, 0x53, 0x90, 0x58, 0x99, 0x93, 0x9f, 0x98, 0x22, 0xc1, 0xac, 0xc0,
	0xa8, 0xc1, 0x13, 0x04, 0xe3, 0x2a, 0xe9, 0x71, 0xb1, 0x80, 0x4c, 0x10, 0xe2, 0xe5, 0xe2, 0x0c,
	0xf5, 0x73, 0x71, 0x75, 0xf3, 0xf4, 0x73, 0x75, 0x11, 0x60, 0x00, 0x71, 0x5d, 0x3c, 0x83, 0x9d,
	0xfd, 0xc3, 0x5c, 0x83, 0x22, 0x05, 0x18, 0x85, 0x38, 0xb8, 0x58, 0x82, 0x23, 0xfd, 0x9c, 0x05,
	0x98, 0x9c, 0xb4, 0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x33, 0x2a, 0x0b, 0x52, 0x8b, 0x72, 0x52, 0x53, 0xd2, 0xe1, 0x1e, 0xd2, 0x87, 0x38, 0x4f, 0x1f,
	0xe4, 0xc7, 0x24, 0x88, 0x77, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x59, 0x84, 0x5f,
	0xf2, 0x00, 0x00, 0x00,
}