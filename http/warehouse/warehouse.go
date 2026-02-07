package warehouse

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type Warehouse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Location_Code string `json:"location_code"`
	Is_Active     bool `json:"is_active"`
	Created_At    string `json:"created_at"`
}

var r = mux.NewRouter()

func GetWarehouse(db *sql.DB, r *http.Response) (string, error) {
	return "", nil
}

func CreateWarehouse(db *sql.DB, r *http.Response) (string, error) {
	return "", nil
}

func UpdateWarehouse(db *sql.DB, r *http.Response) (string, error) {
	return "", nil
}

func DeleteWarehouse(db *sql.DB, r *http.Response) (string, error) {
	return "", nil
}
