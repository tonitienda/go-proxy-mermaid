package templates

import (
	"testing"

	"github.com/tonitienda/go-proxy-mermaid/models"
)

func TestEmptyList(t *testing.T) {
	nodes := []models.Node{}

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n" + Styles

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}

func TestListOneNode(t *testing.T) {
	nodes := []models.Node{}
	nodes = append(nodes, models.Node{ID: "123", Name: "Service"})

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n\t123[\"Service\n\t(123)\"]\n" + Styles

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}

func TestListTwoNodes(t *testing.T) {
	nodes := []models.Node{}
	nodes = append(nodes, models.Node{ID: "123", Name: "S1"})
	nodes = append(nodes, models.Node{ID: "456", Name: "S2"})

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n\t123[\"S1\n\t(123)\"]\n\t456[\"S2\n\t(456)\"]\n" + Styles

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}

func TestListTwoConnectedNodes(t *testing.T) {
	nodes := []models.Node{}
	nodes = append(nodes, models.Node{ID: "123", Name: "S1"})
	nodes = append(nodes, models.Node{ID: "456", Name: "S2", RequestorID: "123"})

	diagram := GetMermaidDiagram(nodes)
	expected := "graph LR\n\t123[\"S1\n\t(123)\"]\n\t456[\"S2\n\t(456)\"]\n\t123 --> 456\n" + Styles

	if diagram != expected {
		t.Errorf("Expected %s, got %s", expected, diagram)
	}
}
