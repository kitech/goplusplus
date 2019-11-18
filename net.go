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
	*http.Client
}

func NewHttpClient() *HttpClient {
	cli := &HttpClient{}
	c := NewHttpClient2(0)
	cli.Client = c
	return cli
}

// timeoms == 0, use default value
func NewHttpClient2(timeoms int) *http.Client {
	cli := &http.Client{}
	// tp1 := (http.DefaultTransport.(*http.Transport))
	// tp := tp1.Clone() // go1.13
	tp := &http.Transport{}
	tp.DisableCompression = false
	if timeoms > 0 {
		todur := time.Duration(timeoms) * time.Millisecond
		cli.Timeout = todur
		tp.TLSHandshakeTimeout = todur
	}

	tlscfg := &tls.Config{}
	tlscfg.InsecureSkipVerify = true
	tp.TLSClientConfig = tlscfg
	tp.DisableKeepAlives = true
	tp.MaxIdleConns = 3

	cli.Transport = tp

	return cli
}
func CloseHttpIdles() {
	tpx := http.DefaultClient.Transport
	if tpx == nil {
		return
	}
	tp := tpx.(*http.Transport)
	tp.CloseIdleConnections()
	// http.DefaultClient.CloseIdleConnections() // go1.12
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
	donec12 := make(chan bool, 1)
	donec21 := make(chan bool, 1)
	go func() {
		cpn, err := io.Copy(c2, c1)
		if ErrHave(err, "use of closed network connection") {
		} else {
			ErrPrint(err, cpn)
		}
		donec12 <- true
	}()
	go func() {
		cpn, err := io.Copy(c1, c2)
		if ErrHave(err, "use of closed network connection") {
		} else {
			ErrPrint(err, cpn)
		}
		donec21 <- true
	}()

	// 双向转发的终止方法
	// first cycle, any done
	select {
	case <-donec12:
	case <-donec21:
	}
	// second cycle, all done or timeout
	select {
	case <-donec12:
	case <-donec21:
	case <-time.After(135 * time.Second):
	}

	c1.Close()
	c2.Close()
	return 0, 0, nil
}
