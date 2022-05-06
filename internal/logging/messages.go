package logging

import "fmt"

func ErrorCannotCreateLogRecord(baseError string, err error) {
	fmt.Printf("%sErrCode: 801. Error: %s.", baseError, err.Error())
}
