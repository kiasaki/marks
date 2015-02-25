package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kiasaki/batbelt/http/middlewares"
	"github.com/kiasaki/batbelt/http/mm"

	"github.com/kiasaki/marks/data"
)

var db *sql.DB

func SetDB(d *sql.DB) {
	db = d
}

func DB() *sql.DB {
	return db
}

// Entrypoint
func main() {
	SetCfg(ParseFlag())
	SetDB(data.Connect(Cfg().PostgresURL))

	if Cfg().CreateDb {
		createDb(DB())
	} else {
		handleHTTP()
	}
}

func handleHTTP() {
	registerHandlers()

	var stringPort = strconv.Itoa(Cfg().Port)
	log.Println("Started listening on port " + stringPort)
	log.Fatal(http.ListenAndServe(":"+stringPort, nil))
}

func registerHandlers() {
	mmChain := mm.New(middlewares.LogWithTiming)

	// Add basic auth middleware if configuration specifies user & pass
	u, p := Cfg().BasicAuthUser, Cfg().BasicAuthPass
	if u != "" && p != "" {
		mmChain = mmChain.Append(middlewares.BasicAuth(u, p))
	}

	http.Handle("/", mmChain.Then(http.HandlerFunc(routeRequest)))
}

func routeRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/new", 302)
	} else if r.URL.Path == "/styles.css" {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(stylesContents))
	} else if r.URL.Path == "/new" {
		if r.Method != "POST" {
			handleNewPage(w, r)
		} else {
			handleNewPagePost(w, r)
		}
	} else if r.URL.Path[:3] == "/n/" && strings.HasSuffix(r.URL.Path, "/delete") {
		handleDeletePage(w, r)
	} else if r.URL.Path[:3] == "/n/" && strings.HasSuffix(r.URL.Path, "/edit") {
		if r.Method != "POST" {
			handleEditPage(w, r)
		} else {
			handleEditPagePost(w, r)
		}
	} else if r.URL.Path[:3] == "/n/" {
		handleViewPage(w, r)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Page not found"))
	}
}
