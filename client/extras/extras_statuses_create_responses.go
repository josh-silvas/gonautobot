package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasStatusesCreateReader is a Reader for the ExtrasStatusesCreate structure.
type ExtrasStatusesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasStatusesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesCreateCreated creates a ExtrasStatusesCreateCreated with default headers values
func NewExtrasStatusesCreateCreated() *ExtrasStatusesCreateCreated {
	return &ExtrasStatusesCreateCreated{}
}

/* ExtrasStatusesCreateCreated describes a response with status code 201, with default header values.

ExtrasStatusesCreateCreated extras statuses create created
*/
type ExtrasStatusesCreateCreated struct {
	Payload *models.Status
}

func (o *ExtrasStatusesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/statuses/][%d] extrasStatusesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasStatusesCreateCreated) GetPayload() *models.Status {
	return o.Payload
}

func (o *ExtrasStatusesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
