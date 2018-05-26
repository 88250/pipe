package stm

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3Adapter provides implementation for the Adapter interface.
type S3Adapter struct {
	Region string
	Bucket string
	ACL    string
	Creds  *credentials.Credentials
}

// Write will create sitemap xml file into the s3.
func (adp *S3Adapter) Write(loc *Location, data []byte) {
	var reader io.Reader = bytes.NewReader(data)

	if GzipPtn.MatchString(loc.Filename()) {
		var writer *io.PipeWriter

		reader, writer = io.Pipe()
		go func() {
			gz := gzip.NewWriter(writer)
			io.Copy(gz, bytes.NewReader(data))

			gz.Close()
			writer.Close()
		}()
	}

	creds := adp.Creds
	if creds == nil {
		creds = credentials.NewEnvCredentials()
	}
	creds.Get()

	sess := session.New(&aws.Config{
		Credentials: creds, Region: &adp.Region})

	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(adp.Bucket),
		Key:    aws.String(loc.PathInPublic()),
		ACL:    aws.String(adp.ACL),
		Body:   reader,
	})

	if err != nil {
		log.Fatal("[F] S3 Upload file Error:", err)
	}
}
