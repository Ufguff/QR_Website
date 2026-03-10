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

func (s *Store) GetQRImage(qr types.QR) ([]byte, error) {
	var cBack, cFore color.Color
	cBack, errBack := hexcolor.Parse(qr.BackgroundColor)
	cFore, errFore := hexcolor.Parse(qr.ForegroundColor)

	if errBack != nil {
		return nil, errBack
	}

	if errFore != nil {
		return nil, errFore
	}

	// err := qrcode.WriteColorFile(qr.Url, qrcode.Highest, qr.Size, cBack, cFore, pathQr)
	// if err != nil {
	// 	return err
	// }

	qrRaw, err := qrcode.New(qr.Url, qrcode.Highest)

	if err != nil {
		return nil, err
	}

	qrRaw.BackgroundColor = cBack
	qrRaw.ForegroundColor = cFore

	qrByte, err := qrRaw.PNG(qr.Size)

	if err != nil {
		return nil, err
	}

	return qrByte, nil
}
