package connector

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitSqlDataBase() (*sql.DB, error) {

	// Set up the MySQL connection
	// var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to open database: %v", err)
	}
	return db, nil
}

func setupDatabase(t *testing.T, db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INT PRIMARY KEY,
        name VARCHAR(255)
    )`)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}
}

func teardownDatabase(t *testing.T, db *sql.DB, table string) {
	_, err := db.Exec("DROP TABLE IF EXISTS " + table)
	if err != nil {
		t.Fatalf("Failed to drop table: %v", err)
	}
}

func TestSQLCRUDOperations(t *testing.T) {
	db, err := InitSqlDataBase()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v. Skipping remaining test cases.", err)
	}
	defer db.Close()

	setupDatabase(t, db)
	defer teardownDatabase(t, db, "users")

	t.Run("CreateRecord", func(t *testing.T) {
		err := CreateRecord(db, 1, "John Doe")
		assert.NoError(t, err)
	})

	t.Run("ReadRecord", func(t *testing.T) {
		name, err := ReadRecord(db, 1)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", name)
	})

	t.Run("UpdateRecord", func(t *testing.T) {
		err := UpdateRecord(db, 1, "Jane Doe")
		assert.NoError(t, err)
	})

	t.Run("ReadUpdatedRecord", func(t *testing.T) {
		name, err := ReadRecord(db, 1)
		assert.NoError(t, err)
		assert.Equal(t, "Jane Doe", name)
	})

	t.Run("DeleteRecord", func(t *testing.T) {
		err := DeleteRecord(db, 1)
		assert.NoError(t, err)
	})

	t.Run("ReadDeletedRecord", func(t *testing.T) {
		name, err := ReadRecord(db, 1)
		assert.Error(t, err)
		assert.Empty(t, name)
	})
}

func loadSQLData(db *sql.DB, sqlFilePath string) error {

	sqlFileContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Failed to read .sql file: %v", err)
		return err
	}

	// Convert the file content to a string
	sqlCommands := string(sqlFileContent)

	// Split the commands by semicolon
	commands := strings.Split(sqlCommands, ";")

	// Execute each command separately
	for _, command := range commands {
		trimmedCommand := strings.TrimSpace(command)
		if trimmedCommand == "" {
			continue
		}
		_, err := db.Exec(trimmedCommand)
		if err != nil {
			log.Fatalf("Failed to execute SQL command: %v\nCommand: %s", err, trimmedCommand)
		}
	}
	return nil
}

func TestSQLQueriesOnMockData(t *testing.T) {
	// Initialize the database
	db, err := InitSqlDataBase()
	if err != nil {
		t.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close()

	teardownDatabase(t, db, "dummy_users")

	// Load SQL data into the actual database
	sqlFilePath := "../../mockdata.sql"
	err = loadSQLData(db, sqlFilePath)
	if err != nil {
		t.Fatalf("Failed to load SQL data: %v", err)
	}

	// Define test cases
	tests := []struct {
		name     string
		query    string
		args     []interface{}
		expected interface{}
	}{
		{
			name:     "SelectUserByID",
			query:    "SELECT name FROM dummy_users WHERE id = ?",
			args:     []interface{}{1},
			expected: "John Doe",
		},
		{
			name:     "SelectUserByEmail",
			query:    "SELECT name FROM dummy_users WHERE email = ?",
			args:     []interface{}{"jane.doe@example.com"},
			expected: "Jane Doe",
		},
		{
			name:     "SelectUserByAge",
			query:    "SELECT name FROM dummy_users WHERE age = ?",
			args:     []interface{}{35},
			expected: "Alice Smith",
		},
		{
			name:     "CountUsers",
			query:    "SELECT COUNT(*) FROM dummy_users",
			args:     []interface{}{},
			expected: 20, // Assuming there are 20 users in the mock data
		},
		{
			name:     "AverageAge",
			query:    "SELECT AVG(age) FROM dummy_users",
			args:     []interface{}{},
			expected: 32.6, // Assuming the average age is 32.5
		},
		{
			name:     "GroupByAge",
			query:    "SELECT age, COUNT(*) FROM dummy_users GROUP BY age",
			args:     []interface{}{},
			expected: map[int]int{30: 1, 28: 1, 35: 1, 40: 1, 25: 1, 45: 1, 32: 1, 29: 1, 27: 1, 33: 1, 31: 1, 38: 1, 26: 1, 34: 1, 37: 1, 24: 1, 39: 1, 36: 1, 41: 1, 22: 1}, // Assuming these age groups
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch expected := tt.expected.(type) {
			case string:
				var result string
				err := db.QueryRow(tt.query, tt.args...).Scan(&result)
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			case int:
				var result int
				err := db.QueryRow(tt.query, tt.args...).Scan(&result)
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			case float64:
				var result float64
				err := db.QueryRow(tt.query, tt.args...).Scan(&result)
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			case struct {
				Name    string
				OrderID int
			}:
				var result struct {
					Name    string
					OrderID int
				}
				err := db.QueryRow(tt.query, tt.args...).Scan(&result.Name, &result.OrderID)
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			case map[int]int:
				rows, err := db.Query(tt.query, tt.args...)
				assert.NoError(t, err)
				defer rows.Close()

				result := make(map[int]int)
				for rows.Next() {
					var age, count int
					err := rows.Scan(&age, &count)
					assert.NoError(t, err)
					result[age] = count
				}
				assert.NoError(t, rows.Err())
				assert.Equal(t, expected, result)
			}
		})
	}
}
