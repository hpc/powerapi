/*
 * Power API v1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lowell@lanl.gov
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

// AggregationResetBody - Aggregation reset request
type AggregationResetBody struct {

	ResetType ResetType `json:"ResetType,omitempty"`

	// A list of system URIs to apply the reset to
	TargetURIs []string `json:"TargetURIs,omitempty"`
}
