package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/sayuen0/go-to-gym/config"
	authmock "github.com/sayuen0/go-to-gym/internal/auth/mock"
	"github.com/sayuen0/go-to-gym/internal/models"
	sessmock "github.com/sayuen0/go-to-gym/internal/session/mock"
	"go.uber.org/zap/zaptest"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_authHandlers_Register(t *testing.T) {
	type args struct {
		body string
	}
	type dependency struct {
		RegisterFunc func(ctx context.Context, user *models.UserCreateRequest) (*models.UserWithToken, error)
		CreateFunc   func(ctx context.Context, sess *models.Session, expires int) (string, error)
	}

	type want struct {
		status int
		body   *models.UserWithToken
	}

	tests := []struct {
		name       string
		args       args
		dependency dependency
		want       want
	}{
		{
			args: args{
				body: `{"name": "John Doe",
"email": "john.doe@example.com",
"password": "password"
}`,
			},
			dependency: dependency{
				RegisterFunc: func(ctx context.Context, user *models.UserCreateRequest) (*models.UserWithToken, error) {
					return &models.UserWithToken{
						User: &models.User{
							UserID: "1",
							Name:   "John Doe",
							Email:  "john.doe@example.com",
						}}, nil
				},
				CreateFunc: func(ctx context.Context, sess *models.Session, expires int) (string, error) {
					return "", nil
				},
			},
			want: want{
				status: http.StatusCreated,
				body: &models.UserWithToken{
					User: &models.User{
						UserID: "1",
						Name:   "John Doe",
						Email:  "john.doe@example.com",
					}},
			},
		},
	}
	for _, tt := range tests {
		gin.SetMode(gin.TestMode)

		t.Run(tt.name, func(t *testing.T) {
			lg := zaptest.NewLogger(t)
			uc := &authmock.UseCaseMock{
				RegisterFunc: tt.dependency.RegisterFunc,
			}
			sessUC := &sessmock.UseCaseMock{
				CreateFunc: tt.dependency.CreateFunc,
			}

			h := NewAuthHandlers(
				&config.Config{},
				lg,
				uc,
				sessUC,
			)

			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)

			r.POST("/", h.Register())
			req, _ := http.NewRequest("POST", "/", strings.NewReader(tt.args.body))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.want.status, w.Code,
				fmt.Sprintf("expected status code %d, but got %d", tt.want.status, w.Code))

			var got *models.UserWithToken
			err := json.Unmarshal(w.Body.Bytes(), &got)
			assert.Equal(t, err, nil)

			assert.Equal(t, got, tt.want.body)

		})
	}
}
