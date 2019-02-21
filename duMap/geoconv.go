package duMap

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Coord struct {
	Lon float64
	Lat float64
}

func (c Coord) String() string {
	return fmt.Sprintf("%g,%g", c.Lon, c.Lat)
}

type Coords struct {
	Coords []Coord
}

func (c *Coords) Add(lon, lat float64) {
	c.Coords = append(c.Coords, Coord{
		Lon: lon,
		Lat: lat,
	})
}

func (c Coords) String() string {
	s := make([]string, 0)
	for _, v := range c.Coords {
		s = append(s, v.String())
	}
	return strings.Join(s, ";")
}

//单次请求可批量解析100个坐标
func (c *DuMapClient) GeoConv(coords Coords, from, to uint) (rCoords *Coords, err error) {
	params := &url.Values{}
	params.Add("coords", coords.String())
	params.Add("from", fmt.Sprint(from))
	params.Add("to", fmt.Sprint(to))
	params.Add("output", "json")

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s%s?%s", "http://lbs.baidubce.com", "/geoconv/v1/", params.Encode()), nil)
	if err != nil {
		return
	}

	req.Header.Add("x-app-id", c.appId)

	resp, err := c.bceClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result := jsoniter.Get(body)

	if result.Get("status").ToInt() == 0 {
		rCoords = &Coords{}
		array := result.Get("result")
		for i := 0; i < array.Size(); i++ {
			rCoords.Add(array.Get(0).Get("x").ToFloat64(),
				array.Get(0).Get("y").ToFloat64())
		}
	} else {
		err = fmt.Errorf("%d - %s", result.Get("status").ToInt(), result.Get("message").ToString())
	}

	return
}
