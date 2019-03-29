// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf-spec/mmlogic.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("api/protobuf-spec/mmlogic.proto", fileDescriptor_5b986081864e12b4) }

var fileDescriptor_5b986081864e12b4 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x40, 0xf7, 0xfb, 0x04, 0x85, 0x80, 0xa2, 0x61, 0xbd, 0xf4, 0xa2, 0xec, 0x7d, 0x1b, 0x51,
	0xc4, 0x83, 0x8a, 0xe8, 0x1e, 0x4a, 0x61, 0x17, 0x8b, 0x47, 0x6f, 0x69, 0x9d, 0x76, 0x23, 0x93,
	0x4e, 0x48, 0x26, 0x07, 0x7f, 0x9e, 0xff, 0x4c, 0xd2, 0x05, 0x8b, 0x50, 0x2f, 0x5e, 0x5f, 0xde,
	0x9b, 0x09, 0x89, 0x38, 0xd3, 0xce, 0x28, 0xe7, 0x89, 0xa9, 0x8e, 0xed, 0x32, 0x38, 0x68, 0x94,
	0xb5, 0x48, 0x9d, 0x69, 0xf2, 0x81, 0xca, 0x3d, 0xed, 0x4c, 0x76, 0x3e, 0x61, 0x41, 0x08, 0xba,
	0x83, 0xb0, 0xd3, 0x2e, 0x3f, 0xff, 0x8b, 0x83, 0x8d, 0x5d, 0xa7, 0x50, 0xde, 0x09, 0x51, 0x00,
	0x57, 0x9e, 0x5a, 0x83, 0x20, 0x4f, 0xf3, 0x6f, 0x75, 0xa3, 0xb9, 0xd9, 0x3e, 0xd7, 0xef, 0xd0,
	0x70, 0x36, 0x8d, 0x17, 0x33, 0x79, 0x2b, 0x8e, 0x56, 0x1e, 0x34, 0x43, 0xe5, 0xc9, 0x51, 0xd0,
	0xf8, 0xdb, 0x84, 0xe3, 0x11, 0xbf, 0x40, 0x88, 0x98, 0xe2, 0x07, 0x71, 0x98, 0x56, 0xa3, 0xfe,
	0x00, 0x5f, 0x11, 0xa1, 0x9c, 0x8f, 0xd2, 0x48, 0xb3, 0x49, 0xba, 0x98, 0x5d, 0xfc, 0x93, 0xf7,
	0x62, 0x5e, 0x00, 0x3f, 0x22, 0x96, 0x5d, 0x4f, 0x1e, 0xde, 0x76, 0xc7, 0x41, 0x9e, 0x8c, 0x45,
	0x89, 0x65, 0xef, 0xe2, 0xcf, 0xfd, 0x14, 0x18, 0xfc, 0x70, 0x79, 0xb9, 0x36, 0x81, 0xff, 0x14,
	0x3f, 0xdd, 0xbc, 0x5e, 0x77, 0x86, 0xb7, 0xb1, 0xce, 0x1b, 0xb2, 0xaa, 0x20, 0xea, 0x10, 0x56,
	0x48, 0x31, 0xcd, 0xe1, 0x96, 0xbc, 0x55, 0xe4, 0xa0, 0x5f, 0xda, 0xf4, 0x06, 0xca, 0xf4, 0x0c,
	0xbe, 0xd7, 0xa8, 0x5c, 0x5d, 0xef, 0x0f, 0x7f, 0x70, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0xe5, 0x29, 0xa6, 0xcd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MmLogicClient is the client API for MmLogic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MmLogicClient interface {
	//  Send GetProfile a match object with the ID field populated, it will return a
	//  'filled' one.
	//  Note: filters are assumed to have been checked for validity by the
	//  backendapi  when accepting a profile
	GetProfile(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*MatchObject, error)
	// CreateProposal is called by MMFs that wish to write their results to
	// a proposed MatchObject, that can be sent out the Backend API once it has
	// been approved (by default, by the evaluator process).
	//  - adds all players in all Rosters to the proposed player ignore list
	//  - writes the proposed match to the provided key
	//  - adds that key to the list of proposals to be considered
	// INPUT:
	//  * TO RETURN A MATCHOBJECT AFTER A SUCCESSFUL MMF RUN
	//    To create a match MatchObject message with these fields populated:
	//      - id, set to the value of the MMF_PROPOSAL_ID env var
	//      - properties
	//      - error.  You must explicitly set this to an empty string if your MMF
	//      - roster, with the playerIDs filled in the 'players' repeated field.
	//      - [optional] pools, set to the output from the 'GetPlayerPools' call,
	//        will populate the pools with stats about how many players the filters
	//        matched and how long the filters took to run, which will be sent out
	//        the backend api along with your match results.
	//        was successful.
	//  * TO RETURN AN ERROR
	//    To report a failure or error, send a MatchObject message with these
	//    these fields populated:
	//      - id, set to the value of the MMF_ERROR_ID env var.
	//      - error, set to a string value describing the error your MMF encountered.
	//      - [optional] properties, anything you put here is returned to the
	//        backend along with your error.
	//      - [optional] rosters, anything you put here is returned to the
	//        backend along with your error.
	//      - [optional] pools, set to the output from the 'GetPlayerPools' call,
	//        will populate the pools with stats about how many players the filters
	//        matched and how long the filters took to run, which will be sent out
	//        the backend api along with your match results.
	// OUTPUT: a Result message with a boolean success value and an error string
	// if an error was encountered
	CreateProposal(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error)
	// Player listing and filtering functions
	//
	// RetrievePlayerPool gets the list of players that match every Filter in the
	// PlayerPool, .excluding players in any configured ignore lists.  It
	// combines the results, and returns the resulting player pool.
	GetPlayerPool(ctx context.Context, in *PlayerPool, opts ...grpc.CallOption) (MmLogic_GetPlayerPoolClient, error)
	// Ignore List functions
	//
	// IlInput is an empty message reserved for future use.
	GetAllIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error)
	// ListIgnoredPlayers retrieves players from the ignore list specified in the
	// config file under 'ignoreLists.proposed.name'.
	ListIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error)
}

type mmLogicClient struct {
	cc *grpc.ClientConn
}

func NewMmLogicClient(cc *grpc.ClientConn) MmLogicClient {
	return &mmLogicClient{cc}
}

func (c *mmLogicClient) GetProfile(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*MatchObject, error) {
	out := new(MatchObject)
	err := c.cc.Invoke(ctx, "/api.MmLogic/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mmLogicClient) CreateProposal(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.MmLogic/CreateProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mmLogicClient) GetPlayerPool(ctx context.Context, in *PlayerPool, opts ...grpc.CallOption) (MmLogic_GetPlayerPoolClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MmLogic_serviceDesc.Streams[0], "/api.MmLogic/GetPlayerPool", opts...)
	if err != nil {
		return nil, err
	}
	x := &mmLogicGetPlayerPoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MmLogic_GetPlayerPoolClient interface {
	Recv() (*PlayerPool, error)
	grpc.ClientStream
}

type mmLogicGetPlayerPoolClient struct {
	grpc.ClientStream
}

func (x *mmLogicGetPlayerPoolClient) Recv() (*PlayerPool, error) {
	m := new(PlayerPool)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mmLogicClient) GetAllIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error) {
	out := new(Roster)
	err := c.cc.Invoke(ctx, "/api.MmLogic/GetAllIgnoredPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mmLogicClient) ListIgnoredPlayers(ctx context.Context, in *IlInput, opts ...grpc.CallOption) (*Roster, error) {
	out := new(Roster)
	err := c.cc.Invoke(ctx, "/api.MmLogic/ListIgnoredPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MmLogicServer is the server API for MmLogic service.
type MmLogicServer interface {
	//  Send GetProfile a match object with the ID field populated, it will return a
	//  'filled' one.
	//  Note: filters are assumed to have been checked for validity by the
	//  backendapi  when accepting a profile
	GetProfile(context.Context, *MatchObject) (*MatchObject, error)
	// CreateProposal is called by MMFs that wish to write their results to
	// a proposed MatchObject, that can be sent out the Backend API once it has
	// been approved (by default, by the evaluator process).
	//  - adds all players in all Rosters to the proposed player ignore list
	//  - writes the proposed match to the provided key
	//  - adds that key to the list of proposals to be considered
	// INPUT:
	//  * TO RETURN A MATCHOBJECT AFTER A SUCCESSFUL MMF RUN
	//    To create a match MatchObject message with these fields populated:
	//      - id, set to the value of the MMF_PROPOSAL_ID env var
	//      - properties
	//      - error.  You must explicitly set this to an empty string if your MMF
	//      - roster, with the playerIDs filled in the 'players' repeated field.
	//      - [optional] pools, set to the output from the 'GetPlayerPools' call,
	//        will populate the pools with stats about how many players the filters
	//        matched and how long the filters took to run, which will be sent out
	//        the backend api along with your match results.
	//        was successful.
	//  * TO RETURN AN ERROR
	//    To report a failure or error, send a MatchObject message with these
	//    these fields populated:
	//      - id, set to the value of the MMF_ERROR_ID env var.
	//      - error, set to a string value describing the error your MMF encountered.
	//      - [optional] properties, anything you put here is returned to the
	//        backend along with your error.
	//      - [optional] rosters, anything you put here is returned to the
	//        backend along with your error.
	//      - [optional] pools, set to the output from the 'GetPlayerPools' call,
	//        will populate the pools with stats about how many players the filters
	//        matched and how long the filters took to run, which will be sent out
	//        the backend api along with your match results.
	// OUTPUT: a Result message with a boolean success value and an error string
	// if an error was encountered
	CreateProposal(context.Context, *MatchObject) (*Result, error)
	// Player listing and filtering functions
	//
	// RetrievePlayerPool gets the list of players that match every Filter in the
	// PlayerPool, .excluding players in any configured ignore lists.  It
	// combines the results, and returns the resulting player pool.
	GetPlayerPool(*PlayerPool, MmLogic_GetPlayerPoolServer) error
	// Ignore List functions
	//
	// IlInput is an empty message reserved for future use.
	GetAllIgnoredPlayers(context.Context, *IlInput) (*Roster, error)
	// ListIgnoredPlayers retrieves players from the ignore list specified in the
	// config file under 'ignoreLists.proposed.name'.
	ListIgnoredPlayers(context.Context, *IlInput) (*Roster, error)
}

func RegisterMmLogicServer(s *grpc.Server, srv MmLogicServer) {
	s.RegisterService(&_MmLogic_serviceDesc, srv)
}

func _MmLogic_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmLogicServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MmLogic/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmLogicServer).GetProfile(ctx, req.(*MatchObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _MmLogic_CreateProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmLogicServer).CreateProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MmLogic/CreateProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmLogicServer).CreateProposal(ctx, req.(*MatchObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _MmLogic_GetPlayerPool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayerPool)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MmLogicServer).GetPlayerPool(m, &mmLogicGetPlayerPoolServer{stream})
}

type MmLogic_GetPlayerPoolServer interface {
	Send(*PlayerPool) error
	grpc.ServerStream
}

type mmLogicGetPlayerPoolServer struct {
	grpc.ServerStream
}

func (x *mmLogicGetPlayerPoolServer) Send(m *PlayerPool) error {
	return x.ServerStream.SendMsg(m)
}

func _MmLogic_GetAllIgnoredPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmLogicServer).GetAllIgnoredPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MmLogic/GetAllIgnoredPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmLogicServer).GetAllIgnoredPlayers(ctx, req.(*IlInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _MmLogic_ListIgnoredPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IlInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmLogicServer).ListIgnoredPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MmLogic/ListIgnoredPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmLogicServer).ListIgnoredPlayers(ctx, req.(*IlInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _MmLogic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MmLogic",
	HandlerType: (*MmLogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _MmLogic_GetProfile_Handler,
		},
		{
			MethodName: "CreateProposal",
			Handler:    _MmLogic_CreateProposal_Handler,
		},
		{
			MethodName: "GetAllIgnoredPlayers",
			Handler:    _MmLogic_GetAllIgnoredPlayers_Handler,
		},
		{
			MethodName: "ListIgnoredPlayers",
			Handler:    _MmLogic_ListIgnoredPlayers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPlayerPool",
			Handler:       _MmLogic_GetPlayerPool_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/protobuf-spec/mmlogic.proto",
}
