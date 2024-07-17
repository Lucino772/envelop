package steamcm

import (
	"encoding/binary"
	"io"

	"github.com/Lucino772/envelop/pkg/steam"
	"github.com/Lucino772/envelop/pkg/steam/steamlang"
)

type MsgClientLogOnResponse struct {
	Result                    steamlang.EResult
	OutOfGameHeartbeatRateSec int32
	InGameHeartbeatRateSec    int32
	ClientSuppliedSteamId     steam.SteamId
	IpPublic                  uint32
	ServerRealTime            uint32
}

func (m *MsgClientLogOnResponse) ReadFrom(r io.Reader) (int64, error) {
	if err := binary.Read(r, binary.LittleEndian, m); err != nil {
		return -1, err
	}
	return int64(binary.Size(m)), nil
}

func (m *MsgClientLogOnResponse) WriteTo(w io.Writer) (int64, error) {
	if err := binary.Write(w, binary.LittleEndian, m); err != nil {
		return -1, err
	}
	return int64(binary.Size(m)), nil
}
