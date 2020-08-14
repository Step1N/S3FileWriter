package main

import (
	"fmt"
	"os"

	cl "S3FileWriter/client"
	fr "S3FileWriter/fileReader"
	ur "S3FileWriter/uploadRoutine"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_DEFAULT_REGION = "us-west-2"
	U_CHANNEL_COUNT    = 4
)

var uChannel = make([]*ur.UploadChannel, U_CHANNEL_COUNT)

func main() {
	//Read input files
	path := "/inputfile/files/data_2018-02-08 12_04_49 AM.csv"
	fInfo := fr.FileInfo{}
	err := fInfo.ReadFile(path)
	if err != nil {
		fmt.Println("error while reading file :", err)
		os.Exit(1)
	}
	fmt.Println("read file done")

	//Create s3 client
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(AWS_DEFAULT_REGION)})
	s3Client := &cl.Client{
		S3Client:   *svc,
		FileReader: fInfo.FileBytes,
	}
	fmt.Println("created client")
	//Write file on S3
	bucket := "test-kinesis-backup"
	object := "gofileUpload" + "/" + fInfo.FileName
	for i := 0; i < U_CHANNEL_COUNT; i++ {
		ch := make(chan string, 3)
		uChannel[i] = ur.NewUploadChannel(i, s3Client, bucket, object)
		fmt.Println("new upload channel for chan: ", i)
		go uChannel[i].Upload(ch)
		fmt.Println("upload done for chan: ", <-ch)
	}

}
