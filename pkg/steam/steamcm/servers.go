package steamcm

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/Lucino772/envelop/pkg/steam/steamweb"
)

type ServerRecord struct {
	Host string
	Port uint16
}

type Servers struct {
	records []*ServerRecord
}

func (s *Servers) Records() []*ServerRecord {
	return s.records
}

func (s *Servers) Update() error {
	if err := s.fetchFromWebApi(0, 20); err != nil {
		log.Println(err)
		if err := s.fetchFromDns(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Servers) fetchFromWebApi(cellId int, maxCnt int) error {
	client := steamweb.NewClient()
	type serverResponse struct {
		Response struct {
			ServerList   []string `json:"serverlist,omitempty"`
			ServerListWs []string `json:"serverlist_websockets,omitempty"`
			Result       int      `json:"result,omitempty"`
			Message      string   `json:"message,omitempty"`
		} `json:"response,omitempty"`
	}

	var params = make(url.Values)
	params.Add("cellid", fmt.Sprint(cellId))
	params.Add("maxcount", fmt.Sprint(maxCnt))
	var u = client.Url("ISteamDirectory", "GetCMList", 1, params)

	var serverResp serverResponse
	if err := client.CallJson(u, &serverResp); err != nil {
		return err
	}

	if serverResp.Response.Result != 1 {
		return errors.New("invalid response type")
	}

	records := make([]*ServerRecord, 0)
	for _, value := range serverResp.Response.ServerList {
		record, err := s.parseIpEndpoint(value)
		if err != nil {
			return err
		}
		records = append(records, record)
	}
	s.records = records
	return nil
}

func (s *Servers) fetchFromDns() error {
	ips, err := net.LookupHost("cm0.steampowered.com")
	if err != nil {
		return err
	}

	records := make([]*ServerRecord, 0)
	for _, ip := range ips {
		s.records = append(s.records, &ServerRecord{
			Host: ip,
			Port: 27017,
		})
	}
	s.records = records
	return nil
}

func (s *Servers) parseIpEndpoint(addr string) (*ServerRecord, error) {
	i := strings.LastIndex(addr, ":")
	if i == -1 {
		return nil, errors.New("invalid ip endpoint format")
	}

	host := addr[:i]
	port, err := strconv.Atoi(addr[i+1:])
	if err != nil {
		return nil, err
	}
	return &ServerRecord{
		Host: host,
		Port: uint16(port),
	}, nil
}
