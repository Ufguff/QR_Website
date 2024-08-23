package qr

import (
	"github.com/skip2/go-qrcode"
	"github.com/ufguff/types"
)

type Store struct {
}

func CreateQR() *types.QR {
	return &types.QR{}
}

func (s *Store) GetQRImage(url string, size int) error {
	qrCode, _ := qrcode.New(url, qrcode.Medium)

	err := qrCode.WriteFile(size, "./static/qr.png")

	if err != nil {
		return err
	}
	return nil
}
