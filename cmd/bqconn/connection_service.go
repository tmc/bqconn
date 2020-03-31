// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	gapic "github.com/tmc/mkbigqueryconn/bqconn"
)

var ConnectionConfig *viper.Viper
var ConnectionClient *gapic.ConnectionClient
var ConnectionSubCommands []string = []string{
	"create-connection",
	"get-connection",
	"list-connections",
	"update-connection",
	"update-connection-credential",
	"delete-connection",
	"get-iam-policy",
	"set-iam-policy",
	"test-iam-permissions",
}

func init() {
	rootCmd.AddCommand(ConnectionServiceCmd)

	ConnectionConfig = viper.New()
	ConnectionConfig.SetEnvPrefix("BQCONN_CONNECTION")
	ConnectionConfig.AutomaticEnv()

	ConnectionServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use BQCONN_CONNECTION_INSECURE. Must be used with \"address\" option")
	ConnectionConfig.BindPFlag("insecure", ConnectionServiceCmd.PersistentFlags().Lookup("insecure"))
	ConnectionConfig.BindEnv("insecure")

	ConnectionServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use BQCONN_CONNECTION_ADDRESS.")
	ConnectionConfig.BindPFlag("address", ConnectionServiceCmd.PersistentFlags().Lookup("address"))
	ConnectionConfig.BindEnv("address")

	ConnectionServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use BQCONN_CONNECTION_TOKEN.")
	ConnectionConfig.BindPFlag("token", ConnectionServiceCmd.PersistentFlags().Lookup("token"))
	ConnectionConfig.BindEnv("token")

	ConnectionServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use BQCONN_CONNECTION_API_KEY.")
	ConnectionConfig.BindPFlag("api_key", ConnectionServiceCmd.PersistentFlags().Lookup("api_key"))
	ConnectionConfig.BindEnv("api_key")
}

var ConnectionServiceCmd = &cobra.Command{
	Use:       "connection",
	Short:     "Manages external data source connections and...",
	Long:      "Manages external data source connections and credentials.",
	ValidArgs: ConnectionSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := ConnectionConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if ConnectionConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := ConnectionConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := ConnectionConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		ConnectionClient, err = gapic.NewConnectionClient(ctx, opts...)
		return
	},
}
