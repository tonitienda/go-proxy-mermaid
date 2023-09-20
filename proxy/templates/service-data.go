package templates

import (
	"os"
	"strings"

	"github.com/google/uuid"
)

type ServiceData struct {
	ID   string
	Name string
}

func init() {
	// Short ID is based on the first 8 chars of the UUID
	shortID := uuid.NewString()[0:8]
	os.Setenv("SERVICE_ID", shortID)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetServiceData() ServiceData {
	return ServiceData{
		ID:   getEnv("SERVICE_ID", "<no id>"),
		Name: getEnv("SERVICE_NAME", "Service"),
	}
}

func (h ServiceData) GetMermaidNode() string {
	return strings.ReplaceAll(strings.ReplaceAll(`#ID#["#NAME#
	(#ID#)"]`, "#ID#", h.ID), "#NAME#", h.Name)
}

func (h ServiceData) GetPage() string {
	return `<html>
	<head>
		<title>` + h.Name + `</title>
		<script language="javascript" src="https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.4.0/mermaid.min.js"></script>
	</head>
	<body>
		<pre class="mermaid">
			graph LR
		    ` + h.GetMermaidNode() + `


			classDef Level0 stroke-width:2px,stroke:#000,fill:#fff;

			class ` + h.ID + ` Level0
		</pre>
	</body>
	</html>`
}
