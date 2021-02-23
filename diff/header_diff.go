package diff

import "github.com/getkin/kin-openapi/openapi3"

// HeaderDiff is a diff between two OAS headers
type HeaderDiff struct {

	// ExtensionProps
	DescriptionDiff *ValueDiff   `json:"description,omitempty"` // diff of 'description' property
	DeprecatedDiff  *ValueDiff   `json:"deprecated,omitempty"`  // diff of 'deprecated' property
	RequiredDiff    *ValueDiff   `json:"required,omitempty"`    // diff of 'required' property
	ExampleDiff     *ValueDiff   `json:"example,omitempty"`     // diff of 'example' property
	SchemaDiff      *SchemaDiff  `json:"schema,omitempty"`      // diff of 'schema' property
	ContentDiff     *ContentDiff `json:"content,omitempty"`     // diff of 'content' property
}

func (headerDiff HeaderDiff) empty() bool {
	return headerDiff == HeaderDiff{}
}

func diffHeaderValues(header1, header2 *openapi3.Header) HeaderDiff {
	result := HeaderDiff{}

	result.DescriptionDiff = getValueDiff(header1.Description, header2.Description)
	result.DeprecatedDiff = getValueDiff(header1.Deprecated, header2.Deprecated)
	result.RequiredDiff = getValueDiff(header1.Required, header2.Required)

	if schemaDiff := getSchemaDiff(header1.Schema, header2.Schema); !schemaDiff.empty() {
		result.SchemaDiff = &schemaDiff
	}

	result.ExampleDiff = getValueDiff(header1.Example, header2.Example)

	if contentDiff := getContentDiff(header1.Content, header2.Content); !contentDiff.empty() {
		result.ContentDiff = &contentDiff
	}

	return result
}
