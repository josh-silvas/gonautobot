package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasObjectChangesReadReader is a Reader for the ExtrasObjectChangesRead structure.
type ExtrasObjectChangesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasObjectChangesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasObjectChangesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasObjectChangesReadOK creates a ExtrasObjectChangesReadOK with default headers values
func NewExtrasObjectChangesReadOK() *ExtrasObjectChangesReadOK {
	return &ExtrasObjectChangesReadOK{}
}

/* ExtrasObjectChangesReadOK describes a response with status code 200, with default header values.

ExtrasObjectChangesReadOK extras object changes read o k
*/
type ExtrasObjectChangesReadOK struct {
	Payload *models.ObjectChange
}

func (o *ExtrasObjectChangesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasObjectChangesReadOK) GetPayload() *models.ObjectChange {
	return o.Payload
}

func (o *ExtrasObjectChangesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectChange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
