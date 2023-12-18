package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntp_time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal()
	}

	formatterd_ntp_time := ntp_time.Format("02.01.2006 15:04:05")
	local_time := time.Now().Format("02.01.2006 15:04:05")

	fmt.Println("Локальное время: ", local_time)
	fmt.Println("Точное время: ", formatterd_ntp_time)
}
