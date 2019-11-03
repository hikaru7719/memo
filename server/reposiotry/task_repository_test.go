package repository

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hikaru7719/memo/server/entity"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	repo := NewTaskRepository()
	id := fmt.Sprintf("ID%s", uuid.New())
	data := &entity.Done{
		Name:        "TestDoneTask",
		Explanation: "Explanation For Task",
		TagID:       1,
		UserID:      id,
	}
	handleTest(t, func(t *testing.T) {
		actual, err := repo.Find(id)
		if err != nil {
			t.Errorf(fmt.Sprintf("error: error in Find(%s) method because of %s", id, err))
		}
		assert.Equal(t, 1, len(actual))
	}, data, repo)
}

func handleTest(t *testing.T, f func(t *testing.T), done *entity.Done, repository *TaskRepository) {
	repository.Create(done)
	defer repository.Delete(done)
	f(t)
}
