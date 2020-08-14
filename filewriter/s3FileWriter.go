package filewriter

import (
	c "S3FileWriter/client"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func WriteFile(cl *c.Client, bucketName, objectName string) error {
	input := &s3.PutObjectInput{
		Body:   cl.FileReader,
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}

	_, err := cl.S3Client.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}
