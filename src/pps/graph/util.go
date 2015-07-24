package graph

import (
	"fmt"
	"log"

	"github.com/pachyderm/pachyderm/src/pps"
)

func getNameToNodeInfo(nodes map[string]*pps.Node) (map[string]*NodeInfo, error) {
	nodeToInputs := getNodeNameToInputStrings(nodes)
	outputToNodes := getOutputStringToNodeNames(nodes)
	nodeInfos := make(map[string](*NodeInfo))
	for name := range nodes {
		nodeInfo := &NodeInfo{
			Parents: make([]string, 0),
		}
		parents := make(map[string]bool)
		for input := range nodeToInputs[name] {
			for parent := range outputToNodes[input] {
				if parent != name {
					parents[parent] = true
				}
			}
		}
		for parent := range parents {
			nodeInfo.Parents = append(nodeInfo.Parents, parent)
		}
		nodeInfos[name] = nodeInfo
	}
	log.Printf("got node infos %v\n", nodeInfos)
	return nodeInfos, nil
}

func getNodeNameToInputStrings(nodes map[string]*pps.Node) map[string]map[string]bool {
	m := make(map[string]map[string]bool)
	for name, node := range nodes {
		n := make(map[string]bool)
		if node.Input != nil {
			for hostDir := range node.Input.Host {
				// just need a differentiating string between types
				n[fmt.Sprintf("host://%s", hostDir)] = true
			}
			for pfsRepo := range node.Input.Pfs {
				n[fmt.Sprintf("pfs://%s", pfsRepo)] = true
			}
		}
		m[name] = n
	}
	return m
}

func getOutputStringToNodeNames(nodes map[string]*pps.Node) map[string]map[string]bool {
	m := make(map[string]map[string]bool)
	for name, node := range nodes {
		if node.Output != nil {
			for hostDir := range node.Output.Host {
				// just need a differentiating string between types
				s := fmt.Sprintf("host://%s", hostDir)
				if _, ok := m[s]; !ok {
					m[s] = make(map[string]bool)
				}
				m[s][name] = true
			}
			for pfsRepo := range node.Output.Pfs {
				s := fmt.Sprintf("pfs://%s", pfsRepo)
				if _, ok := m[s]; !ok {
					m[s] = make(map[string]bool)
				}
				m[s][name] = true
			}
		}
	}
	return m
}