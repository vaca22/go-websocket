package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
)

func main() {
	log.Println("fafa")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error

		}
		fmt.Println("fuck")
		ax := r.URL.Query()
		fmt.Println(ax["fuck"][0])
		go func() {
			defer conn.Close()
			s := "hello, world!"

			b := []byte(s)
			for {
				_, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				err = wsutil.WriteServerMessage(conn, op, b)
				if err != nil {
					// handle error
				}
			}
		}()
	}))
}
