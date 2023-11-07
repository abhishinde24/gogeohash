package gogeohash

const base32 = "0123456789bcdefghjkmnpqrstuvwxyz" // (geohash-specific) Base32 map

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
