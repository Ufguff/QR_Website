package qr

import (
	"image/color"

	"github.com/g4s8/hexcolor"
	"github.com/skip2/go-qrcode"
	"github.com/ufguff/types"
)

type Store struct {
}

func CreateQR() *types.QR {
	return &types.QR{}
}

func (s *Store) GetQRImage(qr types.QR) error {
	var cBack, cFore color.Color
	cBack, errBack := hexcolor.Parse(qr.BackgroundColor)
	cFore, errFore := hexcolor.Parse(qr.ForegroundColor)

	if errBack != nil {
		return errBack
	}

	if errFore != nil {
		return errFore
	}

	err := qrcode.WriteColorFile(qr.Url, qrcode.Highest, qr.Size, cBack, cFore, "./static/qr.png")

	if err != nil {
		return err
	}
	return nil
}
