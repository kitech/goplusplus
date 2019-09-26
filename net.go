package gopp

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

// u: http://ip:port
func ProxyHttpClient(u string) *http.Client {
	tp := &http.Transport{}
	pxyurl := u
	urlo, err := url.Parse(pxyurl)
	if err != nil {
		log.Panicln(err, pxyurl)
	}

	tp.Proxy = http.ProxyURL(urlo)
	cli := &http.Client{}
	cli.Transport = tp

	return cli
}

type HttpClient struct {
	c *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

// timeoms == 0, use default value
func NewHttpClient2(timeoms int) *http.Client {
	cli := &http.Client{}
	if timeoms != 0 {
		cli.Timeout = time.Duration(timeoms) * time.Millisecond
	}
	tp := &http.Transport{}
	tp.DisableCompression = false

	tlscfg := &tls.Config{}
	tlscfg.InsecureSkipVerify = true
	tp.TLSClientConfig = tlscfg

	cli.Transport = tp

	return cli
}

// only ip:port, if ip part is a domain name, this will fail
func ParseUdpAddr(address string) *net.UDPAddr {
	ta := ParseTcpAddr(address)
	ua := &net.UDPAddr{}
	ua.Port = ta.Port
	ua.IP = ta.IP
	return ua
}
func ParseTcpAddr(address string) *net.TCPAddr {
	ao := ParseAddr(address)
	return ao.(*net.TCPAddr)
}
func ParseAddr(address string) net.Addr {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil
	}
	iport, err := strconv.Atoi(port)
	if err != nil {
		return nil
	}
	ip := net.ParseIP(host)
	ao := &net.TCPAddr{}
	ao.IP = ip
	ao.Port = iport
	return ao
}

type FmtWriter struct {
	io.Writer
}

func NewFmtwriter(w io.Writer) *FmtWriter {
	this := &FmtWriter{w}
	return this
}
func (this *FmtWriter) Print(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprint(this.Writer, a...)
	return
}
func (this *FmtWriter) Printf(format string, a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(this.Writer, format, a...)
	return
}
func (this *FmtWriter) Println(a ...interface{}) (n int, err error) {
	n, err = fmt.Fprintln(this.Writer, a...)
	return
}

func HTWFlush(w http.ResponseWriter) bool {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
		return true
	}
	return false
}

func XferCopy(c1, c2 net.Conn, close bool) (int64, int64, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		cpn, err := io.Copy(c2, c1)
		ErrPrint(err, cpn)
		// which error?
		c2.Close()
		c1.Close()
		wg.Done()
	}()
	go func() {
		cpn, err := io.Copy(c1, c2)
		ErrPrint(err, cpn)
		c1.Close()
		c2.Close()
		wg.Done()
	}()
	wg.Wait()
	c1.Close()
	c2.Close()
	return 0, 0, nil
}
