package database

type DatastoreClient interface {
	Client() interface{}
	Connect() error
	Close() error
}

type SearchField struct {
	Field string
	Value interface{}
}

type Datastore[T any] interface {
	GetById(string) (*T, error)
	Get(...SearchField) ([]T, error)
	Save(T) error
}
