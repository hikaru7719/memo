package repository

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hikaru7719/memo/server/domain"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	repo := NewTaskRepository()
	id := fmt.Sprintf("ID%s", uuid.New())
	data := &domain.Task{
		Name:        "TestDoneTask",
		Explanation: "Explanation For Task",
		Status:      domain.Done,
		UserID:      id,
	}
	handleTest(t, func(t *testing.T) {
		actual, err := repo.Find(id, 0)
		if err != nil {
			t.Errorf(fmt.Sprintf("error: error in Find(%s) method because of %s", id, err))
		}
		assert.Equal(t, 1, len(actual))
	}, data, repo)
}

func handleTest(t *testing.T, f func(t *testing.T), task *domain.Task, repository *TaskRepository) {
	repository.Create(task)
	defer repository.Delete(task)
	f(t)
}
