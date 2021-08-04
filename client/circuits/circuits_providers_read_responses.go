package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsProvidersReadReader is a Reader for the CircuitsProvidersRead structure.
type CircuitsProvidersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersReadOK creates a CircuitsProvidersReadOK with default headers values
func NewCircuitsProvidersReadOK() *CircuitsProvidersReadOK {
	return &CircuitsProvidersReadOK{}
}

/* CircuitsProvidersReadOK describes a response with status code 200, with default header values.

CircuitsProvidersReadOK circuits providers read o k
*/
type CircuitsProvidersReadOK struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadOK  %+v", 200, o.Payload)
}
func (o *CircuitsProvidersReadOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
