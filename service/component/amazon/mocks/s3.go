package mocks

import (
	"time"

	"github.com/latolukasz/beeorm"
	"github.com/stretchr/testify/mock"

	s3 "github.com/coretrix/hitrix/service/component/amazon/storage"
)

type FakeS3Client struct {
	mock.Mock
}

func (t *FakeS3Client) GetObjectCachedURL(bucket string, object *s3.Object) string {
	return t.Called(bucket, object).String(0)
}

func (t *FakeS3Client) GetObjectSignedURL(bucket string, object *s3.Object, expires time.Duration) string {
	return t.Called(bucket, object, expires).String(0)
}

func (t *FakeS3Client) UploadObjectFromFile(_ *beeorm.Engine, bucket, localFile string) s3.Object {
	return t.Called(bucket, localFile).Get(0).(s3.Object)
}

func (t *FakeS3Client) UploadObjectFromBase64(_ *beeorm.Engine, bucket, content, extension string) s3.Object {
	return t.Called(bucket, content, extension).Get(0).(s3.Object)
}

func (t *FakeS3Client) UploadObjectFromByte(_ *beeorm.Engine, bucket string, byteData []byte, extension string) s3.Object {
	return t.Called(bucket, byteData, extension).Get(0).(s3.Object)
}

func (t *FakeS3Client) UploadImageFromFile(_ *beeorm.Engine, bucket, localFile string) s3.Object {
	return t.Called(bucket, localFile).Get(0).(s3.Object)
}

func (t *FakeS3Client) UploadImageFromBase64(_ *beeorm.Engine, bucket, image, extension string) s3.Object {
	return t.Called(bucket, image, extension).Get(0).(s3.Object)
}

func (t *FakeS3Client) DeleteObject(bucket string, objects ...*s3.Object) bool {
	return t.Called(bucket, objects).Get(0).(bool)
}

func (t *FakeS3Client) GetClient() interface{} {
	return t.Called().Get(0)
}

func (t *FakeS3Client) CreateObjectFromKey(_ *beeorm.Engine, bucket, key string) s3.Object {
	return t.Called(bucket, key).Get(0).(s3.Object)
}

func (t *FakeS3Client) GetBucketName(bucket string) string {
	return t.Called(bucket).Get(0).(string)
}
