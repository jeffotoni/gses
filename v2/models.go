package v2

type DataEmail struct {
	// Required at least 1
	ToAddresses []string

	// Required
	From string

	// Required
	FromMsg string

	// Required
	Title string

	// Required
	MsgHTML string

	Charset      string
	BccAddresses []string
	CcAddresses  []string
	Attachments  []Attachment
}

type Attachment struct {
	Data []byte
	Name string
}

func (d *DataEmail) Validate() error {
	switch 0 {
	case len(d.ToAddresses):
		return ErrInvalidTo
	case len(d.From):
		return ErrInvalidFrom
	case len(d.FromMsg):
		return ErrInvalidMessage
	case len(d.Title):
		return ErrInvalidTitle
	case len(d.MsgHTML):
		return ErrInvalidMessageHTML
	}

	return nil
}

type MsgEmailBody struct {
	Status   string `json:"status"`
	Msg      string `json:"msg"`
	Duration string `json:"duration"`
}
