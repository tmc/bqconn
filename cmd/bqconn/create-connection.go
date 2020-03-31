// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	"strings"
)

var CreateConnectionInput connectionpb.CreateConnectionRequest

var CreateConnectionFromFile string

var CreateConnectionInputConnectionProperties string

var CreateConnectionInputConnectionPropertiesCloudSql connectionpb.Connection_CloudSql

var CreateConnectionInputConnectionPropertiesCloudSqlType string

func init() {
	ConnectionServiceCmd.AddCommand(CreateConnectionCmd)

	CreateConnectionInput.Connection = new(connectionpb.Connection)

	CreateConnectionInputConnectionPropertiesCloudSql.CloudSql = new(connectionpb.CloudSqlProperties)

	CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential = new(connectionpb.CloudSqlCredential)

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInput.Parent, "parent", "", "Required. Required. Parent resource name.  Must be in the...")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInput.ConnectionId, "connection_id", "", "Optional. Connection id that should be assigned...")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInput.Connection.Name, "connection.name", "", "The resource name of the connection in the form...")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInput.Connection.FriendlyName, "connection.friendly_name", "", "User provided display name for the connection.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInput.Connection.Description, "connection.description", "", "User provided description.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.InstanceId, "connection.properties.cloud_sql.instance_id", "", "Cloud SQL instance ID in the form...")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.Database, "connection.properties.cloud_sql.database", "", "Database name.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionPropertiesCloudSqlType, "connection.properties.cloud_sql.type", "", "Type of the Cloud SQL database.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential.Username, "connection.properties.cloud_sql.credential.username", "", "The username for the credential.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential.Password, "connection.properties.cloud_sql.credential.password", "", "The password for the credential.")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionInputConnectionProperties, "connection.properties", "", "Choices: cloud_sql")

	CreateConnectionCmd.Flags().StringVar(&CreateConnectionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a new connection.",
	Long:  "Creates a new connection.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateConnectionFromFile == "" {

			cmd.MarkFlagRequired("parent")

			cmd.MarkFlagRequired("connection.properties")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if CreateConnectionFromFile != "" {
			in, err = os.Open(CreateConnectionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &CreateConnectionInput)
			if err != nil {
				return err
			}

		} else {

			switch CreateConnectionInputConnectionProperties {

			case "cloud_sql":
				CreateConnectionInput.Connection.Properties = &CreateConnectionInputConnectionPropertiesCloudSql

			default:
				return fmt.Errorf("Missing oneof choice for connection.properties")
			}

			CreateConnectionInputConnectionPropertiesCloudSql.CloudSql.Type = connectionpb.CloudSqlProperties_DatabaseType(connectionpb.CloudSqlProperties_DatabaseType_value[strings.ToUpper(CreateConnectionInputConnectionPropertiesCloudSqlType)])

		}

		if Verbose {
			printVerboseInput("Connection", "CreateConnection", &CreateConnectionInput)
		}
		resp, err := ConnectionClient.CreateConnection(ctx, &CreateConnectionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
