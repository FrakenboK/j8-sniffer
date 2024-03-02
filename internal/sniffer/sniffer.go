package sniffer

import (
	"fmt"
	"log"

	grpcclient "github.com/FrakenboK/j8-sniffer/internal/grpc_client"
	"github.com/FrakenboK/j8-sniffer/internal/sniffer/config"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

type Sniffer struct {
	cfg    *config.Config
	client *grpcclient.Client
}

func Init(cfg *config.Config, client *grpcclient.Client) *Sniffer {
	sniffer := &Sniffer{
		cfg:    cfg,
		client: client,
	}
	return sniffer
}

func (sn *Sniffer) Run(cmd *cobra.Command, args []string) {
	handle, err := pcap.OpenLive(sn.cfg.Api.Interface, 65536, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("Error: invalid interface %s\n", sn.cfg.Api.Interface)
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

		data := packet.ApplicationLayer().Payload()
		srcPort := packet.Layer(layers.LayerTypeTCP).(*layers.TCP).SrcPort.String()
		dstPort := packet.Layer(layers.LayerTypeTCP).(*layers.TCP).DstPort.String()

		sn.client.Send(data, dstPort, srcPort)
	}
}
