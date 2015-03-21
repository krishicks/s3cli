package client

import (
	"io"
	"net/http"
	"os"

	"github.com/rlmcpherson/s3gof3r"
)

type client struct {
	bucket *s3gof3r.Bucket
	conf   *s3gof3r.Config
}

func New(config Config) (client, error) {
	conf := new(s3gof3r.Config)
	*conf = *s3gof3r.DefaultConfig
	conf.Concurrency = 10
	conf.PartSize = 20971520
	conf.Md5Check = true

	// path style causes "The request signature we calculated does not match..."
	// conf.PathStyle = true

	conf.Scheme = config.Scheme()

	keys := s3gof3r.Keys{
		AccessKey: config.AccessKeyID,
		SecretKey: config.SecretAccessKey,
	}

	bucket := s3gof3r.New(config.HostWithPort(), keys).Bucket(config.BucketName)

	return client{bucket, conf}, nil
}

func (c client) Put(key string, file *os.File) error {
	header := make(http.Header)
	header.Set("x-amz-acl", "bucket-owner-full-control")

	uploader, err := c.bucket.PutWriter(key, header, c.conf)
	if err != nil {
		return err
	}

	_, err = io.Copy(uploader, file)
	if err != nil {
		return err
	}

	return uploader.Close()
}
