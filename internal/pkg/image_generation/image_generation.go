package image_generation

import (
	"github.com/google/uuid"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
)

func GenerateAvatar() (string, error) {
	// Создаем новое изображение размером 400x400 пикселей
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	// Генерируем случайный цвет для фона
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	color := color.RGBA{r, g, b, 255}

	// Заполняем изображение цветом
	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)

	// Загружаем изображение из файла
	file, err := os.Open("avatars/background.png")
	if err != nil {
		return "", err
	}
	defer file.Close()

	img2, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Вычисляем смещение для расположения второго изображения по центру
	offset := image.Point{
		X: (img.Bounds().Dx() - img2.Bounds().Dx()) / 2,
		Y: (img.Bounds().Dy() - img2.Bounds().Dy()) / 2,
	}

	// Рисуем изображение на основном изображении
	draw.Draw(img, img2.Bounds().Add(offset), img2, image.Point{}, draw.Src)

	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			if r == 0 && g == 0 && b == 0 {
				img.Set(x, y, color)
			}
		}
	}

	hash := uuid.New().String()
	fileOnDisk, err := os.Create("../../avatars/" + hash + ".png")
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = png.Encode(fileOnDisk, img)
	if err != nil {
		return "", err
	}

	url := "https://technogramm.ru/avatars/" + hash + ".png"
	return url, nil
}

func GenerateGroupAvatar() string {
	url := "https://technogramm.ru/avatars/group_avatar.png"
	return url
}
