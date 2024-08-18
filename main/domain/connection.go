package domain

type Connection interface {
	GetID() string
	GetName() string
	SetID(id string)
	SetName(name string)
}

type BaseConnection struct {
	ID   string
	Name string
}

func (bc *BaseConnection) GetID() string {
	return bc.ID
}

func (bc *BaseConnection) GetName() string {
	return bc.Name
}

func (bc *BaseConnection) SetID(id string) {
	bc.ID = id
}

func (bc *BaseConnection) SetName(name string) {
	bc.Name = name
}
