// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"
)

var GetConnectionInput connectionpb.GetConnectionRequest

var GetConnectionFromFile string

func init() {
	ConnectionServiceCmd.AddCommand(GetConnectionCmd)

	GetConnectionCmd.Flags().StringVar(&GetConnectionInput.Name, "name", "", "Required. Required. Name of the requested connection, for...")

	GetConnectionCmd.Flags().StringVar(&GetConnectionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Returns specified connection.",
	Long:  "Returns specified connection.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetConnectionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetConnectionFromFile != "" {
			in, err = os.Open(GetConnectionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetConnectionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Connection", "GetConnection", &GetConnectionInput)
		}
		resp, err := ConnectionClient.GetConnection(ctx, &GetConnectionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
