package utils

import (
	"errors"
	"net/http"
	"os"
)

func FileExists(file_path string) (bool, error) {
	_, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func CheckFileExist(file_path string) (string, error, int) {
	exists, err := FileExists(TrimString(file_path))
	if err != nil {
		return file_path, err, http.StatusInternalServerError
	}
	if !exists {
		error_string := "File not founded : " + file_path
		return file_path, errors.New(error_string), http.StatusNotFound
	}
	file_path = file_path[1:]
	return file_path, nil, http.StatusOK
}
