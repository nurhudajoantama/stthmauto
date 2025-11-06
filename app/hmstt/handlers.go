package hmstt

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurhudajoantama/stthmauto/app/server"
)

type hmsttHandler struct {
	service   *hmsttService
	templates *template.Template
}

func RegisterHandlers(s *server.Server, svc *hmsttService) {
	templates := template.Must(template.ParseGlob("app/hmstt/views/*.html"))

	h := &hmsttHandler{
		service:   svc,
		templates: templates,
	}
	srv := s.GetRouter()

	hmsttGroup := srv.PathPrefix("/hmstt").Subrouter()
	hmsttGroup.HandleFunc("/", h.handleIndex)
	hmsttGroup.HandleFunc("/statehtml/{type}/{key}", h.handleGetStateHTML).Methods("GET")
	hmsttGroup.HandleFunc("/setstatehtml/{type}/{key}", h.handleSetStateHTML).Methods("POST")

	hmsttGroup.HandleFunc("/getstatevalue/{type}/{key}", h.handleGetState).Methods("GET")
}

func (h *hmsttHandler) handleGetState(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	key := p["key"]
	tipe := p["type"]

	result, err := h.service.GetState(r.Context(), tipe, key)
	if err != nil {
		returnErrorState(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

// handleIndex serves the main (and only) HTML page
func (h *hmsttHandler) handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// handleState is another HTMX endpoint returning an HTML string
func (h *hmsttHandler) handleGetStateHTML(w http.ResponseWriter, r *http.Request) {

	p := mux.Vars(r)
	key := p["key"]
	tipe := p["type"]

	var templateData = map[string]string{
		"tipe": tipe,
		"key":  key,
	}

	results, err := h.service.GetState(r.Context(), tipe, key)
	if err != nil {
		templateData["state"] = ERR_STRING
	} else {
		templateData["state"] = results
	}

	templateFileName, ok := TYPE_TEMPLATES[tipe]
	if !ok {
		templateFileName = HTML_TEMPLATE_NOTFOUND_TYPE
	}

	if err := h.templates.ExecuteTemplate(w, templateFileName, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *hmsttHandler) handleSetStateHTML(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	key := p["key"]
	tipe := p["type"]

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR PARSE FORM"))
		return
	}

	value := r.FormValue("value")
	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("STATE VALUE EMPTY"))
		return
	}

	if err := h.service.SetState(r.Context(), tipe, key, value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR SET STATE"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func returnErrorState(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(ERR_STRING))
}
