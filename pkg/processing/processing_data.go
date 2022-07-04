package processing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"repoInfo/pkg/models"
)

const (
	url = "https://api.github.com/users/francolautaro2/repos"
)

// get data of repositories
func getJSON() {
	// get data from url https://api.github.com/users/francolautaro2/repos
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// parsing data of url
	parserData(body)

}

func parserData(data []byte) string {
	//converting the data of bytes array  toresult
	byteBuffer := bytes.NewBuffer(data).String()

	k := make([]models.Info, 0)
	json.Unmarshal(data, &k)
	fmt.Printf("%#v", k)

	// writing a file with data converted
	file, _ := json.MarshalIndent(k, "", " ")
	_ = ioutil.WriteFile("./data/data.json", file, 0644)

	return byteBuffer
}

func Run() {
	getJSON()
}
