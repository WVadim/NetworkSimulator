package constants

import "fmt"

var BoldStart string = "\u001B[1m"
var BoldEnd string = "\u001B[0m"

func MakeBoldInt(value int) string {
	return fmt.Sprintf("%s%d%s", BoldStart, value, BoldEnd)
}

func MakeBoldString(value string) string {
	return fmt.Sprintf("%s%s%s", BoldStart, value, BoldEnd)
}
