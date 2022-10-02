package stupkg

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger
	Info  *log.Logger
	Warning  *log.Logger
	Error  *log.Logger
)

func init() {
	file, error := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if error !=nil {
		log.Fatalln("fail to open error file:", error)
	}
	infoFile, error := os.OpenFile("info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if error !=nil {
		log.Fatalln("fail to open error infoFile:", error)
	}
	Trace = log.New(ioutil.Discard, "Trace:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(infoFile, os.Stdout), "Info:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func TestLog() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("有些事,你得知道")
	Error.Println("你这做事是错的!")
}