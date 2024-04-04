package servidor

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

type Planet struct {
    Name   string `json:"name"`
    Booty  int    `json:"booty"`
    Captain string `json:"captain"`
}

var planets = make(map[string]int)
var mutex sync.Mutex

func main() {
    initializePlanets()
    http.HandleFunc("/assign-booty", handleAssignBooty)
    fmt.Println("Servidor en ejecución en el puerto 8080...")
    http.ListenAndServe(":8080", nil)
}

func initializePlanets() {
    // Inicializar planetas con botines aleatorios
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 6; i++ {
        planetName := string('A' + i)
        planets[planetName] = rand.Intn(10)
    }
}

func handleAssignBooty(w http.ResponseWriter, r *http.Request) {
    // Parsear la solicitud del cliente
    var planet Planet
    err := json.NewDecoder(r.Body).Decode(&planet)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Asignar botín al planeta adecuado
    assignBooty(&planet)

    // Actualizar registro de botines asignados
    mutex.Lock()
    defer mutex.Unlock()
    planets[planet.Name] += planet.Booty

    // Responder al cliente con el planeta asignado
    response := fmt.Sprintf("Botín asignado a planeta %s para capitán %s", planet.Name, planet.Captain)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(response))
}

func assignBooty(planet *Planet) {
    // Lógica para asignar botín al planeta adecuado
    // (en este caso, simplemente se asigna al planeta con menos botines)
    minBooty := 9999
    var targetPlanet string
    for name, booty := range planets {
        if booty < minBooty {
            minBooty = booty
            targetPlanet = name
        }
    }
    planet.Name = targetPlanet
}
