package utils

import (
	"bytes"
	"encoding/gob"
	"github.com/Dev-ManavSethi/url-shortener/models"
	"os"
)

func LoadMapBackup()(map[string]string, error){

	file, err := os.OpenFile("map.backup", os.O_RDONLY|os.O_CREATE, 0755)
	if err!=nil{
		return  nil, err
	}

	Map := make(map[string]string)

	decoder := gob.NewDecoder(file)
	err2 :=  decoder.Decode(&Map)
	if err2!=nil{
		return nil, err
	}

	return Map, nil



}

func SetMapValue(key, value string){

	models.GlobalMutex.Lock()
	models.Map[key] = value
	models.GlobalMutex.Unlock()



}

func TakeMapBackup() error{
	file, err := os.OpenFile("map.backup", os.O_CREATE|os.O_RDWR, 0755)
	if err!=nil{
		return err
	}
	defer file.Close()
	var Buffer bytes.Buffer

	encoder := gob.NewEncoder(&Buffer)
	err2 := encoder.Encode(models.Map)
	if err2!=nil{
		return err2
	}

	_, err3 := file.Write(Buffer.Bytes())
	if err3!=nil{
		return err3
	}

	return nil
}
