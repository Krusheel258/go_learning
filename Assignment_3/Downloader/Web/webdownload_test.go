package Web

import (
	"testing"
)

func Test_webdownload_UrlDownload(t *testing.T) {

	wd := &webdownload{}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "T01",
			args: args{
				url: "http://www.yahoo.com/image_to_read.jpg",
			},
			wantErr: false,
		},
		{
			name: "T02",
			args: args{
				url: "ghgjkh",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {

		_, err := wd.UrlDownload(tt.args.url)
		if (err != nil) != tt.wantErr {
			t.Errorf("webdownload.UrlDownload() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

	}
}
