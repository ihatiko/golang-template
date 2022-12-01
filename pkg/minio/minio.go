package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

type Config struct {
	User           string
	Password       string
	Host           string
	SSL            bool
	MaxImageSizeMb int64
}

type Client struct {
	client *minio.Client
	cfg    *Config
}

func NewClient(cfg *Config) (*Client, error) {
	client := &Client{cfg: cfg}
	return client, client.connect()
}

func (m *Client) connect() error {
	var err error
	m.client, err = minio.New(m.cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(m.cfg.User, m.cfg.Password, ""),
		Secure: m.cfg.SSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (m *Client) Put(ctx context.Context, bucket, name, contentType string, buffer io.Reader, size int64) error {
	_, err := m.client.PutObject(ctx, bucket, name, buffer, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}