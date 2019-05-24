package main

import (
    "flag"
    "client"
    "log"
    "fmt"
    "time"
    "os"
)

func main()  {

    var server, output_file string
    var port, rep int

    flag.StringVar(&server, "host", "localhost", "Dirección del servidor")
    flag.StringVar(&output_file, "output", "result.csv", "Nombre del archivo donde escribir el resultado")
    flag.IntVar(&port, "port", 9000, "Puerto del servidor")
    flag.IntVar(&rep, "rep", 100, "Cantidad de repticiones")

    flag.Parse()

    file, err := os.Create(output_file)
    if err != nil {
        log.Fatal("Fatal error while opening file", err)
    }
    defer file.Close()

    log.Println("Iniciando prueba")
    var size int
    var time time.Duration

    file.WriteString("Size (B),Time (ns)\n")

    status := 20
    for i := 0; i < rep; i++ {
        size, time = client.GetFile(server, port)
        line := fmt.Sprintf("%d,%d\n", size, time)
        file.WriteString(line)

        if i % status == 0 {
            fmt.Println("...repeticion", i)
        }
    }
    file.WriteString("\n")
    log.Println("Fin, resultados escritos en", output_file)
}
