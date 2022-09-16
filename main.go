package main

//easy ping util for go built over the net/icmp package
import (
	"flag"
	"log"
	"os/user"
)

var (
	//connection data
	dstAddr = flag.String("d", "", "Provide destination address (or list of addresses) to ping")
	srcAddr = flag.String("l", "0.0.0.0", "Provide listen address")
	count   = flag.Int("c", 1, "Provide amount of pings to send")
	flood   = flag.Bool("f", false, "Set this mode to flood the target with icmp packets")

	//var to hold current user to verify if the root user is running the program or not
	currentUser, _ = user.Current()
)

func main() {
	flag.Parse()
	//User should be root to run this program
	if currentUser.Username != "root" {
		log.Fatal("Run program as root")
	}

	if *dstAddr == "" {
		log.Fatal("No destination address provided run program again with -h for help")
	}

	if *flood {
		FloodHost(dstAddr, srcAddr)
	} else {
		SingleHost(dstAddr, srcAddr, count)
	}

}
