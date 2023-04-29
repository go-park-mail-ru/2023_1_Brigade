package image_generation

import (
	"github.com/fogleman/gg"
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strings"
)

func GenerateAvatar(firstCharacterName string) (string, error) {
	firstCharacterName = strings.ToUpper(firstCharacterName)
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	color := color.RGBA{r, g, b, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)

	file, err := os.Create("background.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := png.Encode(file, img); err != nil {
		return "", err
	}

	const S = 1024
	im, err := gg.LoadImage("background.png")
	if err != nil {
		return "", err
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("Go-Mono.ttf", 728); err != nil {
		return "", err
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(firstCharacterName, S/2, S/2, 0.5, 0.5)
	dc.Clip()

	hash := uuid.New().String()

	dc.SavePNG("../../avatars/" + hash + ".png")

	url := "https://technogramm.ru/avatars/" + hash + ".png"
	return url, nil
}
