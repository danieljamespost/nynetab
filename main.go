package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"9fans.net/go/acme"
)

func main() {
	wId, err := strconv.Atoi(os.Getenv("winid"))
	if err != nil {
		log.Print(err)
	}
	w, err := acme.Open(wId, nil)
	if err != nil {
		log.Print(err)
	}
	for e := range events(w) {
		fmt.Printf("%+v\n", e)

	}
}

func events(w *acme.Win) <-chan string {
	c := make(chan string, 10)
	go func() {
		for e := range w.EventChan() {
			evtType := fmt.Sprintf("%d%d", e.C1, e.C2)
			switch (evtType) {
			case "7573":	// KI
    				if string(e.Text) == "	" {
					fmt.Println("TAB")
	// 				fmt.Printf("%+v\n", e)
					err := w.Addr("#%d+#1", e.Q0)
					if err != nil {
						log.Print(err)
					}
					w.Write("data", []byte("    "))
					break
				}
			default:
				w.WriteEvent(e)
			}
		}
		w.CloseFiles()
		close(c)
	}()
	return c
}