package gcloud

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"log"

	"cloud.google.com/go/storage"
	"github.com/clickpop/aww-rats-caching-service/env"
	"golang.org/x/image/draw"
	"google.golang.org/api/option"
)

func getStorage() (*storage.Client, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsJSON([]byte(env.GCLOUD_AUTH_JSON)))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func StoreFile(path string, data []byte) {
	log.Println("Checking image cache ", path)
	client, err := getStorage()
	if err != nil {
		log.Println(err)
	}

	obj := client.Bucket("aww-rats").Object(path)
	_, err = obj.NewReader(context.Background())
	switch {
	case err == storage.ErrObjectNotExist:
		log.Println("Caching image ", path)
		resized := resizeImage(data)
		writer := obj.NewWriter(context.Background())
		_, err = writer.Write(resized)
		if err != nil {
			log.Println(err)
		}
		writer.Close()
	case err != nil:
		log.Println(err)
	}
}

func resizeImage(data []byte) []byte {
	buf := bytes.NewBuffer(data)
	src, err := png.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}
	resized := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/4, src.Bounds().Max.Y/4))
	draw.BiLinear.Scale(resized, resized.Rect, src, src.Bounds(), draw.Over, nil)
	png.Encode(buf, resized)
	return buf.Bytes()
}
