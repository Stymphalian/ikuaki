// Copyright © 2018 Jordan Yu <saturnslight@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"plugin"

	"github.com/Stymphalian/ikuaki/tools/grpcc/common"
	grpcc "github.com/Stymphalian/ikuaki/tools/grpcc/plugins"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmdFlags = struct {
	Ffile        string
	FbuildPlugin bool
}{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grpcc",
	Short: "grpc client - make RPC requests to a server",
	Long: `GRPCC - GRPC Client. A tool which uses proto reflection to allow
you to query any ip:port which is serving a GRPC service and make requests to 
it.`,
	Example: `grpcc localhost:8080 ikuaki.World CreateWorld  <<EOF
name: "jordan"
EOF
`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		hostport := args[0]
		serviceName := args[1]
		methodName := args[2]

		// dial to get the service
		conn, err := grpc.Dial(hostport, grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		// construct info from the connection
		info, err := grpcc.NewInfo(hostport, serviceName, methodName, conn)
		if err != nil {
			fmt.Println(err)
			return
		}

		if rootCmdFlags.Ffile != "" {
			// If a file is provided then run using the plugin file
			err = loadAndRunPlugin(rootCmdFlags.Ffile, info,
				rootCmdFlags.FbuildPlugin)
			return
		} else {
			switch info.StreamType {
			case grpcc.K_UNARY:
				err = makeSingleRequest(info)
			case grpcc.K_STREAM_REQUEST:
				err = makeMultipleRequests(info)
			case grpcc.K_STREAM_RESPONSE:
				err = makeRequestReceiveStream(info)
			case grpcc.K_STREAM_BIDI:
				if rootCmdFlags.Ffile == "" {
					log.Printf(`grpcc doesn't support streaming requests from just CLI, please supply a plugin file with --file`)
					return
				}
				err = loadAndRunPlugin(rootCmdFlags.Ffile, info,
					rootCmdFlags.FbuildPlugin)
			default:
				log.Fatal("Unknown stream type")
			}
		}

		if err != nil {
			log.Println(err)
			return
		}
	},
}

func makeSingleRequest(info *grpcc.Info) error {
	// Read from stdin the request text proto
	req, err := common.ReadTextprotoFromStdin(info.RequestDesc)
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

func makeMultipleRequests(info *grpcc.Info) error {
	// Read from stdin the request text proto
	stream, err := info.Stub.InvokeRpcClientStream(context.Background(),
		info.MethodDesc)
	if err != nil {
		return err
	}

	reqs, err := common.ReadTextprotosFromStdin(info.RequestDesc)
	if err != nil {
		return fmt.Errorf("Failed to parse textproto %s", err)
	}

	for _, req := range reqs {
		err = stream.SendMsg(req)
		if err != nil {
			return err
		}
	}

	// show the response
	resp, err := stream.CloseAndReceive()
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(resp)
	return nil
}

func makeRequestReceiveStream(info *grpcc.Info) error {
	// Read the request
	req, err := common.ReadTextprotoFromStdin(info.RequestDesc)
	if err != nil {
		return fmt.Errorf("Failed to parse textproto %s", err)
	}

	// create stream with request
	stream, err := info.Stub.InvokeRpcServerStream(context.Background(),
		info.MethodDesc, req)
	if err != nil {
		return err
	}

	// Get the stream of responses
	for {
		resp, err := stream.RecvMsg()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println(resp)
	}

	return nil
}

func loadAndRunPlugin(filepath string, info *grpcc.Info, buildplugin bool) error {
	if buildplugin {
		if err := common.BuildPlugin(filepath); err != nil {
			return err
		}
	}

	p, err := plugin.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open plugin")
		return err
	}

	fn, err := p.Lookup("Run")
	if err != nil {
		fmt.Println("Failed to find Run function")
		return err
	}

	return fn.(func(*grpcc.Info) error)(info)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootCmdFlags.Ffile, "file", "f",
		"", `If provided then load this file as a plugin which is a program for 
	running your own code with a stub to the server`)
	rootCmd.Flags().BoolVarP(&rootCmdFlags.FbuildPlugin, "build_plugin", "",
		true, `Whether to build the plugin --file before invoking it.`)
}
