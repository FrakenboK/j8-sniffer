package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// TODO: add conda cli
func main() {
	device := "lo"

	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// No filter??
	filter := "tcp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		app := packet.ApplicationLayer()

		if app == nil {
			continue
		}
		fmt.Println("----------------------------")
		fmt.Println(string(packet.ApplicationLayer().Payload()))
		// send to api
	}
}
