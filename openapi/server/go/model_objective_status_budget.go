/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ObjectiveStatusBudget struct {
	Total float64 `json:"total"`

	Remaining float64 `json:"remaining"`

	Max float64 `json:"max"`
}

// AssertObjectiveStatusBudgetRequired checks if the required fields are not zero-ed
func AssertObjectiveStatusBudgetRequired(obj ObjectiveStatusBudget) error {
	elements := map[string]interface{}{
		"total":     obj.Total,
		"remaining": obj.Remaining,
		"max":       obj.Max,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseObjectiveStatusBudgetRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ObjectiveStatusBudget (e.g. [][]ObjectiveStatusBudget), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseObjectiveStatusBudgetRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aObjectiveStatusBudget, ok := obj.(ObjectiveStatusBudget)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertObjectiveStatusBudgetRequired(aObjectiveStatusBudget)
	})
}
