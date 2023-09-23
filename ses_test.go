package ses

import (
	"fmt"
	"os"
	"testing"
)

var (
	//AWS_REGION            = os.Getenv("AWS_REGION")
	//AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	//AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_FROM = os.Getenv("AWS_FROM")

	AWS_TO1 = os.Getenv("AWS_TO1")
)

func TestSendEmail(t *testing.T) {
	type args struct {
		To      string
		From    string
		FromMsg string
		Titulo  string
		MsgHTML string
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantErr bool
	}{
		{"send_1",
			args{
				To:      AWS_TO1,
				From:    AWS_FROM,
				FromMsg: "message",
				Titulo:  "Your Title Here v0.0.5 gses",
				MsgHTML: "<h1>Body HTML here to Send Email with gses v0.0.5</h1>",
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args)
			err := SendEmail(tt.args.To, tt.args.From, tt.args.FromMsg, tt.args.Titulo, tt.args.MsgHTML)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
