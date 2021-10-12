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
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	response, _ := netClient.Get(url)
	dataB, err := ioutil.ReadAll(response.Body)
	data = &dataB
	return
}
