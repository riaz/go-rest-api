package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id") // this helps not exposing db specific details to client
	ErrNotImplemented  = errors.New("not implented")
)

// Comment - a representation of the comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(ctx context.Context, id string) error
}

// Service is the struct on which our logic will be built
type Service struct {
	Store Store
}

// this is a alternative to using constructor in go
// this returns a pointer to service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")

	//ctx = context.WithValue(ctx, "request-id", "unique-string")

	//fmt.Println(ctx.Value("request-id")) // unique string

	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err) // we can use the original error in datadog / grapfana
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
