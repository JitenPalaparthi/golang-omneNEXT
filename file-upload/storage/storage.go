// type Storage interface {
// 	// Get gets the value for the given key.
// 	// `nil, nil` is returned when the key does not exist
// 	Get(key string) ([]byte, error)

// 	// Set stores the given value for the given key along
// 	// with an expiration value, 0 means no expiration.
// 	// Empty key or value will be ignored without an error.
// 	Set(key string, val []byte, exp time.Duration) error

// 	// Delete deletes the value for the given key.
// 	// It returns no error if the storage does not contain the key,
// 	Delete(key string) error

// 	// Reset resets the storage and delete all keys.
// 	Reset() error

// 	// Close closes the storage and will stop any running garbage
// 	// collectors and open connections.
// 	Close() error
// }

package storage

import (
	"bytes"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	EndPoint, AccessKey, SecretKey, UseSSL, Region, Bucket, ObjectName string
	Size                                                               int64
}

func NewStorage(endpoint, accesKey, secretKey, useSSL, region, bucket, objectName string, size int64) (fiber.Storage, error) {
	return &Storage{endpoint, accesKey, secretKey, useSSL, region, bucket, objectName, size}, nil
}

func (s *Storage) Get(key string) ([]byte, error) {
	return nil, nil
}

func (s *Storage) Set(key string, val []byte, exp time.Duration) error {
	s3Client, err := minio.New(s.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.AccessKey, s.SecretKey, ""),
		Secure: false,
		Region: s.Region, // optional but nice to set
	})

	if err != nil {
		return err
	}
	//s3Client.FPutObject()
	data := bytes.NewReader(val)
	_, err = s3Client.PutObject(context.Background(), s.Bucket, s.ObjectName, data, s.Size, minio.PutObjectOptions{})
	return err
}

func (s *Storage) Delete(key string) error {
	return nil
}

func (s *Storage) Reset() error {
	return nil
}

func (s *Storage) Close() error {
	return nil
}
