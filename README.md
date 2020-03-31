# bqconn

This tool can manage bigquery external connections.

## Usage

    Manages external data source connections and credentials.
    
    Usage:
      bqconn connection [command]
    
    Available Commands:
      create-connection            Creates a new connection.
      delete-connection            Deletes connection and associated credential.
      get-connection               Returns specified connection.
      get-iam-policy               Gets the access control policy for a resource. ...
      list-connections             Returns a list of connections in the given...
      set-iam-policy               Sets the access control policy on the specified...
      test-iam-permissions         Returns permissions that a caller has on the...
      update-connection            Updates the specified connection. For security...
      update-connection-credential Sets the credential for the specified connection.
    
    Flags:
          --address string   Set API address used by client. Or use BQCONN_CONNECTION_ADDRESS.
          --api_key string   Set API Key used by the client. Or use BQCONN_CONNECTION_API_KEY.
      -h, --help             help for connection
          --insecure         Make insecure client connection. Or use BQCONN_CONNECTION_INSECURE. Must be used with "address" option
          --token string     Set Bearer token used by the client. Or use BQCONN_CONNECTION_TOKEN.
    
    Global Flags:
      -j, --json      Print JSON output
      -v, --verbose   Print verbose output
    
    Use "bqconn connection [command] --help" for more information about a command.

## Example

example-input.json:
```json
{
	"parent": "projects/gcp-project-name-here/locations/us",
	"connection_id": "friendly-connection-id-here",
	"connection": {
		"cloud_sql": {
			"instance_id": "gcpproject-name-here:us-central1:cloud-sql-name-here",
			"database": "db-name-here",
			"type": "POSTGRES",
			"credential": {
				"username": "user-name-here",
				"password": "password-here"
			}
		}
	}
}
```

```sh
bqconn connection create-connection --from_file example-input.json
```
