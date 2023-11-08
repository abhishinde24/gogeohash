# gogeohash

gogeohash is go package to get geohash from latitute and logitute and vise versa.

## Installation

```bash
# Install the core.
go get github.com/abshinde24/gogeohash
```

## Usage

```go
import "github.com/abshinde24/gogeohash"

// generate geohash from Lat and Lon
g := geohash.GeoHash{}
hash := g.encode(48.8566,2.35222)

// getting lat and log from geohash
lat, log, err := g.Decode("u09tvw0f")


```

## Reference
- pakage code is mostly referenced code from below git repository.\
<https://github.com/davetroy/geohash-js>\
<https://github.com/chrisveness/latlon-geohash/blob/master/latlon-geohash.js>
