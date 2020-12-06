package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrocProc struct {
	keep map[string]int
}

func NewErrocProc() *ErrocProc {
	e := ErrocProc{}
	e.keep = make(map[string]int)
	return &e
}

func (t *ErrocProc) GetInfo(x string) (string, int, error) {

	line := -1
	filename := ""
	tmp := ""

	i := strings.LastIndexAny(x, "\\")

	//fmt.Println("Index: ", i)
	if i > -1 {
		tmp = x[i+1:]
	} else {
		tmp = x
	}

	r := strings.SplitAfterN(tmp, " ", 2)

	if len(r) < 2 {
		return filename, line, fmt.Errorf("format error")
	}

	filename = r[0]
	filename = strings.Replace(filename, " ", "", -1)
	n := 16 //限制最長檔名
	if len(filename) > n {
		filename = filename[0:n]
	}

	tmp = r[1]
	tmp = strings.Replace(tmp, " ", "", -1)

	line, err := strconv.Atoi(tmp)

	if err != nil {
		return filename, line, err
	}

	return filename, line, nil

}

func (t *ErrocProc) Process(ErrorInfo string) (string, error) {
	count := 0
	v, ok := t.keep[ErrorInfo]
	if ok == true {
		count = v + 1
	}
	t.keep[ErrorInfo] = count

	name, line, err := t.GetInfo(ErrorInfo)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%s %d %d", name, line, count) //檔名,行數,出現次數

	return s, nil
}

func main() {

	a := NewErrocProc()
	in := "E:\\V1R2\\product\\fpgadrive.c   1325"
	s, err := a.Process(in)
	s, err = a.Process(in)
	s, err = a.Process(in)
	if err != nil {
		return
	}
	fmt.Println(s)
	fmt.Println("done")

}
