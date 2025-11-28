package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
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
}

func main() {
	interval := 50 * time.Second // définition d'un interval de 50 secondes
	startTime := time.Now()      // heure de démarrage de l'agent

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	hostname, _ := os.Hostname() // recupération hostname

	fmt.Println("Agent démarré. Envoi toutes les", interval)

	for {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		uptime := time.Since(startTime).Seconds() // recupération uptime

		metrics := Metrics{ // création d'un objet avec ses attributs
			IdAgent:    1,
			Hostname:   hostname,
			OS:         runtime.GOOS,
			Arch:       runtime.GOARCH,
			CPUs:       runtime.NumCPU(),
			Goroutines: runtime.NumGoroutine(),
			Alloc:      mem.Alloc,
			TotalAlloc: mem.TotalAlloc,
			Sys:        mem.Sys,
			UptimeSec:  uptime,
			Timestamp:  time.Now(),
		}

		jsonData, err := json.Marshal(metrics)
		if err != nil {
			fmt.Println("Erreur JSON:", err)
			continue
		}
		// envoi des données au serveur central
		resp, err := client.Post("https://localhost:8443/RPC",
			"application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Erreur HTTP:", err)
		} else {
			resp.Body.Close()
			fmt.Println("Données envoyées OK")
		}

		time.Sleep(interval) // attente de la durée de l'interval
	}
}
