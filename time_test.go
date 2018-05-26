package gopp

import (
	"log"
	"testing"
	"time"
)

func TestDur2hum0(t *testing.T) {
	layOut := "02/01/2006 15:04:05" // dd/mm/yyyy hh:mm:ss
	future, _ := time.Parse(layOut, "07/05/2118 15:12:10")

	diff := time.Until(future)
	log.Println(Dur2hum(diff))
}
