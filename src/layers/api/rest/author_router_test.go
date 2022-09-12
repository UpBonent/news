package rest

import (
	"bytes"
	"context"
	"fmt"
	"github.com/UpBonent/news/src/common/services"
	mockServices "github.com/UpBonent/news/src/common/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
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
			h := &HandlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.All(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handlerAuthor_create(t *testing.T) {
	//ctx := context.Background()
	type mockBehavior func(s *mockServices.MockAuthorHandler, e *echo.Echo, c echo.Context)

	tests := []struct {
		name               string
		inputBody          string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "Positive",
			inputBody: `{"name": "Bob", "surname": "Seger"}`,
			mockBehavior: func(s *mockServices.MockAuthorHandler, e *echo.Echo, c echo.Context) {
				s.EXPECT().Register(e)

				s.EXPECT().Create(c).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "yeah, author has been created",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Init
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockHandler := mockServices.NewMockAuthorHandler(controller)

			//Test Server
			e := echo.New()

			e.POST("/create", mockHandler.Create)

			//Test Request
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", wayToCreate, bytes.NewBufferString(tt.inputBody))

			//Perform Request
			contx := e.NewContext(req, w)
			contx.SetPath(wayToCreate)

			fmt.Println("HERE !!!!!!", contx.Response())
			tt.mockBehavior(mockHandler, e, contx)

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
