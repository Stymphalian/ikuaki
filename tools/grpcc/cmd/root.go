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
	"log"
	"os"
	"plugin"

	grpcc "github.com/Stymphalian/ikuaki/tools/grpcc/plugin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmdFlags = struct {
	Ffile string
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

		// Run the requests/responses either a direct request or by running the
		// user plugin file.
		if rootCmdFlags.Ffile != "" {
			err = loadAndRunPlugin(rootCmdFlags.Ffile, info)
		} else {
			if info.MethodDesc.IsClientStreaming() ||
				info.MethodDesc.IsServerStreaming() {
				log.Printf(`grpcc doesn't support streaming requests from just CLI,
please supply a plugin file with --file`)
				return
			}
			err = makeSingleRequest(info)
		}
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func loadAndRunPlugin(filepath string, info *grpcc.Info) error {
	p, err := plugin.Open(filepath)
	if err != nil {
		return err
	}

	fn, err := p.Lookup("Run")
	if err != nil {
		return err
	}

	return fn.(func(*grpcc.Info) error)(info)
}

func makeSingleRequest(info *grpcc.Info) error {
	// Read from stdin the request text proto
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootCmdFlags.Ffile, "file", "f",
		"", `If provided then load this file as a plugin which is a program for 
	running your own code with a stub to the server`)
}
