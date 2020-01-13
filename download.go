package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Required arguments
var region = flag.String("region", "", "[REQUIRED] AWS Region where your bucket resides")
var baseDir = flag.String("baseDir", "", "[REQUIRED] Directory to copy s3 contents to.")
var bucket = flag.String("bucket", "", "[REQUIRED] S3 Bucket to copy contents from.")
var key = flag.String("key", "", "S3 object key to download")

//Optional arguments
var concurrency = flag.Int("concurrency", 3000, "Number of concurrent connections.")
var partSize = flag.Int("partSize", 100, "Part size of objects")

func main() {
	//Parse command-line arguments
	flag.Parse()
	if len(*baseDir) == 0 || len(*bucket) == 0 || len(*key) == 0 || len(*region) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	DownloadObject(region, baseDir, bucket, key, concurrency, partSize)
}

func DownloadObject(region, baseDir, bucket, key *string, concurrency, partSize *int) {
	//Create S3 session/client
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(*region)},
	)

	if err != nil {
		log.Fatalf("Failed to create a new session. %v", err)
	}

	file, err := os.Create(*baseDir)
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess, func(d *s3manager.Downloader) {
		d.PartSize = *partSize * 1024 * 1024
		d.Concurrency = *concurrency
	})

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(*bucket),
			Key:    aws.String(*key),
		})
	if err != nil {
		log.Fatalf("Unable to download item %q, %v", *key, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
