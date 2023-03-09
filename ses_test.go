package ses

import "testing"

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
		// TODO: Add test cases.
		{"send_1", args{"", "", "send email msg here", "test send email", "<h1>Body HTML here</h1>"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, err := SendEmail(tt.args.To, tt.args.From, tt.args.FromMsg, tt.args.Titulo, tt.args.MsgHTML)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOk != tt.wantOk {
				t.Errorf("SendEmail() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
