// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"
)

var UpdateConnectionCredentialInput connectionpb.UpdateConnectionCredentialRequest

var UpdateConnectionCredentialFromFile string

var UpdateConnectionCredentialInputCredentialCredential string

var UpdateConnectionCredentialInputCredentialCredentialCloudSql connectionpb.ConnectionCredential_CloudSql

func init() {
	ConnectionServiceCmd.AddCommand(UpdateConnectionCredentialCmd)

	UpdateConnectionCredentialInput.Credential = new(connectionpb.ConnectionCredential)

	UpdateConnectionCredentialInputCredentialCredentialCloudSql.CloudSql = new(connectionpb.CloudSqlCredential)

	UpdateConnectionCredentialCmd.Flags().StringVar(&UpdateConnectionCredentialInput.Name, "name", "", "Required. Required. Name of the connection, for example: ...")

	UpdateConnectionCredentialCmd.Flags().StringVar(&UpdateConnectionCredentialInputCredentialCredentialCloudSql.CloudSql.Username, "credential.credential.cloud_sql.username", "", "The username for the credential.")

	UpdateConnectionCredentialCmd.Flags().StringVar(&UpdateConnectionCredentialInputCredentialCredentialCloudSql.CloudSql.Password, "credential.credential.cloud_sql.password", "", "The password for the credential.")

	UpdateConnectionCredentialCmd.Flags().StringVar(&UpdateConnectionCredentialInputCredentialCredential, "credential.credential", "", "Choices: cloud_sql")

	UpdateConnectionCredentialCmd.Flags().StringVar(&UpdateConnectionCredentialFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var UpdateConnectionCredentialCmd = &cobra.Command{
	Use:   "update-connection-credential",
	Short: "Sets the credential for the specified connection.",
	Long:  "Sets the credential for the specified connection.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if UpdateConnectionCredentialFromFile == "" {

			cmd.MarkFlagRequired("name")

			cmd.MarkFlagRequired("credential.credential")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if UpdateConnectionCredentialFromFile != "" {
			in, err = os.Open(UpdateConnectionCredentialFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &UpdateConnectionCredentialInput)
			if err != nil {
				return err
			}

		} else {

			switch UpdateConnectionCredentialInputCredentialCredential {

			case "cloud_sql":
				UpdateConnectionCredentialInput.Credential.Credential = &UpdateConnectionCredentialInputCredentialCredentialCloudSql

			default:
				return fmt.Errorf("Missing oneof choice for credential.credential")
			}

		}

		if Verbose {
			printVerboseInput("Connection", "UpdateConnectionCredential", &UpdateConnectionCredentialInput)
		}
		err = ConnectionClient.UpdateConnectionCredential(ctx, &UpdateConnectionCredentialInput)

		return err
	},
}
