package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_Get(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		shortLink string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get from empty cache",
			fields: fields{
				store: map[string]string{},
			},
			args: args{
				shortLink: "shortLink",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Short link not found",
			fields: fields{
				store: map[string]string{
					"shortLink": "longLink",
				},
			},
			args: args{
				shortLink: "shortLinkNotFound",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Short link found",
			fields: fields{
				store: map[string]string{
					"shortLink": "longLink",
				},
			},
			args: args{
				shortLink: "shortLink",
			},
			want:    "longLink",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				store: tt.fields.store,
			}
			got, err := c.Get(tt.args.shortLink)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "return test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			assert.NotNil(t, got)
		})
	}
}

func TestCache_Put(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		shortLink string
		link      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "put not existing long link",
			fields: fields{
				store: map[string]string{
					"someShortLink": "someLongLink",
				},
			},
			args: args{
				shortLink: "anotherShortLink",
				link:      "anotherLongLink",
			},
			wantErr: false,
		},
		{
			name: "put existing long link",
			fields: fields{
				store: map[string]string{
					"shortLink": "longLink",
				},
			},
			args: args{
				shortLink: "shortLink",
				link:      "longLink",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				store: tt.fields.store,
			}
			if err := c.Put(tt.args.shortLink, tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCache_IsShortLinkExists(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		shortLink string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "short link not existing",
			fields: fields{
				store: map[string]string{
					"someShortLink": "someLongLink",
				},
			},
			args: args{
				shortLink: "anotherShortLink",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "short link existing",
			fields: fields{
				store: map[string]string{
					"shortLink": "longLink",
				},
			},
			args: args{
				shortLink: "shortLink",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				store: tt.fields.store,
			}
			got, err := c.IsShortLinkExists(tt.args.shortLink)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.IsShortLinkExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cache.IsShortLinkExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
