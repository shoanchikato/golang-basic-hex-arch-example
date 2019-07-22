package toy

import (
	"github.com/google/uuid"
	"time"
)

// Service Interface for Toy Service
type Service interface {
	CreateToy(toy *Toy) error
	FindToyByID(id string) (*Toy, error)
	FindAllToys() ([]*Toy, error)
}

// NewToyService init function for creating a new toy service
func NewToyService(repo Repository) Service {
	return &toyService{
		repo,
	}
}

// Implements Toy Service Interface
type toyService struct {
	repo Repository
}

// CreateToy, setting uuid id, createAt and update time
func (s *toyService) CreateToy(toy *Toy) error {
	toy.ID = uuid.New().String()
	toy.CreatedAt = time.Now()
	toy.UpdatedAt = time.Now()
	return s.repo.Create(toy)
}

// FindToyById for finding toy by id
func (s *toyService) FindToyByID(id string) (*Toy, error) {
	return s.repo.FindByID(id)
}

// FindAllToys for find all toys
func (s *toyService) FindAllToys() ([]*Toy, error) {
	return s.repo.FindAll()
}
