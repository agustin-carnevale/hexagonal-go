package ports

import "github.com/agustin-carnevale/hexagonal-go/internal/user/domain"

// UserInputPort defines use cases the application exposes (input port).
// Other style/options for calling this interface could be: 'ForUser' (like a For-Noun style)
// Or also in the fashion of a service like 'UserService'
type UserInputPort interface {
	GetUser(id string) (*domain.User, error)
}
