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
		// export
		// AWS_REGION= AWS_IDENTITY=
		// AWS_ACCESS_KEY_ID= AWS_SECRET_ACCESS_KEY=
		{"send_1", args{"<your-email>", "noreply@<your-email-from>", "send email msg here - é açentuação", "test send email - é açentuação", "<h1>Body HTML here</h1>"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SendEmail(tt.args.To, tt.args.From, tt.args.FromMsg, tt.args.Titulo, tt.args.MsgHTML)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
