package swagger

import (
	"net/http"

	"github.com/flowchartsman/swaggerui"

	_ "embed"
)

const (
	// Path is a path for swagger
	Path = "/docs"
	// Pattern is a pattern for mux that handles requests
	Pattern = "/docs/"
)

//go:embed ramme_template.swagger.json
var spec []byte

// Handler is required to route
var Handler = http.StripPrefix(Path, swaggerui.Handler(spec))
