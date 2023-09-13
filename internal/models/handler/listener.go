package handler

import "sync"

var exitCh = make(chan struct{}, 1)
var exitedCh = make(chan struct{}, 1)
var wg = &sync.WaitGroup{}

func Close() {
	exitCh <- struct{}{}
	<-exitedCh
	wg.Wait()
}