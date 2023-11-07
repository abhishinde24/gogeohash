package gogeohash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
    g := GeoHash{}
	lat1 := 48.8566
	log1 := 2.35222
	prec := 8
	exp_hash := "u09tvw0f"
	hash, err := g.encode(lat1, log1, prec)
	if err != nil {fmt.Printf(hash);}
    if hash != exp_hash {
        t.Errorf("Expected 5, but got %s", hash)
    }
}
