package templates

import (
	"testing"

	"github.com/tonitienda/go-proxy-mermaid/models"
)

var Styles = "\tclassDef Level0 stroke-width:2px,stroke:#000,fill:#fff;"

func TestEmptyList(t *testing.T) {
	nodes := []models.Node{}

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n" + Styles + "\n"

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}

func TestListOneNode(t *testing.T) {
	nodes := []models.Node{}
	nodes = append(nodes, models.Node{ID: "123", Name: "Service"})

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n\t123[\"Service\n\t(123)\"]\n" + Styles + "\n"

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}
