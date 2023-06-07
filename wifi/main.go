package main

// import (
// "fmt"
// "log"
// "time"
// "github.com/google/gopacket"
// "github.com/google/gopacket/pcap"
// // "github.com/google/gopacket/pfring"

// "github.com/google/gopacket/layers"
// )

// func main() {
// // Get a list of interfaces
// ifs, err := pcap.FindAllDevs()
// if err != nil {
// log.Fatal(err)
// }
// // Find the interface that is connected to the network we want to sniff
// var iface *pcap.Interface
// for _, i := range ifs {
// 	if i.Name == "en0" {
// 		iface = &i
// 		break
// 	}
// }

// // Open the device for capturing packets
// handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
// if err != nil {
// 	log.Fatal(err)
// }
// defer handle.Close()

// // Set the filter
// filter := "wlan[0:4] == 0x" + stringToHex("Search")
// err = handle.SetBPFFilter(filter)
// if err != nil {
// 	log.Fatal(err)
// }

// // Loop through the packets
// packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// for packet := range packetSource.Packets() {
// 	// Check if the packet contains a WPA handshake
// 	if handshake := packet.Layer(layers.LayerTypeEAPOL); handshake != nil {
// 		// Print the captured password
// 		fmt.Println("Captured password:", string(handshake.LayerPayload()[83:]))
// 		return
// 	}

// 	// Sleep for a bit to avoid hogging the CPU
// 	time.Sleep(time.Millisecond * 10)
// }

// // If we got here, we didn't find a password
// fmt.Println("Couldn't find the password :(")
// }

// // Helper function to convert a string to its hex representation
// func stringToHex(s string) string {
// 	hex := ""
// 	for _, c := range s {
// 		hex += fmt.Sprintf("%x", uint32(c))
// 	}
// 	return hex
// }



// package main

// import (
// 	"fmt"
// 	"log"
// 	// "time"

// 	"github.com/google/gopacket"
// 	"github.com/google/gopacket/layers"
// 	"github.com/google/gopacket/pcap"
// )

// func main() {
// 	// Get a list of interfaces
// 	ifs, err := pcap.FindAllDevs()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Find the interface that is connected to the network we want to sniff
// 	var iface *pcap.Interface
// 	for _, i := range ifs {
// 		if i.Name == "en0" {
// 			iface = &i
// 			break
// 		}
// 	}

// 	if iface == nil {
// 		log.Fatal("Unable to find en0 interface")
// 	}

// 	// Open the device for capturing packets
// 	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer handle.Close()

// 	// Set the filter to capture packets from a particular Wi-Fi SSID
// 	ssidFilter := "wlan[0:4] == 0x" + stringToHex("Mori")
// 	err = handle.SetBPFFilter(ssidFilter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Loop through the packets
// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// 	for packet := range packetSource.Packets() {
// 		// Check if the packet contains a WPA handshake
// 		if handshake := packet.Layer(layers.LayerTypeEAPOL); handshake != nil {
// 			// Print the captured password
// 			fmt.Println("Captured password:", string(handshake.LayerPayload()[83:]))
// 			return
// 		}

// 		// Sleep for a bit to avoid hogging the CPU
// 		// time.Sleep(time.Millisecond * 10)
// 	}

// 	// If we got here, we didn't find a password
// 	fmt.Println("Couldn't find the password :(")
// }

// // Helper function to convert a string to its hex representation
// func stringToHex(s string) string {
// 	hex := ""
// 	for _, c := range s {
// 		hex += fmt.Sprintf("%x", uint32(c))
// 	}
// 	return hex
// }
