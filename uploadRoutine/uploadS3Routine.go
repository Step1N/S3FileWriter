package uploadRoutine

import (
	"fmt"
	"os"

	cl "S3FileWriter/client"
	fw "S3FileWriter/filewriter"
)

func NewUploadChannel(id int, cl *cl.Client, bucket, object string) (w *UploadChannel) {
	return &UploadChannel{
		s3Client: cl,
		bucket:   bucket,
		object:   object,
	}
}

type UploadChannel struct {
	s3Client *cl.Client
	bucket   string
	object   string
}

func (w *UploadChannel) Upload(channel chan string) {
	err := fw.WriteFile(w.s3Client, w.bucket, w.object)
	if err != nil {
		fmt.Println("error while writing file on s3 :", err)
		os.Exit(1)
	}
	channel <- "done"
}
