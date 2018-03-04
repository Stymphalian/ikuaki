package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

// Return a string in the format
// rpc CreateWorld(CreateWorldReq) returns (CreateWorldRes) {}
func GetMethodString(m *desc.MethodDescriptor, multiLine bool) string {
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

func ReadTextprotoFromStdin(desc *desc.MessageDescriptor) (*dynamic.Message, error) {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	dmsg := dynamic.NewMessage(desc)
	err = dmsg.UnmarshalText(all)
	if err != nil {
		return nil, err
	}
	return dmsg, nil
}

// Read protos from stdin, they should be separate by two newlines.
func ReadTextprotosFromStdin(desc *desc.MessageDescriptor) ([]*dynamic.Message, error) {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	protoStrs := strings.Split(string(all), "\n")
	msgs := make([]*dynamic.Message, 0)
	for _, s := range protoStrs {
		dmsg := dynamic.NewMessage(desc)
		err = dmsg.UnmarshalText([]byte(s))
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, dmsg)
	}
	return msgs, nil
}

func BuildPlugin(filepath string) error {
	// Build the .so file
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(originalDir)

	fmt.Printf("Building '%s' ...\n", filepath)
	os.Chdir(path.Dir(filepath))
	cmd := exec.Command("go", "build", "-buildmode=plugin")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("Done building '%s' ...\n", filepath)
	return nil
}
