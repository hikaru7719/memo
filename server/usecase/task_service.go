package usecase

import "github.com/hikaru7719/memo/server/domain"

// TaskUsecase is an interface
type TaskUsecase interface {
	CreateTask(*domain.Task, string, string) (*domain.Task, error)
}

type taskUsecase struct {
	taskRepositoty domain.TaskRepositoy
	tagRepository  domain.TagRepository
}

func (t *taskUsecase) CreateTask(task *domain.Task, tagName string, userID string) (*domain.Task, error) {
	tag, err := t.tagRepository.FindByName(task.UserID, tagName)
	if err != nil {
		return nil, err
	}
	if tag.ID == 0 {
		tag.Name = tagName
		tag.UserID = userID
		tag, err = t.tagRepository.Create(tag)
		if err != nil {
			return nil, err
		}
	}
	task.TagID = tag.ID
	task, err = t.taskRepositoty.Create(task)
	return task, nil
}
