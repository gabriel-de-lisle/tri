package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/gabriel-de-lisle/tri/protocol"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done task...",
	Short: "Mark a task as done",
	Long:  "When a task is completed, use this command to mark it as done",
	Run: func(cmd *cobra.Command, args []string) {
		CallWithClientAndContext(runDone, cmd, args)
	},
}

func runDone(client pb.TaskHandlerClient, ctx context.Context, cmd *cobra.Command, args []string) {
	var taskNumbers []int64
	for _, arg := range args {
		taskNumber, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			log.Fatalf("%v", err)
		}
		taskNumbers = append(taskNumbers, taskNumber)
	}
	if len(args) == 0 {
		taskNumbers = append(taskNumbers, 1)
	}

	r, err := client.SetDoneTask(ctx, &pb.SetDoneTaskRequest{Positions: taskNumbers})
	if err != nil {
		log.Fatalf("Could set tasks to done: %v", err)
	}
	fmt.Println(r.GetMessage())
	displayTasks(client, ctx, len(taskNumbers), true, false)
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
