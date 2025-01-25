package handlers

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type ReadHandler struct {
    db *sql.DB
}

func NewReadHandler(dsn string) *ReadHandler {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Erro ao conectar no DB de leitura: %v", err)
    }
    return &ReadHandler{db: db}
}

func (h *ReadHandler) ReadData(w http.ResponseWriter, r *http.Request) {
    // Exemplo de leitura simples
    query := "SELECT id, content FROM messages ORDER BY id DESC LIMIT 10"
    rows, err := h.db.Query(query)
    if err != nil {
        http.Error(w, "Erro ao ler dados do DB", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var results []map[string]interface{}
    for rows.Next() {
        var id int
        var content string
        err := rows.Scan(&id, &content)
        if err != nil {
            http.Error(w, "Erro ao iterar no resultado", http.StatusInternalServerError)
            return
        }
        results = append(results, map[string]interface{}{
            "id":      id,
            "content": content,
        })
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}