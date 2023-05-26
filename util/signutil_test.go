package util

import (
	"com.gientech/selection/pkg/config"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestCreateSign(t *testing.T) {
	config.NewConfig(".././config")
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("timestamp:", timestamp)
	sign := CreateSign(config.GetProfile().Server.SignToken, timestamp)
	log.Println("sign:", sign)
}
