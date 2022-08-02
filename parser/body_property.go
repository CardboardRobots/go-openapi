package parser

import (
	"github.com/cardboardrobots/go-openapi/entity"
	"github.com/getkin/kin-openapi/openapi3"
)

func (p *SchemaParser) GetBody(operation *openapi3.Operation) entity.BodyProperty {
	if operation.RequestBody == nil {
		return entity.BodyProperty{}
	}

	// Try using this value directly
	if operation.RequestBody.Value != nil {
		requestBody := operation.RequestBody.Value

		json := requestBody.Content.Get(CONTENT_JSON)
		if json != nil {
			schema := p.GetBySchema(json.Schema.Value)
			if schema != nil {
				return entity.BodyProperty{
					Schema:   schema,
					Encoding: entity.ENCODING_JSON,
				}
			}
		}

		xml := requestBody.Content.Get(CONTENT_XML)
		if xml != nil {
			schema := p.GetBySchema(xml.Schema.Value)
			if schema != nil {
				return entity.BodyProperty{
					Schema:   schema,
					Encoding: entity.ENCODING_XML,
				}
			}
		}

		url := requestBody.Content.Get(CONTENT_URL)
		if url != nil {
			schema := p.GetBySchema(url.Schema.Value)
			if schema != nil {
				return entity.BodyProperty{
					Schema:   schema,
					Encoding: entity.ENCODING_URL,
				}
			}
		}

		text := requestBody.Content.Get(CONTENT_TEXT)
		if text != nil {
			schema := p.GetBySchema(text.Schema.Value)
			if schema != nil {
				return entity.BodyProperty{
					Schema:   schema,
					Encoding: entity.ENCODING_TEXT,
				}
			}
		}
	}

	return entity.BodyProperty{}
}

const CONTENT_JSON = "application/json"
const CONTENT_XML = "application/xml"
const CONTENT_URL = "application/x-www-form-urlencoded"
const CONTENT_TEXT = "text/plain"
