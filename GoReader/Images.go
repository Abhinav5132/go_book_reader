package main

import (
	"errors"
	"os"

)

func BlobFromPath(path string) ([]byte, error) {
	blob, err := os.ReadFile(path)
	if path == ""{
		
		return nil, errors.New("Inavlid Path")
	}
	if err != nil {
		print(err)
		return nil, err
	}
	return blob, nil 
}

func CreateAnImageFromContentsOfLibrary() {
	//TODO
}