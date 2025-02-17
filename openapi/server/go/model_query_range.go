/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type QueryRange struct {
	Query string `json:"query"`

	Labels []string `json:"labels"`

	Values [][]float64 `json:"values"`
}

// AssertQueryRangeRequired checks if the required fields are not zero-ed
func AssertQueryRangeRequired(obj QueryRange) error {
	elements := map[string]interface{}{
		"query":  obj.Query,
		"labels": obj.Labels,
		"values": obj.Values,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseQueryRangeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QueryRange (e.g. [][]QueryRange), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQueryRangeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQueryRange, ok := obj.(QueryRange)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQueryRangeRequired(aQueryRange)
	})
}
