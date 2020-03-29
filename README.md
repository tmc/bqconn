# mkbigqueryconn

This tool can create and update bigquery external connections.

## Usage

    Usage:
      mkbigqueryconn [command]
    
    Available Commands:
      create      Create a BigQuery external connection.
      help        Help about any command
      update      Update a BigQuery external connection.
    
    Flags:
      -h, --help   help for mkbigqueryconn
    
    Use "mkbigqueryconn [command] --help" for more information about a command.

## Example

example-input.json:
```json
{
	"parent": "projects/gcp-project-name-here/locations/us",
	"connection_id": "friends-of-heidi-briones",
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
mkbigqueryconn create example-input.json
```
