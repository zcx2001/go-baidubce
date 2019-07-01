package duMap

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/url"
)

// 地理编码
//
// address 待解析的地址。最多支持84个字节。
//
// city 地址所在的城市名。用于指定上述地址所在的城市，当多个城市都有上述地址时，该参数起到过滤作用，但不限制坐标召回城市。
//
// rCoordtype 添加后返回国测局经纬度坐标或百度米制坐标
func (c *DuMapClient) GeoCoderByAddress(address, city string, rCoordtype CoordType) {
	params := &url.Values{}
	params.Add("address", address)
	params.Add("city", city)
	params.Add("ret_coordtype", string(rCoordtype))
	params.Add("output", "json")

	body, _ := c.do("GET",
		fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)

	fmt.Println(string(body))
}

// 逆向地理编码
func (c *DuMapClient) GeoCoderByLocation(coord Coord, coordtype, rCoordtype CoordType) (country, province, city, address string, err error) {
	params := &url.Values{}
	params.Add("location", coord.LatLng())
	params.Add("coordtype", string(coordtype))
	params.Add("ret_coordtype", string(rCoordtype))
	params.Add("output", "json")

	body, err := c.do("GET",
		fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)
	if err != nil {
		return
	}

	//fmt.Println(string(body))

	result := jsoniter.Get(body)

	if result.Get("status").ToInt() == 0 {
		country = result.Get("result", "addressComponent", "country").ToString()
		province = result.Get("result", "addressComponent", "province").ToString()
		city = result.Get("result", "addressComponent", "city").ToString()
		address = result.Get("result", "formatted_address").ToString()
	} else {
		err = fmt.Errorf("%d - %s", result.Get("status").ToInt(), result.Get("message").ToString())
	}

	return
}
