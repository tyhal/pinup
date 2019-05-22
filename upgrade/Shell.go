package upgrade

import (
	"fmt"
	"strings"
)

func Shell() {}

func ShellCall (in string) {
	cmd := strings.Split(in," ")
	switch cmd[0] {
	case "apk":
		fmt.Println("Alpine pkg manager check")
	case "apt":
		fmt.Println("Ubuntu pkg manager check")
	case "yum":
		fmt.Println("Centos/RHEL pkg manager check")
	case "pip3":
		fmt.Println("Python 3 pkg manager check")
	case "pip":
		fmt.Println("Python pkg manager check")
	case "go":
		fmt.Println("Go pkg manager... no check")
	default:
		fmt.Println("No pkg manager: " + in)
	}
}
