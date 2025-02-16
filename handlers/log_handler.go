
package handlers

import (
	"secure-bank-api2/database"
	"secure-bank-api2/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All Log Aktivitas
func GetLogs(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, id_jurnal, id_karyawan, input_by, aksi, timestamp, keterangan FROM log_aktivitas")
	if err != nil {
		log.Println("Error fetching logs:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error fetching data"})
	}
	defer rows.Close()

	var logs []models.LogAktivitas
	for rows.Next() {
		var logEntry models.LogAktivitas
		if err := rows.Scan(&logEntry.ID, &logEntry.IDJurnal, &logEntry.IDKaryawan, &logEntry.InputBy, &logEntry.Aksi, &logEntry.Timestamp, &logEntry.Keterangan); err != nil {
			log.Println("Error scanning logs:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Error reading data"})
		}
		logs = append(logs, logEntry)
	}

	return c.JSON(http.StatusOK, logs)
}
