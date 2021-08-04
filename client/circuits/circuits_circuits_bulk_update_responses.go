package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitsBulkUpdateReader is a Reader for the CircuitsCircuitsBulkUpdate structure.
type CircuitsCircuitsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsBulkUpdateOK creates a CircuitsCircuitsBulkUpdateOK with default headers values
func NewCircuitsCircuitsBulkUpdateOK() *CircuitsCircuitsBulkUpdateOK {
	return &CircuitsCircuitsBulkUpdateOK{}
}

/* CircuitsCircuitsBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsBulkUpdateOK circuits circuits bulk update o k
*/
type CircuitsCircuitsBulkUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuits/][%d] circuitsCircuitsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsBulkUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
