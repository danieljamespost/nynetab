package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"

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
	tabWidth, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Print(err)
	}
	expandtab(w, tabWidth)
}

func expandtab(w *acme.Win, width int) {
	var tab []byte
	for i := 0; i < width; i++ {
		tab = append(tab, ' ')
	}

	for e := range w.EventChan() {
		evtType := fmt.Sprintf("%s%s", string(e.C1), string(e.C2))
		switch (evtType) {
		case "KI":
    			if string(e.Text) == "	" {
				err := w.Addr("#%d;+#1", e.Q0)
				if err != nil {
					log.Print(err)
				}
				w.Write("data", tab)

				e.C1 = 70
				e.C2 = 73
				e.Q1 = e.Q0 + utf8.RuneCount(tab)
				e.OrigQ1 = e.Q0 + utf8.RuneCount(tab)
				e.Nr = utf8.RuneCount(tab)
				e.Text = tab
				w.WriteEvent(e)
			}
		default:
			w.WriteEvent(e)
		}
	}
}