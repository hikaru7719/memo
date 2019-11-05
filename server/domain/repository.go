package domain

type TaskRepositoy interface {
	Find(string) ([]Task, error)
	Create(*Task) (*Task, error)
	Delete(*Task) error
}

type TagRepository interface {
	Find(string) ([]Tag, error)
	Create(*Tag) (*Tag, error)
	Delete(*Tag) error
	FindByName(string, string) (*Tag, error)
}
