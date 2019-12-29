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
	expandtab(w)
}

func expandtab(w *acme.Win) {
	for e := range w.EventChan() {
		fmt.Printf("%+v\n", e)
		fmt.Printf("'%s'\n", string(e.Text))

		evtType := fmt.Sprintf("%s%s", string(e.C1), string(e.C2))
		switch (evtType) {
		case "KI":
    			if string(e.Text) == "	" {
				tab := []byte("    ")
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

				fmt.Println("TAB")
				fmt.Printf("%+v\n", e)
			}
		case "MX":
			fallthrough
		case "Mx":
			fallthrough
		case "Ml":
			fallthrough
		case "ML":
			w.WriteEvent(e)
		}
	}
}