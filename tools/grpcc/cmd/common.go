package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

func splitHostServiceMethod(s string) (
	hostport string, service string, method string) {
	ss := strings.Split(s, "/")
	if len(ss) == 1 {
		return ss[0], "", ""
	} else if len(ss) == 2 {
		return ss[0], ss[1], ""
	} else if len(ss) == 3 {
		return ss[0], ss[1], ss[2]
	} else {
		log.Fatal("Unsupported argument ", s)
	}
	return "", "", ""
}

// Return a string in the format
// rpc CreateWorld(CreateWorldReq) returns (CreateWorldRes) {}
func getMethodString(m *desc.MethodDescriptor, multiLine bool) string {
	clientStream := ""
	if m.IsClientStreaming() {
		clientStream = "stream "
	}
	serverStream := ""
	if m.IsServerStreaming() {
		serverStream = "stream "
	}
	if multiLine {
		spacing := ""
		llen := len(fmt.Sprintf("rpc %s", m.GetName()))
		for i := 0; i < llen; i++ {
			spacing += " "
		}

		return fmt.Sprintf("rpc %s(%s%s) returns\n%s(%s%s)\n",
			m.GetName(), clientStream, m.GetInputType().GetFullyQualifiedName(),
			spacing,
			serverStream, m.GetOutputType().GetFullyQualifiedName())

	} else {
		return fmt.Sprintf("rpc %s(%s%s) returns (%s%s)\n",
			m.GetName(), clientStream, m.GetInputType().GetFullyQualifiedName(),
			serverStream, m.GetOutputType().GetFullyQualifiedName())
	}
}

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
