package repository

import (
	"github.com/hikaru7719/memo/server/domain"
)

// NewTaskRepository create TaskRepository Object
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

// TaskRepository returns resouces about task
type TaskRepository struct{}

// Find gets 20 Tasks matching user id
func (t *TaskRepository) Find(userID string, startAt int) ([]domain.Task, error) {
	tasks := make([]domain.Task, 0, 20)
	if result := db.Offset(startAt).Where("user_id = ?", userID).Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}

// Create inserts new Task Record
func (t *TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	if result := db.Create(task); result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

// Delete deletes record matching task id
func (t *TaskRepository) Delete(task *domain.Task) error {
	if result := db.Delete(task); result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateAssociation inserts relation record
func (t *TaskRepository) CreateAssociation(task *domain.Task, tag *domain.Tag) error {
	if result := db.Model(task).Association("Tags").Append(tag); result != nil {
		return result.Error
	}
	return nil
}
