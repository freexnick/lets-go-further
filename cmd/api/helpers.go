package main

import (
	"errors"
	"net/http"
	"strconv"
)

func (app *application) readIDParam(r *http.Request) (uint64, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id number")
	}
	return id, nil
}
