//by: isca
package main

import (
  "fmt"
  "net/http"
  "github.com/mgutz/ansi"
)

//Show Splash on start
func splash() {
  pink := ansi.ColorFunc("magenta+bh:black")
  cyan := ansi.ColorCode("cyan+h:black")
  reset := ansi.ColorCode("reset")
  green := ansi.ColorFunc("cyan+Bbh:blue+h")
  help := pink("Listening on: 3000\nHousing: /www\n")
  start := green("Started...")
  fmt.Print(start)
  fmt.Println(cyan,
`
       __________________
      /\  ______________ \
     /::\ \ZZZZZZZZZZZZ/\ \
    /:/\.\ \        /:/\:\ \
   /:/Z/\:\ \      /:/Z/\:\ \
  /:/Z/__\:\ \____/:/Z/  \:\ \
 /:/Z/____\:\ \___\/Z/    \:\ \
 \:\ \ZZZZZ\:\ \ZZ/\ \     \:\ \
  \:\ \     \:\ \ \:\ \     \:\ \
   \:\ \     \:\ \_\;\_\_____\;\ \
    \:\ \     \:\_________________\
     \:\ \    /:/ZZZZZZZZZZZZZZZZZ/
      \:\ \  /:/Z/    \:\ \  /:/Z/
       \:\ \/:/Z/      \:\ \/:/Z/
        \:\/:/Z/________\;\/:/Z/
         \::/Z/_____isca___\/Z/
          \/ZGOZZFILEZZSERVEZ/
 `,reset)
 fmt.Print(help,reset)

 }


//Create a log outputs for default nil server 
func Log(handler http.Handler) http.Handler {
  lime := ansi.ColorCode("green+h:black")
  reset := ansi.ColorCode("reset")
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf(lime)
    fmt.Printf("Req: %s %s %s %s \n", r.Host, r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
    fmt.Printf(reset)
  })
}

func main() {
  splash()
  http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/www"))))
  http.ListenAndServe(":3000", Log(http.DefaultServeMux))
}

