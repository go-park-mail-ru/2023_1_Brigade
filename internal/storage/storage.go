package main

import (
	"net/http"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func Connect() (*minio.Client, error) {
	return minio.New("localhost:9000", &minio.Options{ // TODO config
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""), // TODO config
		Secure: false,
	})
}

func (s *Server) uploadPhoto(w http.ResponseWriter, r *http.Request) {
	// Убеждаемся, что к нам в ручку идут нужным методом
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID сессии аренды, чтобы знать, в каком контексте это фото
	rentID, err := strconv.Atoi(r.Header.Get(HEADER_RENT_ID))
	if err != nil {
		logrus.Errorf("Can`t get rent id: %v\n", err)
		http.Error(w, "Wrong request!", http.StatusBadRequest)
		return
	}

	// Забираем фото из тела запроса
	src, hdr, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Wrong request!", http.StatusBadRequest)
		return
	}

	// Получаем информацию о сессии аренды
	session, err := s.database.GetRentStatus(rentID)
	if err != nil {
		logrus.Errorf("Can`t get session: %v\n", err)
		http.Error(w, "Can`t upload photo!", http.StatusInternalServerError)
		return
	}

	// Складываем данные в объект, который является своего рода контрактом
	// между хранилищем изображений и нашей бизнес-логикой
	object := models.ImageUnit{
		Payload:     src,
		PayloadSize: hdr.Size,
		User:        session.User,
	}
	defer src.Close()

	// Отправляем фото в хранилище
	img, err := s.storage.UploadFile(r.Context(), object)
	if err != nil {
		logrus.Errorf("Fail update img in image strorage: %v\n", err)
		http.Error(w, "Can`t upload photo!", http.StatusInternalServerError)
		return
	}

	// Добавляем запись в БД с привязкой фото к сессии
	err = s.database.AddImageRecord(img, rentID)
	if err != nil {
		logrus.Errorf("Fail update img in database: %v\n", err)
		http.Error(w, "Can`t upload photo!", http.StatusInternalServerError)
	}
}

// func main() {
// 	endpoint := "your-minio-endpoint"
// 	accessKeyID := "your-access-key-id"
// 	secretAccessKey := "your-secret-access-key"
// 	useSSL := false // set to true if your minio server uses SSL
// 	bucketName := "your-bucket-name"
// 	objectName := "your-object-name"
// 	filePath := "image.jpg"

// 	// Initialize minio client object.
// 	client, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer file.Close()

// 	// Get the file stats
// 	stat, err := file.Stat()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Upload the image file to minio
// 	ctx := context.Background()
// 	uploadInfo, err := client.PutObject(ctx, bucketName, objectName, file, stat.Size(), minio.PutObjectOptions{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	log.Printf("Successfully uploaded %s, info: %#v", objectName, uploadInfo)
// }

// package main

// import (
//  "log"
//  "os"

//  "github.com/minio/minio-go"
// )

// func main() {
//  endpoint := "your-minio-endpoint"
//  accessKeyID := "your-access-key-id"
//  secretAccessKey := "your-secret-access-key"
//  useSSL := false // set to true if your minio server uses SSL
//  bucketName := "your-bucket-name"
//  objectName := "your-object-name"
//  filePath := "path/to/your/image.jpg"

//  // Initialize minio client object
//  client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
//  if err != nil {
//   log.Fatalln(err)
//  }

//  // Open the image file
//  file, err := os.Open(filePath)
//  if err != nil {
//   log.Fatalln(err)
//  }
//  defer file.Close()

//  // Get the file stats
//  stat, err := file.Stat()
//  if err != nil {
//   log.Fatalln(err)
//  }

//  // Upload the image file to minio
//  n, err := client.PutObject(bucketName, objectName, file, stat.Size(), minio.PutObjectOptions{})
//  if err != nil {
//   log.Fatalln(err)
//  }

//  log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
// }
