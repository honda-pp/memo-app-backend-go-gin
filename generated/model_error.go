/*
 * Memo App API
 *
 * API description in Markdown.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package generated

type Error struct {

	Message string `json:"message"`

	Code int32 `json:"code"`

	Data map[string]interface{} `json:"data,omitempty"`
}
