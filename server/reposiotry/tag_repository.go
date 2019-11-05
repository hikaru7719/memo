package repository

import "github.com/hikaru7719/memo/server/domain"

func NewTagRepository() *TagRepository {
	return &TagRepository{}
}

type TagRepository struct{}

func (t *TagRepository) Find(userID string) ([]domain.Tag, error) {
	tags := make([]domain.Tag, 0, 20)
	if result := db.Where("user_id = ? ", userID).Find(&tags); result.Error != nil {
		return tags, result.Error
	}
	return tags, nil
}

func (t *TagRepository) Create(tag *domain.Tag) (*domain.Tag, error) {
	if result := db.Create(tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

func (t *TagRepository) Delete(tag *domain.Tag) error {
	if result := db.Delete(tag); result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TagRepository) FindByName(userID string, name string) (*domain.Tag, error) {
	tag := &domain.Tag{}
	if result := db.Where("user_id = ? and name = ?", userID, name).Find(tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}
