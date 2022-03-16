package logo

import (
	"bufio"
	"fmt"
	"os"
)

var GOPATH string = os.Getenv("GOPATH")

//LogoFile 开始时显示的Logo
var LogoFile string = fmt.Sprintf("%s/bin/arp-spoofing/logo/logo.txt", GOPATH)

//Show 显示Logo
func Show(logofile string) {
	file, err := os.Open(logofile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
