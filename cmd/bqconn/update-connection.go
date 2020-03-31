// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"

	field_maskpb "google.golang.org/genproto/protobuf/field_mask"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	"strings"
)

var UpdateConnectionInput connectionpb.UpdateConnectionRequest

var UpdateConnectionFromFile string

var UpdateConnectionInputConnectionProperties string

var UpdateConnectionInputConnectionPropertiesCloudSql connectionpb.Connection_CloudSql

var UpdateConnectionInputConnectionPropertiesCloudSqlType string

func init() {
	ConnectionServiceCmd.AddCommand(UpdateConnectionCmd)

	UpdateConnectionInput.Connection = new(connectionpb.Connection)

	UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql = new(connectionpb.CloudSqlProperties)

	UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential = new(connectionpb.CloudSqlCredential)

	UpdateConnectionInput.UpdateMask = new(field_maskpb.FieldMask)

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInput.Name, "name", "", "Required. Required. Name of the connection to update, for...")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInput.Connection.Name, "connection.name", "", "The resource name of the connection in the form...")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInput.Connection.FriendlyName, "connection.friendly_name", "", "User provided display name for the connection.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInput.Connection.Description, "connection.description", "", "User provided description.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.InstanceId, "connection.properties.cloud_sql.instance_id", "", "Cloud SQL instance ID in the form...")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.Database, "connection.properties.cloud_sql.database", "", "Database name.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionPropertiesCloudSqlType, "connection.properties.cloud_sql.type", "", "Type of the Cloud SQL database.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential.Username, "connection.properties.cloud_sql.credential.username", "", "The username for the credential.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.Credential.Password, "connection.properties.cloud_sql.credential.password", "", "The password for the credential.")

	UpdateConnectionCmd.Flags().StringSliceVar(&UpdateConnectionInput.UpdateMask.Paths, "update_mask.paths", []string{}, "The set of field mask paths.")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionInputConnectionProperties, "connection.properties", "", "Choices: cloud_sql")

	UpdateConnectionCmd.Flags().StringVar(&UpdateConnectionFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var UpdateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates the specified connection. For security...",
	Long:  "Updates the specified connection. For security reasons, also resets  credential if connection properties are in the update field mask.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if UpdateConnectionFromFile == "" {

			cmd.MarkFlagRequired("name")

			cmd.MarkFlagRequired("connection.properties")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if UpdateConnectionFromFile != "" {
			in, err = os.Open(UpdateConnectionFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &UpdateConnectionInput)
			if err != nil {
				return err
			}

		} else {

			switch UpdateConnectionInputConnectionProperties {

			case "cloud_sql":
				UpdateConnectionInput.Connection.Properties = &UpdateConnectionInputConnectionPropertiesCloudSql

			default:
				return fmt.Errorf("Missing oneof choice for connection.properties")
			}

			UpdateConnectionInputConnectionPropertiesCloudSql.CloudSql.Type = connectionpb.CloudSqlProperties_DatabaseType(connectionpb.CloudSqlProperties_DatabaseType_value[strings.ToUpper(UpdateConnectionInputConnectionPropertiesCloudSqlType)])

		}

		if Verbose {
			printVerboseInput("Connection", "UpdateConnection", &UpdateConnectionInput)
		}
		resp, err := ConnectionClient.UpdateConnection(ctx, &UpdateConnectionInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
