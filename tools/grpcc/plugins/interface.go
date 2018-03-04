package grpcc

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type StreamTypeEnum int

const (
	K_UNARY           StreamTypeEnum = iota
	K_STREAM_REQUEST  StreamTypeEnum = iota
	K_STREAM_RESPONSE StreamTypeEnum = iota
	K_STREAM_BIDI     StreamTypeEnum = iota
)

type Info struct {
	// The original hostport, servicename and method name
	Hostport    string
	ServiceName string
	MethodName  string

	// Connection to the server
	Conn   *grpc.ClientConn
	Client *grpcreflect.Client

	// Descriptors for the service, method and request and responses
	Service      *desc.ServiceDescriptor
	MethodDesc   *desc.MethodDescriptor
	RequestDesc  *desc.MessageDescriptor
	ResponseDesc *desc.MessageDescriptor

	// The stub where you can dynamically make calls
	Stub grpcdynamic.Stub

	// Useful enum to help determine if it is Unary, streaming RPC
	StreamType StreamTypeEnum
}

func NewInfo(hostport string, serviceName string, methodName string, conn *grpc.ClientConn) (*Info, error) {
	// dial to get the service
	conn, err := grpc.Dial(hostport, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// Get grpcreflect client and resolve to the service
	client := grpcreflect.NewClient(context.Background(),
		rpb.NewServerReflectionClient(conn))
	service, err := client.ResolveService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("Couldn't find service %s at %s\n%s", serviceName,
			hostport, err)
	}

	// get the method descriptor
	method := service.FindMethodByName(methodName)
	if method == nil {
		return nil, fmt.Errorf("Couldn't find %s/%s on %s\n", serviceName,
			methodName, hostport)
	}
	inputDesc := method.GetInputType()
	outputDesc := method.GetOutputType()

	streamType := K_UNARY
	if method.IsClientStreaming() && method.IsServerStreaming() {
		streamType = K_STREAM_BIDI
	} else if method.IsClientStreaming() {
		streamType = K_STREAM_REQUEST
	} else if method.IsServerStreaming() {
		streamType = K_STREAM_RESPONSE
	}

	stub := grpcdynamic.NewStub(conn)

	return &Info{
		Hostport:    hostport,
		ServiceName: serviceName,
		MethodName:  methodName,

		Conn:   conn,
		Client: client,

		Service:      service,
		MethodDesc:   method,
		RequestDesc:  inputDesc,
		ResponseDesc: outputDesc,

		Stub:       stub,
		StreamType: streamType,
	}, nil
}

// When the plugin is loaded we call a function with the signature:
// Run(info* Info) error
