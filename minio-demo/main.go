//go:build example
// +build example

/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2015-2017 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"log"
	"os"

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
	bucket    = getenv("MINIO_BUCKET", "demo")
	//object    = getenv("OBJECT_NAME", "hello.txt")
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY, my-testfile, my-bucketname and
	// my-objectname are dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	// s3Client, err := minio.New("s3.amazonaws.com", &minio.Options{
	// 	Creds:  credentials.NewStaticV4("YOUR-ACCESSKEYID", "YOUR-SECRETACCESSKEY", ""),
	// 	Secure: true,
	// })

	//ctx := context.Background()

	// ---- Client --------------------------------------------------------------
	s3Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
		Region: region, // optional but nice to set
	})

	if err != nil {
		log.Fatalln(err)
	}

	object, err := os.Open("hello.txt")
	if err != nil {
		log.Println("Seems something went wrong here .......")
		log.Fatalln(err)
	}
	defer object.Close()
	objectStat, err := object.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	info, err := s3Client.PutObject(context.Background(), bucket, "hello.txt", object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/text"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uploaded", "demo", " of size: ", info.Size, "Successfully.")
}
