package mongodb

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/shahzaibaziz/user_operations/db"
	domain "github.com/shahzaibaziz/user_operations/models"
	"github.com/thanhpk/randstr"
)

func Test_client_DeleteUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := NewClient(db.Option{})

	user := getUser()
	_, _ = cli.SaveUser(context.TODO(), user)

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
			name:    "success case of delete user",
			args:    args{ctx: context.Background(), id: user.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cli.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_FindUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := NewClient(db.Option{})

	user := getUser()
	dbUser, _ := cli.SaveUser(context.TODO(), user)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name:    "success case of find user",
			args:    args{ctx: context.Background(), id: user.ID},
			want:    dbUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cli.FindUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_FindUserByFilter(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := NewClient(db.Option{})

	user := getUser()
	dbUser, _ := cli.SaveUser(context.TODO(), user)
	type args struct {
		ctx    context.Context
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name:    "success case of find user by filter",
			args:    args{ctx: context.Background(), filter: map[string]interface{}{nameRef: user.Name}},
			wantErr: false,
			want:    dbUser,
		},
		{
			name:    "success case of find user by filter",
			args:    args{ctx: context.Background(), filter: map[string]interface{}{emailRef: user.Email}},
			wantErr: false,
			want:    dbUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cli.FindUserByFilter(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetUsers(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := NewClient(db.Option{})

	user := getUser()
	dbUser, _ := cli.SaveUser(context.TODO(), user)
	users := make([]*domain.User, 0)
	users = append(users, dbUser)
	type args struct {
		ctx    context.Context
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "success case of get Users",
			args: args{
				ctx:    context.Background(),
				filter: map[string]interface{}{nameRef: user.Name},
			},
			want:    users,
			wantErr: false,
		},
		{
			name: "fail case of get Users",
			args: args{
				ctx:    context.Background(),
				filter: map[string]interface{}{nameRef: "wrong name"},
			},
			want:    users,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cli.GetUsers(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr && len(got) == 0 {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_SaveUser(t *testing.T) {
	os.Setenv("MONGO_DB_HOSTS", "mongodb_dev")
	cli, _ := NewClient(db.Option{})

	user := getUser()

	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "success case of save user",
			args: args{
				ctx:  context.Background(),
				user: user,
			},
			wantErr: false,
			want:    user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cli.SaveUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
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
