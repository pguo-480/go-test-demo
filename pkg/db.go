package pkg

//go:generate mockery --inpackage --name=Db
type Db interface {
	Put(input int) error
}

type db struct{}

func NewDb() Db {
	return &db{}
}

func (d *db) Put(input int) error {
	return nil
}
