package article

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/UpBonent/news/src/common/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepository_Insert(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type args struct {
		art      models.Article
		authorID int
	}
	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantErr      bool
	}{
		{
			name: "Positive",
			args: args{
				art: models.Article{
					Header:      "There's test header",
					Text:        "There's long long long long long test text",
					DatePublish: time.Now(),
				},
				authorID: 1,
			},
			mockBehavior: func(args args) {
				dateCreate := time.Now().Round(time.Minute)
				parseDatePublish, err := time.Parse("02.01.06 15:04", args.art.DatePublish)
				assert.NoError(t, err)

				mock.ExpectExec("INSERT INTO articles").
					WithArgs(args.art.Header, args.art.Text, dateCreate, parseDatePublish, args.authorID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Empty articles fields",
			args: args{
				art: models.Article{
					Header:      "",
					Text:        "",
					DatePublish: "01.01.22 12:00",
				},
				authorID: 1,
			},
			mockBehavior: func(args args) {
				dateCreate := time.Now().Round(time.Minute)
				parseDatePublish, err := time.Parse("02.01.06 15:04", args.art.DatePublish)
				assert.NoError(t, err)

				mock.ExpectExec("INSERT INTO articles").
					WithArgs(args.art.Header, args.art.Text, dateCreate, parseDatePublish, args.authorID).
					WillReturnResult(sqlmock.NewResult(1, 1)).
					WillReturnError(errors.New("insert error"))
			},
			wantErr: true,
		},
		{
			name: "Incorrect date format",
			args: args{
				art: models.Article{
					DatePublish: "1.1.22 12:00",
				},
			},
			mockBehavior: func(args args) {},
			wantErr:      true,
		},
		{
			name: "Empty authors fields",
			args: args{
				art: models.Article{
					Header:      "There's test header",
					Text:        "There's long long long long long test text",
					DatePublish: "01.01.22 12:00",
				},
				authorID: 0,
			},
			mockBehavior: func(args args) {
				dateCreate := time.Now().Round(time.Minute)
				parseDatePublish, err := time.Parse("02.01.06 15:04", args.art.DatePublish)
				assert.NoError(t, err)

				mock.ExpectExec("INSERT INTO articles").
					WithArgs(args.art.Header, args.art.Text, dateCreate, parseDatePublish, args.authorID).
					WillReturnResult(sqlmock.NewResult(1, 1)).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			err = r.CreateNew(ctx, tt.args.art, tt.args.authorID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	mock.ExpectClose()

	err = sqlDB.Close()
	if err != nil {
		panic(errors.Wrap(err, "can't close connection"))
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

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		wantArticles []models.Article
		wantErr      bool
	}{
		{
			name: "Positive",
			mockBehavior: func() {
				datePublishFirst := time.Date(2022, 01, 01, 12, 00, 00, 00, time.UTC)
				datePublishSecond := time.Date(2022, 01, 01, 15, 00, 00, 00, time.UTC)

				rows := mock.NewRows([]string{"id", "header", "text", "date_publish", "id_authors"}).
					AddRow(1, "There's first test header", "There's first test text", datePublishFirst, 1).
					AddRow(2, "There's second test header", "There's second test text", datePublishSecond, 1)

				mock.ExpectQuery("SELECT id, header, text, date_publish, author_id FROM articles").
					WillReturnRows(rows)
			},
			wantArticles: []models.Article{
				{
					Id:          1,
					Header:      "There's first test header",
					Text:        "There's first test text",
					DatePublish: "01.01.22 12:00",
					AuthorID:    1,
				},
				{
					Id:          2,
					Header:      "There's second test header",
					Text:        "There's second test text",
					DatePublish: "01.01.22 15:00",
					AuthorID:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "Empty DB table",
			mockBehavior: func() {
				rows := mock.NewRows([]string{"id", "header", "text", "date_publish", "id_authors"})

				mock.ExpectQuery(all).WillReturnRows(rows)
			},
			wantErr: false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			gotArticles, err := r.GetAll(ctx)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.wantArticles, gotArticles)
			}

		})
	}

	mock.ExpectClose()

	err = sqlDB.Close()
	if err != nil {
		panic(errors.Wrap(err, "can't close connection"))
	}
}

// what about "not found" test case???
func TestRepository_Update(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	r := NewRepository(sqlx.NewDb(sqlDB, "postgres"))

	type args struct {
		existArticle int
		article      models.Article
	}
	type mockBehavior func(args args)

	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantErr      bool
	}{
		{
			name: "Positive",
			args: args{
				existArticle: 1,
				article: models.Article{
					Header:      "There's new test header",
					Text:        "There's new test text",
					DatePublish: "01.01.22 20:00",
				},
			},
			mockBehavior: func(args args) {
				date, err := time.Parse("02.01.06 15:04", args.article.DatePublish)
				if err != nil {
					assert.NoError(t, err)
				}

				mock.ExpectExec("UPDATE articles SET (.+) WHERE (.+)").
					WithArgs(args.article.Header, args.existArticle).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE articles SET (.+) WHERE (.+)").
					WithArgs(args.article.Text, args.existArticle).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE articles SET (.+) WHERE (.+)").
					WithArgs(date, args.existArticle).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Empty fields",
			args: args{
				existArticle: 1,
				article: models.Article{
					Id:          1,
					Header:      "",
					Text:        "",
					DatePublish: "",
				},
			},
			mockBehavior: func(args args) {},
			wantErr:      true,
		},
		{
			name: "Not found",
			args: args{
				existArticle: 2,
				article: models.Article{
					Header:      "There's new test header",
					Text:        "There's new test text",
					DatePublish: "01.01.22 20:00",
				},
			},
			mockBehavior: func(args args) {
				mock.NewRows([]string{"id", "name", "surname"}).AddRow(1, "Bob", "Seger")

				mock.ExpectExec("UPDATE articles SET (.+) WHERE (.+)").
					WithArgs(args.article.Header, args.existArticle).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			err = r.Update(ctx, tt.args.existArticle, tt.args.article)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	mock.ExpectClose()

	err = sqlDB.Close()
	if err != nil {
		panic(errors.Wrap(err, "can't close connection"))
	}
}
