package ports

import "github.com/agustin-carnevale/hexagonal-go/internal/user/domain"

// UserOutputPort defines the behavior the application expects from repositories (output port).
// Other style/options for calling this could be: 'ForQueryingUser' or 'ForLoadingUser' (like a For-Verb-Noun style)
type UserOutputPort interface {
	GetByID(id string) (*domain.User, error)
}
