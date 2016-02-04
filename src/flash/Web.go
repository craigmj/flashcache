package flash

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/craigmj/commander"
	"github.com/golang/glog"
)

var TheCache Cache

func getKey(r *http.Request) string {
	return r.URL.Path
}

func CacheWeb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cacheWebGet(w, r)
		return
	}
	// POST or PUT or any other method
	var raw json.RawMessage
	js := json.NewDecoder(r.Body)
	if err := js.Decode(&raw); nil != err {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	TheCache.Add(getKey(r), raw, time.Minute)
	out := json.NewEncoder(w)
	out.Encode(map[string]string{"status": "ok"})
}

func cacheWebGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	js := json.NewEncoder(w)
	err := js.Encode(TheCache.Find(getKey(r)))
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetEnvPort() int {
	DEFAULT := 16021
	env := os.Getenv("FLASHCACHE_SERVER")
	if "" == env {
		return DEFAULT
	}
	u, err := url.Parse(env)
	if nil != err {
		return DEFAULT
	}
	i := strings.Index(u.Host, ":")
	if -1 == i {
		return DEFAULT
	}
	p, err := strconv.Atoi(u.Host[i+1:])
	if nil != err {
		return DEFAULT
	}
	return p
}

func WebCommand() *commander.Command {
	fs := flag.NewFlagSet("web", flag.ExitOnError)
	port := fs.Int("port", GetEnvPort(), "Port on which webserver should run")
	bind := fs.String("bind", "", "IP on which to bind webserver")

	return commander.NewCommand("web", "Run the cacheweb webserver",
		fs,
		func(args []string) error {
			http.HandleFunc("/__version", func(w http.ResponseWriter, r *http.Request) {
				js := json.NewEncoder(w)
				js.Encode(map[string]string{
					"version": "1.0.0",
				})
			})
			http.HandleFunc("/__dump", TheCache.WebDump)
			http.HandleFunc("/", CacheWeb)

			glog.Infof("Starting flashcache webserver on port %d", *port)
			return http.ListenAndServe(fmt.Sprintf("%s:%d", *bind, *port), nil)
		})
}
