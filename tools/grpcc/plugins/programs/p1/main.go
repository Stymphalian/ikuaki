package main

import (
	"context"
	"fmt"
	"io"
	"time"

	grpcc "github.com/Stymphalian/ikuaki/tools/grpcc/plugins"
	pb "github.com/Stymphalian/ikuaki/tools/grpcc/testing/server"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
)

type State struct {
	info   *grpcc.Info
	stream *grpcdynamic.BidiStream
}

func (this *State) ClientStreamHandler(done chan error) {
	for i := 0; i < 10; i++ {
		req := &pb.CreateWorldReq{Name: fmt.Sprintf("jordan-%d", i)}
		err := this.stream.SendMsg(req)
		if err != nil {
			done <- err
			return
		}

		select {
		case <-context.Background().Done():
			done <- context.Background().Err()
			return
		default:
			time.Sleep(time.Second * 2)
		}
	}
	done <- this.stream.CloseSend()
	return
}

func (this *State) ServerStreamHandler(done chan error) {
	for {
		select {
		case <-context.Background().Done():
			done <- context.Background().Err()
			return
		default:
		}

		res, err := this.stream.RecvMsg()
		if err == io.EOF {
			done <- nil
			return
		}
		if err != nil {
			done <- err
			return
		}
		fmt.Println(res)
	}
}

func Run(info *grpcc.Info) error {
	stream, err := info.Stub.InvokeRpcBidiStream(context.Background(), info.MethodDesc)
	if err != nil {
		return err
	}
	state := &State{info, stream}

	clientDone := make(chan error)
	serverDone := make(chan error)
	go state.ClientStreamHandler(clientDone)
	go state.ServerStreamHandler(serverDone)

	for i := 0; i < 2; i++ {
		var e error
		select {
		case e = <-clientDone:
		case e = <-serverDone:
		}
		if e != nil {
			return e
		}
	}
	return nil
}

func main() {}
