package epoller

import (
	"net"
	"reflect"
)

type Poller interface {
	Add(conn net.Conn) (int, error)
	Remove(conn net.Conn) error
	Wait(count int) ([]net.Conn, error)
	WaitWithBuffer() ([]net.Conn, error)
	WaitChan(count int) <-chan []net.Conn
	GetConnectionByFD(fd int) (net.Conn, error)
	GetFDByConnection(conn net.Conn) (int, error)
	Close() error
}

func SocketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
