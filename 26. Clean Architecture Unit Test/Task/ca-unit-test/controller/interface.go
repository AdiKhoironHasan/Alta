package controller

type User interface {
	Create(user User) error
}

type Business interface {
	Create(user User) error
}
