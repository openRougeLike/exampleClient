package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max+1-min) + min
}

func (coords Coords) String() string {
	return fmt.Sprintf("%v,%v", coords[0], coords[1])
}

func StringToCoords(str string) Coords {
	split := strings.SplitN(str, ",", 2)
	out := Coords{}

	for i, inf := range split {
		ret, err := strconv.ParseInt(inf, 10, 64)
		PanicIfErr(err)
		out[i] = int(ret)
	}

	return out
}

func LocOk(loc int) bool {
	return 0 <= loc && loc <= DIMENSIONS-1
}
