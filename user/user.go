package user

import (
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) GetUsers(ctx context.Context, e *EmptyParams) (*Users, error) {
	return nil, nil
}

func (s *Server) GetUser(ctx context.Context, e *UserID) (*User, error) {
	return nil, nil
}
