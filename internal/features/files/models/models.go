package models

import "io"

type UploadingFile struct {
	ContentType string
	Stream      io.Reader
	Extension   string
	Name        string
	Bucket      string
	Size        int64
}
