package controller

import (
	"fmt"
	"testing"
	"time"
)

func TestIsOnAir(t *testing.T) {
	var loc, _ = time.LoadLocation("Asia/Tokyo")
	time := time.Date(2017, 12, 11, 0, 0, 0, 0, loc)
	fmt.Println(isOnAir(time))
}
