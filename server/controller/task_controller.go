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

// CreateTask is contoroller for endpoint of /tasks
func (t *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	task, tags := t.convertDoamin(req)
	task, err := t.taskUsecase.CreateTask(task, task.UserID, tags)
	if err != nil {
		return nil, status.Error(codes.Internal, "CreateTaskError")
	}
	return t.convertResponse(task), nil
}

func (t *TaskService) convertDoamin(req *pb.CreateTaskRequest) (*domain.Task, []string) {
	task := &domain.Task{}
	task.Name = req.Task.Name
	task.Explanation = req.Task.Explanation
	task.Status = convertDomainStatus(int32(req.Task.Status))
	task.UserID = req.Task.UserID
	tags := make([]string, 0, len(req.Task.Tags))
	for _, tag := range req.Task.Tags {
		tags = append(tags, tag.Name)
	}
	return task, tags
}

func (t *TaskService) convertResponse(task *domain.Task) *pb.Task {
	tags := make([]*pb.Tag, len(task.Tags))
	for _, tag := range task.Tags {
		tags = append(tags, &pb.Tag{
			Id:   uint32(tag.ID),
			Name: tag.Name,
		})
	}
	return &pb.Task{
		Id:          uint32(task.ID),
		Name:        task.Name,
		Explanation: task.Explanation,
		UserID:      task.UserID,
		Status:      pb.Status(convertProtoStatus(task.Status)),
		Tags:        tags,
	}
}
