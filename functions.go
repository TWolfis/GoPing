/*
Contains functions to elevate potential clutter from main.go
*/
package main

import (
	"goping/Ping"
	"log"
)

func SingleHost(dstAddr *string, srcAddr *string, count *int) {
	//Ping a single host
	//Setup, send out and listen for an ICMP echo

	succeeded := 0

	ping := Ping.SetPingSetup(*dstAddr, *srcAddr)
	for i := 0; i < *count; i++ {
		echo := ping.SendEcho()
		if ping.ListenForEcho(echo) {
			succeeded++
		}
	}
	log.Printf("Send out %v packets\t%v got a reply from %v\t%v%v success rate\n",
		*count, succeeded, *dstAddr, succeeded/(*count)*100, "%")
}

func FloodHost(dstAddr *string, srcAddr *string) {
	log.Println("Going to flood target with pings press ctrl+c to stop")
	ping := Ping.SetPingSetup(*dstAddr, *srcAddr)
	for {
		ping.SendEcho()
	}

}
