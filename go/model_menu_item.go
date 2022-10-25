/*
 * CFD
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_ "github.com/hellofresh/health-go/v4/checks/postgres"
	"gorm.io/gorm"
)

type MenuItem struct {
	gorm.Model

	Id int32 `json:"id"`

	Description string `json:"description"`

	Name string `json:"name"`

	Price float32 `json:"price"`

	// URL to an image of the menu item. This should be the image from the /image endpoint
	ImageId int32 `json:"imageId"`
}

// AssertMenuItemRequired checks if the required fields are not zero-ed
func AssertMenuItemRequired(obj MenuItem) error {
	elements := map[string]interface{}{
		"id":          obj.Id,
		"description": obj.Description,
		"name":        obj.Name,
		"price":       obj.Price,
		"imageId":     obj.ImageId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseMenuItemRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MenuItem (e.g. [][]MenuItem), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMenuItemRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMenuItem, ok := obj.(MenuItem)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMenuItemRequired(aMenuItem)
	})
}
