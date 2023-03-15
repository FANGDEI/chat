package obser

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
)

func (m *Manager) GetURL(name string) string {
	return "http://" + C.Address + "/" + C.Bucket + "/" + name
}

func (m *Manager) Upload(ctx context.Context, objectName, suffix string, data []byte) (minio.UploadInfo, error) {
	objName := strconv.Itoa(int(time.Now().Unix())) + objectName + "." + suffix

	reader := bytes.NewReader(data)
	info, err := m.handler.PutObject(ctx, C.Bucket, objName, reader, int64(len(data)), minio.PutObjectOptions{
		// ContentType: "vedio/mp4",
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

// func (m *Manager) UploadImage(ctx context.Context, objectName string, data []byte) (minio.UploadInfo, error) {
// 	objName := strconv.Itoa(int(time.Now().Unix())) + objectName + ".jpg"

// 	reader := bytes.NewReader(data)
// 	info, err := m.handler.PutObject(ctx, C.Bucket, objName, reader, int64(len(data)), minio.PutObjectOptions{})
// 	if err != nil {
// 		return minio.UploadInfo{}, err
// 	}
// 	return info, err
// }

// func (m *Manager) GetVedioURL(ctx context.Context, fileName string) (*url.URL, error) {
// 	reqParams, expires := make(url.Values), time.Second*60*60*24*7
// 	URL, err := m.handler.PresignedGetObject(ctx, C.Bucket, fileName, expires, reqParams)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return URL, nil
// }

// func (m *Manager) GetImageURL(ctx context.Context, fileName string) (*url.URL, error) {
// 	reqParams, expires := make(url.Values), time.Second*60*60*24*7
// 	URL, err := m.handler.PresignedGetObject(ctx, C.Bucket, fileName, expires, reqParams)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return URL, nil
// }
