package main

import (
	apiHandlers "PostgreSQL-Project/apiHandler"
	"PostgreSQL-Project/dbConfig"
	"errors"
	"fmt"
	"log"
	"net"

	//"PostgreSQL-Project/apiHandler"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load environment file")
	}
	// integrations.SetEnvironmentVariables()
}

func main() {
    fmt.Println("Starting application")

	app := fiber.New(fiber.Config{
		AppName: "PostgreSQL-Project",
		BodyLimit: 4000 * 1024,
	})

    dbConfig.ConnectDB()

    apiHandlers.Router(app)

    log.Fatal(app.Listen(":8888"))

}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network")
}