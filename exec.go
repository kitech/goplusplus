package gopp

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
)

type CmdInout struct {
	In  io.WriteCloser
	Out io.ReadCloser
	Err io.ReadCloser
}

func NewCmdInout(cmo *exec.Cmd) *CmdInout {
	var err error
	cio := &CmdInout{}
	cio.In, err = cmo.StdinPipe()
	ErrPrint(err)
	cio.Out, err = cmo.StdoutPipe()
	ErrPrint(err)
	cio.Err, err = cmo.StderrPipe()
	ErrPrint(err)

	return cio
}

// start and wait
func CmdRun(cmdo *exec.Cmd) error {
	err := cmdo.Start()
	if err != nil {
		return err
	}
	err = cmdo.Wait()
	return err
}

// background dump exec.Cmd output, stdout/stderr
// usage:
// cmdo:=exec.Command(...)
// gopp.DumpCmdout(...)
// cmdo.Start()
// cmdo.Wait()
func DumpCmdout(cmdo *exec.Cmd, prefix string, incout bool, incerr bool) (outch, errch chan string) {
	outch = make(chan string, 1)
	errch = make(chan string, 1)
	cmdoutfp, err := cmdo.StdoutPipe()
	ErrPrint(err)
	if err == nil {
		go cmdiocpline(cmdoutfp, outch, prefix, incout)
	}

	cmderrfp, err := cmdo.StderrPipe()
	ErrPrint(err)
	if err == nil {
		go cmdiocpline(cmderrfp, errch, prefix, incerr)
	}
	return
}

func cmdiocpline(cmdiofp io.ReadCloser, ioch chan string, prefix string, incstdio bool) {
	r := bufio.NewReader(cmdiofp)
	defer cmdiofp.Close()
	stop := false
	for !stop {
		lineb, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		ErrPrint(err)
		if err != nil {
			break
		}
		line := string(lineb)
		if incstdio {
			fmt.Println(prefix, line)
		}
		select {
		case ioch <- line:
		default:
			stop = true
			break
		}
	}
}

// 一般命令行都是按字符串的，按行的，所以返回值就特化一点吧
func TeeCmdout(cmdo *exec.Cmd, prefix string, incout bool, incerr bool) ([]string, error) {
	outch, errch := DumpCmdout(cmdo, prefix, incout, incerr)
	stopch := make(chan bool, 1)
	lines := []string{}

	go func() {
		stop := false
		for !stop {
			select {
			case line := <-outch:
				lines = append(lines, line)
			case line := <-errch:
				lines = append(lines, line)
			case <-stopch:
				stop = true
			}
		}
	}()
	runtime.Gosched()

	var err error
	err = cmdo.Start()
	if err != nil {
	} else {
		err = cmdo.Wait()
	}

	select {
	case stopch <- true:

	}
	return lines, err
}

// verbose
func TeeCmdoutV(cmdo *exec.Cmd, prefix string, incout bool, incerr bool) ([]string, error) {
	cddir := IfElseStr(cmdo.Dir == "", "", "cd "+cmdo.Dir+" && ")
	fmt.Printf("> %s %s %v\n", cmdo.Path, cddir, cmdo.Args)
	return TeeCmdout(cmdo, prefix, incout, incerr)
}

func CombinedLines(cmdo *exec.Cmd) ([]string, error) {
	allout, err := cmdo.CombinedOutput()
	lines := strings.Split(string(allout), "\n")
	return lines, err
}
