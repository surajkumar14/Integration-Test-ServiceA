package connector

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAerospikeCRUDOperations(t *testing.T) {
	aerospikeHost := os.Getenv("AEROSPIKE_HOST")
	aerospikePort := os.Getenv("AEROSPIKE_PORT")
	aerospikePortInt, err := strconv.Atoi(aerospikePort)
	if err != nil {
		t.Fatalf("Invalid port number: %v", err)
	}
	err = InitAerospikeClient(aerospikeHost, aerospikePortInt)
	if err != nil {
		t.Fatalf("Failed to initialize Aerospike client: %v. Skipping remaining test cases.", err)
	}

	// Register cleanup function to close the Aerospike client
	t.Cleanup(func() {
		CloseAerospikeClient()
	})

	namespace := "test"
	setName := "users"

	t.Run("CreateRecord", func(t *testing.T) {
		err := CreateRecordAerospike(namespace, setName, 1, "John Doe")
		assert.NoError(t, err)
	})

	t.Run("ReadRecord", func(t *testing.T) {
		name, err := ReadRecordAerospike(namespace, setName, 1)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", name)
	})

	t.Run("UpdateRecord", func(t *testing.T) {
		err := UpdateRecordAerospike(namespace, setName, 1, "Jane Doe")
		assert.NoError(t, err)
	})

	t.Run("ReadUpdatedRecord", func(t *testing.T) {
		name, err := ReadRecordAerospike(namespace, setName, 1)
		assert.NoError(t, err)
		assert.Equal(t, "Jane Doe", name)
	})

	t.Run("DeleteRecord", func(t *testing.T) {
		err := DeleteRecordAerospike(namespace, setName, 1)
		assert.NoError(t, err)
	})

	t.Run("ReadDeletedRecord", func(t *testing.T) {
		name, err := ReadRecordAerospike(namespace, setName, 1)
		assert.Error(t, err)
		assert.Empty(t, name)
	})
}
