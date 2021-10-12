package webclient

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func HTTPResponse(url string) (data *[]byte, err error) {
	var netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 6 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 6 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 6,
		Transport: netTransport,
	}
	response, _ := netClient.Get(url)
	dataB, err := ioutil.ReadAll(response.Body)
	data = &dataB
	return
}
