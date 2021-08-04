package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitsUpdateReader is a Reader for the CircuitsCircuitsUpdate structure.
type CircuitsCircuitsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsUpdateOK creates a CircuitsCircuitsUpdateOK with default headers values
func NewCircuitsCircuitsUpdateOK() *CircuitsCircuitsUpdateOK {
	return &CircuitsCircuitsUpdateOK{}
}

/* CircuitsCircuitsUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsUpdateOK circuits circuits update o k
*/
type CircuitsCircuitsUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuits/{id}/][%d] circuitsCircuitsUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
