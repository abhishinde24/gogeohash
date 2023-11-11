package gogeohash

import (
	"fmt"
	"testing"
)

func TestEncodeHash(t *testing.T) {
    g := GeoHash{}
	lat1 := 48.8566
	lon1 := 2.35222
	prec := 8
	exp_hash := "u09tvw0f"
	hash, err := g.Encode(lat1, lon1, prec)
	if err != nil {fmt.Printf(hash);}
    if hash != exp_hash {
        t.Errorf("Expected 5, but got %s", hash)
    }
}
func TestDecodeHash(t *testing.T) {
    g := GeoHash{}
	lat1 := 48.85663
	lon1 := 2.35228
	geo_hash := "u09tvw0f"
	lat,lon, err := g.Decode(geo_hash)
	if err != nil {fmt.Printf("lat %g , log %g Error %s",lat,lon,err);}
    if lat1 != lat && lon1 != lon {
        t.Errorf("Expected %g & %g, but got %g & %g",lat1,lon1,lat,lon)
    }
}

func TestHashNeighbours(t *testing.T) {
    g := GeoHash{}
	geohash := "u09tvw0f"
	wantValue := map[string]string{
		"n": "u09tvw0g",
		"ne": "u09tvw15",
		"e": "u09tvw14",
		"se": "u09tvw11",
		"s": "u09tvw0c",
		"sw": "u09tvw09",
		"w": "u09tvw0d",
		"nw": "u09tvw0e",
	}
	gotValue := g.Neighbours(geohash)
    if fmt.Sprint(wantValue) != fmt.Sprint(gotValue) {
        t.Errorf("Test failed")
    }
}

func BenchmarkEncodeHash(b *testing.B) {
    g := GeoHash{}
	lat1 := 48.8566
	lon1 := 2.35222
	prec := 8
	for i :=0; i < b.N ; i++{
		g.Encode(lat1, lon1, prec)
	}
}

func BenchmarkDecodeHash(b *testing.B) {
    g := GeoHash{}
	geo_hash := "u09tvw0f"
	for i :=0; i < b.N ; i++{
		g.Decode(geo_hash)
	}
}

func BenchmarkHashNeighbours(b *testing.B) {
    g := GeoHash{}
	geo_hash := "u09tvw0f"
	for i :=0; i < b.N ; i++{
		g.Neighbours(geo_hash)
	}
}