package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//TODO use structured logging
	//TODO cmd line arguments
	//TODO take a go.mod file as input
	fileName := "/tmp/go_graph.txt"
	maxDepth := 3
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var root string
	depsMap := make(map[string][]string)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}

		tok := strings.Split(line, " ")
		if len(tok) != 2 {
			panic("Invalid line: " + line)
		}
		if root == "" {
			root = tok[0]
		}
		var deps []string
		if d, ok := depsMap[tok[0]]; ok {
			deps = d
		} else {
			deps = make([]string, 0)
		}
		t1 := strings.TrimSuffix(tok[1], "\n")
		deps = append(deps, t1)
		depsMap[tok[0]] = deps
	}
	// jsonData, err := json.Marshal(depsMap)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	//fmt.Println(string(jsonData))

	printDeps(root, depsMap, 0, maxDepth)
}

func printDeps(node string, depsMap map[string][]string, depth, maxDepth int) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), node)

	deps, ok := depsMap[node]
	if !ok {
		return
	}
	if depth == maxDepth-1 {
		return
	}
	for _, dep := range deps {
		printDeps(dep, depsMap, depth+1, maxDepth)
	}
}
