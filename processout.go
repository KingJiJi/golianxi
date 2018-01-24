package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {

	//pt-online-schema-change --user=root --password=000000 --host=127.0.0.1 --alter="add column age int;" D=yml_school,t=t_yml --nodrop-old-table --chunk-size=2000 --chunk-time=0.5 --max-lag=2 --check-interval=2 --charset=utf8 --critical-load="Threads_running=64" --max-load="Threads_running=32" --execute
	//cmd := exec.Command("echo", `{"Name": "Bob", "Age": 32}`)
	cmd := exec.Command("/bin/bash", "-c", `pt-online-schema-change --user=root --password=000000 --host=127.0.0.1 --alter="add column weight555 int;" D=yml_school,t=t_yml --nodrop-old-table --chunk-size=2000 --chunk-time=0.5 --max-lag=2 --check-interval=2 --charset=utf8 --critical-load="Threads_running=64" --max-load="Threads_running=32" --execute`)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	stdout, oerr := cmd.StdoutPipe()
	if oerr != nil {
		fmt.Println("cmd.Output pipe err:%v ", oerr)
		return
	}
	defer stdout.Close()

	/*stderr, eerr := cmd.StderrPipe()
	if eerr != nil {
		fmt.Println("cmd.Output pipe err:%v ", eerr)
		return
	}*/
	r := bufio.NewReader(stdout)
	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	fmt.Printf("cmd start -->\n")

	//buf := make([]byte, 1000)

	for {
		//fmt.Printf("wait read\n")
		line, e := r.ReadString('\n')
		//line, e := r.Read(buf)
		if e != nil {
			fmt.Printf("read err:%v", e)
			break
		}
		//fmt.Printf("-->%d:%s\n", line, string(buf))
		fmt.Printf("%s", line)
	}
	if exerr := cmd.Wait(); exerr != nil {
		fmt.Printf("%s ---err--- %v\n", "EXEC", exerr)
	} else {
		fmt.Printf("%s ---OK---\n", "EXEC")
	}
}
