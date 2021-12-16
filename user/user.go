package user

import (
	"golang.org/x/net/context"

	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedUserServiceServer
}

var users = &Users{
	Users: []*User{
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
	for _, user := range users.Users {
		if user.Id == u.UserID {
			userData = user
			break
		}
	}

	if userData == nil {
		return &User{}, status.Error(5, "UserID does not exists")
	}
	return userData, nil
}

func (s *Server) GetUsersWithIds(ctx context.Context, ids *UserIDList) (*Users, error) {
	var userWithIds *Users
	if len(ids.UserIDList) == 0 {
		return &Users{}, status.Error(3, "UserIDList is empty")
	}
	for _, user := range users.Users {
		for _, id := range ids.UserIDList {
			if user.Id == id {
				userWithIds.Users = append(userWithIds.Users, user)
				break
			}
		}
	}

	if userWithIds == nil {
		return &Users{}, status.Error(5, "Provided UserIDs not found")
	}

	return userWithIds, nil
}
