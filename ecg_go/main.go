package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"
)

type ECGData struct {
	Timestamp string  `json:"timestamp"`
	SensorID  int     `json:"sensorID"`
	ECG       float64 `json:"ECG"`
}

// SimularECG simula una señal de ECG y la imprime en la terminal
func SimularECG() {
	const frecuenciaMuestreo = 250.0 // Hz
	const duracion = 10.0            // segundos
	totalMuestras := int(frecuenciaMuestreo * duracion)

	for i := 0; i < totalMuestras; i++ {
		t := float64(i) / frecuenciaMuestreo
		valorECG := señalECG(t)
		fmt.Printf("Tiempo: %.3f s, ECG: %.3f mV\n", t, valorECG)
		escribirJSON(valorECG)
		time.Sleep(time.Second / time.Duration(frecuenciaMuestreo))
	}
}

func señalECG(t float64) float64 {
	// Modelo simplificado de una señal de ECG utilizando ondas sinusoidales
	frecuenciaCardiaca := 60.0 // latidos por minuto
	latidosPorSegundo := frecuenciaCardiaca / 60.0

	// Onda principal para simular el latido
	valor := math.Sin(2 * math.Pi * latidosPorSegundo * t)

	// Agregar un pico para simular el complejo QRS
	valor += 0.5 * math.Exp(-math.Pow((t*latidosPorSegundo*2)-math.Floor(t*latidosPorSegundo*2)-0.5, 2)/0.02)

	return valor
}

func escribirJSON(valorECG float64) {
	data := ECGData{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		SensorID:  721,
		ECG:       valorECG,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al codificar JSON:", err)
		return
	}

	file, err := os.Create("/tmp/ecg_data.json")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

func main() {
	SimularECG()
}
