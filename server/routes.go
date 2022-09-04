package main

import (
	"context"
	"sort"

	pb "github.com/gabriel-de-lisle/tri/protocol"
)

func (s *server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.AddTaskReply, error) {

	task := Task{Text: in.GetDescription(), Done: false}
	task.SetPriority(in.GetPriority())
	task.SetDate()

	tasks, err := ReadTasks(*datafile)
	if err != nil {
		return &pb.AddTaskReply{Message: "Error"}, err
	}
	tasks = append(tasks, task)
	err = SaveTasks(*datafile, tasks)
	if err != nil {
		return &pb.AddTaskReply{Message: "Error"}, err
	}

	return &pb.AddTaskReply{Message: "Task successfully added"}, nil
}

func (s *server) SetDoneTask(ctx context.Context, in *pb.SetDoneTaskRequest) (*pb.SetDoneTaskReply, error) {

	positions := in.GetPositions()

	tasks, err := ReadTasks(*datafile)
	if err != nil {
		return &pb.SetDoneTaskReply{Message: "Error"}, err
	}

	sort.Sort(ByPriority(tasks))

	for _, position := range positions {
		tasks[position-1].Done = true
		tasks[position-1].SetDate()
	}

	err = SaveTasks(*datafile, tasks)
	if err != nil {
		return &pb.SetDoneTaskReply{Message: "Error"}, err
	}

	return &pb.SetDoneTaskReply{Message: "Tasks successfully set to done"}, nil
}

func (s *server) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	showDone := in.GetShowDone()
	showAll := in.GetShowAll()
	top := int(in.GetTop())

	var requestedTasks []*pb.Task
	tasks, err := ReadTasks(*datafile)
	if err != nil {
		return &pb.GetTaskReply{Tasks: requestedTasks}, err
	}

	sort.Sort(ByPriority(tasks))
	displayedTask := 0
	for position, task := range tasks {
		if (showAll || (showDone && task.Done) || (!showDone && !task.Done)) && displayedTask < top {
			requestedTask := pb.Task{
				Description: task.Text,
				Done:        task.Done,
				Priority:    task.Priority,
				Position:    int32(position + 1),
			}
			requestedTasks = append(requestedTasks, &requestedTask)
			displayedTask += 1
		}

	}

	return &pb.GetTaskReply{Tasks: requestedTasks}, nil
}
