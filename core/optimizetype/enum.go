package optimizetype

import (
	"strings"
)

type Enum int

const (
	Logo Enum = iota
	Banner
	Ad
	Thumb
)

var optimizeTypes = [...]string{
	"Logo",
	"Banner",
	"Ad",
	"Thumb",
}

func (r Enum) String() string {
	return optimizeTypes[r]
}

func GetOptimizeType(name string) Enum {
	var result Enum

	for k, v := range optimizeTypes {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = Enum(k)
			break
		}
	}

	return result
}
