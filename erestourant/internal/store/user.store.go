package store

import (
	"context"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/google/uuid"
)

type UserStore struct {
	qry *repository.Queries
}

func NewUserStore(qry *repository.Queries) *UserStore {
	return &UserStore{qry: qry}
}

func (s *UserStore) CreateUser(userParams dto.RegisterDto) (repository.User, error) {
	ctx := context.Background()
	return s.qry.CreateUser(ctx, repository.CreateUserParams{
		Email:    userParams.Email,
		Name:     userParams.Name,
		Role:     userParams.Role,
		Password: userParams.Password,
	})
}

func (s *UserStore) GetUserByEmail(email string) (repository.User, error) {
	ctx := context.Background()

	user, err := s.qry.GetUserByEmail(ctx, email)
	// log.Printf("The user is %v user and the error is %v\n", user, err)
	return user, err
}

func (s *UserStore) GetUserById(id string) (repository.User, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetUserById(ctx, uid)
}

func (s *UserStore) GetAllUsers() ([]repository.User, error) {
	ctx := context.Background()
	return s.qry.GetUsers(ctx)
}

func (s *UserStore) DeleteUser(id string) error {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.DeleteUserById(ctx, uid)
}
