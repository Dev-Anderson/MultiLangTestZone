package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func hostName() (string, error) {
	h, err := os.Hostname()
	if err != nil {
		return "", nil
	}
	return h, nil
}

func getMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.HardwareAddr != nil {
			return iface.HardwareAddr.String(), nil
		}
	}
	return "", fmt.Errorf("não foi possível encontrar o endereço MAC")
}

func main() {
	h, err := hostName()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hostname: %s\n", h)

	mac, err := getMACAddress()
	if err != nil {
		log.Fatal("Erro pega ro MAC: ", err)
	}

	fmt.Printf("MAC: %s\n", mac)
}
