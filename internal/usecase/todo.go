package usecase

type TodoInterface interface {
	Create() error
	Get() error
	Getlist() error
	Update() error
	Delete() error
}
