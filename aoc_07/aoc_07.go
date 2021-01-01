package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type bag struct {
	childs  []string
	parents []string
}
type bags map[string]*bag

func createBagGraph(path string) (*simple.WeightedDirectedGraph, map[string]graph.Node) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines [][]string
	lineRegexp := regexp.MustCompile(`^(\S+ \S+) bags contain (.*)\.$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, lineRegexp.FindAllStringSubmatch(scanner.Text(), -1)[0])
	}

	g := simple.NewWeightedDirectedGraph(1, 1)
	nodeNames := make(map[string]graph.Node)
	for _, line := range lines {
		n := g.NewNode()
		g.AddNode(n)
		nodeNames[line[1]] = n
	}

	innerRegexp := regexp.MustCompile(`^(\d+) (\S+ \S+) bags?$`)
	for _, line := range lines {
		for _, innerBag := range strings.Split(line[2], ", ") {
			if innerBag == "no other bags" {
				break
			}
			innerBagParsed := innerRegexp.FindAllStringSubmatch(innerBag, -1)
			weight, _ := strconv.Atoi(innerBagParsed[0][1])
			g.SetWeightedEdge(
				g.NewWeightedEdge(
					nodeNames[line[1]],
					nodeNames[innerBagParsed[0][2]],
					float64(weight),
				))
		}
	}

	return g, nodeNames
}

func allParents(g graph.Directed, start graph.Node) []graph.Node {
	var result []graph.Node
	parents := g.To(start.ID())
	for parents.Next() {
		p := parents.Node()
		result = append(result, p)
		result = append(result, allParents(g, p)...)
	}
	return result
}

func totalBags(g graph.WeightedDirected, start graph.Node) int {
	result := 1
	childs := g.From(start.ID())
	for childs.Next() {
		c := childs.Node()
		e := g.WeightedEdge(start.ID(), c.ID())
		result += int(e.Weight()) * totalBags(g, c)
	}
	return result
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing file path!")
	}
	g, nodeMap := createBagGraph(os.Args[1])
	parentsMap := make(map[graph.Node]bool)
	start := nodeMap["shiny gold"]
	for _, p := range allParents(g, start) {
		parentsMap[p] = true
	}
	fmt.Println("Part 1:", len(parentsMap))
	fmt.Println("Part 2:", totalBags(g, start)-1)
}
