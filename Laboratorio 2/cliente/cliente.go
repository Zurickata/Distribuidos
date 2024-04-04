package main

import (    
	"fmt"
	"net"
    "strings"
)

type Planet struct {
	Name    string 
	Booty   int    
	Captain string 
}

func main() {
    planet := "Pluton"
    booty := 10
    captain := "C1"

    message := fmt.Sprintf("%s:%d:%s", planet, booty, captain)
    sendRequest(message)
}

func sendRequest(message string) {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error al conectar con el servidor:", err)
        return
    }
    defer conn.Close()

    // Enviar mensaje al servidor
    _, err = conn.Write([]byte(message))

	parts := strings.Split(message, ":")
	captain := parts[2]
	planetName := parts[0]
	fmt.Printf("Capitán %s encontró botín en Planeta %s, enviando solicitud de asignación...\n", captain, planetName)

    if err != nil {
        fmt.Println("Error al enviar el mensaje:", err)
        return
    }

    // Leer respuesta del servidor
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error al leer la respuesta del servidor:", err)
        return
    }
    response := string(buffer[:n])

    // Imprimir respuesta del servidor
    fmt.Println("Respuesta del servidor:", response)
}