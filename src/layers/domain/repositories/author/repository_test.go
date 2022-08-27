package author

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/UpBonent/news/src/common/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bung = sqlmock.NewResult(1, 1)

func TestRepository_Insert(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type args struct {
		author models.Author
	}
	type mockBehavior func(args args)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "Positive",
			args: args{
				author: models.Author{
					Name:    "Bob",
					Surname: "Seger",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("INSERT INTO authors").
					WithArgs(args.author.Name, args.author.Surname).
					WillReturnResult(bung)
			},
			wantErr: false,
		},
		{
			name: "Empty fields",
			args: args{
				author: models.Author{
					Name:    "",
					Surname: "",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("INSERT INTO authors").
					WithArgs(args.author.Name, args.author.Surname).
					WillReturnError(errors.New("empty fields"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			err = r.Insert(ctx, tt.args.author)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_All(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type mockBehavior func()

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		wantAuthors  []models.Author
		wantErr      bool
	}{
		{
			name: "Positive",
			mockBehavior: func() {
				rows := mock.NewRows([]string{"id", "name", "surname"}).
					AddRow(1, "Bob", "Seger").
					AddRow(2, "Jimi", "Hendrix")

				mock.ExpectQuery("SELECT id, name, surname FROM authors").
					WillReturnRows(rows)
			},
			wantAuthors: []models.Author{
				{
					Id:      1,
					Name:    "Bob",
					Surname: "Seger",
				},
				{
					Id:      2,
					Name:    "Jimi",
					Surname: "Hendrix",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			gotAuthors, err := r.All(ctx)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.wantAuthors, gotAuthors)
			}

		})
	}
}

func TestRepository_Delete(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type args struct {
		author models.Author
	}
	type mockBehavior func(id int)

	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantErr      bool
	}{
		{
			name: "Positive",
			args: args{
				author: models.Author{
					Id: 1,
				},
			},
			mockBehavior: func(id int) {
				mock.ExpectExec("DELETE FROM authors WHERE (.+)").WithArgs(id).WillReturnResult(bung)
			},
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				author: models.Author{
					Id: 1,
				},
			},
			mockBehavior: func(id int) {
				mock.ExpectExec("DELETE FROM authors WHERE (.+)").WithArgs(id).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args.author.Id)

			err = r.Delete(ctx, tt.args.author.Id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_GetByID(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type args struct {
		id int
	}

	type mockBehavior func(id int)

	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantAuthor   models.Author
		wantErr      bool
	}{
		{
			name: "Positive",
			args: args{
				id: 1,
			},
			mockBehavior: func(id int) {
				rows := mock.NewRows([]string{"name", "surname"}).
					AddRow("Bob", "Seger")

				mock.ExpectQuery("SELECT name, surname FROM authors WHERE(.+)").WithArgs(1).
					WillReturnRows(rows)
			},
			wantAuthor: models.Author{
				Name:    "Bob",
				Surname: "Seger",
			},
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				id: 1,
			},
			mockBehavior: func(id int) {
				rows := mock.NewRows([]string{"name", "surname"})

				mock.ExpectQuery("SELECT name, surname FROM authors WHERE(.+)").WithArgs(1).
					WillReturnRows(rows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args.id)

			gotAuthor, err := r.GetByID(ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantAuthor, gotAuthor)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
