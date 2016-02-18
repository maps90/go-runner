package action

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	FILE_PATH               = "resources/images/"
	FILE_FOLDER_PERMISSIONS = 0777
)

type Image struct {
	ParentSku     string `json:"parent_or_vendor_sku"`
	ImageSequence string `json:"image_sequence"`
	ImagePath     string `json:"image_path"`
}

func LoadImages(args string) {
	file, err := ioutil.ReadFile(args)
	if err != nil {
		log.Fatal(err)
	}
	_, fn := filepath.Split(args)
	if ext := path.Ext(args); ext != ".json" {
		log.Fatal("not a json file.")
		return
	}

	var p = FILE_PATH + strings.Trim(fn, ".json")
	if _, err := os.Stat(p); os.IsNotExist(err) {
		log.Println(p)
		os.MkdirAll(p, FILE_FOLDER_PERMISSIONS)
	}
	var imgs []Image
	if err := json.Unmarshal(file, &imgs); err != nil {
		log.Fatal("not a valid json file.")
		return
	}
	for _, img := range imgs {
		response, err := http.Get(img.ImagePath)
		if err != nil {
			log.Fatal("error while donwloading", img.ImagePath, "-", err)
		}
		if response.Header.Get("Content-Type") != "image/jpeg" {
			log.Fatal("path is not recognized.")
		}
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("error while reading response: ", "-", err)
		}

		file := path.Base(img.ImagePath)
		err = ioutil.WriteFile(p+"/"+file, contents, FILE_FOLDER_PERMISSIONS)
		if err != nil {
			log.Fatal("trouble creating file -- ", err)
		}

	}

}
