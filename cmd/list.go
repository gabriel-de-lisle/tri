/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gabriel-de-lisle/tri/todo"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [top]",
	Short: "List all tasks",
	Long:  "List all tasks",
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	var top int
	if len(args) == 0 {
		top = 10
	} else {
		value, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("%v", err)
		}
		top = value
	}

	showDone, err := cmd.Flags().GetBool("done")
	if err != nil {
		fmt.Println("%v", err)
	}
	showAll, err := cmd.Flags().GetBool("all")
	if err != nil {
		fmt.Println("%v", err)
	}
	displayTasks(top, showDone, showAll)
}

func displayTasks(top int, showDone bool, showAll bool) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 1)

	items, err := todo.ReadItems(datafile)

	sort.Sort(todo.ByPriority(items))
	if err != nil {
		fmt.Println("%v", err)
	}

	displayedTask := 0
	for i, item := range items {
		doneState := " "
		if item.Done {
			doneState = "x"
		}
		if (showAll || (showDone && item.Done) || (!showDone && !item.Done)) && displayedTask < top {
			line := fmt.Sprintf("%v.\t%s\t(%v)\t%s\n", i+1, doneState, item.Priority, item.Text)
			fmt.Fprint(w, line)
			displayedTask += 1
		}

	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks")
	listCmd.Flags().BoolP("done", "d", false, "Show done tasks only")
}
