package memory

import (
	"errors"

	"github.com/agustin-carnevale/hexagonal-go/internal/user/domain"
	driven_ports "github.com/agustin-carnevale/hexagonal-go/internal/user/ports/drivens"
)

// InMemoryUserRepo implements the UserOutputPort using in-memory storage.
type InMemoryUserRepo struct {
	data map[string]*domain.User
}

// New creates a new in-memory repository.
func New() driven_ports.UserOutputPort {
	return &InMemoryUserRepo{
		data: map[string]*domain.User{
			"1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
		},
	}
}

// GetByID retrieves a user by ID from in-memory data.
func (r *InMemoryUserRepo) GetByID(id string) (*domain.User, error) {
	user, ok := r.data[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
