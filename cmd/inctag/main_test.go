package main

import "testing"

func Test_updateTag(t *testing.T) {
	type args struct {
		tag  string
		part string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				tag:  "",
				part: "minor",
			},
			wantErr: true,
		},
		{
			name: "error 1",
			args: args{
				tag:  "0",
				part: "minor",
			},
			wantErr: true,
		},
		{
			name: "error 2",
			args: args{
				tag:  "0.0",
				part: "minor",
			},
			wantErr: true,
		},
		{
			name: "error 3",
			args: args{
				tag:  "x0.0.0",
				part: "minor",
			},
			wantErr: true,
		},
		{
			name: "no prefix 1",
			args: args{
				tag:  "0.0.0",
				part: "patch",
			},
			want: "0.0.1",
		},
		{
			name: "no prefix 2",
			args: args{
				tag:  "0.0.0",
				part: "minor",
			},
			want: "0.1.0",
		},
		{
			name: "prefix 1",
			args: args{
				tag:  "v0.0.0",
				part: "patch",
			},
			want: "v0.0.1",
		},
		{
			name: "prefix 2",
			args: args{
				tag:  "golang/v0.0.0",
				part: "patch",
			},
			want: "golang/v0.0.1",
		},
		{
			name: "prefix 3",
			args: args{
				tag:  "golang/v1.2.3",
				part: "minor",
			},
			want: "golang/v1.3.3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateTag(tt.args.tag, tt.args.part)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("updateTag() = %v, want %v", got, tt.want)
			}
		})
	}
}
