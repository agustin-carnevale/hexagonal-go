package usecases

import (
	"github.com/agustin-carnevale/hexagonal-go/internal/user/domain"
	driven_ports "github.com/agustin-carnevale/hexagonal-go/internal/user/ports/drivens"
	driver_ports "github.com/agustin-carnevale/hexagonal-go/internal/user/ports/drivers"
)

// UserService is the implementation of application logic (use cases).
// Application logic struct (often named with a "Service" suffix)
// Some teams could create one struct per useCase like: 'GetUserUseCase' struct
type UserService struct {
	repo driven_ports.UserOutputPort
}

// NewUserService creates a new use case handler that conforms to the input port.
// Is a Constructor and returns an interface (the input port)
func NewUserService(repo driven_ports.UserOutputPort) driver_ports.UserInputPort {
	return &UserService{repo: repo}
}

// GetUser retrieves a user using the output port.
// Is an application method (implements the interface which is a input port)
func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

//alternative
// type GetUserUseCase struct {
// 	repo driven_ports.UserOutputPort
// }

// func (uc *GetUserUseCase) Execute(id string) (*domain.User, error) {
// 	return uc.repo.GetByID(id)
// }
