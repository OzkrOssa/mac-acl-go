package devices

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/go-routeros/routeros"
	"github.com/joho/godotenv"
)

func Mikrotik(ipAddress string, macAddress string, comment string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		address  = flag.String("address", fmt.Sprintf("%s:8728", ipAddress), "Address")
		username = flag.String("username", os.Getenv("USER"), "Username")
		password = flag.String("password", os.Getenv("PASSWORD"), "Password")
	)
	flag.Parse()

	parseMac, err := net.ParseMAC(macAddress)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: check ip address in database

	c, err := routeros.Dial(*address, *username, *password)
	if err != nil {
		log.Fatal(err)
	}

	reply, err := c.Run("/interface/wireless/access-list/add", "=interface=wlan1", fmt.Sprintf("=mac-address=%s", parseMac), fmt.Sprintf("=comment=%s", comment))

	if err != nil {
		log.Fatal(err)
	}

	return reply.Done.Word, nil

}
