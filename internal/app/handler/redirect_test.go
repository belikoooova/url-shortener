package handler

import (
	"errors"
	mock "github.com/belikoooova/url-shortener/internal/app/mock"
	"github.com/belikoooova/url-shortener/internal/app/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type redirectWant struct {
	code           int
	locationHeader string
}

func TestRedirectHandler_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)

	tests := []struct {
		number   int
		name     string
		method   string
		shortUrl string
		storage  *mock.MockStorage
		want     redirectWant
	}{
		{
			number:   1,
			name:     "passed right parameters expected correct answer",
			method:   http.MethodGet,
			shortUrl: "/good",
			storage:  mockStorage,
			want: redirectWant{
				code:           http.StatusTemporaryRedirect,
				locationHeader: "https://example.com",
			},
		},
		{
			number:   2,
			name:     "passed unknown url expected status not found",
			method:   http.MethodGet,
			shortUrl: "/unknown",
			storage:  mockStorage,
			want: redirectWant{
				code:           http.StatusNotFound,
				locationHeader: "",
			},
		},
		{
			number:   3,
			name:     "passed wrong method expected status method not allowed",
			method:   http.MethodPost,
			shortUrl: "/good",
			storage:  mockStorage,
			want: redirectWant{
				code:           http.StatusMethodNotAllowed,
				locationHeader: "",
			},
		},
	}

	for _, tt := range tests {
		switch tt.number {
		case 1:
			tt.storage.EXPECT().FindByID("good").Return(&model.URL{ID: "good", OriginalURL: "https://example.com", ShortURL: "http://localhost:8080/good"}, nil)
		case 2:
			tt.storage.EXPECT().FindByID("unknown").Return(&model.URL{ID: "unknown", OriginalURL: "", ShortURL: ""}, errors.New("url does not exist"))
		}

		handler := NewRedirectHandler(tt.storage)

		req, err := http.NewRequest(tt.method, tt.shortUrl, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, tt.want.code, rec.Code, "Failed test '%s'", tt.name)
		assert.Equal(t, tt.want.locationHeader, rec.Header().Get("Location"), "Failed test '%s'", tt.name)
	}
}
