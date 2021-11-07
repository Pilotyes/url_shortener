package apiserver

import "testing"

func Test_isUrl(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid url",
			args: args{
				str: "invalid url",
			},
			want: false,
		},
		{
			name: "invalid url",
			args: args{
				str: "http://invalid url",
			},
			want: false,
		},
		{
			name: "invalid url",
			args: args{
				str: "http://invalid url",
			},
			want: false,
		},
		{
			name: "valid url",
			args: args{
				str: "http://validurl:123",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUrl(tt.args.str); got != tt.want {
				t.Errorf("isUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
