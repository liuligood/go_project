package parser

const (
	JSON = "json"
	YAML = "yaml"
)

type Parser struct {
	FileParser map[string]fileType
}

type fileType interface {
	Load(content []byte, config *map[string]interface{}) error
}

func NewFileParser() *Parser {
	fileTypeMap := make(map[string]fileType)
	fileTypeMap[YAML] = &Yaml{}
	fileTypeMap[JSON] = &Json{}
	return &Parser{
		FileParser: fileTypeMap,
	}
}
