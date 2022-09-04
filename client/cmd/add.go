package cmd

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gabriel-de-lisle/tri/protocol"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add task...",
	Short: "Add a task",
	Long:  "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		CallWithClientAndContext(addRun, cmd, args)
	},
}

func addRun(client pb.TaskHandlerClient, ctx context.Context, cmd *cobra.Command, args []string) {
	for _, arg := range args {
		_, err := client.AddTask(ctx, &pb.AddTaskRequest{Description: arg, Priority: 1})
		if err != nil {
			log.Fatalf("Could not add tasks: %v", err)
		}
	}
	fmt.Println("Tasks successfully added")
	displayTasks(client, ctx, len(args), false, false)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority:1,2,3")
}
