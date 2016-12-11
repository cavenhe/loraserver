// Code generated by protoc-gen-go.
// source: as.proto
// DO NOT EDIT!

/*
Package as is a generated protocol buffer package.

It is generated from these files:
	as.proto

It has these top-level messages:
	DataRate
	RXInfo
	TXInfo
	JoinRequestRequest
	JoinRequestResponse
	HandleDataUpRequest
	GetDataDownRequest
	GetDataDownResponse
	HandleDataUpResponse
	HandleDataDownACKRequest
	HandleDataDownACKResponse
	HandleErrorRequest
	HandleErrorResponse
*/
package as

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorType int32

const (
	ErrorType_Generic      ErrorType = 0
	ErrorType_OTAA         ErrorType = 1
	ErrorType_DATA_UP_FCNT ErrorType = 2
	ErrorType_DATA_UP_MIC  ErrorType = 3
)

var ErrorType_name = map[int32]string{
	0: "Generic",
	1: "OTAA",
	2: "DATA_UP_FCNT",
	3: "DATA_UP_MIC",
}
var ErrorType_value = map[string]int32{
	"Generic":      0,
	"OTAA":         1,
	"DATA_UP_FCNT": 2,
	"DATA_UP_MIC":  3,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DataRate struct {
	Modulation   string `protobuf:"bytes,1,opt,name=modulation" json:"modulation,omitempty"`
	BandWidth    uint32 `protobuf:"varint,2,opt,name=bandWidth" json:"bandWidth,omitempty"`
	SpreadFactor uint32 `protobuf:"varint,3,opt,name=spreadFactor" json:"spreadFactor,omitempty"`
	Bitrate      uint32 `protobuf:"varint,4,opt,name=bitrate" json:"bitrate,omitempty"`
}

func (m *DataRate) Reset()                    { *m = DataRate{} }
func (m *DataRate) String() string            { return proto.CompactTextString(m) }
func (*DataRate) ProtoMessage()               {}
func (*DataRate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DataRate) GetModulation() string {
	if m != nil {
		return m.Modulation
	}
	return ""
}

func (m *DataRate) GetBandWidth() uint32 {
	if m != nil {
		return m.BandWidth
	}
	return 0
}

func (m *DataRate) GetSpreadFactor() uint32 {
	if m != nil {
		return m.SpreadFactor
	}
	return 0
}

func (m *DataRate) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

type RXInfo struct {
	Mac     []byte  `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Time    string  `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Rssi    int32   `protobuf:"varint,3,opt,name=rssi" json:"rssi,omitempty"`
	LoRaSNR float64 `protobuf:"fixed64,4,opt,name=loRaSNR" json:"loRaSNR,omitempty"`
}

func (m *RXInfo) Reset()                    { *m = RXInfo{} }
func (m *RXInfo) String() string            { return proto.CompactTextString(m) }
func (*RXInfo) ProtoMessage()               {}
func (*RXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RXInfo) GetMac() []byte {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *RXInfo) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *RXInfo) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *RXInfo) GetLoRaSNR() float64 {
	if m != nil {
		return m.LoRaSNR
	}
	return 0
}

type TXInfo struct {
	Frequency int64     `protobuf:"varint,1,opt,name=frequency" json:"frequency,omitempty"`
	DataRate  *DataRate `protobuf:"bytes,2,opt,name=dataRate" json:"dataRate,omitempty"`
	Adr       bool      `protobuf:"varint,3,opt,name=adr" json:"adr,omitempty"`
	CodeRate  string    `protobuf:"bytes,4,opt,name=codeRate" json:"codeRate,omitempty"`
}

func (m *TXInfo) Reset()                    { *m = TXInfo{} }
func (m *TXInfo) String() string            { return proto.CompactTextString(m) }
func (*TXInfo) ProtoMessage()               {}
func (*TXInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TXInfo) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *TXInfo) GetDataRate() *DataRate {
	if m != nil {
		return m.DataRate
	}
	return nil
}

func (m *TXInfo) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *TXInfo) GetCodeRate() string {
	if m != nil {
		return m.CodeRate
	}
	return ""
}

type JoinRequestRequest struct {
	PhyPayload []byte `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	DevAddr    []byte `protobuf:"bytes,2,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	NetID      []byte `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
}

func (m *JoinRequestRequest) Reset()                    { *m = JoinRequestRequest{} }
func (m *JoinRequestRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestRequest) ProtoMessage()               {}
func (*JoinRequestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JoinRequestRequest) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *JoinRequestRequest) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *JoinRequestRequest) GetNetID() []byte {
	if m != nil {
		return m.NetID
	}
	return nil
}

type JoinRequestResponse struct {
	PhyPayload         []byte   `protobuf:"bytes,1,opt,name=phyPayload,proto3" json:"phyPayload,omitempty"`
	NwkSKey            []byte   `protobuf:"bytes,2,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	RxDelay            uint32   `protobuf:"varint,3,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset        uint32   `protobuf:"varint,4,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList             []uint32 `protobuf:"varint,5,rep,packed,name=cFList" json:"cFList,omitempty"`
	RxWindow           RXWindow `protobuf:"varint,6,opt,name=rxWindow,enum=as.RXWindow" json:"rxWindow,omitempty"`
	Rx2DR              uint32   `protobuf:"varint,7,opt,name=rx2DR" json:"rx2DR,omitempty"`
	RelaxFCnt          bool     `protobuf:"varint,8,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	AdrInterval        uint32   `protobuf:"varint,9,opt,name=adrInterval" json:"adrInterval,omitempty"`
	InstallationMargin float64  `protobuf:"fixed64,10,opt,name=installationMargin" json:"installationMargin,omitempty"`
}

func (m *JoinRequestResponse) Reset()                    { *m = JoinRequestResponse{} }
func (m *JoinRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinRequestResponse) ProtoMessage()               {}
func (*JoinRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *JoinRequestResponse) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *JoinRequestResponse) GetNwkSKey() []byte {
	if m != nil {
		return m.NwkSKey
	}
	return nil
}

func (m *JoinRequestResponse) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *JoinRequestResponse) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *JoinRequestResponse) GetCFList() []uint32 {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *JoinRequestResponse) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *JoinRequestResponse) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *JoinRequestResponse) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *JoinRequestResponse) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *JoinRequestResponse) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

type HandleDataUpRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte    `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	FCnt   uint32    `protobuf:"varint,3,opt,name=fCnt" json:"fCnt,omitempty"`
	FPort  uint32    `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	Data   []byte    `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	TxInfo *TXInfo   `protobuf:"bytes,6,opt,name=txInfo" json:"txInfo,omitempty"`
	RxInfo []*RXInfo `protobuf:"bytes,7,rep,name=rxInfo" json:"rxInfo,omitempty"`
}

func (m *HandleDataUpRequest) Reset()                    { *m = HandleDataUpRequest{} }
func (m *HandleDataUpRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleDataUpRequest) ProtoMessage()               {}
func (*HandleDataUpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HandleDataUpRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *HandleDataUpRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *HandleDataUpRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *HandleDataUpRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *HandleDataUpRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *HandleDataUpRequest) GetTxInfo() *TXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleDataUpRequest) GetRxInfo() []*RXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type GetDataDownRequest struct {
	DevEUI         []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI         []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	MaxPayloadSize uint32 `protobuf:"varint,3,opt,name=maxPayloadSize" json:"maxPayloadSize,omitempty"`
	FCnt           uint32 `protobuf:"varint,4,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *GetDataDownRequest) Reset()                    { *m = GetDataDownRequest{} }
func (m *GetDataDownRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDataDownRequest) ProtoMessage()               {}
func (*GetDataDownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetDataDownRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *GetDataDownRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *GetDataDownRequest) GetMaxPayloadSize() uint32 {
	if m != nil {
		return m.MaxPayloadSize
	}
	return 0
}

func (m *GetDataDownRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type GetDataDownResponse struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Confirmed bool   `protobuf:"varint,2,opt,name=confirmed" json:"confirmed,omitempty"`
	FPort     uint32 `protobuf:"varint,3,opt,name=fPort" json:"fPort,omitempty"`
	MoreData  bool   `protobuf:"varint,4,opt,name=moreData" json:"moreData,omitempty"`
}

func (m *GetDataDownResponse) Reset()                    { *m = GetDataDownResponse{} }
func (m *GetDataDownResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDataDownResponse) ProtoMessage()               {}
func (*GetDataDownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetDataDownResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetDataDownResponse) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *GetDataDownResponse) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *GetDataDownResponse) GetMoreData() bool {
	if m != nil {
		return m.MoreData
	}
	return false
}

type HandleDataUpResponse struct {
}

func (m *HandleDataUpResponse) Reset()                    { *m = HandleDataUpResponse{} }
func (m *HandleDataUpResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleDataUpResponse) ProtoMessage()               {}
func (*HandleDataUpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type HandleDataDownACKRequest struct {
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	FCnt   uint32 `protobuf:"varint,3,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *HandleDataDownACKRequest) Reset()                    { *m = HandleDataDownACKRequest{} }
func (m *HandleDataDownACKRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleDataDownACKRequest) ProtoMessage()               {}
func (*HandleDataDownACKRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *HandleDataDownACKRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *HandleDataDownACKRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *HandleDataDownACKRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type HandleDataDownACKResponse struct {
}

func (m *HandleDataDownACKResponse) Reset()                    { *m = HandleDataDownACKResponse{} }
func (m *HandleDataDownACKResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleDataDownACKResponse) ProtoMessage()               {}
func (*HandleDataDownACKResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type HandleErrorRequest struct {
	DevEUI []byte    `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	AppEUI []byte    `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	Type   ErrorType `protobuf:"varint,3,opt,name=type,enum=as.ErrorType" json:"type,omitempty"`
	Error  string    `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *HandleErrorRequest) Reset()                    { *m = HandleErrorRequest{} }
func (m *HandleErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*HandleErrorRequest) ProtoMessage()               {}
func (*HandleErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *HandleErrorRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *HandleErrorRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *HandleErrorRequest) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_Generic
}

func (m *HandleErrorRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type HandleErrorResponse struct {
}

func (m *HandleErrorResponse) Reset()                    { *m = HandleErrorResponse{} }
func (m *HandleErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*HandleErrorResponse) ProtoMessage()               {}
func (*HandleErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*DataRate)(nil), "as.DataRate")
	proto.RegisterType((*RXInfo)(nil), "as.RXInfo")
	proto.RegisterType((*TXInfo)(nil), "as.TXInfo")
	proto.RegisterType((*JoinRequestRequest)(nil), "as.JoinRequestRequest")
	proto.RegisterType((*JoinRequestResponse)(nil), "as.JoinRequestResponse")
	proto.RegisterType((*HandleDataUpRequest)(nil), "as.HandleDataUpRequest")
	proto.RegisterType((*GetDataDownRequest)(nil), "as.GetDataDownRequest")
	proto.RegisterType((*GetDataDownResponse)(nil), "as.GetDataDownResponse")
	proto.RegisterType((*HandleDataUpResponse)(nil), "as.HandleDataUpResponse")
	proto.RegisterType((*HandleDataDownACKRequest)(nil), "as.HandleDataDownACKRequest")
	proto.RegisterType((*HandleDataDownACKResponse)(nil), "as.HandleDataDownACKResponse")
	proto.RegisterType((*HandleErrorRequest)(nil), "as.HandleErrorRequest")
	proto.RegisterType((*HandleErrorResponse)(nil), "as.HandleErrorResponse")
	proto.RegisterEnum("as.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterEnum("as.ErrorType", ErrorType_name, ErrorType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApplicationServer service

type ApplicationServerClient interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error)
	// HandleDataUp publishes data received from an end-device.
	HandleDataUp(ctx context.Context, in *HandleDataUpRequest, opts ...grpc.CallOption) (*HandleDataUpResponse, error)
	// GetDataDown gets data from the downlink queue.
	GetDataDown(ctx context.Context, in *GetDataDownRequest, opts ...grpc.CallOption) (*GetDataDownResponse, error)
	// HandleDataDownACK publishes a data-down ack response.
	HandleDataDownACK(ctx context.Context, in *HandleDataDownACKRequest, opts ...grpc.CallOption) (*HandleDataDownACKResponse, error)
	// HandleError publishes an error message.
	HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*HandleErrorResponse, error)
}

type applicationServerClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServerClient(cc *grpc.ClientConn) ApplicationServerClient {
	return &applicationServerClient{cc}
}

func (c *applicationServerClient) JoinRequest(ctx context.Context, in *JoinRequestRequest, opts ...grpc.CallOption) (*JoinRequestResponse, error) {
	out := new(JoinRequestResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/JoinRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleDataUp(ctx context.Context, in *HandleDataUpRequest, opts ...grpc.CallOption) (*HandleDataUpResponse, error) {
	out := new(HandleDataUpResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleDataUp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) GetDataDown(ctx context.Context, in *GetDataDownRequest, opts ...grpc.CallOption) (*GetDataDownResponse, error) {
	out := new(GetDataDownResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/GetDataDown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleDataDownACK(ctx context.Context, in *HandleDataDownACKRequest, opts ...grpc.CallOption) (*HandleDataDownACKResponse, error) {
	out := new(HandleDataDownACKResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleDataDownACK", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerClient) HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*HandleErrorResponse, error) {
	out := new(HandleErrorResponse)
	err := grpc.Invoke(ctx, "/as.ApplicationServer/HandleError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApplicationServer service

type ApplicationServerServer interface {
	// JoinRequest requests the application server to validate the join-request and return an encrypted join-accept.
	JoinRequest(context.Context, *JoinRequestRequest) (*JoinRequestResponse, error)
	// HandleDataUp publishes data received from an end-device.
	HandleDataUp(context.Context, *HandleDataUpRequest) (*HandleDataUpResponse, error)
	// GetDataDown gets data from the downlink queue.
	GetDataDown(context.Context, *GetDataDownRequest) (*GetDataDownResponse, error)
	// HandleDataDownACK publishes a data-down ack response.
	HandleDataDownACK(context.Context, *HandleDataDownACKRequest) (*HandleDataDownACKResponse, error)
	// HandleError publishes an error message.
	HandleError(context.Context, *HandleErrorRequest) (*HandleErrorResponse, error)
}

func RegisterApplicationServerServer(s *grpc.Server, srv ApplicationServerServer) {
	s.RegisterService(&_ApplicationServer_serviceDesc, srv)
}

func _ApplicationServer_JoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).JoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/JoinRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).JoinRequest(ctx, req.(*JoinRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleDataUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDataUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleDataUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleDataUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleDataUp(ctx, req.(*HandleDataUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_GetDataDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).GetDataDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/GetDataDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).GetDataDown(ctx, req.(*GetDataDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleDataDownACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDataDownACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleDataDownACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleDataDownACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleDataDownACK(ctx, req.(*HandleDataDownACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServer_HandleError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServer).HandleError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServer/HandleError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServer).HandleError(ctx, req.(*HandleErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "as.ApplicationServer",
	HandlerType: (*ApplicationServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinRequest",
			Handler:    _ApplicationServer_JoinRequest_Handler,
		},
		{
			MethodName: "HandleDataUp",
			Handler:    _ApplicationServer_HandleDataUp_Handler,
		},
		{
			MethodName: "GetDataDown",
			Handler:    _ApplicationServer_GetDataDown_Handler,
		},
		{
			MethodName: "HandleDataDownACK",
			Handler:    _ApplicationServer_HandleDataDownACK_Handler,
		},
		{
			MethodName: "HandleError",
			Handler:    _ApplicationServer_HandleError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as.proto",
}

func init() { proto.RegisterFile("as.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xce, 0xda, 0x8e, 0xbd, 0x3e, 0x76, 0xc2, 0x76, 0x52, 0xd2, 0xc5, 0x18, 0x64, 0xf6, 0x02,
	0x59, 0xbd, 0x88, 0x54, 0xf3, 0x02, 0xb5, 0xec, 0xa4, 0x98, 0xd2, 0x36, 0x9a, 0x38, 0x6a, 0x2e,
	0x10, 0xd5, 0x64, 0x67, 0x4c, 0x57, 0xac, 0x67, 0x96, 0xd9, 0x69, 0xe2, 0x45, 0x02, 0x71, 0xc5,
	0xcb, 0x71, 0xc9, 0x9b, 0xf0, 0x04, 0x68, 0x7e, 0xd6, 0x5e, 0x63, 0x23, 0xa1, 0x88, 0xab, 0xcc,
	0xf9, 0xce, 0xe4, 0x7c, 0xe7, 0x7c, 0xe7, 0x9b, 0x35, 0xf8, 0x24, 0x3f, 0xcb, 0xa4, 0x50, 0x02,
	0xd5, 0x48, 0x1e, 0xfd, 0xee, 0x81, 0x3f, 0x25, 0x8a, 0x60, 0xa2, 0x18, 0xfa, 0x1c, 0x60, 0x29,
	0xe8, 0x87, 0x94, 0xa8, 0x44, 0xf0, 0xd0, 0x1b, 0x78, 0xc3, 0x36, 0xae, 0x20, 0xa8, 0x0f, 0xed,
	0x5b, 0xc2, 0xe9, 0xdb, 0x84, 0xaa, 0xf7, 0x61, 0x6d, 0xe0, 0x0d, 0x8f, 0xf0, 0x06, 0x40, 0x11,
	0x74, 0xf3, 0x4c, 0x32, 0x42, 0x2f, 0x48, 0xac, 0x84, 0x0c, 0xeb, 0xe6, 0xc2, 0x16, 0x86, 0x42,
	0x68, 0xdd, 0x26, 0x4a, 0x12, 0xc5, 0xc2, 0x86, 0x49, 0x97, 0x61, 0xf4, 0x1d, 0x34, 0xf1, 0xcd,
	0x8c, 0x2f, 0x04, 0x0a, 0xa0, 0xbe, 0x24, 0xb1, 0xa1, 0xef, 0x62, 0x7d, 0x44, 0x08, 0x1a, 0x2a,
	0x59, 0x32, 0x43, 0xd9, 0xc6, 0xe6, 0xac, 0x31, 0x99, 0xe7, 0x89, 0x61, 0x39, 0xc4, 0xe6, 0xac,
	0xab, 0xa7, 0x02, 0x93, 0xab, 0xd7, 0xd8, 0x54, 0xf7, 0x70, 0x19, 0x46, 0xbf, 0x42, 0x73, 0x6e,
	0xab, 0xf7, 0xa1, 0xbd, 0x90, 0xec, 0xa7, 0x0f, 0x8c, 0xc7, 0x85, 0xe1, 0xa8, 0xe3, 0x0d, 0x80,
	0x86, 0xe0, 0x53, 0xa7, 0x86, 0x61, 0xeb, 0x8c, 0xba, 0x67, 0x24, 0x3f, 0x2b, 0x15, 0xc2, 0xeb,
	0xac, 0xee, 0x92, 0x50, 0x3b, 0xa4, 0x8f, 0xf5, 0x11, 0xf5, 0xc0, 0x8f, 0x05, 0x65, 0xb8, 0x1c,
	0xae, 0x8d, 0xd7, 0x71, 0x44, 0x01, 0x7d, 0x23, 0x12, 0x8e, 0x35, 0x4f, 0xae, 0xdc, 0x1f, 0xad,
	0x77, 0xf6, 0xbe, 0xb8, 0x24, 0x45, 0x2a, 0x08, 0x75, 0x03, 0x57, 0x10, 0x3d, 0x0f, 0x65, 0x77,
	0x63, 0x4a, 0xa5, 0x69, 0xa6, 0x8b, 0xcb, 0x10, 0x3d, 0x86, 0x43, 0xce, 0xd4, 0x6c, 0x6a, 0xf8,
	0xbb, 0xd8, 0x06, 0xd1, 0x9f, 0x35, 0x38, 0xd9, 0xa2, 0xc9, 0x33, 0xc1, 0x73, 0xf6, 0x5f, 0x78,
	0xf8, 0xfd, 0x8f, 0x57, 0x2f, 0x59, 0x51, 0xf2, 0xb8, 0x50, 0x67, 0xe4, 0x6a, 0xca, 0x52, 0x52,
	0xb8, 0x75, 0x96, 0x21, 0x1a, 0x40, 0x47, 0xae, 0x9e, 0x4d, 0xf1, 0x9b, 0xc5, 0x22, 0x67, 0xca,
	0x6d, 0xb3, 0x0a, 0xa1, 0x53, 0x68, 0xc6, 0x17, 0xdf, 0x26, 0xb9, 0x0a, 0x0f, 0x07, 0xf5, 0xe1,
	0x11, 0x76, 0x91, 0xd6, 0x58, 0xae, 0xde, 0x26, 0x9c, 0x8a, 0xfb, 0xb0, 0x39, 0xf0, 0x86, 0xc7,
	0x56, 0x63, 0x7c, 0x63, 0x31, 0xbc, 0xce, 0xea, 0x29, 0xe5, 0x6a, 0x34, 0xc5, 0x61, 0xcb, 0x54,
	0xb7, 0x81, 0xde, 0xa0, 0x64, 0x29, 0x59, 0x5d, 0x4c, 0xb8, 0x0a, 0x7d, 0xa3, 0xff, 0x06, 0xd0,
	0x7d, 0x11, 0x2a, 0x67, 0x5c, 0x31, 0x79, 0x47, 0xd2, 0xb0, 0x6d, 0xfb, 0xaa, 0x40, 0xe8, 0x0c,
	0x50, 0xc2, 0x73, 0x45, 0x52, 0xeb, 0xea, 0x57, 0x44, 0xfe, 0x90, 0xf0, 0x10, 0x8c, 0x61, 0xf6,
	0x64, 0xa2, 0x3f, 0x3c, 0x38, 0xf9, 0x9a, 0x70, 0x9a, 0x32, 0x6d, 0x83, 0xeb, 0xac, 0xdc, 0xde,
	0x29, 0x34, 0x29, 0xbb, 0x3b, 0xbf, 0x9e, 0x39, 0x45, 0x5d, 0xa4, 0x71, 0x92, 0x65, 0x1a, 0xb7,
	0x62, 0xba, 0x48, 0x3b, 0x76, 0xa1, 0x5b, 0xb6, 0x42, 0x9a, 0xb3, 0x9e, 0x70, 0x71, 0x29, 0x64,
	0xa9, 0x9f, 0x0d, 0xf4, 0x4d, 0xed, 0xb3, 0xf0, 0xd0, 0xfc, 0xbf, 0x39, 0xa3, 0x08, 0x9a, 0x6a,
	0xa5, 0x1d, 0x6c, 0x34, 0xeb, 0x8c, 0x40, 0x6b, 0x66, 0x3d, 0x8d, 0x5d, 0x46, 0xdf, 0x91, 0xf6,
	0x4e, 0x6b, 0x50, 0x2f, 0xef, 0x60, 0x77, 0xc7, 0x66, 0xa2, 0xdf, 0x3c, 0x40, 0x2f, 0x98, 0xd2,
	0xa3, 0x4c, 0xc5, 0x3d, 0x7f, 0xe8, 0x30, 0x5f, 0xc2, 0xf1, 0x92, 0xac, 0x9c, 0x81, 0xae, 0x92,
	0x9f, 0x99, 0x1b, 0xeb, 0x1f, 0xe8, 0x7a, 0xe8, 0xc6, 0x66, 0xe8, 0xa8, 0x80, 0x93, 0xad, 0x0e,
	0x9c, 0x4b, 0xcb, 0xa9, 0xbd, 0xca, 0xd4, 0x7d, 0x68, 0xc7, 0x82, 0x2f, 0x12, 0xb9, 0x64, 0xd4,
	0x74, 0xe0, 0xe3, 0x0d, 0xb0, 0x51, 0xaf, 0x5e, 0x55, 0xaf, 0x07, 0xfe, 0x52, 0x48, 0xb3, 0x2c,
	0x43, 0xeb, 0xe3, 0x75, 0x1c, 0x9d, 0xc2, 0xe3, 0xed, 0x55, 0x5a, 0xee, 0xe8, 0x7b, 0x08, 0x37,
	0xb8, 0xee, 0x6a, 0x3c, 0x79, 0xf9, 0x3f, 0xee, 0x39, 0xfa, 0x14, 0x3e, 0xd9, 0x53, 0xdf, 0x91,
	0xff, 0x02, 0xc8, 0x26, 0xcf, 0xa5, 0x14, 0xf2, 0xa1, 0xb4, 0x5f, 0x40, 0x43, 0x15, 0x99, 0xdd,
	0xc3, 0xf1, 0xe8, 0x48, 0xaf, 0xde, 0xd4, 0x9b, 0x17, 0x19, 0xc3, 0x26, 0xa5, 0xf5, 0x62, 0x1a,
	0x72, 0x9f, 0x27, 0x1b, 0x44, 0x1f, 0x97, 0xf6, 0x76, 0xf4, 0xb6, 0xab, 0xa7, 0x7d, 0xf0, 0xcb,
	0x27, 0x89, 0x5a, 0x50, 0xc7, 0x37, 0xcf, 0x82, 0x03, 0x7b, 0x18, 0x05, 0xde, 0xd3, 0x73, 0x68,
	0xaf, 0xab, 0xa3, 0x0e, 0xb4, 0x5e, 0x30, 0xce, 0x64, 0x12, 0x07, 0x07, 0xc8, 0x87, 0xc6, 0x9b,
	0xf9, 0x78, 0x1c, 0x78, 0x28, 0x80, 0xee, 0x74, 0x3c, 0x1f, 0xbf, 0xbb, 0xbe, 0x7c, 0x77, 0x31,
	0x79, 0x3d, 0x0f, 0x6a, 0xe8, 0x23, 0xe8, 0x94, 0xc8, 0xab, 0xd9, 0x24, 0xa8, 0x8f, 0xfe, 0xaa,
	0xc1, 0xa3, 0x71, 0x96, 0xa5, 0x49, 0x6c, 0x5e, 0xdc, 0x15, 0x93, 0x77, 0x4c, 0xa2, 0xe7, 0xd0,
	0xa9, 0x7c, 0xc6, 0xd0, 0xa9, 0x9e, 0x65, 0xf7, 0xf3, 0xd9, 0x7b, 0xb2, 0x83, 0x3b, 0x41, 0x0f,
	0xd0, 0x04, 0xba, 0xd5, 0x3d, 0x23, 0x73, 0x75, 0xcf, 0x23, 0xee, 0x85, 0xbb, 0x89, 0x75, 0x91,
	0xe7, 0xd0, 0xa9, 0xf8, 0xd4, 0xb6, 0xb1, 0xfb, 0x74, 0x6c, 0x1b, 0x7b, 0x0c, 0x1d, 0x1d, 0x20,
	0x0c, 0x8f, 0x76, 0xd6, 0x8e, 0xfa, 0xdb, 0x94, 0xdb, 0x6e, 0xeb, 0x7d, 0xf6, 0x2f, 0xd9, 0x6a,
	0x57, 0x95, 0x75, 0xd9, 0xae, 0x76, 0xed, 0xd3, 0x7b, 0xb2, 0x83, 0x97, 0x15, 0x6e, 0x9b, 0xe6,
	0xe7, 0xff, 0xab, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x8a, 0xb5, 0xa6, 0x0a, 0x08, 0x00,
	0x00,
}