package image_generation

import (
	"crypto/rand"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"math/big"
	"os"
	"strings"
)

func GenerateAvatar(firstCharacterName string) error {
	firstCharacterName = strings.ToUpper(firstCharacterName)
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	rBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return err
	}

	gBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return err
	}

	bBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return err
	}

	color := color.RGBA{uint8(rBig.Uint64()), uint8(gBig.Uint64()), uint8(bBig.Uint64()), 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)

	file, err := os.Create("background.png")
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	const S = 1024
	im, err := gg.LoadImage("background.png")
	if err != nil {
		return err
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("str.ttf", 728); err != nil {
		return err
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(firstCharacterName, S/2, S/2, 0.5, 0.5)
	dc.Clip()

	err = dc.SavePNG("background.png")
	if err != nil {
		return err
	}

	return nil
}
