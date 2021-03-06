// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	iampb "google.golang.org/genproto/googleapis/iam/v1"

	"os"
)

var TestIamPermissionsInput iampb.TestIamPermissionsRequest

var TestIamPermissionsFromFile string

func init() {
	ConnectionServiceCmd.AddCommand(TestIamPermissionsCmd)

	TestIamPermissionsCmd.Flags().StringVar(&TestIamPermissionsInput.Resource, "resource", "", "Required. REQUIRED: The resource for which the policy...")

	TestIamPermissionsCmd.Flags().StringSliceVar(&TestIamPermissionsInput.Permissions, "permissions", []string{}, "Required. The set of permissions to check for the...")

	TestIamPermissionsCmd.Flags().StringVar(&TestIamPermissionsFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var TestIamPermissionsCmd = &cobra.Command{
	Use:   "test-iam-permissions",
	Short: "Returns permissions that a caller has on the...",
	Long:  "Returns permissions that a caller has on the specified resource.  If the resource does not exist, this will return an empty set of  permissions, not...",
	PreRun: func(cmd *cobra.Command, args []string) {

		if TestIamPermissionsFromFile == "" {

			cmd.MarkFlagRequired("resource")

			cmd.MarkFlagRequired("permissions")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if TestIamPermissionsFromFile != "" {
			in, err = os.Open(TestIamPermissionsFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &TestIamPermissionsInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Connection", "TestIamPermissions", &TestIamPermissionsInput)
		}
		resp, err := ConnectionClient.TestIamPermissions(ctx, &TestIamPermissionsInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
