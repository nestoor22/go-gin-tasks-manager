package repositories

type BaseRepository[Model any, CreateSchema any, UpdateSchema any] interface {
	GetById(id string) (*Model, error)
	Create(inputData CreateSchema) (*Model, error)
	Update(instance Model, inputData UpdateSchema) (*Model, error)
	DeleteById(id string) error
	DeleteObj(instance Model) error
}
