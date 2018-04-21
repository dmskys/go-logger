package logger

import "os"
import "fmt"

//	"log"
const (
	//go-logger version
	_VER string = "1.0.3"
)

type LEVEL int32
type UNIT int64
type ROLLTYPE int //dailyRolling ,rollingFile

const _DATEFORMAT = "2006-01-02"

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

const (
	_DAILY    ROLLTYPE = iota
	_ROLLFILE
)

var LogLevel LEVEL = DEBUG
var LogMaxNumber int32 = 10
var LogMaxSize int64 = 10
var LogUnit = MB
var LogDir string

// 判断文件夹是否存在
func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createDirImpl(name string) bool {
	err := os.MkdirAll(name, 0777)
	if err == nil {
		return true
	} else {
		fmt.Println("Error: ", err)
		return false
	}
}

func GetLogger(fileName string) (log *Logger) {
	log = getLogBean()
	if LogDir == ""{
		panic("log Dir is not setting")
	}
	if !DirExists(LogDir) {
		createDirImpl(LogDir)
	}
	log.SetRollingFile(LogDir, fileName, LogMaxNumber, LogMaxSize, LogUnit)
	log.SetLevel(LogLevel)
	return
}