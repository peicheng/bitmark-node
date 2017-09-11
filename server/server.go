package server

import (
	"crypto/tls"
	"net/rpc/jsonrpc"

	"github.com/bitmark-inc/bitmark-node/services"
	"github.com/bitmark-inc/bitmarkd/rpc"
	"github.com/gin-gonic/gin"
)

type ServiceOptionRequest struct {
	Option string `json:"option"`
}

type WebServer struct {
	Bitmarkd services.Service
	Prooferd services.Service
}

func NewWebServer(bitmarkd, prooferd services.Service) *WebServer {
	return &WebServer{
		Bitmarkd: bitmarkd,
		Prooferd: prooferd,
	}
}

func (ws *WebServer) BitmarkdStartStop(c *gin.Context) {
	var req ServiceOptionRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.String(400, "can not parse action option")
		return
	}

	err = nil
	switch req.Option {
	case "start":
		err = ws.Bitmarkd.Start()
	case "stop":
		err = ws.Bitmarkd.Stop()
	case "status":
		c.JSON(200, map[string]interface{}{
			"ok":     1,
			"result": ws.Bitmarkd.Status(),
		})
		return
	case "info":
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}

		conn, err := tls.Dial("tcp", "127.0.0.1:2130", tlsConfig)
		if nil != err {
			c.String(500, "can not parse action option")
			return
		}

		client := jsonrpc.NewClient(conn)
		defer client.Close()

		var reply rpc.InfoReply
		if err := client.Call("Node.Info", rpc.InfoArguments{}, &reply); err != nil {
			c.String(500, "Node.Info error: %s\n", err.Error())
			return
		}

		c.JSON(200, map[string]interface{}{
			"ok":     1,
			"result": reply,
		})
		return
	default:
		c.String(400, "invalid option")
		return
	}

	if err != nil {
		c.JSON(500, map[string]interface{}{
			"ok":  0,
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"ok": 1,
		})
	}
}

func (ws *WebServer) ProoferdStartStop(c *gin.Context) {
	var req ServiceOptionRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.String(400, "can not parse action option")
		return
	}

	err = nil
	switch req.Option {
	case "start":
		err = ws.Prooferd.Start()
	case "stop":
		err = ws.Prooferd.Stop()
	case "status":
		c.JSON(200, map[string]interface{}{
			"ok":     1,
			"result": ws.Prooferd.Status(),
		})
		return
	default:
		c.String(400, "invalid option")
		return
	}

	if err != nil {
		c.JSON(500, map[string]interface{}{
			"ok":  0,
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"ok": 1,
		})
	}
}

func (ws *WebServer) DiscoveryStartStop(c *gin.Context) {

}
