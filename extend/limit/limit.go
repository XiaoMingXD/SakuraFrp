package limit

import (
	"io"

	frpNet "github.com/fatedier/frp/utils/net"
)

const (
	B uint64 = 1 << (10 * (iota))
	KB
	MB
	GB
	TB
	PB
	EB
)

const burstLimit = 1024 * 1024 * 1024

type Conn struct {
	frpNet.Conn

	lr io.Reader
	lw io.Writer
}

func NewLimitConn(maxread, maxwrite uint64, c frpNet.Conn) Conn {
	// 这里不知道为什么要 49 才能对的上真实速度
	// 49 是根据 wget 速度来取的，测试了 512、1024、2048、4096、8192 等多种速度下都很准确
	return Conn{
		lr:   NewReaderWithLimit(c, maxread),
		lw:   NewWriterWithLimit(c, maxwrite),
		Conn: c,
	}
}

func (c Conn) Read(p []byte) (n int, err error) {
	return c.lr.Read(p)
}

func (c Conn) Write(p []byte) (n int, err error) {
	return c.lw.Write(p)
}
