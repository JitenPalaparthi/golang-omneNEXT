// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

var (
	endpoint  = getenv("MINIO_ENDPOINT", "localhost:9000") // your in-cluster DNS:port
	accessKey = getenv("MINIO_ACCESS_KEY", "minioadmin")
	secretKey = getenv("MINIO_SECRET_KEY", "minioadmin")
	useSSL    = getenv("MINIO_USE_SSL", "false") == "true"
	region    = getenv("MINIO_REGION", "us-east-1") // not useful but still keep some value
	bucket    = getenv("MINIO_BUCKET", "uploads")
	//object    = getenv("OBJECT_NAME", "hello.txt")
)

func main() {
	// Fiber instance
	app := fiber.New()

	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
		Region: region, // optional but nice to set
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Routes
	app.Post("/upload", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}

		f, err := file.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		info, err := s3Client.PutObject(context.Background(), bucket, file.Filename, f, file.Size, minio.PutObjectOptions{ContentType: "simple/text"})
		if err != nil {
			return err
		}
		log.Println("Uploaded", "demo", " of size: ", info.Size, "Successfully.")

		return nil
		// Save file to root directory:
		//return c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
