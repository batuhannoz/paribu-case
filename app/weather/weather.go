package weather

import "math/rand"

func Api1(city string) float32 {
	return rand.Float32() * 30
}

func Api2(city string) float32 {
	return rand.Float32() * 30
}
