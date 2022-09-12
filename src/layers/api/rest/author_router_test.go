package rest

import (
	"context"
	"github.com/UpBonent/news/src/common/services"
	"github.com/labstack/echo"
	"reflect"
	"testing"
)

func TestNewHandlerAuthor(t *testing.T) {
	type args struct {
		ctx context.Context
		s   string
		r   services.AuthorRepository
	}
	tests := []struct {
		name string
		args args
		want services.RESTMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlerAuthor(tt.args.ctx, tt.args.s, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlerAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handlerAuthor_Register(t *testing.T) {
	type fields struct {
		ctx        context.Context
		way        string
		repository services.AuthorRepository
	}
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			h.Register(tt.args.e)
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.all(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("all() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.articlesByAuthor(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("articlesByAuthor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handlerAuthor_create(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("create() error = %v, wantErr %v", err, tt.wantErr)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handlerAuthor{
				ctx:        tt.fields.ctx,
				way:        tt.fields.way,
				repository: tt.fields.repository,
			}
			if err := h.delete(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
