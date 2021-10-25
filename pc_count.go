package main

import (
  "fmt"
  "log"
  "github.com/google/gopacket"
  "github.com/google/gopacket/layers"
  "github.com/google/gopacket/pcap"
  "time"
//  "encoding/json"
//  "bytes"
)

type Port_Number_Count struct {
  PortNumber int `json:"PortNumber"`
  Count      int `json:"Count"`
}

type Array []Port_Number_Count

var (
//  device       string = "ens33"
  device       string = "lo"
  snapshot_len int32  = 1024
  promiscuous  bool   = false
  err          error
  timeout      time.Duration = 30 * time.Second
  handle       *pcap.Handle

  ethLayer     layers.Ethernet
  ipLayer      layers.IPv4
  tcpLayer     layers.TCP
  udpLayer     layers.UDP
  payloadLayer gopacket.Payload

  port_number_count = Port_Number_Count{}
  array        Array
)
func main() {

  handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
  if err != nil {
    log.Fatal(err)
  }
  defer handle.Close()

//  var filter string = "udp"
//  setFilter(filter, handle)

  fmt.Println("Stert capturing packets.")

  parser := gopacket.NewDecodingLayerParser(
    layers.LayerTypeEthernet,
    &ethLayer,
    &ipLayer,
    &tcpLayer,
    &udpLayer,
    &payloadLayer,
  )
  var source gopacket.PacketDataSource = handle
  decodedLayers := make([]gopacket.LayerType, 0, 10)
  for {
    data, _, err := source.ReadPacketData()
    if err != nil {
      fmt.Println("Error reading packet data: ", err)
      continue
    }
    fmt.Println("Decodeing packet")
    err = parser.DecodeLayers(data, &decodedLayers)
    for _, layerType := range decodedLayers { if layerType == layers.LayerTypeIPv4 {
        fmt.Println("IPv4: ", ipLayer.SrcIP, "->", ipLayer.DstIP)
      }
      if layerType == layers.LayerTypeTCP {
        fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
        fmt.Println("")
      }
      if layerType == layers.LayerTypeUDP {
        fmt.Println("UDP Port: ", udpLayer.SrcPort, "->", udpLayer.DstPort)
        fmt.Println("")
      }
    }
    if decodedLayers.Truncated {
      fmt.Println(" Packet has been truncated")
    }
    if err != nil {
      fmt.Println(" Error encontered:", err)
    }
  }
}

func setFilter(filter string, handle *pcap.Handle) {
  err = handle.SetBPFFilter(filter)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Filter Set: ", filter)
}

//var count int = 0
//func packetCount(udpDstPort) {
func packetCount() {
//  for i=0 i<len(port_number_count.PortNumber); i++{
//    }
//  count = count + 1
//  println ("cont: ", count)
  buf :=  Port_Number_Count{PortNumber: 5001, Count: 0}
  array = append(array, buf)
  buf =  Port_Number_Count{PortNumber: 5002, Count: 0}
  array = append(array, buf)

//  array[0].Count++

  for i:=0; i<len(array); i++ {
    //fmt.Println(array[i].PortNumber)
    fmt.Println(array[i].Count)
  }
}