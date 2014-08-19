package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "github.com/diebels727/logbot"
  "strings" 
)

func IRCHandler(response http.ResponseWriter,request *http.Request) {
  params := mux.Vars(request)
  go func(params map[string]string) {
    server := params["server"]
    channel := params["channel"]

    server = strings.Replace(server,"-",".",-1)

    fmt.Println("Launching logbot for: ",server,"/",channel)

    log_bot := logbot.New(server,channel,"6667","redbot","lukewarm")
    log_bot.RunAndLoop()
  }(params)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/{server}/{channel}",IRCHandler)
  http.Handle("/",router)
  http.ListenAndServe(":8080",nil)
}
