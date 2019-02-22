package duMap

import (
	"github.com/zcx2001/go-baidubce/bceClient"
	"io"
	"io/ioutil"
	"net/http"
)

type DuMapClient struct {
	appId     string
	bceClient *bceClient.BceClient
}

func NewDuMapClient(appId string, bceClient *bceClient.BceClient) *DuMapClient {
	return &DuMapClient{
		appId:     appId,
		bceClient: bceClient,
	}
}

func (c *DuMapClient) GetAppId() string {
	return c.appId
}

func (c *DuMapClient) do(method, url string, body io.Reader) (rbody []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req.Header.Add("x-app-id", c.appId)

	resp, err := c.bceClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	rbody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
