package errors

import (
	"fmt"
	"log"
)

const (
	Tag = "fts:ERROR:"
)

func baseErrorLog(message string, err error) error {
	text := fmt.Sprint(Tag, message)
	log.Println(text)
	return fmt.Errorf(text)
}

func GrpcDialError(err error) error {
	return baseErrorLog("cannnot dial fts server", err)
}

func GetFilesError(dir string, err error) error {
	message := fmt.Sprintf("cannot get files in dir:%s", dir)
	return baseErrorLog(message, err)
}

func OpenFileError(filename string, err error) error {
	message := fmt.Sprintf("cannot open file:%s", filename)
	return baseErrorLog(message, err)
}

func GetStreamError(err error) error {
	return baseErrorLog("cannot get file transfer stream", err)
}

func FileReadError(err error) error {
	return baseErrorLog("cannot read file", err)
}

func StreamCloseError(err error) error {
	return baseErrorLog("error occuerd when closing stream", err)
}

func StreamSendError(err error) error {
	return baseErrorLog("cannot send filedata", err)
}
