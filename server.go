package main

import (
	"log"
	"net/http"
	"strconv"
)

var cycles, size, nframes, delay = 5, 100, 64, 8
var res = 0.001

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, size, nframes, delay, res = parse(r)
		Lissajous(w, cycles, size, nframes, delay, res)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parse(r *http.Request) (int, int, int, int, float64) {
	q := r.URL.Query()
	if q.Has("cycles") {
		var err error
		cycles, err = strconv.Atoi(q["cycles"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("size") {
		var err error
		size, err = strconv.Atoi(q["size"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("nframes") {
		var err error
		nframes, err = strconv.Atoi(q["nframes"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("delay") {
		var err error
		delay, err = strconv.Atoi(q["delay"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("res") {
		var err error
		res, err = strconv.ParseFloat(q["res"][0], 64)
		if err != nil {
			panic(err)
		}
	}
	return cycles, size, nframes, delay, res
}
