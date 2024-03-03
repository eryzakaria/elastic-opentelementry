package main

import (
    "log"
    "os"
    "time"
)

func main() {
    // Buat file untuk menulis log
    file, err := os.OpenFile("/root/docker-compose/elk-stack/apps-logstash/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Inisialisasi logger untuk menulis ke file
    logger := log.New(file, "", log.LstdFlags)

    // Infinite loop untuk menghasilkan log setiap detik
    for {
        logger.Println("Ini adalah contoh log -", time.Now().Format(time.RFC3339))
        time.Sleep(time.Second)
    }
}
