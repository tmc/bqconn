// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	wrapperspb "github.com/golang/protobuf/ptypes/wrappers"
)

var ListConnectionsInput connectionpb.ListConnectionsRequest

var ListConnectionsFromFile string

func init() {
	ConnectionServiceCmd.AddCommand(ListConnectionsCmd)

	ListConnectionsInput.MaxResults = new(wrapperspb.UInt32Value)

	ListConnectionsCmd.Flags().StringVar(&ListConnectionsInput.Parent, "parent", "", "Required. Required. Parent resource name.  Must be in the...")

	ListConnectionsCmd.Flags().Uint32Var(&ListConnectionsInput.MaxResults.Value, "max_results.value", 0, "The uint32 value.")

	ListConnectionsCmd.Flags().StringVar(&ListConnectionsInput.PageToken, "page_token", "", "Page token.")

	ListConnectionsCmd.Flags().StringVar(&ListConnectionsFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var ListConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Returns a list of connections in the given...",
	Long:  "Returns a list of connections in the given project.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if ListConnectionsFromFile == "" {

			cmd.MarkFlagRequired("parent")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if ListConnectionsFromFile != "" {
			in, err = os.Open(ListConnectionsFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &ListConnectionsInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Connection", "ListConnections", &ListConnectionsInput)
		}
		resp, err := ConnectionClient.ListConnections(ctx, &ListConnectionsInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
