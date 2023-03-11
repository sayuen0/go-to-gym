package usecase

import (
	"context"
	"errors"
	"github.com/sayuen0/go-to-gym/internal/session/mock"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/models"
)

func Test_sessionUC_CreateSession(t *testing.T) {
	type args struct {
		sess    *models.Session
		expires int
	}
	type dependency struct {
		CreateFunc func(ctx context.Context, sess *models.Session, expires int) (string, error)
	}

	tests := []struct {
		name       string
		args       args
		dependency dependency
		want       string
		wantErr    bool
	}{
		{
			name: "セッションを登録",
			args: args{
				sess:    &models.Session{SessionID: "1", UserID: "1"},
				expires: 100,
			},
			dependency: dependency{
				CreateFunc: func(ctx context.Context, sess *models.Session, expires int) (string, error) {
					return "", nil
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "セッション登録失敗",
			args: args{
				sess:    &models.Session{SessionID: "1", UserID: "1"},
				expires: 100,
			},
			dependency: dependency{
				CreateFunc: func(ctx context.Context, sess *models.Session, expires int) (string, error) {
					return "", errors.New("error")
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			rp := &mock.RepositoryMock{
				CreateFunc: tt.dependency.CreateFunc,
			}
			u := NewSessionUseCase(&config.Config{}, rp)

			got, err := u.Create(ctx, tt.args.sess, tt.args.expires)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result (-got +want):\n%s", diff)
			}

		})
	}
}

func Test_sessionUC_DeleteByID(t *testing.T) {
	type args struct {
		id string
	}
	type dependency struct {
		DeleteByIDFunc func(ctx context.Context, id string) error
	}

	tests := []struct {
		name       string
		args       args
		dependency dependency
		wantErr    bool
	}{
		{
			name: "セッションを削除",
			args: args{
				id: "1",
			},
			dependency: dependency{
				DeleteByIDFunc: func(ctx context.Context, id string) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "セッション削除失敗",
			args: args{
				id: "1",
			},
			dependency: dependency{
				DeleteByIDFunc: func(ctx context.Context, id string) error { return errors.New("error") },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			rp := &mock.RepositoryMock{
				DeleteByIDFunc: tt.dependency.DeleteByIDFunc,
			}
			u := NewSessionUseCase(&config.Config{}, rp)

			err := u.DeleteByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_sessionUC_GetSessionByID(t *testing.T) {
	type args struct {
		id string
	}
	type dependency struct {
		GetByIDFunc func(ctx context.Context, id string) (*models.Session, error)
	}

	tests := []struct {
		name       string
		args       args
		dependency dependency
		want       *models.Session
		wantErr    bool
	}{
		{
			name: "セッションを取得",
			args: args{
				id: "1",
			},
			dependency: dependency{
				GetByIDFunc: func(ctx context.Context, id string) (*models.Session, error) {
					return &models.Session{SessionID: "1", UserID: "1"}, nil
				},
			},
			want:    &models.Session{SessionID: "1", UserID: "1"},
			wantErr: false,
		},
		{
			name: "セッション取得失敗",
			args: args{
				id: "1",
			},
			dependency: dependency{
				GetByIDFunc: func(ctx context.Context, id string) (*models.Session, error) {
					return nil, errors.New("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			rp := &mock.RepositoryMock{
				GetByIDFunc: tt.dependency.GetByIDFunc,
			}
			u := NewSessionUseCase(&config.Config{}, rp)

			got, err := u.GetByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Unexpected result (-got +want):\n%s", diff)
			}
		})
	}
}
