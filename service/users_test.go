package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"

	"github.com/shahzaibaziz/user_operations/db"
	"github.com/shahzaibaziz/user_operations/db/mongodb"
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	domain "github.com/shahzaibaziz/user_operations/models"
)

func Test_service_DeleteUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := mongodb.NewClient(db.Option{})
	s := NewService(&cli)
	ctx := context.Background()
	user := getUser()
	_, _ = s.SaveUser(ctx, user)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success case of save user",
			args:    args{ctx: ctx, id: user.ID},
			wantErr: false,
		},
		{
			name:    "fail case of save user",
			args:    args{ctx: ctx, id: uuid.New().String()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_service_FindUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := mongodb.NewClient(db.Option{})
	s := NewService(&cli)
	ctx := context.Background()
	user := getUser()
	userDB, _ := s.SaveUser(ctx, user)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *genModel.UserResponse
		wantErr bool
	}{
		{
			name:    "success case of find user",
			args:    args{ctx: ctx, id: user.ID},
			wantErr: false,
			want:    userDB,
		},
		{
			name:    "fail case of find user",
			args:    args{ctx: ctx, id: uuid.New().String()},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.FindUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_service_GetUsers(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := mongodb.NewClient(db.Option{})
	s := NewService(&cli)
	ctx := context.Background()
	user := getUser()
	userDB, _ := s.SaveUser(ctx, user)
	users := make([]*genModel.UserResponse, 0)
	users = append(users, userDB)

	type args struct {
		ctx    context.Context
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*genModel.UserResponse
		wantErr bool
	}{
		{
			name:    "success case of get users",
			args:    args{ctx: ctx, filter: map[string]interface{}{"name": user.Name}},
			want:    users,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetUsers(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_service_SaveUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := mongodb.NewClient(db.Option{})
	s := NewService(&cli)
	ctx := context.Background()
	user := getUser()
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success case of save user",
			args:    args{ctx: ctx, user: user},
			wantErr: false,
		},
		{
			name:    "fail case of save user",
			args:    args{ctx: ctx, user: user},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.SaveUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_service_UpdateUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := mongodb.NewClient(db.Option{})
	s := NewService(&cli)
	ctx := context.Background()
	user := getUser()
	userDB, _ := s.SaveUser(ctx, user)

	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *genModel.UserResponse
		wantErr bool
	}{
		{
			name:    "success case of user update",
			args:    args{ctx: ctx, user: user},
			wantErr: false,
			want:    userDB,
		},
		{
			name:    "fail case of user update",
			args:    args{ctx: ctx, user: getUser()},
			wantErr: true,
			want:    userDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func getUser() *domain.User {
	return &domain.User{
		ID:      uuid.New().String(),
		Name:    randstr.String(12),
		Address: "testing address",
		Email:   fmt.Sprintf("%s@case.com", randstr.String(8)),
	}
}
