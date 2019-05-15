package tun2socks

import (
	"fmt"
	"github.com/google/gopacket/layers"
	"io"
	"log"
	"net"
	"time"
)

type Session struct {
	ByPass     bool
	Pipe       *ProxyPipe
	UPTime     time.Time
	RemoteIP   net.IP
	RemotePort int
}

func (s *Session) ToString() string {
	return fmt.Sprintf("(%t)%s:%d t=%s", s.ByPass, s.RemoteIP, s.RemotePort,
		s.UPTime.Format("2006-01-02 15:04:05"))
}

func newSession(ip4 *layers.IPv4, tcp *layers.TCP) *Session {
	s := &Session{
		UPTime:     time.Now(),
		RemoteIP:   ip4.DstIP,
		RemotePort: int(tcp.DstPort),
		ByPass:     SysConfig.ByPass(ip4.DstIP),
	}
	return s
}

type ProxyPipe struct {
	Left  *net.TCPConn
	Right *net.TCPConn
}

func (pp *ProxyPipe) Right2Left() {
	no, err := io.Copy(pp.Left, pp.Right)
	log.Println("Proxy pipe right 2 left finished:", no, err)
}

func (pp *ProxyPipe) WriteTunnel(buf []byte) {
	if _, e := pp.Right.Write(buf); e != nil {
		log.Println("Proxy pipe left 2 right err:", e)
		pp.Close()
	}
}

func (pp *ProxyPipe) Close() {

}
