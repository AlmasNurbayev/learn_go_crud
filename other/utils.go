package other

import (
	"encoding/json"
	"path"
	"runtime"
)

func GetLogPath() string {
	_, filename, _, _ := runtime.Caller(0)
	logpath := path.Join(path.Dir(filename), "../logs/log.txt")

	return logpath
}

func GetJWTKey() string {
	_, filename, _, _ := runtime.Caller(0)
	logpath := path.Join(path.Dir(filename), "../logs/log.txt")

	return logpath
}

func ToJSON(obj interface{}) string {
	res, err := json.Marshal(obj)
	if err != nil {
		panic("error with json serialization " + err.Error())
	}
	return string(res)
}
