package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextsReadReader is a Reader for the ExtrasConfigContextsRead structure.
type ExtrasConfigContextsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsReadOK creates a ExtrasConfigContextsReadOK with default headers values
func NewExtrasConfigContextsReadOK() *ExtrasConfigContextsReadOK {
	return &ExtrasConfigContextsReadOK{}
}

/* ExtrasConfigContextsReadOK describes a response with status code 200, with default header values.

ExtrasConfigContextsReadOK extras config contexts read o k
*/
type ExtrasConfigContextsReadOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/config-contexts/{id}/][%d] extrasConfigContextsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsReadOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
