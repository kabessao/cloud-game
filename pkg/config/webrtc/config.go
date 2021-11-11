package webrtc

import "github.com/giongto35/cloud-game/v2/pkg/config/encoder"

type Webrtc struct {
	DisableDefaultInterceptors bool
	IceServers                 []IceServer
	IcePorts                   struct {
		Min uint16
		Max uint16
	}
	IceIpMap   string
	SinglePort int
	LogLevel   int
}

type IceServer struct {
	Urls       string `json:"urls,omitempty"`
	Username   string `json:"username,omitempty"`
	Credential string `json:"credential,omitempty"`
}

type Config struct {
	Encoder encoder.Encoder
	Webrtc  Webrtc
}

func (w Webrtc) HasPortRange() bool  { return w.IcePorts.Min > 0 && w.IcePorts.Max > 0 }
func (w Webrtc) HasSinglePort() bool { return w.SinglePort > 0 }
func (w Webrtc) HasIceIpMap() bool   { return w.IceIpMap != "" }
