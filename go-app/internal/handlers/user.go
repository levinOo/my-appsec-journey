package handlers

import (
	"fmt"
	"net/http"

	"vulnerable-app/internal/db"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// SQL Injection - намеренное склеивание строк через fmt.Sprintf (CWE-89)
	query := fmt.Sprintf("SELECT id, username, role FROM users WHERE username = '%s'", name)
	
	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var id int
		var username, role string
		if err := rows.Scan(&id, &username, &role); err != nil {
			continue
		}
		result += fmt.Sprintf("ID: %d, Username: %s, Role: %s\n", id, username, role)
	}

	w.Write([]byte(result))
}
