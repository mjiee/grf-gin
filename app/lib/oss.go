package lib

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"github.com/mjiee/grf-gin/app/pkg/conf"
)

// Credentials 临时证书文件

type Credentials struct {
	SecurityToken   string
	AccessKeyId     string
	AccessKeySecret string
}

// 提供Ali OSS相关服务
type OssService struct {
	conf   *conf.Oss
	redis  *redis.Client
	bucket *oss.Bucket
}

func NewOssService(cfg *conf.Config, redis *redis.Client, bucket *oss.Bucket) *OssService {
	return &OssService{&cfg.Oss, redis, bucket}
}

// GenStsToken 获取oss临时上传权限
func (s *OssService) GenStsToken() (*sts.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var result = &sts.Credentials{}

	val, err := s.redis.Get(ctx, "oss_sts_token").Result()
	if val == "" || err != nil {
		client, err := sts.NewClientWithAccessKey(s.conf.Region, s.conf.AccessKeyId, s.conf.AccessKeySecret)
		if err != nil {
			return result, err
		}
		request := sts.CreateAssumeRoleRequest()
		request.Scheme = "https"
		request.RoleArn = s.conf.RoleArn
		request.RoleSessionName = s.conf.RoleSessionName

		response, err := client.AssumeRole(request)
		if err != nil {
			return result, err
		}
		result = &response.Credentials

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		data, err := json.Marshal(result)
		if err != nil {
			return result, err
		}
		s.redis.SetEX(ctx, "oss_sts_token", data, 30*time.Minute)

		return result, nil
	}

	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CopyOssObject 拷贝oss目标文件到指定目录
func (s *OssService) CopyOssObject(source, dest string) error {
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.CacheControl("max-age=30758400"),
	}

	_, err := s.bucket.CopyObject(source, dest, options...)
	return err
}

// DeleteOssObject 删除oss目标文件
func (s *OssService) DeleteOssObject(source string) error {
	err := s.bucket.DeleteObject(source)
	return err
}
