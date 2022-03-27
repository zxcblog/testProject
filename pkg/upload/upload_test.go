package upload

import (
	"new-project/pkg/config"
	"testing"
)

func TestUpload(t *testing.T) {
	client, err := NewClient(&config.Oss{
		Endpoint:        "oss-cn-chengdu.aliyuncs.com",
		AccessKeyID:     "LTAI5tG4p5UuLHA33W4ijBnT",
		AccessKeySecret: "JAnZJDI6KZMDUQGRQfAgPEpyJu0PwJ",
		BucketName:      "bucketts-filestore",
	})

	if err != nil {
		t.Error(err)
		return
	}

	//t.Log("查看客户端", client)
	//t.Log(client.FilePathUpload("a3AKotsqnm1JYm1sUTm3zqEgGXCww0Ys.mp4", "./a3AKotsqnm1JYm1sUTm3zqEgGXCww0Ys.mp4"))

	t.Log(client.DeleteFile("a3AKotsqnm1JYm1sUTm3zqEgGXCww0Ys.mp4"))
}
