package Ping

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type PacketSetup struct {
	dstAddr    *net.IPAddr
	listenAddr string
	msg        *icmp.Message
}

func SetPingSetup(dstAddr string, srcAddr string) *PacketSetup {

	//see if the destination is a valid destination
	destAddr, err := net.ResolveIPAddr("ip4", dstAddr)
	if err != nil {
		log.Fatal("Error while setting up PingSetup struct: ", err)
	}

	//fill up the PingSetup struct with the destination addr, listen addr and the msg of the packet
	pingPacket := PacketSetup{
		dstAddr:    destAddr,
		listenAddr: srcAddr,

		//message contains the type, content and sequence number of the packet that is being sent
		msg: &icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  1,
				Data: []byte(""),
			},
		},
	}

	return &pingPacket
}

func (PacketSetup *PacketSetup) SendEcho() *icmp.PacketConn {
	//set type of connection to listen for
	const network = "ip4:icmp"

	//setup listener for the ICMP message/reply
	c, err := icmp.ListenPacket(network, PacketSetup.listenAddr)
	if err != nil {
		log.Fatal("Error when sending message:", err)
	}

	//encode ICMP message into a binary format
	encodedMessage, err := PacketSetup.msg.Marshal(nil)
	if err != nil {
		log.Fatal("Error when encoding message:", err)
	}

	//Write message to our tx/rx queue to be sent out to the destination target
	if _, err := c.WriteTo(encodedMessage, PacketSetup.dstAddr); err != nil {
		return &icmp.PacketConn{}
	}

	return c
}

func (PacketSetup *PacketSetup) ListenForEcho(conn *icmp.PacketConn) bool {
	//Setup listener for an echo reply
	rx := make([]byte, 1500)
	proto := ipv4.ICMPTypeEchoReply.Protocol()

	//close conn when operations are done
	defer conn.Close()

	err := conn.SetReadDeadline(time.Now().Add(time.Second * 1))
	if err != nil {
		log.Fatal("Error while setting Read deadline:", err)
	}
	//Read reply from rx buffer as well as size of the reply
	nBytes, _, err := conn.ReadFrom(rx)
	if err != nil {
		return false
	}

	//Parse reply from buffer into and icmp message
	reply, err := icmp.ParseMessage(proto, rx[:nBytes])
	if err != nil {
		log.Fatal("Error while parsing message:", err)
		return false
	}

	switch reply.Code {
	case 0:
		//		log.Printf("Got reply back from:%v\tSize:%v bytes read\n", dstAddr, nBytes)
		return true
	case 3:
		return false
	case 11:
		return false
	default:
		log.Printf("Got %+v want echo reply\n", reply)
		return false
	}
}
