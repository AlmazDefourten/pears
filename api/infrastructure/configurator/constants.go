package configurator

import "github.com/iris-contrib/swagger"

const ConfigName = "appconfig"
const ConfigType = "json"
const ConfigPath = "./api"

var SwaggerConfig = swagger.Config{
	// The url pointing to API definition.
	URL:          "http://localhost:8080/swagger/doc.json", // TODO: localhost:8080 refactor
	DeepLinking:  true,
	DocExpansion: "list",
	DomID:        "#swagger-ui",
	// The UI prefix URL (see route).
	Prefix: "/swagger",
}
