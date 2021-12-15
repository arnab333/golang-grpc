# Golang gRPC

- This is a project where gRPC has been used to create three endpoints `GetUser(id)`, `GetUsers()` & `GetUsersWithIds([]id)`
- Dummy data has been used for the "Users"
- `GetUser(id)` - Get user data based on a `id`
- `GetUsers()` - Get all users data
- `GetUsersWithIds([]id)` - Get user data based on list of ids

## Run

- To run the server - `go run server.go`
- To run the client - `cd client && go run client.go`
