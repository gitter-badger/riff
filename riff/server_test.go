package riff

import (
	"fmt"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	//riff.AddService()
	//ch := make(chan bool)
	st := time.Now()
	c := DefaultConfig()
	s, _ := NewServer(c)
	//for i := 0; i < 100; i++ {
	//	s.AddNode(&Node{
	//		Name: "node" + strconv.Itoa(i),
	//	})
	//}
	s.Shutter()
	e := time.Now()
	fmt.Printf("time: %v,snap: %s\n", e.Sub(st), s.SnapShot)
	//<-ch
}
