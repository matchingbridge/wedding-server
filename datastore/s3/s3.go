package s3

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION = "ap-northeast-2"
	S3_BUCKET = "wedding-profile"
	S3_ACL    = "private"
)

var Session *session.Session

func ConnectBucket() {
	var err error
	Session, err = session.NewSession(&aws.Config{
		Region:      aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials("AKIAYSPR6RDL2HAXSRCR", "PR1LuyBD3HkKIIYOfQBDTzjLyj6dqJKKOOijGu8l", ""),
	})
	if err != nil {
		log.Fatalf("session.NewSession - err: %v", err)
	}
}

func UploadFile(key string, buffer []byte) error {
	_, err := s3.New(Session).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(S3_BUCKET),
		Key:                aws.String(key),
		ACL:                aws.String(S3_ACL),
		Body:               bytes.NewReader(buffer),
		ContentDisposition: aws.String("attachment"),
		ContentLength:      aws.Int64(int64(len(buffer))),
		ContentType:        aws.String(http.DetectContentType(buffer)),
	})
	return err
}

func ReadFile(key string) ([]byte, error) {
	results, err := s3.New(Session).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DownloadFile(key string) (string, error) {
	request, _ := s3.New(Session).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(key),
	})
	return request.Presign(5 * time.Minute)
}

func DeleteFile(key string) error {
	_, err := s3.New(Session).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(key),
	})
	return err
}

func MakeKey(bucket string, userID string) string {
	return fmt.Sprintf("dev/%s/%s", bucket, userID)
}
