package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/gabriel-de-lisle/tri/protocol"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	url = flag.String("url", "localhost:5001", "The server url")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo list application",
	Long:  "Tri will help you manage your tasks more efficiently",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	flag.Parse()
}

type rawCmd func(*cobra.Command, []string)
type cmdWithClientAndContext func(pb.TaskHandlerClient, context.Context, *cobra.Command, []string)

func AddClientAndContext(call cmdWithClientAndContext) rawCmd {
	return func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(*url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := pb.NewTaskHandlerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		call(client, ctx, cmd, args)
	}
}
