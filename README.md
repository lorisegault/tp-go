# Documentation

## Fichier `agent.go`

Dans ce fichier on retrouve le contenu suivant : 

- Un import des packages
- Un constructeur d'objets `Metrics`
- Une fonction `main` avec : 
    - Un interval utilisé dans la boucle for
    - Un client http
    - Une récupération des métriques et la création d'un objet correspondant
    - Un envoi des métriques en `POST` vers le serveur central `https://localhost:8443/RPC`

## Fichier `cli.go`

Dans ce fichier on retrouve le contenu suivant : 

- Un import des packages
- Un client http 
- Un `GET` vers le serveur central `https://localhost:8443/view`

## Fichier `serv.go`

Dans ce fichier on retrouve le contenu suivant : 

- Un import des packages
- Un constructeur d'objets `Metrics`
- Un map de stockage des objets `Metrics` qui contiendra les métriques reçus
- Un endpoint `POST /RPC` qui lance l'enregistrement des métriques reçus dans le map de stockage
- Un endpoint `GET /view` qui permet de renvoyer les objets `Metrics` enregistré dans le map de stockage
- Un serveur http lançé sur le port `8443`


## Utilisation du projet

Pour lançer le projet il faut utiliser les commandes suivantes : 

- `go run serv.go` -> pour lançer le serveur
- `go run agent.go` -> pour lançer l'agent
- `go run cli.go` -> pour lançer le client