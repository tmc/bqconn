// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	"github.com/golang/protobuf/jsonpb"

	"os"
)

var DeleteConnectionInput connectionpb.DeleteConnectionRequest

var DeleteConnectionFromFile string

func init() {
	ConnectionServiceCmd.AddCommand(DeleteConnectionCmd)

	DeleteConnectionCmd.Flags().StringVar(&DeleteConnectionInput.Name, "name", "", "Required. Required. Name of the deleted connection, for...")

	DeleteConnectionCmd.Flags().StringVar(&DeleteConnectionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DeleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes connection and associated credential.",
	Long:  "Deletes connection and associated credential.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DeleteConnectionFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DeleteConnectionFromFile != "" {
			in, err = os.Open(DeleteConnectionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DeleteConnectionInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Connection", "DeleteConnection", &DeleteConnectionInput)
		}
		err = ConnectionClient.DeleteConnection(ctx, &DeleteConnectionInput)

		return err
	},
}
