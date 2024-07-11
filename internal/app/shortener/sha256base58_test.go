package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type want struct {
	shortened string
	err       error
}

func TestSha256WithBase58Shortener_Shorten(t *testing.T) {
	tests := []struct {
		number int
		name   string
		url    string
		want   want
	}{
		{
			number: 1,
			name:   "passed right parameters expected not empty result",
			url:    "https://example.com",
			want: want{
				shortened: "N4GNaegGCpG",
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		shortener := NewSha256WithBase58Shortener()

		shorten, _ := shortener.Shorten(tt.url)
		assert.Equal(t, tt.want.shortened, shorten)
	}
}
