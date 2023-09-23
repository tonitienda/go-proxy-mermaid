package templates

import (
	"strings"

	"github.com/tonitienda/go-proxy-mermaid/models"
)

func GetMermaidNode(node models.Node) string {
	return strings.ReplaceAll(strings.ReplaceAll(`#ID#["#NAME#
	(#ID#)"]`, "#ID#", node.ID), "#NAME#", node.Name)
}

func GetMermaidNodeList(nodes []models.Node) string {
	nodeList := ""

	for _, node := range nodes {
		nodeList += "\t" + GetMermaidNode(node) + "\n"
	}

	return nodeList
}

func GetMermaidDiagram(nodes []models.Node) string {
	nodeList := GetMermaidNodeList(nodes)

	return "graph LR\n" + nodeList + "\tclassDef Level0 stroke-width:2px,stroke:#000,fill:#fff;\n"
}

func GetPage(nodes []models.Node) string {
	return `<html>
	<head>
		<title></title>
		<script language="javascript" src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.4.0/mermaid.min.js"></script>
	</head>
	<body>
		<pre class="mermaid">` +
		GetMermaidDiagram(nodes) +
		`</pre>
	</body>
	</html>`
}
