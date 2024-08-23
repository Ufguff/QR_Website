package types

import (
	_ "image/png" // ?
)

type QR struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
	// Image
	// color

}

type ICreateQr interface {
	GetQRImage(string, int) error
}
