package db

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/mjiee/grf-gin/app/pkg/conf"
)

func NewOssBucket(cfg *conf.Config) (*oss.Bucket, error) {
	client, err := oss.New(cfg.Oss.Endpoint, cfg.Oss.AccessKeyId, cfg.Oss.AccessKeySecret)
	if err != nil {
		return &oss.Bucket{}, nil
	}

	bucket, err := client.Bucket(cfg.Oss.Bucket)
	return bucket, err
}
