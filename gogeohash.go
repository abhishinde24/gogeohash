package gogeohash

import (
	"errors"
	"math"
	"strings"
)

type point struct {
	lat float64
	lon float64
}
type bounds struct {
	sw point
	ne point
}

const base32 = "0123456789bcdefghjkmnpqrstuvwxyz" // (geohash-specific) Base32 map

func roundToPrecision(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value*multiplier) / multiplier
}

type GeoHash struct{}

func (g GeoHash) Encode(lat float64, lon float64, precision int) (string, error) {
	/* try to get precision by lat and log precision given
	 */
	if precision < 0 || precision > 12 {
		precision = 12
	}

	idx := 0
	bits := 0
	evenbit := true
	geohash := ""

	latMin := -90.0
	latMax := 90.0
	lonMin := -180.0
	lonMax := 180.0

	for len(geohash) < precision {

		if evenbit {
			var lonMid float64 = (lonMin + lonMax) / 2
			if lon >= lonMid {
				idx = 2*idx + 1
				lonMin = lonMid
			} else {
				idx = 2 * idx
				lonMax = lonMid
			}
		} else {
			latMid := (latMin + latMax) / 2
			if lat >= latMid {
				idx = 2*idx + 1
				latMin = latMid
			} else {
				idx = 2 * idx
				latMax = latMid
			}
		}

		evenbit = !evenbit

		bits++
		if bits == 5 {
			geohash += string(base32[idx])
			idx = 0
			bits = 0
		}
	}
	return geohash, nil
}

func (g GeoHash) Decode(geohash string) (float64, float64, error) {
	bound,err := g.bound(geohash)
	if err != nil {
		return 360.0,360.0,err
	}
	latMin := bound.sw.lat
	lonMin := bound.sw.lon
	latMax := bound.ne.lat
	lonMax := bound.ne.lon

	lat := (latMax + latMin) / 2
	lon := (lonMax + lonMin) / 2

	latPrecision := int(2 - math.Log10(latMax-latMin))
	lonPrecision := int(2 - math.Log10(lonMax-lonMin))

	roundedLat := roundToPrecision(lat, latPrecision) 
	roundedLon := roundToPrecision(lon, lonPrecision)

	return roundedLat,roundedLon,nil
}

func (g GeoHash)bound(geohash string) (bounds,error){
	if len(geohash) == 0 {
		return bounds{}, errors.New("Invalid hash")
	}
	geohash = strings.ToLower(geohash)

	evenBit := true
	latMin := -90.0
	latMax := 90.0
	lonMin := -180.0
	lonMax := 180.0

	for i :=0 ; i < len(geohash); i++ {
		chr := geohash[i]
		idx := strings.Index(base32,string(chr)) 
		if 	idx == -1 {
			return bounds{}, errors.New("Invalid char geohash") 
		}

		for k := 4 ; k >= 0 ; k-- {
			setBit := idx >> k & 1
			if evenBit {
				lonMid := (lonMax + lonMin) / 2
				if setBit == 1 {
					lonMin = lonMid	
				} else {
					lonMax = lonMid	
				}
			} else {
				latMid := (latMax + latMin) / 2
				if setBit == 1 {
					latMin = latMid	
				} else {
					latMax = latMid	
				}
			}
			evenBit = !evenBit
		}
	}
	bound := bounds{
		sw: point{latMin,lonMin},
		ne: point{latMax,lonMax},
	}
	return bound,nil
}

func (g GeoHash) adjacent(geohash string,direction string) (string){

	geohash = strings.ToLower(geohash)
	direction = strings.ToLower(direction)

	if len(geohash) == 0 {
		return ""	
	}
	if strings.Index("nsew",direction) == -1 {
		return ""
	} 

	neighbour := map[string][]string{
	"n": {"p0r21436x8zb9dcf5h7kjnmqesgutwvy", "bc01fg45238967deuvhjyznpkmstqrwx"},
	"s": {"14365h7k9dcfesgujnmqp0r2twvyx8zb", "238967debc01fg45kmstqrwxuvhjyznp"},
	"e": {"bc01fg45238967deuvhjyznpkmstqrwx", "p0r21436x8zb9dcf5h7kjnmqesgutwvy"},
	"w": {"238967debc01fg45kmstqrwxuvhjyznp", "14365h7k9dcfesgujnmqp0r2twvyx8zb"},
	}

	border := map[string][]string{
	"n": {"prxz", "bcfguvyz"},
	"s": {"028b", "0145hjnp"},
	"e": {"bcfguvyz", "prxz"},
	"w": {"0145hjnp", "028b"},
	}

	lastCh := geohash[len(geohash) - 1:];    // last character of hash
	parent := geohash[0:len(geohash) - 1]; // hash without last character
	
	Bittype := len(geohash) % 2;

	// check for edge-cases which don't share common prefix
	if (strings.Index(border[direction][Bittype],lastCh) != -1 && parent != "") {
		parent = g.adjacent(parent, direction);
	}

	// append letter for direction to parent
	return parent + string(base32[strings.Index(neighbour[direction][Bittype],lastCh)]);
} 

func (g GeoHash) Neighbours(geohash string) map[string]string{
	return map[string]string{
		"n":g.adjacent(geohash,"n"),
		"ne":g.adjacent(g.adjacent(geohash,"n"),"e"),
		"e":g.adjacent(geohash,"e"),
		"se":g.adjacent(g.adjacent(geohash,"s"),"e"),
		"s":g.adjacent(geohash,"s"),
		"sw":g.adjacent(g.adjacent(geohash,"s"),"w"),
		"w":g.adjacent(geohash,"w"),
		"nw":g.adjacent(g.adjacent(geohash,"n"),"w"),
	}
}