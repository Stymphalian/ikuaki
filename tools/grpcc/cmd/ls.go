// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

var lsCmdFlags = struct {
	FMethodRpcMultiLine bool
}{}

// 127.0.0.1:8080/ikuaki.Lobby/Method
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all the services",
	Long:  `List the services/methods available on this servier`,
	Example: `
grpcc 127.0.0.1:8080
>> ikuaki.Service
>> grpc.reflection.v1alpha.ServerReflection

grpcc 127.0.0.1:8080 ikuaki.Service
>> CreateWorld
>> CreateAgent

grpcc 127.0.0.1:8080 ikuaki.Service CreateWorld
>> 
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostport := args[0]
		conn, err := grpc.Dial(hostport, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Failed to reach server(%s) %v", hostport, err)
		}
		client := grpcreflect.NewClient(context.Background(), rpb.NewServerReflectionClient(conn))

		if len(args) == 1 {
			// Just list the services on the server
			services, err := client.ListServices()
			if err != nil {
				fmt.Printf("Failed to list services %v", err)
			}
			for _, s := range services {
				fmt.Println(s)
			}
		} else if len(args) == 2 {
			// List the methods belonging to the service
			service, err := client.ResolveService(args[1])
			if err != nil {
				fmt.Printf("Service %s not found", service)
			}
			methods := service.GetMethods()
			for _, m := range methods {
				fmt.Println(m.GetName())
			}

		} else if len(args) == 3 {
			methodName := args[2]
			// List the methods belonging to the service
			service, err := client.ResolveService(args[1])
			if err != nil {
				fmt.Printf("Service %s not found", service)
			}

			m := service.FindMethodByName(methodName)
			if m == nil {
				fmt.Printf("Failed to find method: %s\n", methodName)
				return
			}
			fmt.Print(getMethodString(m, lsCmdFlags.FMethodRpcMultiLine))

		} else {
			// Ignoring extra args.
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.
	lsCmd.PersistentFlags().BoolVar(&lsCmdFlags.FMethodRpcMultiLine,
		"method-multi-line", false, "Print out the method rpc on multiple lines")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
