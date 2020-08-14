package client

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/s3"
)

type Client struct {
	S3Client   s3.S3
	FileReader *bytes.Reader
}
