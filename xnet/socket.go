package xnet

/*
#include <errno.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
*/
import "C"
import "unsafe"

type Socket struct {
	fd   int
	eno  int
	emsg string
}

func NewSocket() *Socket {
	fd := C.socket(C.AF_INET, C.SOCK_STREAM, 0)
	fd2 := int(fd)
	sock := &Socket{}
	sock.fd = fd2
	//
	return sock
}

func (sk *Socket) Connect(address string, port int) {
	var sa = &C.struct_sockaddr_in{}
	sa.sin_family = C.AF_INET
	sa.sin_port = C.htons(port)
	C.inet_pton(C.AF_INET, address.ptr, &sa.sin_addr.s_addr)
	var rv = C.connect(sk.fd, sa, unsafe.Sizeof(C.struct_sockaddr_in))
	if rv != 0 {
	}
	eno := *C.__errno_location()
	println(rv, *C.__errno_location())
	emsg := C.strerror(eno)
	println(emsg)
	println(sk.fd, sa.sin_port)

}
