package sniffer

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

func Sniff(cmd *cobra.Command, args []string) {
	netInterface, err := cmd.Flags().GetString("interface")
	if err != nil || len(netInterface) == 0 {
		cmd.Help()
		return
	}

	apiHost, err := cmd.Flags().GetString("host")
	if err != nil || len(apiHost) == 0 {
		cmd.Help()
		return
	}

	handle, err := pcap.OpenLive(netInterface, 65536, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("Error: invalid interface %s\n", netInterface)
		fmt.Printf("Warning: run j8 out of sudo!\n")
		cmd.Help()
		return
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
		fmt.Println(packet.Layer(layers.LayerTypeTCP).(*layers.TCP).SrcPort.String())
		fmt.Println(string(packet.ApplicationLayer().Payload()))
		// send to api
	}
}
