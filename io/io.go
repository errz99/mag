package io

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func TestOne() {
	fmt.Println("Test One working")
}

// WriteJson takes a struct and saves it to a json file
func WriteJson(path string, jdata interface{}) error {
	marshal, err := json.Marshal(&jdata)
	if err != nil {
		return err
	}
	return SaveStrFile(path, string(marshal))
}

// WriteJsonIndent takes a struct and saves it indented to a json file
func WriteJsonIndented(path string, jdata interface{}) error {
	marshal, err := json.MarshalIndent(&jdata, "", "\t")
	if err != nil {
		return err
	}
	return SaveStrFile(path, string(marshal))
}

// ReadJson takes a json file and reads it into a struct
func ReadJson(file string, data interface{}) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return err
	}
	return nil
}

func SaveStrFile(file, data string) error {
	fich, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fich.Close()
	_, err = fich.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadBytes(filedat string) (content []byte, e string) {
	var err error
	content, err = ioutil.ReadFile(filedat)
	if err != nil {
		if os.IsNotExist(err) == true {
			e = "file doesn't exist"
		} else {
			e = "problem loading file"
		}
		return
	}
	return
}

// WriteGob takes a struct and saves it to a gob file
func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		enc_err := encoder.Encode(object)
		if enc_err != nil {
			err = enc_err
		}
	}
	file.Close()
	return err
}

// ReadGob takes a gob file and reads it into a struct
func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileCreate(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}

	defer func() {
		fo.Close()
	}()

	return nil
}

//func Json2File(file string, jdata interface{}) error {
//	var s bytes.Buffer
//
//	marshal, err := json.Marshal(&jdata)
//	if err != nil {
//		return err
//
//	} else {
//		s.WriteString(string(marshal))
//		fich, err := os.Create(file)
//		if err != nil {
//			return err
//
//		} else {
//			defer fich.Close()
//			_, err := fich.WriteString(s.String())
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
