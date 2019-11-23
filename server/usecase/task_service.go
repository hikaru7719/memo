package usecase

import "github.com/hikaru7719/memo/server/domain"

// TaskUsecase is an interface
type TaskUsecase interface {
	CreateTask(*domain.Task, string, []string) (*domain.Task, error)
}

type taskUsecase struct {
	taskRepositoty domain.TaskRepositoy
	tagRepository  domain.TagRepository
}

func (t *taskUsecase) CreateTask(task *domain.Task, tagNames []string, userID string) (*domain.Task, error) {
	tags := make([]*domain.Tag, 0, len(tagNames))
	for _, name := range tagNames {
		tag, err := t.tagRepository.FindByName(name)
		if err != nil {
			return nil, err
		}
		if tag.ID == 0 {
			tag.Name = name
			tag, err = t.tagRepository.Create(tag)
			if err != nil {
				return nil, err
			}
		}
	}

	task, err := t.taskRepositoty.Create(task)
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		err = t.taskRepositoty.CreateAssociation(task, tag)
		if err != nil {
			return nil, err
		}
	}

	task.Tags = tags
	return task, nil
}
