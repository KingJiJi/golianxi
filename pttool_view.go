package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	P_OVER    = "OVER"
	P_ING     = "ING"
	R_SUCC    = "SUCC"
	R_FAILURE = "FAILURE"
	R_UNKNOWN = ""
)

var (
	ERR_NIL_OBJ               = errors.New("not this status")
	NIL_STRING                = ""
	ST_TIMEOUT  time.Duration = time.Duration(60 * time.Second)
)

type Buffer struct {
	b bytes.Buffer
	m sync.Mutex
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Read(p)
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Write(p)
}
func (b *Buffer) WriteString(p string) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteString(p)
}
func (b *Buffer) String() string {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.String()
}

type IOBuffer struct {
	O     *Buffer
	E     *Buffer
	OK    bool
	Over  bool
	Begin time.Time
	End   time.Time
	Name  string
}

type Result struct {
	P string `json:"process"`
	R string `json:"isok"`
	O string `json:"isover"`
}

func (b *IOBuffer) GetStatus() (bool, string, error) {
	var ret Result
	switch {
	case !b.Over: // processing
		ret.P = b.O.String()
		ret.R = R_UNKNOWN
		ret.O = P_ING
	case b.Over && b.OK: //process over, result is succ
		ret.P = b.O.String()
		ret.R = R_SUCC
		ret.O = P_OVER
	case b.Over && !b.OK: //process over, result is failure
		ret.P = b.E.String()
		ret.R = R_FAILURE
		ret.O = P_OVER
	}

	retdata, err := json.Marshal(ret)
	return b.Over, string(retdata), err
}

type PtChan struct {
	ChanChan map[string]*IOBuffer
	ChanLock sync.Mutex
}

func (pt *PtChan) QueryStatus(name string) (string, error) {

	pt.ChanLock.Lock()
	defer pt.ChanLock.Unlock()
	data, ok := pt.ChanChan[name]
	if !ok {
		return NIL_STRING, ERR_NIL_OBJ
	} else {
		isover, status, err := data.GetStatus()
		if isover {
			delete(pt.ChanChan, name)
		}
		return status, err
	}
}

func (pt *PtChan) AddStatus(name string, buf *IOBuffer) {
	pt.ChanLock.Lock()
	defer pt.ChanLock.Unlock()
	pt.ChanChan[name] = buf
	buf.Name = name
}

func (pt *PtChan) ClearStatus() {
	pt.ChanLock.Lock()
	defer pt.ChanLock.Unlock()
	now := time.Now()
	for _, data := range pt.ChanChan {
		if data.Over {
			if now.Sub(data.End) > ST_TIMEOUT {
				delete(pt.ChanChan, data.Name)
			}
		}
	}
}

func WriteBuf(src io.ReadCloser, buf *Buffer) {
	r := bufio.NewReader(src)
	for {
		//fmt.Printf("wait read\n")
		line, e := r.ReadString('\n')
		//line, e := r.Read(buf)
		if e != nil {
			//fmt.Printf("read err:%v\n", e)
			break
		}
		//fmt.Printf("-->%d:%s\n", line, string(buf))
		//fmt.Printf("%s", line)
		buf.WriteString(line)
	}
	return
}

func main() {

	//pt-online-schema-change --user=root --password=000000 --host=127.0.0.1 --alter="add column age int;" D=yml_school,t=t_yml --nodrop-old-table --chunk-size=2000 --chunk-time=0.5 --max-lag=2 --check-interval=2 --charset=utf8 --critical-load="Threads_running=64" --max-load="Threads_running=32" --execute
	//cmd := exec.Command("echo", `{"Name": "Bob", "Age": 32}`)
	cmd := exec.Command("/bin/bash", "-c", "-l", `pt-online-schema-change --user=dbms_w --password=aanLykl1Imsh65XS --host=10.104.105.71 --alter="add column haha int ;" D=yml_school,t=student  --nodrop-old-table --chunk-size=2000 --chunk-time=0.5 --max-lag=2 --check-interval=2 --charset=utf8 --critical-load="Threads_running=64" --max-load="Threads_running=32" --execute
	`)

	var mubuf IOBuffer
	var obuf, ebuf Buffer
	mubuf.E = &ebuf
	mubuf.O = &obuf

	stdout, oerr := cmd.StdoutPipe()
	if oerr != nil {
		fmt.Println("cmd.Output pipe err:%v \n", oerr)
		return
	}
	//defer stdout.Close()

	stderr, eerr := cmd.StderrPipe()
	if eerr != nil {
		fmt.Println("cmd.Output pipe err:%v ", eerr)
		return
	}

	go WriteBuf(stdout, mubuf.O)
	go WriteBuf(stderr, mubuf.E)

	mubuf.Begin = time.Now()
	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
	//_, mydata, _ := mubuf.GetStatus()
	//fmt.Printf("final\n\n\n\n%s\n\n\n\n", mydata)
	_, mydata, _ := mubuf.GetStatus()
	fmt.Printf("ing\n\n\n\n%s\n\n\n\n", mydata)
	//_, mydata, _ = mubuf.GetStatus()
	//fmt.Printf("final\n\n\n\n%s\n\n\n\n", mydata)

	if exerr := cmd.Wait(); exerr != nil {
		fmt.Printf("len(cmd.Args:%d)\n", len(cmd.Args))

		var argslist []string
		if len(cmd.Args) == 4 {
			argshaha := strings.Split(cmd.Args[3], " ")
			for _, arg := range argshaha {
				fmt.Printf("arg: [%s]\n", arg)
				if strings.Contains(arg, "--password=") {
					argslist = append(argslist, "--password=XXXXXXXX")
				} else {
					argslist = append(argslist, arg)
				}
			}
		}
		fmt.Printf("cmd:[%s]\n", strings.Join(argslist, " "))
		fmt.Printf("%s ---err--- %v\n %s", "EXEC", exerr, mubuf.E.String())
		mubuf.OK = false

	} else {
		argshaha := strings.Split(cmd.Args[2], " ")
		var argslist []string
		for _, arg := range argshaha {
			fmt.Printf("arg: [%s]\n", arg)
			if strings.Contains(arg, "--password=") {
				argslist = append(argslist, "--password=XXXXXXXX")
			} else {
				argslist = append(argslist, arg)
			}
		}
		fmt.Printf("cmd:[%s]\n", strings.Join(argslist, " "))
		fmt.Printf("%s ---OK---\n%s\n", "EXEC", mubuf.O.String())
		mubuf.OK = true
	}
	mubuf.Over = true
	mubuf.End = time.Now()

	_, mydata, _ = mubuf.GetStatus()
	fmt.Printf("final\n\n\n\n%s\n\n\n\n", mydata)

	fmt.Printf("closing......\n\n\n")
	time.Sleep(2 * time.Second)
}
