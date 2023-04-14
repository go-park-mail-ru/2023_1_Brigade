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
	file, err := os.Open("../../background.png")
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

	//// Сохраняем изображение в файл
	//file, err = os.Create("image_with_image_centered.png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()

	//err = png.Encode(file, img)
	//if err != nil {
	//	return "", err
	//}

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

	//_, err = io.Copy(fileOnDisk, file)
	//if err != nil {
	//	return "", err
	//}

	url := "https://technogramm.ru/avatars/" + hash + ".png"

	return url, nil
}

func GenerateGroupAvatar() string {

	url := "https://technogramm.ru/avatars/group_avatar.png"
	return url

	////28,28,36
	//// Создаем новое изображение размером 400x400 пикселей
	//img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	//
	//// Генерируем случайный цвет для фона
	//r := uint8(rand.Intn(256))
	//g := uint8(rand.Intn(256))
	//b := uint8(rand.Intn(256))
	//color := color.RGBA{r, g, b, 255}
	//
	//// Заполняем изображение цветом
	//draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	//
	//// Загружаем изображение из файла
	////file, err := os.Open("../../background_group.png")
	//file, err := os.Open("background_group.png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()
	//
	//img2, _, err := image.Decode(file)
	//if err != nil {
	//	return "", err
	//}
	//img3 := image.NewRGBA(img2.Bounds())
	//img2 = img3.
	//// Вычисляем смещение для расположения второго изображения по центру
	//offset := image.Point{
	//	X: (img.Bounds().Dx() - img2.Bounds().Dx()) / 2,
	//	Y: (img.Bounds().Dy() - img2.Bounds().Dy()) / 2,
	//}
	//
	//for x := img3.Bounds().Min.X; x < img3.Bounds().Max.X; x++ {
	//	for y := img3.Bounds().Min.Y; y < img3.Bounds().Max.Y; y++ {
	//		pixel := img3.At(x, y)
	//		r, g, b, _ := pixel.RGBA()
	//		if r > 1 || g > 1 || b > 1 {
	//			img3.Set(x, y, color)
	//		}
	//	}
	//}
	//
	//// Рисуем изображение на основном изображении
	//draw.Draw(img, img2.Bounds().Add(offset), img2, image.Point{}, draw.Src)
	//
	//// Сохраняем изображение в файл
	//file, err = os.Create("image_with_group.png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()
	//
	//err = png.Encode(file, img)
	//if err != nil {
	//	return "", err
	//}
	//
	//return "", nil

	//hash := uuid.New().String()
	//fileOnDisk, err := os.Create("../../avatars/" + hash + ".png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()
	//
	//err = png.Encode(fileOnDisk, img)
	//if err != nil {
	//	return "", err
	//}

	//_, err = io.Copy(fileOnDisk, file)
	//if err != nil {
	//	return "", err
	//}

	//url := "https://technogramm.ru/avatars/" + hash + ".png"

	//return url, nil

	//// Создаем новое изображение размером 400x400 пикселей
	//img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	//
	//// Генерируем случайный цвет для фона
	//r := uint8(rand.Intn(256))
	//g := uint8(rand.Intn(256))
	//b := uint8(rand.Intn(256))
	//color := color.RGBA{r, g, b, 255}
	//
	//// Заполняем изображение цветом
	//draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	//
	//// Загружаем изображение из файла
	////file, err := os.Open("../../background_group.png")
	//file, err := os.Open("background_group.png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()
	//
	//img2, _, err := image.Decode(file)
	//if err != nil {
	//	return "", err
	//}
	//
	//// Вычисляем смещение для расположения второго изображения по центру
	//offset := image.Point{
	//	X: (img.Bounds().Dx() - img2.Bounds().Dx()) / 2,
	//	Y: (img.Bounds().Dy() - img2.Bounds().Dy()) / 2,
	//}
	//
	//// Рисуем изображение на основном изображении
	//draw.Draw(img, img2.Bounds().Add(offset), img2, image.Point{}, draw.Src)
	//
	////img := gocv.IMRead("cut_image.png", gocv.IMReadColor)
	////mask := gocv.IMRead("mask.png", gocv.IMReadGrayScale)
	//
	////cutMask := gocv.NewMat()
	////gocv.InRange(mask, gocv.Scalar{0, 0, 0}, gocv.Scalar{0, 0, 0}, &cutMask)
	////
	////// выборка вырезанных пикселей из исходного изображения
	////cutPixels := gocv.NewMat()
	////gocv.BitwiseAnd(img, img, &cutPixels, cutMask)
	////
	////// покраска вырезанных пикселей в красный цвет
	////cutPixels.SetTo(gocv.Scalar{0, 0, 255}, cutMask)
	////
	////// сохранение результата
	////gocv.IMWrite("result.png", img)
	//
	////cutPixels := img.Region(mask)
	////
	////// покраска вырезанных пикселей в красный цвет
	////cutPixels.Set(gocv.Scalar{0, 0, 255}, mask)
	////
	////// сохранение результата
	////gocv.IMWrite("result.png", img)
	////for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
	////	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
	////		//pixel := img.At(x, y)
	////		//r, g, b, _ := pixel.RGBA()
	////		//if r == 0 && g == 0 && b == 0 {
	////		//	img.Set(x, y, color)
	////		//}
	////		//pixel := img.At(x, y)
	////		//r, g, b, _ := pixel.RGBA()
	////		//if r < 240 && g < 240 && b < 240 {
	////		//img.Set(x, y, color)
	////		//}
	////	}
	////}
	//
	//// Сохраняем изображение в файл
	//file, err = os.Create("image_background_group.png")
	//if err != nil {
	//	return "", err
	//}
	//defer file.Close()
	//
	//err = png.Encode(file, img)
	//if err != nil {
	//	return "", err
	//}
	//
	////hash := uuid.New().String()
	////fileOnDisk, err := os.Create("../../avatars/" + hash + ".png")
	////if err != nil {
	////	return "", err
	////}
	////defer file.Close()
	////
	////err = png.Encode(fileOnDisk, img)
	////if err != nil {
	////	return "", err
	////}
	////
	////url := "https://technogramm.ru/avatars/" + hash + ".png"
	//
	////return url, nil
	//
	//return "", nil
}
