package toy

// Repository for Toys Interface
type Repository interface {
	Create(toy *Toy) error
	FindByID(id string) (*Toy, error)
	FindAll() ([]*Toy, error)
}
