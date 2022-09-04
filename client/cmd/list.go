package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"

	pb "github.com/gabriel-de-lisle/tri/protocol"
)

var listCmd = &cobra.Command{
	Use:   "list [top]",
	Short: "List all tasks",
	Long:  "List all tasks",
	Run:   AddClientAndContext(listRun),
}

func listRun(client pb.TaskHandlerClient, ctx context.Context, cmd *cobra.Command, args []string) {
	var top int
	if len(args) == 0 {
		top = 10
	} else {
		value, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("%v", err)
		}
		top = value
	}

	showDone, err := cmd.Flags().GetBool("done")
	if err != nil {
		log.Fatalf("%v", err)
	}
	showAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		log.Fatalf("%v", err)
	}
	displayTasks(client, ctx, top, showDone, showAll)
}

func displayTasks(client pb.TaskHandlerClient, ctx context.Context, top int, showDone bool, showAll bool) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 1)

	reply, err := client.GetTask(ctx, &pb.GetTaskRequest{Top: int32(top), ShowDone: showDone, ShowAll: showAll})
	if err != nil {
		log.Fatalf("Could list tasks: %v", err)
	}

	for _, task := range reply.GetTasks() {
		doneState := " "
		if task.Done {
			doneState = "x"
		}

		line := fmt.Sprintf("%v.\t%s\t(%v)\t%s\n", task.Id, doneState, task.Priority, task.Description)
		fmt.Fprint(w, line)
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "Show all tasks")
	listCmd.Flags().BoolP("done", "d", false, "Show done tasks only")
}
