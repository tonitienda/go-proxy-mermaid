package templates

import (
	"fmt"
	"strings"

	"github.com/crazy3lf/colorconv"
	"github.com/tonitienda/go-proxy-mermaid/models"
)

var NodeTemplate = "#ID#(\"#ID#\n\t#NAME#\")"
var StyleTemplate string = "classDef Level#IDX# stroke-width:2px,stroke:#DARK#,fill:#LIGHT#,color:#DARK#"
var MainColorHue int = 202
var MinSaturation int = 77
var MaxSaturation int = 87
var MinLightness int = 13
var MaxLightness int = 75
var NumberOfLevels int = 6

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
		"#URL#":  node.Url,
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
			connectionList += "\t" + node.RequestorID + " -- " + node.Url + " --> " + node.ID + "\n"
		}
	}

	return connectionList
}

var StyleLevels int = 6
var ErrorStyle string = "\tclassDef Error stroke-width:2px,stroke:#d00,fill:#f50;\n"

func GetStyles() string {
	saturationStep := (MaxSaturation - MinSaturation) / NumberOfLevels
	lighnessStep := (MaxLightness - MinLightness) / NumberOfLevels

	medSaturation := (MaxSaturation-MinSaturation)/2 + MinSaturation
	medLightness := (MaxLightness-MinLightness)/2 + MinLightness

	styles := ""

	for idx := 0; idx < NumberOfLevels; idx++ {
		styles += "\t" + ReplaceMultiple(StyleTemplate, map[string]string{
			"#IDX#":   fmt.Sprintf("%d", idx),
			"#DARK#":  HSLtoHEX(MainColorHue, MinSaturation+idx*saturationStep, MinLightness+idx*lighnessStep),
			"#LIGHT#": HSLtoHEX(MainColorHue, medSaturation+idx*saturationStep, medLightness+idx*lighnessStep),
		}) + "\n"
	}

	return styles
}

func HSLtoHEX(hue int, saturation int, lighness int) string {
	fmt.Println("HSLtoHEX", hue, saturation, lighness)
	r, g, b, err := colorconv.HSLToRGB(float64(hue), float64(saturation)/100, float64(lighness)/100)

	// Since values are hardcoded, this should be catched in tests
	if err != nil {
		fmt.Println("Error converting HSL to RGB", err)
		panic(err)
	}

	return ReplaceMultiple(colorconv.RGBToHex(r, g, b), map[string]string{
		"0x": "#"})
}

func GetNodesStyles(nodes []models.Node) string {

	nodesStyles := ""

	for idx, node := range nodes {
		if node.HasError {
			nodesStyles += "\tclass " + node.ID + " Error\n"
		} else {
			nodesStyles += "\tclass " + node.ID + " Level" + fmt.Sprintf("%d", (idx%StyleLevels)) + "\n"
		}

	}

	return nodesStyles
}

func GetMermaidDiagram(nodes []models.Node) string {
	nodeList := GetMermaidNodeList(nodes)
	connectionsList := GetMermaidConnections(nodes)
	styles := GetStyles() + ErrorStyle
	applyStyles := GetNodesStyles(nodes)

	fmt.Println("Styles", styles)

	return "graph LR\n" + nodeList + connectionsList + styles + applyStyles
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
		<div>` +
		// ReplaceMultiple(GetMermaidDiagram(nodes), map[string]string{
		// 	"\n": "<br />",
		// 	"\t": "&nbsp;&nbsp;&nbsp;&nbsp;"}) +
		`</div>
	</body>
	</html>`
}
