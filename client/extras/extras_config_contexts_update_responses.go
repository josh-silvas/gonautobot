package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextsUpdateReader is a Reader for the ExtrasConfigContextsUpdate structure.
type ExtrasConfigContextsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsUpdateOK creates a ExtrasConfigContextsUpdateOK with default headers values
func NewExtrasConfigContextsUpdateOK() *ExtrasConfigContextsUpdateOK {
	return &ExtrasConfigContextsUpdateOK{}
}

/* ExtrasConfigContextsUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsUpdateOK extras config contexts update o k
*/
type ExtrasConfigContextsUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/config-contexts/{id}/][%d] extrasConfigContextsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
