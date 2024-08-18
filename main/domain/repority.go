package domain

type ConnectionRepo interface {
	GetConnection() Connection
	CreateConnection() Connection
}
