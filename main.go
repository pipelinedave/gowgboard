package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdlayher/wgctrl"
)

type Interface struct {
	Name    string `json:"name"`
	PrivateKey string `json:"private_key"`
	PublicKey string `json:"public_key"`
	Address string `json:"address"`
	ListenPort int `json:"listen_port"`
}

func main() {
	router := gin.Default()

	router.GET("/interfaces", func(c *gin.Context) {
		interfaces, err := getInterfaces()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, interfaces)
	})

	router.Run()
}

func getInterfaces() ([]Interface, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	links, err := client.Devices()
	if err != nil {
		return nil, err
	}

	var interfaces []Interface
	for _, link := range links {
		iface := Interface{
			Name:    link.Name,
			PrivateKey: link.PrivateKey.String(),
			PublicKey: link.PublicKey.String(),
			Address: link.Addresses[0].String(),
			ListenPort: link.ListenPort,
		}
		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}
