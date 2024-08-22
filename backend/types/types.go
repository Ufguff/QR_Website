package types

import (
	_ "image/png" // ?
)

type QR struct {
	Url  string
	Size int
	// Image
	// color

}

type ICreateQr interface {
	GetQRImage(string, int) error
}
