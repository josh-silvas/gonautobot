package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitsCreateReader is a Reader for the CircuitsCircuitsCreate structure.
type CircuitsCircuitsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsCreateCreated creates a CircuitsCircuitsCreateCreated with default headers values
func NewCircuitsCircuitsCreateCreated() *CircuitsCircuitsCreateCreated {
	return &CircuitsCircuitsCreateCreated{}
}

/* CircuitsCircuitsCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitsCreateCreated circuits circuits create created
*/
type CircuitsCircuitsCreateCreated struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuitsCircuitsCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsCircuitsCreateCreated) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
