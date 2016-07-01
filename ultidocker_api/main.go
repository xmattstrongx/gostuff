package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

type imageInfo struct {
	Options string `json:"options"`
	ImageID string `json:"imageid"`
}

func getImages(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("docker", "images")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	img := string(out)
	fmt.Println(img)
	io.WriteString(w, img)
}

func runImage(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	request := &imageInfo{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		log.Printf("[CreateChallenge] could not parse request: %s", err)
		return
	}

	byteRequest, _ := json.Marshal(request)
	fmt.Printf("original StringRequest: %s", string(byteRequest))

	stringRequest := stringFromJSON(string(byteRequest))
	fmt.Printf("\nnext StringRequest: %s", stringRequest)

	// dockerStringRequest := fmt.Sprint("docker run" + stringRequest)
	// fmt.Printf("\ndocker stringRequest: %s", dockerStringRequest)

	cmdStringRequest := strings.Replace(stringRequest, " ", ",", -1)
	fmt.Printf("\ncmd stringRequest: %s\n", cmdStringRequest)

	//TODO: Fix the docker exec. It is either escaping my double quotes or not inserting
	cmd := exec.Command(fmt.Sprintf("docker "+"run"), cmdStringRequest)
	err := cmd.Run()
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	io.WriteString(w, cmdStringRequest)
}

func stringFromJSON(s string) string {
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, "options:", " ", -1)
	s = strings.Replace(s, "imageid:", " ", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)

	return s
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/images"), getImages)
	mux.HandleFuncC(pat.Post("/images"), runImage)

	http.ListenAndServe(":8000", mux)
}
