package steamcm

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

type SteamClient struct {
	conn *SteamConnection
	errg *errgroup.Group
}

func NewSteamClient(conn *SteamConnection) *SteamClient {
	client := &SteamClient{
		conn: conn,
		errg: new(errgroup.Group),
	}
	return client
}

func (client *SteamClient) Connect() error {
	s := new(Servers)
	if err := s.Update(); err != nil {
		return err
	}
	server := s.Records()[0]

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))
	if err != nil {
		return err
	}
	return client.start(conn)
}

func (client *SteamClient) start(conn net.Conn) error {
	client.errg.Go(func() error {
		defer log.Println("client: done reading")
		for {
			buff := make([]byte, 1024)
			log.Println("client: reading...")
			nr, err := conn.Read(buff)
			if err != nil {
				return err
			}
			log.Println("client: recv data", nr)
			if err := client.conn.ProcessBytes(buff[:nr]); err != nil {
				return err
			}
		}
	})
	client.errg.Go(func() error {
		defer log.Println("client: done writing")
		for {
			buff := make([]byte, 1024)
			log.Println("client: waiting for data to send")
			nr, err := client.conn.Read(buff)
			if err != nil {
				return err
			}
			log.Println("client: send data", len(buff[:nr]))
			if _, err := conn.Write(buff[:nr]); err != nil {
				return err
			}
		}
	})
	return nil
}
