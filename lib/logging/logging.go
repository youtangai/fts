package logging

import (
	"fmt"
	"log"

	pb "github.com/youtangai/dts/proto"
)

const (
	InfoTag = "dts:INFO:"
)

func baseLog(message string) {
	text := fmt.Sprint(InfoTag, message)
	log.Println(text)
}

func NoDirInfo(dir string) {
	message := fmt.Sprintf("create dir because no such dir:%s", dir)
	baseLog(message)
}

func WillStartServerInfo(url string) {
	message := fmt.Sprintf("start grpc server on url:%s", url)
	baseLog(message)
}

func OpenFileInfo(filename string) {
	message := fmt.Sprintf("open file:%s done", filename)
	baseLog(message)
}

func RecvFileData(fileData *pb.FileData) {
	message := fmt.Sprintf("recv fileData %v", fileData)
	baseLog(message)
}
