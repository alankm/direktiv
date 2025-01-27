// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteGlobalPrivateRegistryReader is a Reader for the DeleteGlobalPrivateRegistry structure.
type DeleteGlobalPrivateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGlobalPrivateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGlobalPrivateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGlobalPrivateRegistryOK creates a DeleteGlobalPrivateRegistryOK with default headers values
func NewDeleteGlobalPrivateRegistryOK() *DeleteGlobalPrivateRegistryOK {
	return &DeleteGlobalPrivateRegistryOK{}
}

/* DeleteGlobalPrivateRegistryOK describes a response with status code 200, with default header values.

successfully delete global private registry
*/
type DeleteGlobalPrivateRegistryOK struct {
}

func (o *DeleteGlobalPrivateRegistryOK) Error() string {
	return fmt.Sprintf("[DELETE /api/functions/registries/private][%d] deleteGlobalPrivateRegistryOK ", 200)
}

func (o *DeleteGlobalPrivateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteGlobalPrivateRegistryBody delete global private registry body
// Example: {"data":"admin:8QwFLg%D$qg*","reg":"https://prod.customreg.io"}
swagger:model DeleteGlobalPrivateRegistryBody
*/
type DeleteGlobalPrivateRegistryBody struct {

	// Target registry URL
	// Required: true
	Reg *string `json:"reg"`
}

// Validate validates this delete global private registry body
func (o *DeleteGlobalPrivateRegistryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteGlobalPrivateRegistryBody) validateReg(formats strfmt.Registry) error {

	if err := validate.Required("Registry Payload"+"."+"reg", "body", o.Reg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete global private registry body based on context it is used
func (o *DeleteGlobalPrivateRegistryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteGlobalPrivateRegistryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteGlobalPrivateRegistryBody) UnmarshalBinary(b []byte) error {
	var res DeleteGlobalPrivateRegistryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
