package main

import (
	"context"
	"sort"

	pb "github.com/gabriel-de-lisle/tri/protocol"
)

func (s *server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.AddTaskReply, error) {
	err := CreateTask(in.Description, int(in.Priority))
	if err != nil {
		return &pb.AddTaskReply{Message: "Error"}, err
	}

	return &pb.AddTaskReply{Message: "Task successfully added"}, nil
}

func (s *server) SetDoneTask(ctx context.Context, in *pb.SetDoneTaskRequest) (*pb.SetDoneTaskReply, error) {
	var taskIds []uint
	for _, taskId := range in.Ids {
		taskIds = append(taskIds, uint(taskId))
	}

	err := MarkTasksAsDone(taskIds)
	if err != nil {
		return &pb.SetDoneTaskReply{Message: "Error"}, err
	}

	return &pb.SetDoneTaskReply{Message: "Tasks successfully set to done"}, nil
}

func (s *server) GetTask(ctx context.Context, in *pb.GetTaskRequest) (*pb.GetTaskReply, error) {
	var requestedTasks []*pb.Task

	tasks, err := GetTasks(in.ShowAll, in.ShowDone)
	if err != nil {
		return &pb.GetTaskReply{Tasks: requestedTasks}, err
	}

	sort.Sort(ByPriority(tasks))

	for count, task := range tasks {
		if count >= int(in.Top) {
			break
		}
		requestedTask := pb.Task{
			Description: task.Description,
			Done:        task.Done,
			Priority:    int32(task.Priority),
			Id:          uint32(task.ID),
		}
		requestedTasks = append(requestedTasks, &requestedTask)
	}

	return &pb.GetTaskReply{Tasks: requestedTasks}, nil
}
