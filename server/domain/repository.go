package domain

type TaskRepositoy interface {
	Find(string, int) ([]Task, error)
	Create(*Task) (*Task, error)
	Delete(*Task) error
	CreateAssociation(*Task, *Tag) error
}

type TagRepository interface {
	Find(int) ([]Tag, error)
	Create(*Tag) (*Tag, error)
	Delete(*Tag) error
	FindByName(string) (*Tag, error)
}
