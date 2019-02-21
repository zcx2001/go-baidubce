package duMap

import "github.com/zcx2001/go-baidubce/bceClient"

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
