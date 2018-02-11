// test.go
package main

import (
	"bytes"
	"fmt"
	"lpp"
	"os"
	"strconv"
)

func main() {
	RAWHANDLE, _ := os.Open(os.Args[1])
	RAWIO := lpp.GetBlockRead(RAWHANDLE, "\n", false, 10000)
	max := 0
	all_need := make(map[string]string, 10000)
	all_cov := 0
	END, _ := os.Create(os.Args[2])
	for {
		line, err := RAWIO.Next()

		line_l := bytes.Split(line, []byte("\t"))
		c_name := string(line_l[0])
		_, ok := all_need[c_name]
		if !ok || err == nil {
			t_cov := all_cov / max
			END.WriteString(fmt.Sprintf("%s\t%d\n", c_name, t_cov))
		}
		coord, _ := strconv.Atoi(string(line_l[1]))
		cov, _ := strconv.Atoi(string(line_l[2]))
		max = coord
		all_cov += cov
		if err != nil {
			break
		}

	}
}
