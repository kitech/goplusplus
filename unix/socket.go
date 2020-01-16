package unix

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"syscall"
	"unsafe"

	unixo "golang.org/x/sys/unix"
)

func dupsomefile(netc interface{}) (*os.File, error) {
	var fp *os.File
	var err error
	switch rc := netc.(type) {
	case *net.TCPConn:
		fp, err = rc.File()
	case *net.UDPConn:
		fp, err = rc.File()
	case *net.TCPListener:
		fp, err = rc.File()
	case *net.UnixListener:
		fp, err = rc.File()
	case *net.IPConn:
		fp, err = rc.File()
	case *os.File:
		fd, err1 := unixo.Dup(int(rc.Fd()))
		err = err1
		if err == nil {
			fp = os.NewFile(uintptr(fd), rc.Name())
		}
	case int:
		fd, err1 := unixo.Dup(int(rc))
		err = err1
		if err == nil {
			fp = os.NewFile(uintptr(fd), fmt.Sprintf("dummyfd%d", rc))
		}
	case uintptr:
		fd, err1 := unixo.Dup(int(rc))
		err = err1
		if err == nil {
			fp = os.NewFile(uintptr(fd), fmt.Sprintf("dummyfd%d", rc))
		}
	default:
		ncty := reflect.TypeOf(netc).Elem()
		for i := 0; i < ncty.NumField(); i++ {
			fld := ncty.Field(i)
			if fld.Type.String() == "net.Conn" {
				netc2 := reflect.ValueOf(netc).Elem().Field(i).Interface()
				return dupsomefile(netc2)
			}
		}
		return nil, os.ErrInvalid
	}
	return fp, err
}
func GetSocketOutq(netc interface{}) (int, error) {
	fp, err := dupsomefile(netc)
	if err != nil {
		return 0, err
	}
	// fp is dup of c, need caller close fp, see https://golang.google.cn/pkg/net/#TCPConn.File
	defer fp.Close()

	// syscall.SOL_SOCKET
	// v, err := syscall.GetsockoptInt(int(fp.Fd()), syscall.SOL_, syscall.TIOCOUTQ)
	v, err := unixo.IoctlGetInt(int(fp.Fd()), unixo.TIOCOUTQ)
	// why ??? file tcp 10.0.0.32:41848->72.69.89.8:443: fcntl: too many open files 72.69.89.8:443
	return v, err
}
func GetSockBlocking(tc *net.TCPListener) (bool, error) {
	fd := GetSockFD2(tc)
	return GetSockBlocking0(fd)
}
func GetSockBlocking0(fd int) (bool, error) {
	flags, err := unixo.FcntlInt(uintptr(fd), unixo.F_GETFL, 0)
	if err != nil {
		return false, err
	}
	return flags&unixo.O_NONBLOCK > 0, nil
}
func SetSockBlocking0(fd int, blocking bool) error {
	flags, err := unixo.FcntlInt(uintptr(fd), unixo.F_GETFL, 0)
	if err != nil {
		return err
	}
	if blocking {
		flags = flags & ^unixo.O_NONBLOCK
	} else {
		flags = flags | unixo.O_NONBLOCK
	}
	_, err = unixo.FcntlInt(uintptr(fd), unixo.F_SETFL, flags)
	return err
}

type netFD struct {
	// fdMutex
	a uint64
	b uint32
	c uint32

	sysfd int
}
type netinconn struct {
	fd *netFD // netFD
}

func GetSockFD(uc *net.UDPConn) int {
	inc := (*netinconn)(unsafe.Pointer(uc))
	sysfd := inc.fd.sysfd
	return sysfd
}
func GetSockFD2(uc *net.TCPListener) int {
	inc := (*netinconn)(unsafe.Pointer(uc))
	sysfd := inc.fd.sysfd
	return sysfd
}

func SetSockTTL(uc *net.UDPConn, ttl int) error {
	sysfd := GetSockFD(uc)
	err := syscall.SetsockoptInt(sysfd, syscall.IPPROTO_IP, syscall.IP_TTL, ttl)
	return err
}

func SetSockReuseAddrPort(uc *net.TCPListener, reuse bool) error {
	sysfd := GetSockFD2(uc)
	return SetSockReuseAddrPort2(sysfd, reuse)
}
func SetSockReuseAddrPort2(fd int, reuse bool) error {
	val := 0
	if reuse {
		val = 1
	}
	err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, val)
	if err == nil {
		return err
	}
	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, unixo.SO_REUSEPORT, val)
	return err
}

func init() {
	if false {
		syscall.SetsockoptInt(0, syscall.SOCK_DGRAM, 0, 0)
	}

}
