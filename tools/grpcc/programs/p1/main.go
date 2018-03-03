package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	grpcc "github.com/Stymphalian/ikuaki/tools/grpcc/plugin"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

func ReadTextprotoFromStdin(in *desc.MessageDescriptor) (*dynamic.Message, error) {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	dmsg := dynamic.NewMessage(in)
	err = dmsg.UnmarshalText(all)
	if err != nil {
		return nil, err
	}
	return dmsg, nil
}

func Run(info *grpcc.Info) error {
	req, err := ReadTextprotoFromStdin(info.RequestDesc)
	if err != nil {
		return fmt.Errorf("Failed to parse textproto %s", err)
	}

	// make the request
	resp, err := info.Stub.InvokeRpc(
		context.Background(), info.MethodDesc, req)
	if err != nil {
		return err
	}

	// show the response
	fmt.Println()
	fmt.Println(resp)
	return nil
}

func main() {}
