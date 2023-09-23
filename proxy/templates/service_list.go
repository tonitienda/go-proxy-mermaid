package templates

import (
	"strings"

	"github.com/tonitienda/go-proxy-mermaid/models"
)

var NodeTemplate = "#ID#[\"#NAME#\n\t(#ID#)\"]"

func ReplaceMultiple(text string, replacements map[string]string) string {
	for key, value := range replacements {
		text = strings.ReplaceAll(text, key, value)
	}

	return text
}

func GetMermaidNode(node models.Node) string {
	return ReplaceMultiple(NodeTemplate, map[string]string{
		"#ID#":   node.ID,
		"#NAME#": node.Name,
	})
}

func GetMermaidNodeList(nodes []models.Node) string {
	nodeList := ""

	for _, node := range nodes {
		nodeList += "\t" + GetMermaidNode(node) + "\n"
	}

	return nodeList
}

func GetMermaidConnections(nodes []models.Node) string {
	connectionList := ""

	for _, node := range nodes {
		if node.RequestorID != "" {
			connectionList += "\t" + node.RequestorID + " --> " + node.ID + "\n"
		}
	}

	return connectionList
}

var Styles string = "\tclassDef Level0 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level2 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level3 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level4 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level5 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level6 stroke-width:2px,stroke:#000,fill:#fff;\n" +
	"\tclassDef Level7 stroke-width:2px,stroke:#000,fill:#fff;\n"

func GetNodesStyles(nodes []models.Node) string {

	nodesStyles := ""

	for _, node := range nodes {
		nodesStyles += "\tclass " + node.ID + " Level0\n"

	}

	return nodesStyles
}

func GetMermaidDiagram(nodes []models.Node) string {
	nodeList := GetMermaidNodeList(nodes)
	connectionsList := GetMermaidConnections(nodes)
	applyStyles := GetNodesStyles(nodes)

	return "graph LR\n" + nodeList + connectionsList + Styles + applyStyles
}

func GetPage(nodes []models.Node) string {
	return `<html>
	<head>
		<title></title>
		<script language="javascript" src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.4.0/mermaid.min.js"></script>
	</head>
	<body>
		<div class="mermaid">` +
		GetMermaidDiagram(nodes) +
		`</div>
	</body>
	</html>`
}
