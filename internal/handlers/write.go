package handlers

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type WriteHandler struct {
    db *sql.DB
}

func NewWriteHandler(dsn string) *WriteHandler {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Erro ao conectar no DB de escrita: %v", err)
    }
    return &WriteHandler{db: db}
}

func (h *WriteHandler) WriteData(w http.ResponseWriter, r *http.Request) {
    // Exemplo de escrita simples
    type Payload struct {
        Message string `json:"message"`
    }
    var p Payload

    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, "Erro ao decodificar payload", http.StatusBadRequest)
        return
    }

    // Insert básico
    query := "INSERT INTO messages (content) VALUES (?)"
    _, err := h.db.Exec(query, p.Message)
    if err != nil {
        http.Error(w, "Erro ao inserir dados no DB", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Mensagem inserida com sucesso: %s", p.Message)
}
