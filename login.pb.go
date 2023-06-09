// Code generated by protoc-gen-go. DO NOT EDIT.
// source: douyin-miniapp/login.proto

package douyin_miniapp // import "github.com/dev-openapi/douyin-miniapp"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Code2SessionReq struct {
	Appid                string   `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	AnonymousCode        string   `protobuf:"bytes,4,opt,name=anonymous_code,json=anonymousCode" json:"anonymous_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code2SessionReq) Reset()         { *m = Code2SessionReq{} }
func (m *Code2SessionReq) String() string { return proto.CompactTextString(m) }
func (*Code2SessionReq) ProtoMessage()    {}
func (*Code2SessionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_727f8c19069e1bf7, []int{0}
}
func (m *Code2SessionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code2SessionReq.Unmarshal(m, b)
}
func (m *Code2SessionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code2SessionReq.Marshal(b, m, deterministic)
}
func (dst *Code2SessionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code2SessionReq.Merge(dst, src)
}
func (m *Code2SessionReq) XXX_Size() int {
	return xxx_messageInfo_Code2SessionReq.Size(m)
}
func (m *Code2SessionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Code2SessionReq.DiscardUnknown(m)
}

var xxx_messageInfo_Code2SessionReq proto.InternalMessageInfo

func (m *Code2SessionReq) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *Code2SessionReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Code2SessionReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Code2SessionReq) GetAnonymousCode() string {
	if m != nil {
		return m.AnonymousCode
	}
	return ""
}

type Code2SessionRes struct {
	ErrNo                int64    `protobuf:"varint,1,opt,name=err_no,json=errNo" json:"err_no,omitempty"`
	ErrTips              string   `protobuf:"bytes,2,opt,name=err_tips,json=errTips" json:"err_tips,omitempty"`
	SessionKey           string   `protobuf:"bytes,10,opt,name=session_key,json=sessionKey" json:"session_key,omitempty"`
	Openid               string   `protobuf:"bytes,11,opt,name=openid" json:"openid,omitempty"`
	AnonymousOpenid      string   `protobuf:"bytes,12,opt,name=anonymous_openid,json=anonymousOpenid" json:"anonymous_openid,omitempty"`
	Unionid              string   `protobuf:"bytes,13,opt,name=unionid" json:"unionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code2SessionRes) Reset()         { *m = Code2SessionRes{} }
func (m *Code2SessionRes) String() string { return proto.CompactTextString(m) }
func (*Code2SessionRes) ProtoMessage()    {}
func (*Code2SessionRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_727f8c19069e1bf7, []int{1}
}
func (m *Code2SessionRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code2SessionRes.Unmarshal(m, b)
}
func (m *Code2SessionRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code2SessionRes.Marshal(b, m, deterministic)
}
func (dst *Code2SessionRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code2SessionRes.Merge(dst, src)
}
func (m *Code2SessionRes) XXX_Size() int {
	return xxx_messageInfo_Code2SessionRes.Size(m)
}
func (m *Code2SessionRes) XXX_DiscardUnknown() {
	xxx_messageInfo_Code2SessionRes.DiscardUnknown(m)
}

var xxx_messageInfo_Code2SessionRes proto.InternalMessageInfo

func (m *Code2SessionRes) GetErrNo() int64 {
	if m != nil {
		return m.ErrNo
	}
	return 0
}

func (m *Code2SessionRes) GetErrTips() string {
	if m != nil {
		return m.ErrTips
	}
	return ""
}

func (m *Code2SessionRes) GetSessionKey() string {
	if m != nil {
		return m.SessionKey
	}
	return ""
}

func (m *Code2SessionRes) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *Code2SessionRes) GetAnonymousOpenid() string {
	if m != nil {
		return m.AnonymousOpenid
	}
	return ""
}

func (m *Code2SessionRes) GetUnionid() string {
	if m != nil {
		return m.Unionid
	}
	return ""
}

func init() {
	proto.RegisterType((*Code2SessionReq)(nil), "developer.toutiao.com.Code2SessionReq")
	proto.RegisterType((*Code2SessionRes)(nil), "developer.toutiao.com.Code2SessionRes")
}

func init() { proto.RegisterFile("douyin-miniapp/login.proto", fileDescriptor_login_727f8c19069e1bf7) }

var fileDescriptor_login_727f8c19069e1bf7 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x45, 0x7e, 0xb6, 0x63, 0xbb, 0x2e, 0x43, 0x5d, 0x54, 0xb7, 0x50, 0x23, 0xb0, 0xfb, 0x00,
	0x4b, 0xa0, 0xee, 0xba, 0x6c, 0x97, 0x2d, 0x09, 0xd8, 0x59, 0x65, 0x63, 0x64, 0xe9, 0xa2, 0x4c,
	0x62, 0xcf, 0x9d, 0xcc, 0x8c, 0x04, 0xca, 0x32, 0x8b, 0xfc, 0x40, 0x16, 0xf9, 0x9e, 0x7c, 0x43,
	0x7e, 0x21, 0x1f, 0x12, 0x66, 0x46, 0x38, 0x89, 0xc9, 0x22, 0xbb, 0x39, 0xe7, 0x9e, 0x3b, 0xe7,
	0xbe, 0xc8, 0x38, 0xc3, 0xa2, 0x62, 0x7c, 0xbe, 0x65, 0x9c, 0x25, 0x42, 0x44, 0x1b, 0xcc, 0x19,
	0x0f, 0x85, 0x44, 0x8d, 0x74, 0x94, 0x41, 0x09, 0x1b, 0x14, 0x20, 0x43, 0x8d, 0x85, 0x66, 0x09,
	0x86, 0x29, 0x6e, 0xc7, 0x5f, 0x72, 0xc4, 0x7c, 0x03, 0x51, 0x22, 0x58, 0x94, 0x70, 0x8e, 0x3a,
	0xd1, 0x0c, 0xb9, 0x72, 0x49, 0xc1, 0x05, 0x19, 0xfe, 0xc5, 0x0c, 0xe2, 0x25, 0x28, 0xc5, 0x90,
	0x2f, 0xe0, 0x9c, 0x7e, 0x20, 0xed, 0x44, 0x08, 0x96, 0xf9, 0xde, 0xc4, 0xfb, 0xfe, 0x76, 0xe1,
	0x00, 0xfd, 0x48, 0x3a, 0x0a, 0x52, 0x09, 0xda, 0x6f, 0x58, 0xba, 0x46, 0x94, 0x92, 0x56, 0x8a,
	0x19, 0xf8, 0x4d, 0xcb, 0xda, 0x37, 0x9d, 0x92, 0x77, 0x09, 0x47, 0x5e, 0x6d, 0xb1, 0x50, 0x2b,
	0x1b, 0x6d, 0xd9, 0xe8, 0x60, 0xc7, 0x1a, 0xcf, 0xe0, 0xd6, 0xdb, 0x37, 0x57, 0x74, 0x44, 0x3a,
	0x20, 0xe5, 0x8a, 0xa3, 0x75, 0x6f, 0x2e, 0xda, 0x20, 0xe5, 0x01, 0xd2, 0x4f, 0xe4, 0x8d, 0xa1,
	0x35, 0x13, 0xaa, 0xf6, 0xef, 0x82, 0x94, 0x47, 0x4c, 0x28, 0xfa, 0x95, 0xf4, 0x94, 0xcb, 0x5f,
	0x9d, 0x41, 0xe5, 0x13, 0x1b, 0x25, 0x35, 0xf5, 0x0f, 0x2a, 0x53, 0x39, 0x0a, 0xe0, 0x2c, 0xf3,
	0x7b, 0xae, 0x72, 0x87, 0xe8, 0x0f, 0xf2, 0xfe, 0xb1, 0xca, 0x5a, 0xd1, 0xb7, 0x8a, 0xe1, 0x8e,
	0x3f, 0x74, 0x52, 0x9f, 0x74, 0x0b, 0xce, 0xd0, 0x28, 0x06, 0xce, 0xbd, 0x86, 0xf1, 0x8d, 0x47,
	0xfa, 0xff, 0xcd, 0x12, 0x96, 0x20, 0x4b, 0x96, 0x02, 0xbd, 0xf2, 0x48, 0xff, 0x69, 0x53, 0x74,
	0x16, 0xbe, 0xb8, 0x97, 0x70, 0x6f, 0xec, 0xe3, 0xd7, 0xe9, 0x54, 0x30, 0xbb, 0xbc, 0xbb, 0xbf,
	0x6e, 0x4c, 0x82, 0xcf, 0x6e, 0xa3, 0x42, 0xa8, 0xa8, 0x8c, 0xa3, 0x53, 0x65, 0xe6, 0x1c, 0xd7,
	0x6d, 0xff, 0xf6, 0x7e, 0xfe, 0xf9, 0x76, 0x3c, 0xcd, 0x99, 0x3e, 0x29, 0xd6, 0xe6, 0x97, 0x28,
	0x83, 0x72, 0x6e, 0x7a, 0x34, 0x39, 0xcf, 0x6f, 0x68, 0xdd, 0xb1, 0x97, 0xf0, 0xeb, 0x21, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0x79, 0xc1, 0x40, 0x5c, 0x02, 0x00, 0x00,
}
