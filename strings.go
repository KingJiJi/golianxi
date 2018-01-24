package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetProgess(lie string) int {
	// Copying `ucar_dispatch`.`t_scd_order_dispatch`:  15% 29:27 remain
	//fmt.Printf("GetProgess| %s|\n", line)
	line := strings.TrimRight(lie, "\n")
	if strings.HasPrefix(line, "Copying") && strings.HasSuffix(line, "remain") {
		items := strings.Split(line, " ")
		//fmt.Printf("%d --- 0:%s,  1:%s, 2:'%s', 3:%s, 4:%s\n", len(items), items[0], items[1], items[2], items[3], items[4])
		if len(items) < 4 {
			return -1
		}
		ppp := strings.Trim(items[3], "% ")
		//fmt.Printf("progress-->|%s|\n", ppp)
		ppi, err := strconv.Atoi(ppp)
		if err != nil {
			return -1
		}
		//fmt.Printf("progress-->|%d|\n", ppi)
		return ppi
	}
	return -1
}

var progress0 = `Copying ucar_dispatch.t_scd_order_dispatch:   1% 33:18 remain
Copying ucar_dispatch.t_scd_order_dispatch:   2% 32:26 remain
Copying ucar_dispatch.t_scd_order_dispatch:   4% 30:37 remain
Copying ucar_dispatch.t_scd_order_dispatch:   6% 29:50 remain
Copying ucar_dispatch.t_scd_order_dispatch:   7% 29:24 remain
Copying ucar_dispatch.t_scd_order_dispatch:   9% 29:09 remain
Copying ucar_dispatch.t_scd_order_dispatch:  10% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  12% 29:13 remain
Copying ucar_dispatch.t_scd_order_dispatch:  13% 29:11 remain
Copying ucar_dispatch.t_scd_order_dispatch:  14% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  20% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  30% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  40% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  50% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  60% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  10% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  12% 29:13 remain
Copying ucar_dispatch.t_scd_order_dispatch:  13% 29:11 remain
Copying ucar_dispatch.t_scd_order_dispatch:  14% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  20% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  10% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  12% 29:13 remain
Copying ucar_dispatch.t_scd_order_dispatch:  13% 29:11 remain
Copying ucar_dispatch.t_scd_order_dispatch:  14% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  20% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  30% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  40% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  50% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  60% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  69% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  72% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  30% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  40% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  10% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  12% 29:13 remain
Copying ucar_dispatch.t_scd_order_dispatch:  13% 29:11 remain
Copying ucar_dispatch.t_scd_order_dispatch:  14% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  20% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  30% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  40% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  50% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  60% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  69% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  72% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  50% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  60% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  69% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  72% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  69% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  72% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  80% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  90% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  99% 29:27 remain
Copying ucar_dispatch.t_scd_order_dispatch:  100% 29:27 remain
execute complete ok`
var progress = `Copying ucar_dispatch.t_scd_order_dispatch:   1% 33:18 remain
Copying ucar_dispatch.t_scd_order_dispatch:   10% 32:26 remain
Copying ucar_dispatch.t_scd_order_dispatch:  20% 29:29 remain
Copying ucar_dispatch.t_scd_order_dispatch:  30% 29:27 remain
haha `

func main() {
	buffer := bytes.NewBufferString(progress0)

	var line string
	var ll error
	for ll == nil {
		line, ll = buffer.ReadString('\n')
		if ll != nil {
			//fmt.Printf("buff readstring err:%v", ll)
		}
		dl := strings.TrimRight(line, "\n")
		//fmt.Printf("|%s|\n", dl)
		ppi := GetProgess(dl)
		if ppi > 0 {
			for j := 100 - ppi; j > 0; j-- {
				fmt.Printf(" ")
			}
			for i := 0; i < ppi; i++ {
				fmt.Printf("=")
				time.Sleep(10 * time.Millisecond)
			}
			fmt.Printf("=>%d\n", ppi)

		}
		//fmt.Printf("%s", line)
	}

	hellp := "ABCGDFF:helloworld"
	hell2 := ""
	ii := strings.IndexByte(hellp, ':')
	if ii != -1 {
		hell2 = hellp[ii+1:]
	}
	fmt.Printf("\n\n%s=%s\n", hellp, hell2)

}
