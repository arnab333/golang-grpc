package user

import (
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedUserServiceServer
}

var users = &Users{
	User: []*User{
		{Id: 1, Fname: "User1", City: "City1", Phone: 123, Height: 32.12, Married: false},
		{Id: 2, Fname: "User2", City: "City2", Phone: 1233, Height: 32.14, Married: false},
		{Id: 3, Fname: "User3", City: "City3", Phone: 1234, Height: 32.45, Married: true},
	},
}

func (s *Server) GetUsers(ctx context.Context, e *EmptyParams) (*Users, error) {
	return users, nil
}

func (s *Server) GetUser(ctx context.Context, u *UserID) (*User, error) {
	var userData *User
	for _, user := range users.User {
		if user.Id == u.UserID {
			userData = user
		}
	}
	return userData, nil
}
