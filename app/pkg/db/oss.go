package db

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/mjiee/grf-gin/app/pkg/conf"
)

func NewOssBucket(cfg *conf.Oss) (*oss.Bucket, error) {
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		return &oss.Bucket{}, nil
	}

	bucket, err := client.Bucket(cfg.Bucket)
	return bucket, err
}
