package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	pb "github.com/Stymphalian/ikuaki/athrun/protos"
	"github.com/mgutz/ansi"
)

type Server struct {
	Addr      string
	GoSrc     string
	OutputDir string

	goBinary string
}

// Create a new server initializing to used the address, goPath and outputDir
// This will MkDir on the outputDir to make sure it exists.
func NewServer(addr string, goPath string, outputDir string) (*Server, error) {
	goBin, err := exec.LookPath("go")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find 'go' binary."+
			" Make sure it is in your path. %v", err)
	}

	if outputDir[0] != '/' {
		// We were passed a relative directory, so we must append the PWD
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("Failed to get current working directory,"+
				" couldn't initialize --output_dir. %v", err)
		}
		outputDir = path.Join(cwd, outputDir)
	}

	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return nil, fmt.Errorf("Failed to create --output_dir %s. %v", outputDir,
				err)
		}
	}

	return &Server{
		Addr:      addr,
		GoSrc:     goPath,
		OutputDir: outputDir,
		goBinary:  goBin,
	}, nil
}

// Build effectivly does 'go build -o <outputname> [other args] <input package>
// It will validate the request to make sure all the parameters are provided.
func (this *Server) Build(ctx context.Context, req *pb.BuildRequest) (*pb.BuildResponse, error) {
	// Make sure all REQUIRED fields are provided
	if req.Filepath == "" {
		return nil, fmt.Errorf("filepath in the request must be set")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name in the request must be set")
	}

	// Make sure the input filepath actually points to a directory
	inputFilepath := path.Join(this.GoSrc, req.Filepath)
	if _, err := os.Stat(inputFilepath); err != nil {
		return nil, fmt.Errorf("File %s does not exist %v", inputFilepath, err)
	}

	outputFilepath := path.Join(this.OutputDir, req.Name)
	if _, err := os.Stat(outputFilepath); err == nil {
		if !req.Overwrite {
			return nil, fmt.Errorf("File with name %s already exists. "+
				"not overwriting", outputFilepath)
		}
	}

	args := []string{
		"build", "-o", outputFilepath,
	}
	args = append(args, req.Args...)
	args = append(args, inputFilepath)
	cmd := exec.Command(this.goBinary, args...)

	// Print the command
	log.Println(ansi.Blue, "Beginning build:", ansi.Reset)
	fmt.Printf("%s\n", cmd.Path)
	for _, s := range cmd.Args[1:] {
		fmt.Printf("  %s\n", s)
	}
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(ansi.Green, "Error", ansi.Reset, string(output), err)
		return nil, err
	}
	log.Println(ansi.Green, "Finished build.", ansi.Reset)

	return &pb.BuildResponse{
		OutputFilepath: outputFilepath,
		Stdout:         string(output),
		Command:        cmd.Args,
	}, nil
}

// Env returns information such as the GOSRC, OUTPUT_DIR and the current ADDR
func (this *Server) Env(ctx context.Context, req *pb.EnvRequest) (*pb.EnvResponse, error) {
	if len(req.Args) == 0 {
		req.Args = []string{"ADDR", "GOSRC", "OUTPUT_DIR"}
	}

	resp := &pb.EnvResponse{
		Env: make(map[string]string),
	}
	for _, s := range req.Args {
		switch s {
		case "ADDR":
			resp.Env["ADDR"] = this.Addr
		case "GOSRC":
			resp.Env["GOSRC"] = this.GoSrc
		case "OUTPUT_DIR":
			resp.Env["OUTPUT_DIR"] = this.OutputDir
		default:
			return nil, fmt.Errorf("Arg %s not supported", s)
		}
	}
	return resp, nil
}
