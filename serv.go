package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Metrics struct { // constructeur d'objets metrics
	IdAgent    int       `json:"idagent"`
	Hostname   string    `json:"hostname"`
	OS         string    `json:"os"`
	Arch       string    `json:"arch"`
	CPUs       int       `json:"cpus"`
	Goroutines int       `json:"goroutines"`
	Alloc      uint64    `json:"alloc"`
	TotalAlloc uint64    `json:"total_alloc"`
	Sys        uint64    `json:"sys"`
	UptimeSec  float64   `json:"uptime_sec"`
	Timestamp  time.Time `json:"timestamp"`
	State      string    `json:"state"`
}

var metricsStorage = make(map[int]Metrics) // map de stockage des metriques reçus depuis les agents

func main() {

	http.HandleFunc("POST /RPC", receiveMetrics) // endpoint /RPC en POST pour les agents

	http.HandleFunc("GET /view", func(w http.ResponseWriter, r *http.Request) { // endpoint /view en GET pour les clients
		for id, item := range metricsStorage { // ajout de l'état DOWN ou WARNING ou UP pour chaque agent
			if time.Since(item.Timestamp) > 120*time.Second { // si plus de nouvelles depuis plus de 120 secondes
				item.State = "DOWN"
			} else if time.Since(item.Timestamp) > 80*time.Second { // si plus de nouvelles depuis plus de 80 secondes
				item.State = "WARNING"
			} else {
				item.State = "UP"
			}
			metricsStorage[id] = item // mise à jour de la métrique dans le map
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metricsStorage)
	})

	log.Println("Serveur central démarré sur :8443")
	err := http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", nil) // démarrage du serveur central
	if err != nil {
		log.Fatal(err)
	}
}

func receiveMetrics(w http.ResponseWriter, r *http.Request) { // traitement de la réception des métriques
	var m Metrics

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil { // décodage des métriques reçus
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}
	metricsStorage[m.IdAgent] = m // ajout au map de stockage des métriques
	log.Println(metricsStorage)
	log.Printf("[OK] Reçu : %+v\n", m)

	w.WriteHeader(http.StatusOK)
}
