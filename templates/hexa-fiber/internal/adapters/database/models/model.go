package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Value-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()
	valueString, err := json.Marshal(j)
	return string(valueString), err
}
func (j *JSONB) Scan(value interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Scan-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()

	if data, ok := value.([]byte); ok {
		if err := json.Unmarshal(data, j); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected type for JSONB: %T", value)
	}

	return nil
}
