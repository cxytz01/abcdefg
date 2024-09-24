package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileUploaded struct {
	StoreDir string

	filePath string
	ext      string
	base     string

	content []byte
}

func (f *FileUploaded) CheckUploadFile(handle *multipart.FileHeader) error {
	uploadedFile, err := handle.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	f.content, err = io.ReadAll(uploadedFile)
	if err != nil {
		return err
	}

	f.ext = filepath.Ext(handle.Filename)
	f.base = filepath.Base(handle.Filename)

	hashbyte := sha256.Sum256(f.content)
	f.filePath = fmt.Sprintf("%v/%v_%v_%v%v", f.StoreDir, f.base, hex.EncodeToString(hashbyte[:]), time.Now().Unix(), f.ext)

	return nil
}

func (f *FileUploaded) store() {
	err := os.MkdirAll(f.StoreDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	destFile, err := os.Create(f.filePath)
	if err != nil {
		panic(err)
	}

	_, err = destFile.Write(f.content)
	if err != nil {
		panic(err)
	}

	err = destFile.Sync()
	if err != nil {
		panic(err)
	}
}
