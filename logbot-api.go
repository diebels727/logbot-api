package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "github.com/logbot"
)

func IRCHandler(response http.ResponseWriter,request *http.Request) {
  params := mux.Vars(request)
  fmt.Println("server: ",params["server"]," channel: ",params["channel"])
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/{server}/{channel}",IRCHandler)
  http.Handle("/",router)
  http.ListenAndServe(":8080",nil)
}
