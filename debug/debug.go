package debug

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//Println debug输出
func Println(a ...interface{}) (n int, err error) {
	if true == viper.GetBool("debug") {
		newA := make([]interface{}, 0, len(a)+1)
		newA = append(newA, "[DEBUG]")
		newA = append(newA, a...)
		return fmt.Fprintln(os.Stderr, newA...)
	}
	return 0, nil
}

//Printf 格式化debug输出
func Printf(format string, a ...interface{}) (n int, err error) {
	if true == viper.GetBool("debug") {
		newA := make([]interface{}, 0, len(a)+1)
		newA = append(newA, "[DEBUG]")
		newA = append(newA, a...)
		return fmt.Fprintf(os.Stderr, format, newA...)
	}
	return 0, nil
}
