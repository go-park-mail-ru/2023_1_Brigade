package image_generation

import (
	"bytes"
	"crypto/rand"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"math/big"
	"os"
	"strings"
)

//func imageToReader(img image.Image) (io.Reader, error) {
//	// Создаем буфер для записи содержимого изображения
//	buf := new(bytes.Buffer)
//
//	// Записываем содержимое изображения в буфер
//	err := jpeg.Encode(buf, img, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	// Возвращаем объект io.Reader
//	return buf, nil
//}

func GenerateAvatar(firstCharacterName string) (io.Reader, error) {
	firstCharacterName = strings.ToUpper(firstCharacterName)
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	rBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return nil, err
	}

	gBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return nil, err
	}

	bBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	if err != nil {
		return nil, err
	}

	color := color.RGBA{uint8(rBig.Uint64()), uint8(gBig.Uint64()), uint8(bBig.Uint64()), 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)

	file, err := os.Create("background.png")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := png.Encode(file, img); err != nil {
		return nil, err
	}

	const S = 1024
	im, err := gg.LoadImage("background.png")
	if err != nil {
		return nil, err
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("../../avatars/Go-Mono.ttf", 728); err != nil {
		return nil, err
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(firstCharacterName, S/2, S/2, 0.5, 0.5)
	dc.Clip()

	buf := new(bytes.Buffer)

	// Записываем содержимое изображения в буфер
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}

	// Возвращаем объект io.Reader
	return buf, nil

	//hash := uuid.New().String()
	//
	//dc.("../../avatars/" + hash + ".png")
	//
	//url := "https://technogramm.ru/avatars/" + hash + ".png"
	//return url, nil
}
