package gopp

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

// background dump exec.Cmd output, stdout/stderr
// usage:
// cmdo:=exec.Command(...)
// gopp.DumpCmdout(...)
// cmdo.Start()
// cmdo.Wait()
func DumpCmdout(cmdo *exec.Cmd, prefix string, incout bool, incerr bool) {
	cmdoutfp, err := cmdo.StdoutPipe()
	ErrPrint(err)
	go func() {
		r := bufio.NewReader(cmdoutfp)
		for {
			lineb, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			ErrPrint(err)
			if err != nil {
				break
			}
			line := string(lineb)
			log.Println(prefix, line)
		}
	}()
	cmderrfp, err := cmdo.StderrPipe()
	ErrPrint(err)
	go func() {
		r := bufio.NewReader(cmderrfp)
		for {
			lineb, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			ErrPrint(err)
			if err != nil {
				break
			}
			line := string(lineb)
			log.Println("E", prefix, line)
		}
	}()
}
