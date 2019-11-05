package controller

import (
	"context"

	"github.com/hikaru7719/memo/server/domain"
	pb "github.com/hikaru7719/memo/server/proto"
	"github.com/hikaru7719/memo/server/usecase"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// TaskService implements pb.TaskService
type TaskService struct {
	taskUsecase usecase.TaskUsecase
}

func (t *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	task := t.convertDoamin(req)
	task, err := t.taskUsecase.CreateTask(task, task.UserID, task.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, "CreateTaskError")
	}
	return t.convertResponse(task), nil
}

func (t *TaskService) convertDoamin(req *pb.CreateTaskRequest) *domain.Task {
	task := &domain.Task{}
	task.Name = req.Task.Name
	task.Explanation = req.Task.Explanation
	task.Status = convertDomainStatus(int32(req.Task.Status))
	task.UserID = req.Task.UserID
	return task
}

func (t *TaskService) convertResponse(task *domain.Task) *pb.Task {
	return &pb.Task{
		TaskID:      uint32(task.ID),
		Name:        task.Name,
		Explanation: task.Explanation,
		UserID:      task.UserID,
		TagID:       uint32(task.TagID),
		Status:      pb.Status(convertProtoStatus(task.Status)),
	}
}
