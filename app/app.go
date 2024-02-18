package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando"
	app.Usage = "Busca IPs e nome de servidor na internet"

	local := localIp()

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: local.String(),
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "busca ips de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca nome de servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func localIp() net.IP {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}

func buscarIps(c *cli.Context) {
	// pegando o host
	host := c.String("host")
	// busca de ips
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
