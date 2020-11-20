package main

import (
  "fmt"
  "net"
  "time"
  "os"
)

var udpCount int
var udpPacketsNum int
var traffic int
var total float64
var lastTotal float64
var udpStartTime time.Time

func main() {
  udpPacketsNum = 15000
  fmt.Println("Start UDP Server")
  addr, err := net.ResolveUDPAddr("udp", "192.168.2.54:6000") // 205, 206
  if err != nil {
    fmt.Print(err)
    return
  }

  listenner, err := net.ListenUDP("udp", addr)
  if err != nil {
    fmt.Println("net.Listen err =", err)
    return
  }
  defer listenner.Close()

  RecvFile("test", listenner)
}

func RecvFile(fileName string, conn net.Conn) {
  buf := make([]byte, 1400)

  traffic := (udpPacketsNum*1024)/(1024*1024)*8
  f, err := os.Create("my_test.txt")
  if err != nil {
    fmt.Println("conn.Read err =", err)
  }

  f.Close()
  for {
    udpStartTime = time.Now() // get current time

    for i := 0; i < udpPacketsNum-1; i++ {
      _, err := conn.Read(buf)
      if err != nil {
        fmt.Println("conn.Read err =", err)
      }

    }

    endTime := time.Now() // get current time
    duration := endTime.Sub(udpStartTime)
    durationSec := float64(duration)/1000000000
    //fmt.Printf("Traffic: %d bytes\n", traffic)
    //fmt.Printf("Traffic: %f bps\n", (float64(traffic)/final_s)*8)
    //fmt.Printf("Traffic: %f MBs\n", (float64(traffic)/(1024*1024))/final_s)
    f, err := os.OpenFile("my_test.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
      fmt.Println("conn.Read err =", err)
    }
    item := float64(traffic)/durationSec
    _, err = fmt.Fprintln(f,item)
    if err != nil {
      fmt.Println("conn.Read err =", err)
    }
    f.Close()
    fmt.Printf("%f\n", float64(traffic)/durationSec)
  }
}
