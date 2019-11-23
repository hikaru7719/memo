package repository

import "github.com/hikaru7719/memo/server/domain"

// NewTagRepository create TagRepository Object
func NewTagRepository() *TagRepository {
	return &TagRepository{}
}

// TagRepository returns ressources about tag
type TagRepository struct{}

// Find gets 20 tags
func (t *TagRepository) Find(startAt int) ([]domain.Tag, error) {
	tags := make([]domain.Tag, 0, 20)
	if result := db.Offset(startAt).Find(&tags); result.Error != nil {
		return tags, result.Error
	}
	return tags, nil
}

// Create inserts new Tag Record
func (t *TagRepository) Create(tag *domain.Tag) (*domain.Tag, error) {
	if result := db.Create(tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

// Delete deletes record matching tag id
func (t *TagRepository) Delete(tag *domain.Tag) error {
	if result := db.Delete(tag); result.Error != nil {
		return result.Error
	}
	return nil
}

// FindByName gets tag matching tag name
func (t *TagRepository) FindByName(name string) (*domain.Tag, error) {
	tag := &domain.Tag{}
	if result := db.Where("name = ?", name).Find(tag); result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}
