package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsProvidersBulkUpdateReader is a Reader for the CircuitsProvidersBulkUpdate structure.
type CircuitsProvidersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersBulkUpdateOK creates a CircuitsProvidersBulkUpdateOK with default headers values
func NewCircuitsProvidersBulkUpdateOK() *CircuitsProvidersBulkUpdateOK {
	return &CircuitsProvidersBulkUpdateOK{}
}

/* CircuitsProvidersBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersBulkUpdateOK circuits providers bulk update o k
*/
type CircuitsProvidersBulkUpdateOK struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/][%d] circuitsProvidersBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsProvidersBulkUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
