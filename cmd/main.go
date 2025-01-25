package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/RafaelSSantiago/poc-go-eks/internal/handlers"
)

func main() {
    // Variáveis de ambiente para conexão ao banco.
    // Exemplo de uso: "user:password@tcp(write-service.default.svc.cluster.local:3306)/dbname"
    // Para este PoC, ajustaremos no deploy do Kubernetes.
    writeDSN := os.Getenv("WRITE_DSN")
    readDSN := os.Getenv("READ_DSN") // No caso de 2 leituras, poderíamos usar Round Robin ou algo do tipo.

    // Inicia os handlers com as strings de conexão.
    // Para simplificar, vamos passar somente uma DSN de leitura (ou uma que represente um load balance).
    wh := handlers.NewWriteHandler(writeDSN)
    rh := handlers.NewReadHandler(readDSN)

    r := mux.NewRouter()
    r.HandleFunc("/write", wh.WriteData).Methods("POST")
    r.HandleFunc("/read", rh.ReadData).Methods("GET")

    log.Println("Servidor iniciado na porta 8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Erro ao iniciar servidor: %v", err)
    }
}
