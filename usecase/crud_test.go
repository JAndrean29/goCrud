package usecase

import (
	"context"
	"errors"
	"goCrud/config"
	"goCrud/repository"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

type fakeRepo struct {
	createUserFn func(ctx context.Context, arg repository.CreateUserParams) (repository.User, error)
}

func (f *fakeRepo) CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
	return f.createUserFn(ctx, arg)
}
func (f *fakeRepo) GetAll(ctx context.Context) ([]repository.User, error) { return nil, nil }
func (f *fakeRepo) UpdateUser(ctx context.Context, arg repository.UpdateUserParams) (repository.User, error) {
	return repository.User{}, nil
}
func (f *fakeRepo) DeleteUser(ctx context.Context, id int64) error { return nil }

func Test_crudUsecase_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    repository.CreateUserParams
		repoRes repository.User
		repoErr error
		want    repository.User
		wantErr bool
	}{
		{
			name: "success",
			user: repository.CreateUserParams{
				Name:   pgtype.Text{String: "John Doe", Valid: true},
				Age:    pgtype.Int4{Int32: 34, Valid: true},
				Gender: pgtype.Text{String: "male", Valid: true},
			},
			repoRes: repository.User{ID: 1},
			want:    repository.User{ID: 1},
		},
		{
			name:    "repo fails",
			repoErr: errors.New("db down"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := crudUsecase{
				config: &config.Cfg{}, // empty is fine unless CreateUser reads from it
				crudRepo: &fakeRepo{
					createUserFn: func(ctx context.Context, arg repository.CreateUserParams) (repository.User, error) {
						return tt.repoRes, tt.repoErr
					},
				},
			}

			got, gotErr := c.CreateUser(context.Background(), tt.user)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateUser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateUser() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
