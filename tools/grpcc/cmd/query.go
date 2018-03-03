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
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "get info about protos",
	Long: `Pass in the fully qualified proto name <package>.<message> and this
command will print out the proto message for you.`,
	Example: `grpcc query 127.0.0.1:8080 ikuaki.CreateWorldReq
>> message CreateWorldReq {
	string name = 1;
}`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		hostport := args[0]
		conn, err := grpc.Dial(hostport, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Failed to reach server(%s) %v", hostport, err)
			return
		}
		defer conn.Close()
		client := grpcreflect.NewClient(context.Background(),
			rpb.NewServerReflectionClient(conn))

		msgName := args[1]
		msg, err := client.ResolveMessage(msgName)
		if err != nil {
			fmt.Printf("Message %s not found\n", msgName)
			return
		}

		p := protoprint.Printer{}
		buf := bytes.NewBuffer([]byte{})
		err = p.PrintProtoFile(msg.GetFile(), buf)
		if err != nil {
			fmt.Printf("Failed to print file\n")
			return
		}

		// get the base message name (i.e. remove the package path)
		baseMsgName := ""
		msgParts := strings.Split(msgName, ".")
		if len(msgParts) == 1 {
			baseMsgName = msgName
		} else {
			baseMsgName = msgParts[len(msgParts)-1]
		}

		i := strings.Index(buf.String(), "message "+baseMsgName)
		startString := buf.String()[i:len(buf.String())]
		last := strings.Index(startString, "\n}")
		fullThing := startString[0 : last+len("\n}")]
		fmt.Println(fullThing)

	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
