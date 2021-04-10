package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"

	"github.com/gmhafiz/go8/internal/domain/book"
	"github.com/gmhafiz/go8/internal/domain/book/mock"
	"github.com/gmhafiz/go8/internal/models"
	"github.com/gmhafiz/go8/internal/utility/filter"
)

//go:generate mockgen -package mock -source ../../handler.go -destination=../../mock/mock_handler.go

func TestHandler_Create(t *testing.T) {
	testBookRequest := &models.Book{
		Title:         "test01",
		PublishedDate: now.MustParse("2020-02-02"),
		ImageURL: null.String{
			String: "http://example.com/image.png",
			Valid:  true,
		},
		Description: "test01",
	}
	body, err := json.Marshal(testBookRequest)
	assert.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUseCase(ctrl)

	ctx := context.Background()
	var e error

	var ucResp *models.Book
	uc.EXPECT().Create(ctx, testBookRequest).Return(ucResp, e).AnyTimes()

	router := chi.NewRouter()

	val := validator.New()
	h := NewHandler(uc, val)
	RegisterHTTPEndPoints(router, val, uc)

	ww := httptest.NewRecorder()
	rr, _ := http.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))

	h.Create(ww, rr)

	var gotBook book.Res
	err = json.NewDecoder(ww.Body).Decode(&gotBook)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, ww.Code)
}

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUseCase(ctrl)

	ctx := context.Background()
	var e error
	var id int64
	ucResp := &models.Book{
		BookID:        0,
		Title:         "",
		PublishedDate: time.Time{},
		ImageURL:      null.String{},
		Description:   "",
		CreatedAt:     null.Time{},
		UpdatedAt:     null.Time{},
		DeletedAt:     null.Time{},
	}

	uc.EXPECT().Read(ctx, id).Return(ucResp, e).AnyTimes()

	router := chi.NewRouter()

	h := NewHandler(uc, nil)
	RegisterHTTPEndPoints(router, nil, uc)

	ww := httptest.NewRecorder()
	rr, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	assert.NoError(t, err)

	h.Get(ww, rr)

	var gotBook book.Res
	err = json.NewDecoder(ww.Body).Decode(&gotBook)

	assert.NoError(t, err)
	assert.Equal(t, ucResp.BookID, gotBook.BookID)
	assert.Equal(t, ucResp.Description, gotBook.Description.String)
	assert.Equal(t, ucResp.PublishedDate, gotBook.PublishedDate)
	assert.Equal(t, ucResp.ImageURL, gotBook.ImageURL)
}

func TestHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUseCase(ctrl)

	ctx := context.Background()
	var e error
	f := &book.Filter{
		Base: filter.Filter{
			Page:          0,
			Size:          10,
			DisablePaging: false,
			Search:        false,
		},
	}
	var books []*models.Book

	uc.EXPECT().List(ctx, f).Return(books, e).AnyTimes()

	router := chi.NewRouter()

	val := validator.New()
	h := NewHandler(uc, val)
	RegisterHTTPEndPoints(router, val, uc)

	ww := httptest.NewRecorder()
	rr, err := http.NewRequest(http.MethodGet, "/api/v1/books?page=1&size=10", nil)
	assert.NoError(t, err)

	h.List(ww, rr)

	var gotBook []book.Res
	err = json.NewDecoder(ww.Body).Decode(&gotBook)

	assert.NoError(t, err)
}

func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUseCase(ctrl)

	ctx := context.Background()
	var e error
	bookReq := &models.Book{
		Title:         "test01",
		PublishedDate: now.MustParse("2020-02-02"),
		ImageURL: null.String{
			String: "http://example.com/image.png",
			Valid:  true,
		},
		Description: "test01",
	}
	body, err := json.Marshal(bookReq)
	assert.NoError(t, err)
	var book *models.Book

	uc.EXPECT().Update(ctx, bookReq).Return(book, e).AnyTimes()

	router := chi.NewRouter()

	val := validator.New()
	h := NewHandler(uc, val)
	RegisterHTTPEndPoints(router, val, uc)

	ww := httptest.NewRecorder()
	rr, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", bytes.NewBuffer(body))
	assert.NoError(t, err)

	h.Update(ww, rr)

	var gotBook models.Book
	err = json.NewDecoder(ww.Body).Decode(&gotBook)

	assert.NoError(t, err)

}

func TestHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUseCase(ctrl)

	ctx := context.Background()
	var e error
	var id int64

	uc.EXPECT().Delete(ctx, id).Return(e).AnyTimes()

	router := chi.NewRouter()

	val := validator.New()
	h := NewHandler(uc, val)
	RegisterHTTPEndPoints(router, val, uc)

	ww := httptest.NewRecorder()
	rr, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	assert.NoError(t, err)

	h.Delete(ww, rr)

	assert.NoError(t, err)
}
