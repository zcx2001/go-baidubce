package duMap

import (
	"fmt"
	"strings"
)

const baseUrl = "http://lbs.baidubce.com" //duMap基础url

type CoordType string //坐标的类型

const (
	Bd09ll  CoordType = "bd09ll"  //百度经纬度坐标
	Bd09mc  CoordType = "bd09mc"  //百度米制经纬度坐标
	Gcj02ll CoordType = "gcj02ll" //国测局坐标
	Wgs84ll CoordType = "wgs84ll" //WGS84坐标
)

type Coord struct {
	Lng float64 //经度
	Lat float64 //纬度
}

func (c Coord) String() string {
	return fmt.Sprintf("%g,%g", c.Lng, c.Lat)
}

func (c Coord) LatLng() string {
	return fmt.Sprintf("%g,%g", c.Lat, c.Lng)
}

type Coords struct {
	Coords []Coord
}

func (c *Coords) Add(lng, lat float64) {
	c.Coords = append(c.Coords, Coord{
		Lng: lng,
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
