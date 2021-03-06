package image

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args struct {
		imageUrls []string
	}
	tests := []struct {
		name      string
		args      args
		wantPaths []string
		wantErr   bool
	}{
		{
			name:      "test1",
			args:      args{imageUrls: []string{"https://ae01.alicdn.com/kf/HTB1PSPIw4uTBuNkHFNRq6A9qpXau.jpg"}},
			wantPaths: []string{""},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPaths, err := Reduce(tt.args.imageUrls)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPaths, tt.wantPaths) {
				t.Errorf("Reduce() gotPaths = %v, want %v", gotPaths, tt.wantPaths)
			}
		})
	}
}

func TestReduceImage(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{filePath: ""},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReduceImage(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReduceImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReduceImage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
