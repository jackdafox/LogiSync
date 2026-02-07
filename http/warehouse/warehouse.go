package warehouse

import (
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"net/http"
)


var r = mux.NewRouter()

func GetWarehouse(db *sql.DB, id int) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        r.Body

    })
}

func CreateWarehouse(db *sql.DB, id int, name string, location_code string, is_active bool, created_at *time.Time) {
	
}

func UpdateWarehouse(db *sql.DB, id int) {

}

func DeleteWarehouse(db *sql.DB) {

}