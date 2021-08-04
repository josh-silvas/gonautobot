package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsUpdateReader is a Reader for the ExtrasCustomFieldsUpdate structure.
type ExtrasCustomFieldsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsUpdateOK creates a ExtrasCustomFieldsUpdateOK with default headers values
func NewExtrasCustomFieldsUpdateOK() *ExtrasCustomFieldsUpdateOK {
	return &ExtrasCustomFieldsUpdateOK{}
}

/* ExtrasCustomFieldsUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsUpdateOK extras custom fields update o k
*/
type ExtrasCustomFieldsUpdateOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extrasCustomFieldsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldsUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
