package handler

import (
	"errors"
	"github.com/belikoooova/url-shortener/cmd/shortener/config"
	"github.com/belikoooova/url-shortener/internal/app/mock"
	"github.com/belikoooova/url-shortener/internal/app/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type createWant struct {
	code int
	url  model.URL
}

func TestCreateHandler_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	mockShortener := mock.NewMockShortener(ctrl)

	tests := []struct {
		number            int
		name              string
		method            string
		contentTypeHeader string
		body              string
		storage           *mock.MockStorage
		shortener         *mock.MockShortener
		want              createWant
	}{
		{
			number:            1,
			name:              "passed right parameters expected correct answer",
			method:            http.MethodPost,
			contentTypeHeader: "text/plain",
			body:              "http://example.com/long",
			storage:           mockStorage,
			shortener:         mockShortener,
			want: createWant{
				code: http.StatusCreated,
				url: model.URL{
					OriginalURL: "http://example.com/long",
					ShortURL:    "localhost:8080/short",
					ID:          "short",
				},
			},
		},
		{
			number:            2,
			name:              "passed wrong media type expected status unsupported media type",
			method:            http.MethodPost,
			contentTypeHeader: "application/json",
			body:              "http://example.com/long",
			storage:           mockStorage,
			shortener:         mockShortener,
			want: createWant{
				code: http.StatusUnsupportedMediaType,
				url:  model.URL{},
			},
		},
		{
			number:            3,
			name:              "passed bad shortener expected error while shortening and status bad request",
			method:            http.MethodPost,
			contentTypeHeader: "text/plain",
			body:              "http://example.com/long",
			storage:           mockStorage,
			shortener:         mockShortener,
			want: createWant{
				code: http.StatusBadRequest,
				url:  model.URL{},
			},
		},
		{
			number:            4,
			name:              "passed bad storage expected error while saving and status conflict",
			method:            http.MethodPost,
			contentTypeHeader: "text/plain",
			body:              "http://example.com/long",
			storage:           mockStorage,
			shortener:         mockShortener,
			want: createWant{
				code: http.StatusConflict,
				url: model.URL{
					OriginalURL: "http://example.com/long",
					ShortURL:    "localhost:8080/short",
					ID:          "short",
				},
			},
		},
	}

	for _, tt := range tests {
		switch tt.number {
		case 1:
			tt.shortener.EXPECT().Shorten(tt.body).Return(tt.want.url.ID, nil)
			tt.storage.EXPECT().Save(tt.want.url).Return(&tt.want.url, nil)
		case 3:
			tt.shortener.EXPECT().Shorten(tt.body).Return("", errors.New("error while shortening url"))
		case 4:
			tt.shortener.EXPECT().Shorten(tt.body).Return(tt.want.url.ID, nil)
			tt.storage.EXPECT().Save(tt.want.url).Return(nil, errors.New("error while saving url"))
		}

		handler := NewCreateHandler(tt.storage, tt.shortener, config.Config{RedirectServerAddress: "localhost:8080"})

		req, err := http.NewRequest(tt.method, "http://localhost:8080", strings.NewReader(tt.body))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		req.Header.Set("Content-Type", tt.contentTypeHeader)

		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, tt.want.code, rec.Code)
		if tt.want.code == http.StatusOK {
			assert.Equal(t, tt.want.url.ShortURL, rec.Body.String())
		}
	}
}
