package sc

import (
	"fmt"
	"os/exec"
)

func StartContract(arg string) (outStr string, err error) {
	// example:
	// cmd := exec.Command("java", "-classpath", "/home/yearning/eclipse-workspace/yjs.jar", "com.yancloud.sc.SCAPI", "startContract", "{\"type\":\"Data\",\"id\":\"656564\"}")
	cmd := exec.Command("java", "-classpath", "yjs.jar", "com.yancloud.sc.SCAPI", "startContract", arg)

	fmt.Println(cmd.Path, cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	outStr = string(out)
	fmt.Print("cmd output:", outStr)
	return
}

func ExecContract(arg string) (outStr string, err error) {
	// example:
	// cmd := exec.Command("java", "-classpath", "/home/yearning/eclipse-workspace/yjs.jar", "com.yancloud.sc.SCAPI", "approveContract", "{\"arg\":\"http://www.baidu.com\",\"contractID\":\"656564\"}")
	cmd := exec.Command(
		"java", "-classpath",
		"yjs.jar",
		"com.yancloud.sc.SCAPI", "approveContract",
		arg)

	fmt.Println(cmd.Path, cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	outStr = string(out)
	fmt.Print("cmd output:", string(out))
	return
}
