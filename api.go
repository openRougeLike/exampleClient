package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BASE = "http://localhost:3000/game"

func Init() Map {
	Fetch(http.MethodPost, "/self/new", nil)
	init, _ := Fetch(http.MethodGet, "/map/", nil)
	return BToMap(init)
}

type User struct {
	Map Map `json:"map"`
}

var client = http.Client{
	Timeout: 10 * time.Second,
}

type Direction int8

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
	DirCur
)

var DirectionEnum = map[Direction]Coords{
	DirLeft:  {0, -1},
	DirRight: {0, 1},
	DirUp:    {-1, 0},
	DirDown:  {1, 0},
	DirCur:   {0, 0},
}

func (m *Map) Move(dir Direction, num int) *APIError {
	b, err := Fetch(http.MethodPost, fmt.Sprintf("/map/move/%v/%v", dir, num), nil)
	if err != nil {
		if err.Msg == "there is a wall in the way" {
			fmt.Println(m.Layout[m.User[0]+DirectionEnum[dir][0]][m.User[1]+DirectionEnum[dir][1]])
		}
		return err
	}
	*m = BToMap(b)
	return nil
}

func (m *Map) Action(dir Direction) *APIError {
	b, err := Fetch(http.MethodPost, fmt.Sprintf("/map/action/%v", dir), nil)
	if err != nil {
		return err
	}
	ret := BToUsr(b)
	*m = ret.Map
	return nil
}

func BToMap(b []byte) Map {
	out := Map{}
	err := json.Unmarshal(b, &out)
	PanicIfErr(err)
	return out
}

func BToUsr(b []byte) User {
	out := User{}
	err := json.Unmarshal(b, &out)
	PanicIfErr(err)
	return out
}

type APIError struct {
	Msg string `json:"err"`
}

func Fetch(method string, uri string, body io.Reader) ([]byte, *APIError) {
	req, err := http.NewRequest(method, BASE+uri, body)
	PanicIfErr(err)

	req.Header.Set("Authorization", "Bearer cum")

	res, err := client.Do(req)
	PanicIfErr(err)
	resBody, err := io.ReadAll(res.Body)
	PanicIfErr(err)

	if res.StatusCode != 200 {
		if res.StatusCode == 404 {
			panic("404 while requesting '" + uri + "'")
		} else if res.StatusCode == 401 {
			panic("401!")
		} else if res.StatusCode == 429 {
			time.Sleep(10 * time.Second)
			return Fetch(method, uri, body)
		} else if res.StatusCode == 400 {
			errBody := APIError{}
			json.Unmarshal(resBody, &errBody)
			return nil, &errBody
		} else {
			panic(fmt.Sprintf("Unknown status code %v detected!", res.StatusCode))
		}
	}

	return resBody, nil
}
