package domain

type ConnectionBuilder struct {
	*BaseConnection
	Type Type
}

type Type string

const (
	type1 Type = "type1"
	type2 Type = "type2"
)

func NewConnectionBuilder(t Type) *ConnectionBuilder {
	return &ConnectionBuilder{Type: t, BaseConnection: &BaseConnection{}}
}

func (cb *ConnectionBuilder) SetID(id string) {
	cb.BaseConnection.ID = id
}

func (cb *ConnectionBuilder) SetType(t Type) {
	cb.Type = t
}
