package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func writeToJSONFile(f *os.File, data map[string]AlarmData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	} else if string(jsonData) == "{}" {
		return nil
	}

	_, err = f.Write(append(jsonData, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func ReadByLine(f *os.File) ([]map[string]AlarmData, error) {
	var ms []map[string]AlarmData
	r := bufio.NewReader(f)
	for {
		var m map[string]AlarmData
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else if line == "" {
			if ms == nil {
				return nil, fmt.Errorf("delete uneccessary backspaces") //todo создать тип ошибки и обработать с объяснением
			} else {
				break
			}
		}
		err = json.Unmarshal([]byte(line), &m)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	return ms, nil
}
func DeleteAlarm(f *os.File, name string) error {
	data, err := ReadByLine(f)
	if err != nil {
		return fmt.Errorf("error reading file %v", err)
	}
	for _, val := range data {
		if _, ok := val[name]; ok {
			delete(val, name)
		}
	}
	f.Truncate(0)

	for _, val := range data {
		err := writeToJSONFile(f, val)
		if err != nil {
			return err
		}
	}
	return nil
}
