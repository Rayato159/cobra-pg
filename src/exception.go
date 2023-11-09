package src

import "fmt"

func BuildExceptionMessage(exception string, errMsg string) string {
	return fmt.Sprintf("%s: %s", exception, errMsg)
}
