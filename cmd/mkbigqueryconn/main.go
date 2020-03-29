// Command mkcloudsqlconn creates an external connection to a cloud sql instance.
package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	bigqueryconnection "github.com/tmc/mkbigqueryconn/bqconn"
	connectionpb "google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1beta1"
)

var (
	cmdCreate = &cobra.Command{
		Use:  "create",
		Args: cobra.ExactArgs(1),
		Run:  runCreate,
	}
	cmdUpdate = &cobra.Command{
		Use:  "update",
		Args: cobra.ExactArgs(1),
		Run:  runUpdate,
	}
)

func main() {
	rootCmd := &cobra.Command{Use: os.Args[0]}
	rootCmd.AddCommand(
		cmdCreate,
		cmdUpdate,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func getIn(src string) (io.Reader, error) {
	var in io.Reader = os.Stdin
	var err error
	if src != "-" {
		in, err = os.Open(src)
	}
	return in, err
}

func run(ctx context.Context, src string, dump bool) error {
	c, err := bigqueryconnection.NewConnectionClient(ctx)
	if err != nil {
		return err
	}
	in, err := getIn(src)
	if err != nil {
		return err
	}
	fmt.Println(c, in)
	if err != nil {
		return err
	}
	return nil
}

func errOut(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func maybeErrOut(err error) {
	if err != nil {
		errOut(err)
	}
}

func marshalToStdout(m proto.Message) {
	(&jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "  ",
	}).Marshal(os.Stdout, m)
}

func runCreate(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c, err := bigqueryconnection.NewConnectionClient(ctx)
	maybeErrOut(err)
	in, err := getIn(args[0])
	maybeErrOut(err)

	req := &connectionpb.CreateConnectionRequest{}
	if err := jsonpb.Unmarshal(in, req); err != nil {
		errOut(err)
	}
	marshalToStdout(req)
	conn, err := c.CreateConnection(ctx, req)
	if err != nil {
		errOut(err)
	}
	marshalToStdout(conn)
}

func runUpdate(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c, err := bigqueryconnection.NewConnectionClient(ctx)
	maybeErrOut(err)
	in, err := getIn(args[0])
	maybeErrOut(err)

	req := &connectionpb.UpdateConnectionRequest{}
	if err := jsonpb.Unmarshal(in, req); err != nil {
		errOut(err)
	}
	marshalToStdout(req)
	conn, err := c.UpdateConnection(ctx, req)
	if err != nil {
		errOut(err)
	}
	marshalToStdout(conn)
}
