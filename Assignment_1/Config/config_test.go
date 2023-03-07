package Config

import (
	"testing"
)

func TestApp_Inc(t *testing.T) {
	a := &App{
		Mp:   make(map[string]int),
		Size: 0,
	}
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantdata int
	}{
		{
			name: "T01",
			args: args{
				id: "first",
			},
			wantdata: 1,
		},
		{
			name: "T02",
			args: args{
				id: "first",
			},
			wantdata: 2,
		},
		{
			name: "T03",
			args: args{
				id: "second",
			},
			wantdata: 1,
		},
	}
	for _, tt := range tests {
		a.Inc(tt.args.id)
		if a.V[a.Mp[tt.args.id]].Views != tt.wantdata {
			t.Error("Error in incrementing views")
		}
	}
}

func TestApp_Get(t *testing.T) {
	a := &App{
		Mp:   make(map[string]int),
		Size: 0,
	}
	a.Inc("first")
	a.Inc("first")
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "T01",
			args: args{
				id: "first",
			},
			want: 2,
		},
		{
			name: "T02",
			args: args{
				id: "second",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := a.Get(tt.args.id); got != tt.want {
			t.Errorf("App.Get() = %v, want %v", got, tt.want)
		}

	}
}
