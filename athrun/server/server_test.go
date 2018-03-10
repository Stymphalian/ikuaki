package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	pb "github.com/Stymphalian/ikuaki/athrun/protos"
)

var (
	TEMP_DIR string
)

func newTestServerOrDie(t *testing.T) *Server {
	addr := "localhost:8080"
	gopath, ok := os.LookupEnv("GOPATH")
	if !ok {
		t.Fatal("Couldn't get a valid gopath")
	}
	gopath += "/src"
	s, err := NewServer(addr, gopath, TEMP_DIR)
	if err != nil {
		t.Fatalf("Failed to create server")
	}
	return s
}

func TestNewServer(t *testing.T) {
	tests := []struct {
		outputDir string
	}{
		{"relative/dir"},
		{path.Join(TEMP_DIR, "abs/dir")},
	}

	addr := "localhost:8080"
	gopath := "go/path"
	for tci, tc := range tests {
		fmt.Println(tc)
		s, err := NewServer(addr, gopath, tc.outputDir)
		if err != nil {
			t.Fatalf("[%d] NewServer should have created correctly but didn't: %v",
				tci, err)
		}

		// verify all the members are set correctly
		if s.Addr != addr {
			t.Errorf("[%d] Addr not set correctly got %s, but wanted %s", tci,
				s.Addr, addr)
		}
		if s.GoSrc != "go/path" {
			t.Errorf("[%d] GoPath not set correctly got %s, but wanted %s", tci,
				s.GoSrc, gopath)
		}
		if !strings.HasSuffix(s.OutputDir, tc.outputDir) {
			t.Errorf("[%d] OutputDir not set correctly, expected path to end with "+
				"%s but got %s", tci, tc.outputDir, s.OutputDir)
		}

		// Make sure the output directory was created
		if _, err := os.Stat(s.OutputDir); err != nil {
			t.Errorf("[%d] expected output directory to exist but it doesn't", tci)
		}
	}
}

func TestBuildValidateRequests(t *testing.T) {
	s := newTestServerOrDie(t)
	filepath := "github.com/Stymphalian/ikuaki/athrun/server/testbinary/main.go"
	binaryName := "testbinary"

	// Create the binary in the output directory
	f, err := os.OpenFile(path.Join(s.OutputDir, binaryName), os.O_CREATE, 0755)
	if err != nil {
		t.FailNow()
	}
	f.Close()

	failTests := []struct {
		req pb.BuildRequest
	}{
		{pb.BuildRequest{}},
		{pb.BuildRequest{Filepath: ""}},
		{pb.BuildRequest{Filepath: "path"}},
		{pb.BuildRequest{Filepath: "path", Name: ""}},
		{pb.BuildRequest{Filepath: "path", Name: binaryName}},
		{pb.BuildRequest{Filepath: filepath, Name: binaryName}},
		{pb.BuildRequest{Filepath: filepath, Name: binaryName, Overwrite: false}},
	}

	for tci, tc := range failTests {
		_, err := s.Build(context.Background(), &tc.req)
		if err == nil {
			t.Fatalf("[%d] expected Build to fail but didn't", tci)
		}
	}
}

func TestBuildSuccess(t *testing.T) {
	s := newTestServerOrDie(t)
	filepath := "github.com/Stymphalian/ikuaki/athrun/server/testbinary/main.go"
	binaryName := "testbinary2"
	binaryFilepath := path.Join(s.OutputDir, binaryName)

	passTests := []struct {
		req pb.BuildRequest
	}{
		{
			pb.BuildRequest{Filepath: filepath, Name: binaryName},
		},
		{
			pb.BuildRequest{Filepath: filepath, Name: binaryName,
				Args: []string{"-v"}},
		},
	}

	for tci, tc := range passTests {
		resp, err := s.Build(context.Background(), &tc.req)
		if err != nil {
			t.Fatalf("[%d] expected Build to succeed but didn't %v", tci, err)
		}

		if resp.OutputFilepath != binaryFilepath {
			t.Errorf("[%d] OutputFilepath not expected. got %s, want %s",
				tci, resp.OutputFilepath, binaryFilepath)
		}
		if _, err := os.Stat(resp.OutputFilepath); err != nil {
			t.Errorf("[%d] OutputFilepath %s should exist but doesn't",
				tci, resp.OutputFilepath)
		}
		// delete the file for the next test case.
		if err := os.Remove(binaryFilepath); err != nil {
			panic(err)
		}

		if len(resp.Command) == 0 {
			t.Errorf("[%d] Command is supposed to be filled but is empty", tci)
		}
	}
}

func TestEnv(t *testing.T) {
	addr := "localhost:8080"
	gopath := "gopath"
	s, err := NewServer(addr, gopath, TEMP_DIR)
	if err != nil {
		t.Fatalf("Failed to create server")
	}

	testcases := []struct {
		req  pb.EnvRequest
		want map[string]string
		pass bool
	}{
		{
			pb.EnvRequest{},
			map[string]string{
				"ADDR":       addr,
				"GOSRC":      gopath,
				"OUTPUT_DIR": s.OutputDir},
			true,
		},
		{
			pb.EnvRequest{Args: []string{"ADDR"}},
			map[string]string{
				"ADDR": addr},
			true,
		},
		{
			pb.EnvRequest{Args: []string{"GOSRC"}},
			map[string]string{
				"GOSRC": gopath},
			true,
		},
		{
			pb.EnvRequest{Args: []string{"OUTPUT_DIR"}},
			map[string]string{
				"OUTPUT_DIR": s.OutputDir},
			true,
		},
		{
			pb.EnvRequest{Args: []string{"GOSRC", "OUTPUT_DIR"}},
			map[string]string{
				"GOSRC":      gopath,
				"OUTPUT_DIR": s.OutputDir},
			true,
		},
		{
			pb.EnvRequest{Args: []string{"NOOOOPE"}},
			map[string]string{},
			false,
		},
	}

	for tci, tc := range testcases {
		resp, err := s.Env(context.Background(), &tc.req)
		if tc.pass && err != nil {
			t.Fatalf("[%d] expected Env() to succeed but didn't %v", tci, err)
		} else if !tc.pass && err == nil {
			t.Fatalf("[%d] expected Env() to fail but didn't", tci)
		}

		if err != nil {
			continue
		}

		if len(resp.Env) != len(tc.want) {
			t.Errorf("[%d] expected Env map to be the same but they are not."+
				" got %v, want %v", tci, resp.Env, tc.want)
		}
		for k, _ := range resp.Env {
			if tc.want[k] != resp.Env[k] {
				t.Errorf("[%d] expected Env[%s] elements to be the same but they diff"+
					"got %s, want %s", tci, k, resp.Env[k], tc.want[k])
			}
		}
	}
}

func TestMain(m *testing.M) {
	tempdir, err := ioutil.TempDir("", "server_test.go")
	if err != nil {
		panic(err)
	}
	TEMP_DIR = tempdir
	if err := os.Chmod(TEMP_DIR, 0777); err != nil {
		panic(err)
	}

	originalDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir(TEMP_DIR)

	defer os.Chdir(originalDir)
	defer os.RemoveAll(TEMP_DIR)
	os.Exit(m.Run())
}
