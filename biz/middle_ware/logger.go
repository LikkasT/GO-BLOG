package middle_ware

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func LoggerInit(level logrus.Level, log_file_name string) (*logrus.Logger, error) {
	L := logrus.New()
	L.SetLevel(level)
	writer, err := os.OpenFile(log_file_name, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("create file %s failed: %v", log_file_name, err)
	}
	L.SetOutput(io.MultiWriter(writer))
	return L, nil
}

var Logger_level_map = map[string]logrus.Level{
	"Trace": logrus.TraceLevel,
	"Debug": logrus.DebugLevel,
	"Info":  logrus.InfoLevel,
	"Warn":  logrus.WarnLevel,
	"Error": logrus.ErrorLevel,
	"Fatal": logrus.FatalLevel,
	"Panic": logrus.PanicLevel,
}
