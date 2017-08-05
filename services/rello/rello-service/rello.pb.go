// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rello.proto

/*
Package rello is a generated protocol buffer package.

It is generated from these files:
	rello.proto

It has these top-level messages:
	Empty
	ChecklistUpdate
	Model
	CheckItem
	Action
	Data
	CheckList
	Card
	Board
	MemberCreator
*/
package rello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/TuneLab/truss/deftree/googlethirdparty"

import (
	context "golang.org/x/net/context"
	newcontext "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Response structure from trello webhook
// TODO: Think about watching card in order to get card description and card
// due date for occurrence.Data and occurrence.Datetime respectively.
// TODO: Add references to the trello api documentation for fields
type ChecklistUpdate struct {
	Model  *Model  `protobuf:"bytes,1,opt,name=model" json:"model,omitempty"`
	Action *Action `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *ChecklistUpdate) Reset()                    { *m = ChecklistUpdate{} }
func (m *ChecklistUpdate) String() string            { return proto.CompactTextString(m) }
func (*ChecklistUpdate) ProtoMessage()               {}
func (*ChecklistUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChecklistUpdate) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

func (m *ChecklistUpdate) GetAction() *Action {
	if m != nil {
		return m.Action
	}
	return nil
}

type Model struct {
	Id         string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name       string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IdBoard    string       `protobuf:"bytes,3,opt,name=idBoard" json:"idBoard,omitempty"`
	IdCard     string       `protobuf:"bytes,4,opt,name=idCard" json:"idCard,omitempty"`
	Pos        int32        `protobuf:"varint,5,opt,name=pos" json:"pos,omitempty"`
	CheckItems []*CheckItem `protobuf:"bytes,6,rep,name=checkItems" json:"checkItems,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Model) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetIdBoard() string {
	if m != nil {
		return m.IdBoard
	}
	return ""
}

func (m *Model) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *Model) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Model) GetCheckItems() []*CheckItem {
	if m != nil {
		return m.CheckItems
	}
	return nil
}

type CheckItem struct {
	State       string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	IdChecklist string `protobuf:"bytes,2,opt,name=idChecklist" json:"idChecklist,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Pos         int32  `protobuf:"varint,6,opt,name=pos" json:"pos,omitempty"`
}

func (m *CheckItem) Reset()                    { *m = CheckItem{} }
func (m *CheckItem) String() string            { return proto.CompactTextString(m) }
func (*CheckItem) ProtoMessage()               {}
func (*CheckItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckItem) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CheckItem) GetIdChecklist() string {
	if m != nil {
		return m.IdChecklist
	}
	return ""
}

func (m *CheckItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CheckItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CheckItem) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type Action struct {
	Id              string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	IdMemberCreator string         `protobuf:"bytes,2,opt,name=idMemberCreator" json:"idMemberCreator,omitempty"`
	Data            *Data          `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Type            string         `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Date            string         `protobuf:"bytes,5,opt,name=date" json:"date,omitempty"`
	MemberCreator   *MemberCreator `protobuf:"bytes,6,opt,name=memberCreator" json:"memberCreator,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Action) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Action) GetIdMemberCreator() string {
	if m != nil {
		return m.IdMemberCreator
	}
	return ""
}

func (m *Action) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Action) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Action) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Action) GetMemberCreator() *MemberCreator {
	if m != nil {
		return m.MemberCreator
	}
	return nil
}

type Data struct {
	CheckItem *CheckItem `protobuf:"bytes,1,opt,name=checkItem" json:"checkItem,omitempty"`
	Checklist *CheckList `protobuf:"bytes,2,opt,name=checklist" json:"checklist,omitempty"`
	Card      *Card      `protobuf:"bytes,3,opt,name=card" json:"card,omitempty"`
	Board     *Board     `protobuf:"bytes,4,opt,name=board" json:"board,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Data) GetCheckItem() *CheckItem {
	if m != nil {
		return m.CheckItem
	}
	return nil
}

func (m *Data) GetChecklist() *CheckList {
	if m != nil {
		return m.Checklist
	}
	return nil
}

func (m *Data) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *Data) GetBoard() *Board {
	if m != nil {
		return m.Board
	}
	return nil
}

type CheckList struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *CheckList) Reset()                    { *m = CheckList{} }
func (m *CheckList) String() string            { return proto.CompactTextString(m) }
func (*CheckList) ProtoMessage()               {}
func (*CheckList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CheckList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CheckList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Card struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	ShortLink string `protobuf:"bytes,3,opt,name=shortLink" json:"shortLink,omitempty"`
	IdShort   int32  `protobuf:"varint,4,opt,name=idShort" json:"idShort,omitempty"`
}

func (m *Card) Reset()                    { *m = Card{} }
func (m *Card) String() string            { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()               {}
func (*Card) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Card) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Card) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Card) GetShortLink() string {
	if m != nil {
		return m.ShortLink
	}
	return ""
}

func (m *Card) GetIdShort() int32 {
	if m != nil {
		return m.IdShort
	}
	return 0
}

type Board struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	ShortLink string `protobuf:"bytes,3,opt,name=shortLink" json:"shortLink,omitempty"`
}

func (m *Board) Reset()                    { *m = Board{} }
func (m *Board) String() string            { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()               {}
func (*Board) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Board) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Board) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Board) GetShortLink() string {
	if m != nil {
		return m.ShortLink
	}
	return ""
}

type MemberCreator struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	AvatarHas string `protobuf:"bytes,2,opt,name=avatarHas" json:"avatarHas,omitempty"`
	FullName  string `protobuf:"bytes,3,opt,name=fullName" json:"fullName,omitempty"`
	Initials  string `protobuf:"bytes,4,opt,name=initials" json:"initials,omitempty"`
	Username  string `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
}

func (m *MemberCreator) Reset()                    { *m = MemberCreator{} }
func (m *MemberCreator) String() string            { return proto.CompactTextString(m) }
func (*MemberCreator) ProtoMessage()               {}
func (*MemberCreator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MemberCreator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MemberCreator) GetAvatarHas() string {
	if m != nil {
		return m.AvatarHas
	}
	return ""
}

func (m *MemberCreator) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *MemberCreator) GetInitials() string {
	if m != nil {
		return m.Initials
	}
	return ""
}

func (m *MemberCreator) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "rello.Empty")
	proto.RegisterType((*ChecklistUpdate)(nil), "rello.ChecklistUpdate")
	proto.RegisterType((*Model)(nil), "rello.Model")
	proto.RegisterType((*CheckItem)(nil), "rello.CheckItem")
	proto.RegisterType((*Action)(nil), "rello.Action")
	proto.RegisterType((*Data)(nil), "rello.Data")
	proto.RegisterType((*CheckList)(nil), "rello.CheckList")
	proto.RegisterType((*Card)(nil), "rello.Card")
	proto.RegisterType((*Board)(nil), "rello.Board")
	proto.RegisterType((*MemberCreator)(nil), "rello.MemberCreator")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rello service

type RelloClient interface {
	// CheckListWebhook is an endpoint for trello checklist webhooks that will
	// parse the incoming data. It ensures that a request is from trello by doing
	// a base64 encoding with the users secrete key.
	// it then calls out to ambition-model to create an occurrence
	CheckListWebhook(ctx context.Context, in *ChecklistUpdate, opts ...grpc.CallOption) (*Empty, error)
	EmptyRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type relloClient struct {
	cc *grpc.ClientConn
}

func NewRelloClient(cc *grpc.ClientConn) RelloClient {
	return &relloClient{cc}
}

func (c *relloClient) CheckListWebhook(ctx context.Context, in *ChecklistUpdate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/rello.Rello/CheckListWebhook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relloClient) EmptyRPC(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/rello.Rello/EmptyRPC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rello service

type RelloServer interface {
	// CheckListWebhook is an endpoint for trello checklist webhooks that will
	// parse the incoming data. It ensures that a request is from trello by doing
	// a base64 encoding with the users secrete key.
	// it then calls out to ambition-model to create an occurrence
	CheckListWebhook(newcontext.Context, *ChecklistUpdate) (*Empty, error)
	EmptyRPC(newcontext.Context, *Empty) (*Empty, error)
}

func RegisterRelloServer(s *grpc.Server, srv RelloServer) {
	s.RegisterService(&_Rello_serviceDesc, srv)
}

func _Rello_CheckListWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChecklistUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelloServer).CheckListWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rello.Rello/CheckListWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelloServer).CheckListWebhook(ctx, req.(*ChecklistUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rello_EmptyRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelloServer).EmptyRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rello.Rello/EmptyRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelloServer).EmptyRPC(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rello.Rello",
	HandlerType: (*RelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckListWebhook",
			Handler:    _Rello_CheckListWebhook_Handler,
		},
		{
			MethodName: "EmptyRPC",
			Handler:    _Rello_EmptyRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rello.proto",
}

func init() { proto.RegisterFile("rello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0x93, 0xec, 0xb6, 0x3b, 0x69, 0x69, 0x65, 0x55, 0x55, 0x54, 0x55, 0x22, 0x5a, 0x09,
	0x29, 0xe2, 0x90, 0x45, 0xe1, 0xd6, 0x1b, 0x04, 0x10, 0x95, 0x5a, 0x84, 0x0c, 0x88, 0x0b, 0x17,
	0x6f, 0xec, 0x36, 0x56, 0x77, 0xd7, 0x91, 0xed, 0x20, 0xe5, 0xca, 0x1b, 0x20, 0x9e, 0x80, 0x33,
	0x0f, 0xc2, 0x03, 0xf0, 0x0a, 0x3c, 0x08, 0xf2, 0xcf, 0x7a, 0x37, 0xe9, 0x05, 0x89, 0x9b, 0x67,
	0xbe, 0xcf, 0xe3, 0x6f, 0x66, 0x3e, 0x19, 0x86, 0x92, 0x95, 0xa5, 0x98, 0xae, 0xa4, 0xd0, 0x02,
	0xc5, 0x36, 0x38, 0x9b, 0xdf, 0x72, 0xbd, 0x5c, 0x17, 0xd3, 0x85, 0xa8, 0xf2, 0x0f, 0xeb, 0x9a,
	0x5d, 0x91, 0x22, 0xd7, 0x72, 0xad, 0x54, 0x4e, 0xd9, 0x8d, 0x96, 0x8c, 0xe5, 0xb7, 0x42, 0xdc,
	0x96, 0x4c, 0x2f, 0xb9, 0xa4, 0x2b, 0x22, 0xf5, 0x26, 0x27, 0x75, 0x2d, 0x34, 0xd1, 0x5c, 0xd4,
	0xca, 0xd5, 0xca, 0xf6, 0x20, 0x7e, 0x55, 0xad, 0xf4, 0x26, 0xfb, 0x0c, 0x47, 0xf3, 0x25, 0x5b,
	0xdc, 0x95, 0x5c, 0xe9, 0x8f, 0x2b, 0x4a, 0x34, 0x43, 0x19, 0xc4, 0x95, 0xa0, 0xac, 0x1c, 0x45,
	0xe3, 0x68, 0x32, 0x9c, 0x1d, 0x4c, 0x9d, 0x88, 0x6b, 0x93, 0xc3, 0x0e, 0x42, 0x8f, 0x21, 0x21,
	0x0b, 0x53, 0x70, 0xd4, 0xb3, 0xa4, 0x43, 0x4f, 0x7a, 0x6e, 0x93, 0xd8, 0x83, 0xd9, 0x8f, 0x08,
	0x62, 0x7b, 0x0f, 0x3d, 0x84, 0x1e, 0xa7, 0xb6, 0x62, 0x8a, 0x7b, 0x9c, 0x22, 0x04, 0x83, 0x9a,
	0x54, 0xcc, 0x5e, 0x4f, 0xb1, 0x3d, 0xa3, 0x11, 0xec, 0x71, 0xfa, 0x42, 0x10, 0x49, 0x47, 0x7d,
	0x9b, 0x6e, 0x42, 0x74, 0x0a, 0x09, 0xa7, 0x73, 0x03, 0x0c, 0x2c, 0xe0, 0x23, 0x74, 0x0c, 0xfd,
	0x95, 0x50, 0xa3, 0x78, 0x1c, 0x4d, 0x62, 0x6c, 0x8e, 0xe8, 0x29, 0xc0, 0xc2, 0xf4, 0x73, 0xa9,
	0x59, 0xa5, 0x46, 0xc9, 0xb8, 0x3f, 0x19, 0xce, 0x8e, 0xbd, 0xb8, 0x79, 0x03, 0xe0, 0x0e, 0x27,
	0xdb, 0x40, 0x1a, 0x00, 0x74, 0x02, 0xb1, 0xd2, 0x44, 0x33, 0xaf, 0xd4, 0x05, 0x68, 0x0c, 0x43,
	0x4e, 0xc3, 0x98, 0xbc, 0xe6, 0x6e, 0xca, 0xb7, 0xd7, 0xbf, 0xd7, 0xde, 0xa0, 0xd3, 0x9e, 0x17,
	0x9b, 0x04, 0xb1, 0xd9, 0xaf, 0x08, 0x12, 0x37, 0xb1, 0x7b, 0xf3, 0x99, 0xc0, 0x11, 0xa7, 0xd7,
	0xac, 0x2a, 0x98, 0x9c, 0x4b, 0x46, 0xb4, 0x90, 0xfe, 0xd9, 0xdd, 0x34, 0x7a, 0x04, 0x03, 0x4a,
	0x34, 0xb1, 0x8f, 0x0f, 0x67, 0x43, 0xdf, 0xeb, 0x4b, 0xa2, 0x09, 0xb6, 0x80, 0xd1, 0xa2, 0x37,
	0xab, 0xa0, 0xc5, 0x9c, 0x4d, 0xce, 0xec, 0xda, 0x4e, 0x2e, 0xb5, 0x3c, 0x86, 0x2e, 0xe0, 0xb0,
	0xda, 0x7a, 0x30, 0xb1, 0x15, 0x4f, 0x9a, 0xfd, 0x77, 0x31, 0xbc, 0x4d, 0xcd, 0x7e, 0x46, 0x30,
	0x30, 0x4f, 0xa2, 0x29, 0xa4, 0x61, 0xb6, 0xde, 0x40, 0xf7, 0xc7, 0xdf, 0x52, 0x02, 0x3f, 0x0c,
	0x76, 0x87, 0x7f, 0xc5, 0x95, 0xc6, 0x2d, 0xc5, 0x74, 0xbb, 0x68, 0x0c, 0xd2, 0x76, 0x6b, 0xcc,
	0x80, 0x2d, 0x60, 0xdc, 0x5b, 0x88, 0xc6, 0x29, 0xad, 0x7b, 0xad, 0x8f, 0xb0, 0x83, 0xb2, 0xdc,
	0xaf, 0xdc, 0x14, 0x0f, 0xab, 0x8a, 0x3a, 0xab, 0x72, 0xdb, 0xe8, 0x35, 0xdb, 0xc8, 0x0a, 0x18,
	0x58, 0xbf, 0xfd, 0x03, 0x17, 0x9d, 0x43, 0xaa, 0x96, 0x42, 0xea, 0x2b, 0x5e, 0xdf, 0x79, 0x47,
	0xb4, 0x09, 0xe7, 0xf1, 0xf7, 0x26, 0xb4, 0x02, 0x63, 0xdc, 0x84, 0xd9, 0x25, 0xc4, 0xce, 0xec,
	0xff, 0xfd, 0x48, 0xf6, 0x2d, 0x82, 0xc3, 0x6d, 0x93, 0xec, 0xda, 0xeb, 0x1c, 0x52, 0xf2, 0x85,
	0x68, 0x22, 0xdf, 0x10, 0xe5, 0xcb, 0xb6, 0x09, 0x74, 0x06, 0xfb, 0x37, 0xeb, 0xb2, 0x7c, 0x6b,
	0x54, 0xb8, 0xe2, 0x21, 0x36, 0x18, 0xaf, 0xb9, 0xe6, 0xa4, 0x54, 0xde, 0x51, 0x21, 0x36, 0xd8,
	0x5a, 0x31, 0x69, 0xd5, 0x3b, 0x67, 0x85, 0x78, 0xb6, 0x81, 0x18, 0x9b, 0x4d, 0xa0, 0xd7, 0x70,
	0x1c, 0x86, 0xff, 0x89, 0x15, 0x4b, 0x21, 0xee, 0xd0, 0x69, 0x77, 0xe5, 0xed, 0x57, 0x74, 0xd6,
	0x6c, 0xcf, 0xfd, 0x55, 0x07, 0x5f, 0x7f, 0xff, 0xf9, 0xde, 0x4b, 0xb2, 0x28, 0xbf, 0x88, 0x9e,
	0xa0, 0x09, 0xec, 0xdb, 0x34, 0x7e, 0x37, 0x47, 0x5b, 0xbc, 0x9d, 0x5b, 0x0f, 0x8a, 0xc4, 0xfe,
	0x79, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x2d, 0x87, 0x9e, 0x4e, 0x05, 0x00, 0x00,
}
