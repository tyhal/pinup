package upgrade

import (
	"fmt"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"io"
	"log"
	"strconv"
	"strings"
)

func getLatestImage(name string) string {
	// Send request to dockerhub to lookup latest images
	return "latest"
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
	}
}

func Docker(in io.Reader, out io.Writer) {
	res, err := parser.Parse(in)

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
