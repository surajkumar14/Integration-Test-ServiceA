package connector

import (
	"fmt"

	"github.com/aerospike/aerospike-client-go"
)

var client *aerospike.Client

// InitAerospikeClient initializes the Aerospike client
func InitAerospikeClient(host string, port int) error {
	var err error
	client, err = aerospike.NewClient(host, port)
	if err != nil {
		return fmt.Errorf("failed to connect to Aerospike: %v", err)
	}
	return nil
}

// CreateRecord creates a new record in Aerospike
func CreateRecordAerospike(namespace, setName string, key int, name string) error {
	asKey, err := aerospike.NewKey(namespace, setName, key)
	if err != nil {
		return fmt.Errorf("failed to create key: %v", err)
	}
	bins := aerospike.BinMap{"name": name}
	err = client.Put(nil, asKey, bins)
	if err != nil {
		return fmt.Errorf("failed to create record: %v", err)
	}
	return nil
}

// ReadRecord reads a record from Aerospike
func ReadRecordAerospike(namespace, setName string, key int) (string, error) {
	asKey, err := aerospike.NewKey(namespace, setName, key)
	if err != nil {
		return "", fmt.Errorf("failed to create key: %v", err)
	}
	record, err := client.Get(nil, asKey)
	if err != nil {
		return "", fmt.Errorf("failed to read record: %v", err)
	}
	if record == nil {
		return "", fmt.Errorf("record not found")
	}
	return record.Bins["name"].(string), nil
}

// UpdateRecord updates an existing record in Aerospike
func UpdateRecordAerospike(namespace, setName string, key int, name string) error {
	return CreateRecordAerospike(namespace, setName, key, name)
}

// DeleteRecord deletes a record from Aerospike
func DeleteRecordAerospike(namespace, setName string, key int) error {
	asKey, err := aerospike.NewKey(namespace, setName, key)
	if err != nil {
		return fmt.Errorf("failed to create key: %v", err)
	}
	existed, err := client.Delete(nil, asKey)
	if err != nil {
		return fmt.Errorf("failed to delete record: %v", err)
	}
	if !existed {
		return fmt.Errorf("record not found")
	}
	return nil
}

func CloseAerospikeClient() {
	if client != nil {
		client.Close()
	}
}
