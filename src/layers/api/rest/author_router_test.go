package rest

import (
	"bytes"
	"context"
	"github.com/UpBonent/news/src/common/models"
	"github.com/UpBonent/news/src/common/services"
	mockServices "github.com/UpBonent/news/src/common/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handlerAuthor_all(t *testing.T) {

	type fields struct {
		ctx        context.Context
		way        string
		repository services.AuthorRepository
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.all(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handlerAuthor_create(t *testing.T) {
	ctx := context.Background()
	type mockBehavior func(s *mockServices.MockAuthorRepository, author models.Author)

	tests := []struct {
		name               string
		inputBody          string
		inputAuthor        models.Author
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "Positive",
			inputBody: `{"name": "Bob", "surname": "Seger"}`,
			inputAuthor: models.Author{
				Name:    "Bob",
				Surname: "Seger",
			},
			mockBehavior: func(s *mockServices.MockAuthorRepository, author models.Author) {
				s.EXPECT().Insert(ctx, author).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   "yeah, author has been created",
		},
		{
			name:        "Empty fields",
			inputBody:   `{"name": "", "surname": ""}`,
			inputAuthor: models.Author{},
			mockBehavior: func(s *mockServices.MockAuthorRepository, author models.Author) {
				s.EXPECT().Insert(ctx, author).Return(errors.New("author's fields is empty"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "author's fields is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Init
			controller := gomock.NewController(t)
			defer controller.Finish()

			repo := mockServices.NewMockAuthorRepository(controller)
			tt.mockBehavior(repo, tt.inputAuthor)

			handler := handlerAuthor{
				ctx:        ctx,
				way:        "/create",
				repository: repo,
			}
			//Test Server
			e := echo.New()
			e.POST("/create", handler.create)
			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", wayToCreate, bytes.NewBufferString(tt.inputBody))
			//Perform Request
			e.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
		})
	}
}

func Test_handlerAuthor_delete(t *testing.T) {
	type fields struct {
		ctx        context.Context
		way        string
		repository services.AuthorRepository
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func Test_handlerAuthor_articlesByAuthor(t *testing.T) {
	type fields struct {
		ctx        context.Context
		way        string
		repository services.AuthorRepository
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
