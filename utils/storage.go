package utils

import (
	"os"
)

func WriteFileWithJSON(data map[string]AlarmData) error {
	wd, err := os.UserHomeDir()
	if err != nil {
		return HomeDirErr
	}

	err = os.Chdir(wd)
	if err != nil {
		return ChDirErr
	}

	dirName := "noty-data"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return CreatingDirErr
		}
	}

	err = os.Chdir(dirName)
	if err != nil {
		return ChDirErr
	}

	f, err := os.OpenFile("alarms-data.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return OpenFileErr
	}
	defer f.Close()

	err = writeToJSONFile(f, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDataFromFile(str string) error {
	wd, err := os.UserHomeDir()
	if err != nil {
		return HomeDirErr
	}

	err = os.Chdir(wd)
	if err != nil {
		return ChDirErr
	}

	f, err := os.OpenFile("noty-data/alarms-data", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return OpenFileErr
	}
	f.Close()

	err = DeleteDataFromFile(str)
	if err != nil {
		return err
	}
	return nil
}

func GetAllAlarms() ([]map[string]AlarmData, error) {
	f, err := os.OpenFile("noty-data/alarms-data", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, OpenFileErr
	}
	f.Close()
	dat, err := ReadByLine(f)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
