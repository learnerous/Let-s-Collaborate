package logutil

import (
	// "errors"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	// "io"
	// "os"
	// "path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	// "time"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Logger() *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}
	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return log.WithField("file", filename).WithField("function", fn)
}

// Call this function to initialize a logger that writes to a file
func InitLocalLogger() {
	//TODO: setup log level in configuration
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	path := "./log"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Cannot create log reposiroty: %v", err)
		}
	}
	timeTxt := time.Now().Format("20060102-150405")
	logFileName := "./log/" + filepath.Base(os.Args[0]) + "-" + timeTxt + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Failed to log to file, using default stderr")
	} else {
		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
		gin.DefaultWriter = mw
	}
	exPath := ""
	exPath = filepath.Dir(exPath)
	Logger().Infof("Log file:%s %s", exPath, logFileName)
	Logger().Infof("Go version: %s", runtime.Version())
}

func InitLogger() {
	//TODO: setup log level in configuration
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
	Logger().Infof("Go version: %s", runtime.Version())
}

var (
	//colorReset  = "\033[0m"
	colorRed   = "" // does not look nice in test.out "\033[31m"
	colorGreen = "" //\033[32m"
	//colorYellow = "\033[33m"
	//colorBlue   = "\033[34m"
	//colorPurple = "\033[35m"
	//colorCyan   = "\033[36m"
	//colorWhite  = "\033[37m"
)

func LogTestSuccessGeneric(t *testing.T, successMessage string) {
	success := fmt.Sprintf("%s%s%s", colorGreen, "OK    :", successMessage)
	t.Logf(success)

}

func LogTestSuccessExpected(t *testing.T, successMessage string, resultAsExpected string) {
	success := fmt.Sprintf("%s%s%s%s%s", colorGreen, "OK    :", successMessage, ". Got:", resultAsExpected)
	t.Logf(success)

}

func LogTestErrorGeneric(t *testing.T, errMessage string) {
	error := fmt.Sprintf("%s%s%s", colorRed, "FAILED:", errMessage)
	t.Errorf(error)
}

func LogTestErrorWantVsExpected(t *testing.T, errMessage string, wanted string, got string) {
	error := fmt.Sprintf("%s%s%s%s%s%s%s", colorRed, "FAILED:", errMessage, ". Wanted: ", wanted, " Got:", got)
	t.Errorf(error)
}
