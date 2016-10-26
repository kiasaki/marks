package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kiasaki/batbelt/mst"

	"github.com/kiasaki/marks/data"
)

func handleNewPage(w http.ResponseWriter, r *http.Request) {
	contents, err := RenderNewPage(Page{
		Notes:       data.MustGetAllNotes(DB()),
		Title:       "",
		Body:        "",
		PostbackURL: "/new",
	})
	mst.MustNotErr(err)
	w.Write(contents)
}

func handleNewPagePost(w http.ResponseWriter, r *http.Request) {
	title, body := r.FormValue("title"), r.FormValue("body")
	if savedNote, err := data.CreateNote(DB(), title, body); err != nil {
		contents, rerr := RenderNewPage(Page{
			Notes:       data.MustGetAllNotes(DB()),
			Title:       title,
			Body:        err.Error() + "\n\n" + body,
			PostbackURL: "/new",
		})
		mst.MustNotErr(rerr)
		w.WriteHeader(400)
		w.Write(contents)
	} else {
		http.Redirect(w, r, "/n/"+savedNote.Id, 302)
	}
}

func handleViewPage(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Path[3:]
	// Guard ourselves again matching /n/gibberish (check for uuid length)
	if len(noteId) != 36 {
		w.WriteHeader(404)
		w.Write([]byte("404 - Page not found"))
		return
	}

	note, err := data.GetNote(DB(), noteId)
	if err == sql.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("404 - Page not found"))
	} else if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 - Error fetching note"))
		log.Fatal(err)
	} else {
		contents, err := RenderViewPage(Page{
			Notes:       data.MustGetAllNotes(DB()),
			Id:          note.Id,
			Title:       note.Title,
			Body:        note.Body,
			PostbackURL: "/n/" + note.Id,
		})
		mst.MustNotErr(err)
		w.Write(contents)
	}
}

func handleDeletePage(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Path[3 : len(r.URL.Path)-7]
	err := data.DeleteNote(DB(), noteId)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		http.Redirect(w, r, "/new", 302)
	}
}

func handleEditPage(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Path[3 : len(r.URL.Path)-5]
	note, err := data.GetNote(DB(), noteId)
	if err == sql.ErrNoRows {
		w.WriteHeader(404)
		w.Write([]byte("404 - Page not found"))
	} else if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 - Error fetching note"))
		log.Fatal(err)
	} else {
		contents, err := RenderEditPage(Page{
			Notes:       data.MustGetAllNotes(DB()),
			Id:          note.Id,
			Title:       note.Title,
			Body:        note.Body,
			PostbackURL: "/n/" + note.Id + "/edit",
		})
		mst.MustNotErr(err)
		w.Write(contents)
	}
}

func handleEditPagePost(w http.ResponseWriter, r *http.Request) {
	noteId := r.URL.Path[3 : len(r.URL.Path)-5]
	note := data.Note{
		Id:    noteId,
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}
	_, err := data.UpdateNote(DB(), note)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 - Error: " + err.Error()))
	} else {
		http.Redirect(w, r, "/n/"+noteId, 302)
	}
}
