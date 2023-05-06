package usecase

import (
	"context"
	"io"
	"os"
	"project/internal/images"
	"project/internal/pkg/image_generation"
)

type usecase struct {
	imagesRepo images.Repostiory
}

func NewImagesUsecase(imagesRepo images.Repostiory) images.Usecase {
	return &usecase{imagesRepo: imagesRepo}
}

func (u usecase) UploadGeneratedImage(ctx context.Context, bucketName string, filename string, firstCharacterName string) error {
	err := image_generation.GenerateAvatar(firstCharacterName)
	if err != nil {
		return err
	}

	file, err := os.Open("../../avatars/background.png")
	if err != nil {
		return err
	}

	filename = filename + ".png"
	err = u.UploadImage(ctx, file, bucketName, filename)
	if err != nil {
		return err
	}

	return nil
	//firstCharacterName = strings.ToUpper(firstCharacterName)
	//img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	//
	//rBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	//if err != nil {
	//	return err
	//}
	//
	//gBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	//if err != nil {
	//	return err
	//}
	//
	//bBig, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint32))
	//if err != nil {
	//	return err
	//}
	//
	//color := color.RGBA{uint8(rBig.Uint64()), uint8(gBig.Uint64()), uint8(bBig.Uint64()), 255}
	//
	//draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.Point{}, draw.Src)
	//
	//file, err := os.Create("../../avatars/background.png")
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//if err := png.Encode(file, img); err != nil {
	//	return err
	//}
	//
	//const S = 1024
	//im, err := gg.LoadImage("../../avatars/background.png")
	//if err != nil {
	//	return err
	//}
	//
	//dc := gg.NewContext(S, S)
	//dc.SetRGB(1, 1, 1)
	//dc.Clear()
	//dc.SetRGB(1, 1, 1)
	//if err := dc.LoadFontFace("../../avatars/Go-Mono.ttf", 728); err != nil {
	//	return err
	//}
	//
	//dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	//dc.DrawImage(im, 0, 0)
	//dc.DrawStringAnchored(firstCharacterName, S/2, S/2, 0.5, 0.5)
	//dc.Clip()
	//
	//err = dc.SavePNG("../../avatars/background.png")
	//if err != nil {
	//	return err
	//}
	//
	//return nil
	//err = u.UploadImage(ctx, file, bucketName, filename)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
}

func (u usecase) UploadImage(ctx context.Context, file io.Reader, bucketName string, filename string) error {
	err := u.imagesRepo.UploadImage(ctx, file, bucketName, filename)
	return err
}

func (u usecase) GetImage(ctx context.Context, bucketName string, filename string) (string, error) {
	filename = filename + ".png"
	url, err := u.imagesRepo.GetImage(ctx, bucketName, filename)
	if err != nil {
		return "", err
	}

	return url, nil
}

//func (u usecase) LoadImage(ctx context.Context, file multipart.File, filename string, userID uint64) (string, error) {
//	imageUrl, err := u.imagesRepo.LoadImage(context.Background(), file, filename, userID)
//	return imageUrl, err
//}
