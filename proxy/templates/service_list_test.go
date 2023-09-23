package templates

import (
	"testing"

	"github.com/tonitienda/go-proxy-mermaid/models"
)

func TestEmptyList(t *testing.T) {
	nodes := []models.Node{}

	diagram := GetMermaidDiagram(nodes)
	print(diagram)
	expected := "graph LR\n\tclassDef Level0 stroke-width:2px,stroke:#000,fill:#fff;\n"

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}

}
