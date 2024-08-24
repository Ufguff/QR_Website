package types

type QR struct {
	Url             string `json:"url"`
	Size            int    `json:"size"`
	BackgroundColor string `json:"background"`
	ForegroundColor string `json:"foreground"`
}

type ICreateQr interface {
	GetQRImage(QR) error
}
