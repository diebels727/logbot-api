package main
import (
  "database/sql"
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "github.com/diebels727/logbot"
  "strings"
  // _ "github.com/mattn/go-sqlite3"
  "path"
  "encoding/json"
)

type Event struct {
  Timestamp int
  Username string
  Host string
  Message string
}

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

func EventsIndex(response http.ResponseWriter,request *http.Request) {
  response.Header().Set("Content-Type", "application/json")

  params := mux.Vars(request)
  db_file := "db.sqlite3"
  db_path := path.Join(params["server"],params["channel"],db_file)
  db, err := sql.Open("sqlite3",db_path)
  if err != nil {
    fmt.Println("Cannot get handle to DB, aborting...")
    return
  }
  defer db.Close()


  rows,err := db.Query(`SELECT timestamp,username,host,message FROM events;`)
  if err != nil {
    fmt.Println("Statement invalid, aborting...")
    return
  }

  defer rows.Close()

  events := make([]Event,0)

  for rows.Next() {
    var timestamp int
    var username,host,message string
    rows.Scan(&timestamp,&username,&host,&message)
    event := Event{timestamp,username,host,message}
    events = append(events,event)
  }
  // events = Events(events)
  bytes,err := json.Marshal(events)
  if err != nil {
    fmt.Println("Error marshalling events, aborting...")
    return
  }
  jsonEvent := string(bytes)
  fmt.Fprint(response,"{\"events\": "+jsonEvent+"}")
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/{server}/{channel}",IRCHandler)
  router.HandleFunc("/{server}/{channel}/events",EventsIndex)
  http.Handle("/",router)
  http.ListenAndServe(":8080",nil)
}
