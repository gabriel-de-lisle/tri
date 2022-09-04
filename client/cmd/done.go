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
	Run:   AddClientAndContext(runDone),
}

func runDone(client pb.TaskHandlerClient, ctx context.Context, cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("Please specify the ids of the tasks you want to mark as done")
	}
	var taskNumbers []uint32
	for _, arg := range args {
		taskNumber, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("%v", err)
		}
		taskNumbers = append(taskNumbers, uint32(taskNumber))
	}

	r, err := client.SetDoneTask(ctx, &pb.SetDoneTaskRequest{Ids: taskNumbers})
	if err != nil {
		log.Fatalf("Could set tasks to done: %v", err)
	}
	fmt.Println(r.Message)
	displayTasks(client, ctx, len(taskNumbers), true, false)
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
