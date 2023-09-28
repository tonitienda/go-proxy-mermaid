package models

type ServiceProxy struct {
}

type Node struct {
	ID          string
	Name        string
	Url         string
	RequestorID string
	HasError    bool
}
