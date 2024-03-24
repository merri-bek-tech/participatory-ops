package main

import "time"

type Component struct {
	Uuid string `json:"uuid"`
	At   int64  `json:"at"`
}

func componentCache(cache map[string]Component) {
	time.Sleep(20 * time.Second)

	cache["f08b7172-36d8-447f-85e1-41403d2730c8"] = Component{
		Uuid: "f08b7172-36d8-447f-85e1-41403d2730c8",
		At:   1620000000,
	}
}
