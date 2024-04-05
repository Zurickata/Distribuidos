package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type Planet struct {
	Name    string
	Booty   int
	Captain string
}

func main() {
	// Definir el número máximo de iteraciones
	maxIterations := 2

	// Definir el rango de datos aleatorios para el botín
	minBooty := 1
	maxBooty := 20

	// Establecer una semilla aleatoria
	rand.Seed(time.Now().UnixNano())

	// Bucle para enviar solicitudes repetidamente
	for i := 0; i < maxIterations; i++ {

		planet := string(rune('A' + rand.Intn(6)))
		booty := rand.Intn(maxBooty-minBooty+1) + minBooty
		captain := "C" + strconv.Itoa(rand.Intn(3)+1)

		message := fmt.Sprintf("%s:%d:%s", planet, booty, captain)
		sendRequest(message)

		// Generar un intervalo aleatorio entre 2 y 7 segundos para enviar solicitudes
		interval := time.Duration(rand.Intn(6)+2) * time.Second

		// Esperar el intervalo de tiempo antes de enviar la próxima solicitud
		time.Sleep(interval)
	}
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
}
