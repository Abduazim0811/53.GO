package rpc

import (
	"homework/internal/db"
	"homework/internal/models"
	"net/http"
	"strconv"
)

type RPCService struct{}

type GetUserArgs struct {
	ID string
}

type CreateUserArgs struct {
	User models.User
}

type UpdateUserArgs struct {
	User models.User
}

type DeleteUserArgs struct {
	ID string
}

func (s *RPCService) GetUser(r *http.Request, args *GetUserArgs, reply *models.User) error {
	id, err := strconv.Atoi(args.ID)
	if err != nil {
		return err
	}

	user, err := db.GetUser(id)
	if err != nil {
		return err
	}

	*reply = user
	return nil
}

func (s *RPCService) CreateUser(r *http.Request, args *CreateUserArgs, reply *models.User) error {
	err := db.CreateUser(&args.User)
	if err != nil {
		return err
	}

	*reply = args.User
	return nil
}

func (s *RPCService) UpdateUser(r *http.Request, args *UpdateUserArgs, reply *models.User) error {
	err := db.UpdateUser(args.User)
	if err != nil {
		return err
	}

	*reply = args.User
	return nil
}

func (s *RPCService) DeleteUser(r *http.Request, args *DeleteUserArgs, reply *string) error {
	id, err := strconv.Atoi(args.ID)
	if err != nil {
		return err
	}

	err = db.DeleteUser(id)
	if err != nil {
		return err
	}

	*reply = "User deleted"
	return nil
}
