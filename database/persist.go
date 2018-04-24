package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

const (
	fileSuffix = ".pd"
)

var (
	persistDirecty = "./"
	fileList       map[string]*os.File
)

func init() {
}

// ExportRegion saves all the data from the region in JSON format to a file
func (db *PCDB) ExportRegion(regionName string) error {
	file := fileList[regionName]
	if file == nil {
		// attempt to open file
		fileName := persistDirecty + regionName + fileSuffix
		f, openErr := os.OpenFile(fileName, os.O_WRONLY, 0600)
		file = f
		// create file if it does not exist
		if openErr != nil {
			f, createErr := os.Create(fileName)
			file = f
			if createErr != nil {
				return fmt.Errorf("Failed to create new file: %v | %v", createErr, openErr)
			}
		}
		_ = f
	}
	// get all data then write to file
	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Failed to save region data for region: %v | %v", regionName, err)
	}
	data := r.Data
	jd, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("unable to convert: %v to JSON | %v", jd, err)
	}
	_, err = file.Write(jd)
	if err != nil {
		return fmt.Errorf("Failed to write entry: %v | %v", string(jd), err)
	}
	return nil
}

// ImportRegion imports data from a file into a region
func (db *PCDB) ImportRegion(regionName, filename string) error {
	// attempt to open file
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Failed to read data from file: %v | %v", filename, err)
	}
	var data interface{}
	if err = json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("Failed to unmarshal data: %v | %v", jsonData, err)
	}

	r, err := db.GetRegion(regionName)
	if err != nil {
		return fmt.Errorf("Region: %v does not exist | %v", regionName, err)
	}
	dataMap := reflect.ValueOf(data)
	if dataMap.Kind() == reflect.Map {
		for _, key := range dataMap.MapKeys() {
			d := dataMap.MapIndex(key)
			r.Add(key.String(), d)
		}
	}
	return nil
}

func shutdownPersistence() {
	for _, f := range fileList {
		f.Close()
	}
}
