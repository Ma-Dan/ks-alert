package main

import "alert-kubesphere-plugin/pkg/restapi/api"

func main() {
	api.Run()

}

type user struct {
	signalCh chan int
	alerts   []alert
}

type alert struct {
	alertCollectorChan chan alert
}
