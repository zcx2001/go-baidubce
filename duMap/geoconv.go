package duMap

import (
	"fmt"
	"github.com/json-iterator/go"
	"net/url"
)

// 坐标转换服务
// 单次请求可批量解析100个坐标
func (c *DuMapClient) GeoConv(coords Coords, from, to uint) (rCoords *Coords, err error) {
	params := &url.Values{}
	params.Add("coords", coords.String())
	params.Add("from", fmt.Sprint(from))
	params.Add("to", fmt.Sprint(to))
	params.Add("output", "json")

	body, err := c.do("GET",
		fmt.Sprintf("%s%s?%s", baseUrl, "/geoconv/v1/", params.Encode()), nil)
	if err != nil {
		return
	}

	result := jsoniter.Get(body)

	if result.Get("status").ToInt() == 0 {
		rCoords = &Coords{}
		array := result.Get("result")
		for i := 0; i < array.Size(); i++ {
			rCoords.Add(array.Get(i).Get("x").ToFloat64(),
				array.Get(i).Get("y").ToFloat64())
		}
	} else {
		err = fmt.Errorf("%d - %s", result.Get("status").ToInt(), result.Get("message").ToString())
	}

	return
}
