package gopp

import "time"

/*
unblock features, include chan and socket
*/

// 可能的错误，自动检测chan 是否关闭，是否已满阻塞
// 需要用到reflect，效率上多少折扣呢？
func ChanSend(c interface{}, timeout time.Duration) error {
	return nil
}

func ChanRecv() error {
	return nil
}

// see https://github.com/tidwall/evio
func SocketWrite() error {
	return nil
}

func SocketRead() error {
	return nil
}
