/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gabriel-de-lisle/tri/todo"
	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add task...",
	Short: "Add a task",
	Long:  "Add a task",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(datafile)
	for _, arg := range args {
		item := todo.Item{Text: arg, Done: false}
		item.SetPriority(priority)
		item.SetDate()
		items = append(items, item)
	}
	err := todo.SaveItems(datafile, items)
	if err != nil {
		fmt.Println("%v", err)
	}
	displayTasks(10, false, false)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority:1,2,3")
}
