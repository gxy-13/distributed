package log

import (
	"io"
	stdlog "log"
	"net/http"
	"os"
)

var log *stdlog.Logger

type fileLog string

func (f fileLog) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(f), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(data)
}

func Run(destination string) {
	log = stdlog.New(fileLog(destination), "go: ", stdlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			stdlog.Println(msg)
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	stdlog.Println(message)
	stdlog.Println("start write")
	log.Printf("%v\n", message)

	stdlog.Println("end write")
}
