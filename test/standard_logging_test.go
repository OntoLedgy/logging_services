package test

import (
	"log"
	"os"
	"testing"
)

func TestStandardLogger(t *testing.T) {

	f, _ := os.OpenFile("test", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	log.SetOutput(f)
	log.Println("hello ")
}
