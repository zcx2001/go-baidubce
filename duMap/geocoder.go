package duMap

import (
	"fmt"
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

	//req, err := http.NewRequest("GET",
	//	fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)
	//if err != nil {
	//	return
	//}
	//
	//req.Header.Add("x-app-id", c.appId)
	//
	//resp, err := c.bceClient.Do(req)
	//if err != nil {
	//	return
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}

	body, _ := c.do("GET",
		fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)

	fmt.Println(string(body))
}

// 逆向地理编码
func (c *DuMapClient) GeoCoderByLocation(coord Coord, coordtype, rCoordtype CoordType) {
	params := &url.Values{}
	params.Add("location", coord.LatLng())
	params.Add("coordtype", string(coordtype))
	params.Add("ret_coordtype", string(rCoordtype))
	params.Add("output", "json")

	//req, err := http.NewRequest("GET",
	//	fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)
	//if err != nil {
	//	return
	//}
	//
	//req.Header.Add("x-app-id", c.appId)
	//
	//resp, err := c.bceClient.Do(req)
	//if err != nil {
	//	return
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}

	body, _ := c.do("GET",
		fmt.Sprintf("%s%s?%s", baseUrl, "/geocoder/v2/", params.Encode()), nil)

	fmt.Println(string(body))
}
