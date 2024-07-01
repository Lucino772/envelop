package steamcm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	type serverResponse struct {
		Response struct {
			ServerList   []string `json:"serverlist,omitempty"`
			ServerListWs []string `json:"serverlist_websockets,omitempty"`
			Result       int      `json:"result,omitempty"`
			Message      string   `json:"message,omitempty"`
		} `json:"response,omitempty"`
	}

	var u = url.URL{
		Scheme: "https",
		Host:   "api.steampowered.com",
		Path:   fmt.Sprintf("/ISteamDirectory/GetCMList/v%d/", 1),
	}
	q := u.Query()
	q.Add("cellid", fmt.Sprint(cellId))
	q.Add("maxcount", fmt.Sprint(maxCnt))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var serverResp serverResponse
	if err := json.Unmarshal(data, &serverResp); err != nil {
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
