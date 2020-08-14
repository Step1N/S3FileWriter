package upload

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_ACCESS_KEY_ID     = "ACCESS_KEY"
	AWS_SECRET_ACCESS_KEY = "SECRET__KEY"
	AWS_DEFAULT_REGION    = "us-west-2"
)

func UploadToS3(S3_BUCKET, S3_KEY, fileDir string) error {
	var timeout time.Duration
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(AWS_DEFAULT_REGION)})

	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	defer cancelFn()
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	name := strings.Split(fileDir, "/")
	path := S3_KEY + "/" + name[len(name)-1]
	fmt.Println("upload path : ", path)
	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(path),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	return err
}
