package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitsBulkPartialUpdateReader is a Reader for the CircuitsCircuitsBulkPartialUpdate structure.
type CircuitsCircuitsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsBulkPartialUpdateOK creates a CircuitsCircuitsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitsBulkPartialUpdateOK() *CircuitsCircuitsBulkPartialUpdateOK {
	return &CircuitsCircuitsBulkPartialUpdateOK{}
}

/* CircuitsCircuitsBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsBulkPartialUpdateOK circuits circuits bulk partial update o k
*/
type CircuitsCircuitsBulkPartialUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/][%d] circuitsCircuitsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsBulkPartialUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
