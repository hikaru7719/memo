package repository

import (
	"github.com/hikaru7719/memo/server/entity"
)

// NewTaskRepository create TaskRepository Object
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

// TaskRepository returns resouces about task
type TaskRepository struct {
}

// Find gets all Done task matching user id
func (t *TaskRepository) Find(userID string) ([]entity.Done, error) {
	dones := make([]entity.Done, 0, 20)
	if result := db.Where("user_id = ?", userID).Find(&dones); result.Error != nil {
		return nil, result.Error
	}
	return dones, nil
}

// Create insert new Done Record
func (t *TaskRepository) Create(done *entity.Done) (*entity.Done, error) {
	if result := db.Create(done); result.Error != nil {
		return nil, result.Error
	}
	return done, nil
}

// Delete delete record matching done id
func (t *TaskRepository) Delete(done *entity.Done) error {
	if result := db.Delete(done); result.Error != nil {
		return result.Error
	}
	return nil
}
