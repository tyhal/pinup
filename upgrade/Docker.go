package upgrade

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type imageTag struct {
	Name string `json:"name"`
}

type imageTags struct {
	Results []imageTag `json:"results"`
	Next    string     `json:"next"`
}

var maxPages = 10

func getAllTags(img string) (semver.Versions, error) {

	tags := semver.Versions{}

	data := new(imageTags)

	data.Next = "1"
	i := int(1)

	// Send request to dockerhub to lookup latest images
	for data.Next != "" && i < maxPages {

		fmt.Print(".")

		resp, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/library/%s/tags/?page=%d", img, i))
		i = i + 1

		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, errors.New(resp.Status)
		}

		defer resp.Body.Close()
		data = new(imageTags)
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, err
		}
		for i := range data.Results {
			parsedVers, err := semver.Parse(data.Results[i].Name)
			if err != nil {
				continue
			}
			tags = append(tags, parsedVers)
		}
	}

	return tags, nil
}

func getLatestImage(name string) string {
	image := strings.Split(name, ":")[0]

	tags, err := getAllTags(image)

	if err != nil {
		log.Println(err.Error())
		return name
	}

	semver.Sort(tags)

	fmt.Println()

	return fmt.Sprintf("%s:%s", image, tags[len(tags)-1])
}

func fromStatement(curNode *parser.Node) {
	curImage := curNode.Next.Value

	// XXX Skip if its not using versioning already, makes it easy to avoid
	// tagged layers
	if !strings.Contains(curImage, ":") {
		return
	}

	newImage := getLatestImage(curImage)
	if newImage != curImage {
		fmt.Println("Line " + strconv.Itoa(curNode.StartLine) +
			" can use: " + newImage)
	} else {
		fmt.Println("Line " + strconv.Itoa(curNode.StartLine) + " is correct")
	}
}

func Docker(file *os.File) {
	res, err := parser.Parse(file)

	if err != nil {
		log.Fatal("Could not parse")
	}

	node := res.AST.Children

	for n := range node {
		curNode := node[n]
		switch curNode.Value {
		case "from":
			fromStatement(curNode)
		case "run":
			// Extract the shell part and then give it to our func from the shell upgrader
			ShellCall(curNode.Next.Value)
		}
	}
}
