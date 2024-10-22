package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
	"github.com/sidiqPratomo/DJKI-Pengaduan/database"
	"github.com/sirupsen/logrus"
)

func CreateRouter(log *logrus.Logger, config *config.Config) *gin.Engine {
	// Connect to the database
	db := database.ConnectDB(config, log)

	// Create the router
	router := gin.Default()

	// Define the test route
	router.GET("/test-db", func(c *gin.Context) {
		if err := TestQuery(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Database connected successfully!",
		})
	})

	return router
}

// TestQuery is a simple function to test DB connection
func TestQuery(db *sql.DB) error {
	var (
		name string
	)

	q := `SELECT DATABASE()` // MySQL query to get the current database name

	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			return err
		}
		fmt.Printf("name: %s\n", name)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}