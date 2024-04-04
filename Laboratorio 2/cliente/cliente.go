package cliente

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strings"
    "time"
)

type Planet struct {
    Name   string `json:"name"`
    Booty  int    `json:"booty"`
    Captain string `json:"captain"`
}

func main() {
    planet := Planet{Name: "", Booty: 10, Captain: "C1"} // Ejemplo de solicitud de botín
    sendRequest(planet)
}

func sendRequest(planet Planet) {
    url := "http://localhost:8080/assign-booty"
    jsonValue, _ := json.Marshal(planet)
    response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Println("Error al enviar solicitud:", err)
        return
    }
    defer response.Body.Close()

    // Leer respuesta del servidor
    var result string
    if response.StatusCode == http.StatusOK {
        fmt.Println("Solicitud exitosa.")
        fmt.Println("Respuesta del servidor:")
        _, _ = fmt.Scanln(&result)
    } else {
        fmt.Println("Error:", response.Status)
    }
}
