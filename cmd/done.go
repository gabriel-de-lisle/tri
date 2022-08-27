/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gabriel-de-lisle/tri/todo"
	"github.com/spf13/cobra"
	"sort"
	"strconv"
	"time"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done task...",
	Short: "Mark a task as done",
	Long:  "When a task is completed, use this command to mark it as done",
	Run:   runDone,
}

func runDone(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(datafile)
	if err != nil {
		fmt.Println("%v", err)
	}

	sort.Sort(todo.ByPriority(items))
	for _, arg := range args {
		taskNumber, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("%v", err)
		}
		items[taskNumber-1].Done = true
		items[taskNumber-1].Date = time.Now()
	}

	err = todo.SaveItems(datafile, items)
	if err != nil {
		fmt.Println("%v", err)
	}
	displayTasks(len(args), true, false)
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//doneCmd.Flags().IntVar("", "t", false, "Help message for toggle")
}
