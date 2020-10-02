package main

import (
	"fmt"
	"os"
	"path/filepath"

	mio "github.com/errz99/mag/io"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

var names = []Name{
	Name{"Jorma", "Kaukonen"},
	Name{"Patti", "Smith"},
	Name{"Ian", "Anderson"},
}

var jsonfile = filepath.Join(os.Getenv("HOME"), "__jsonTest.json")
var gobfile = filepath.Join(os.Getenv("HOME"), "__gobTest.gob")

var checks = []string{
	"writejson", "readjson",
	"writegob", "readgob",
	"filexists", "filecreate",
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {

		case checks[0]:
			fmt.Println("\nchecking 'writejson'")
			fmt.Println("file:", jsonfile)
			checkWriteJson(false)

		case checks[1]:
			fmt.Println("\nchecking 'readjson'")
			fmt.Println("file:", jsonfile)
			checkReadJson()

		case checks[2]:
			fmt.Println("\nchecking 'writegob'")
			fmt.Println("file:", gobfile)
			checkWriteGob()

		case checks[3]:
			fmt.Println("\nchecking 'readgob'")
			fmt.Println("file:", gobfile)
			checkReadGob()

		case checks[4]:
			if len(os.Args) == 3 {
				checkFileExists(os.Args[2])
			}

		case checks[5]:
			if len(os.Args) == 3 {
				checkFileCreate(os.Args[2])
			}

		default:
			fmt.Println("available checks:\n", checks)
		}

	} else {
		fmt.Println("available checks:\n", checks)
	}
}

func checkWriteJson(indented bool) {
	var err error
	if indented {
		err = mio.WriteJsonIndented(jsonfile, names)
	} else {
		err = mio.WriteJson(jsonfile, names)
	}

	if err != nil {
		fmt.Println("FAILED:")
		fmt.Println(err, "\n")
	} else {
		fmt.Println("SUCCEED:")
		fmt.Println(names, "\n")
	}
}

func checkReadJson() {
	var namesBis []Name
	err := mio.ReadJson(jsonfile, &namesBis)

	if err != nil {
		fmt.Println("FAILED:")
		fmt.Println(err, "\n")
	} else {
		fmt.Println("SUCCEED:")
		fmt.Println(namesBis, "\n")
	}
}

func checkWriteGob() {
	err := mio.WriteGob(gobfile, names)

	if err != nil {
		fmt.Println("FAILED:")
		fmt.Println(err, "\n")
	} else {
		fmt.Println("SUCCEED:")
		fmt.Println(names, "\n")
	}
}

func checkReadGob() {
	var namesBis []Name
	err := mio.ReadGob(gobfile, &namesBis)

	if err != nil {
		fmt.Println("FAILED:")
		fmt.Println(err, "\n")
	} else {
		fmt.Println("SUCCEED:")
		fmt.Println(namesBis, "\n")
	}
}

func checkFileExists(path string) {
	exists := mio.FileExists(path)
	if exists {
		fmt.Println(path, "DOES EXIST")
	} else {
		fmt.Println(path, "DOES NOT EXIST")
	}
}

func checkFileCreate(path string) {
	if !mio.FileExists(path) {
		if err := mio.FileCreate(path); err == nil {
			fmt.Println(path, "CREATED")
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(path, "NOT CREATED (already exists)")
	}
}
