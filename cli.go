package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get("https://localhost:8443/view") // url de get
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(content)) // affichage du contenu reçu depuis le serveur

}
