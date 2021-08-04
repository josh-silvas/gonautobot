package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackReservationsCreateReader is a Reader for the DcimRackReservationsCreate structure.
type DcimRackReservationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackReservationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsCreateCreated creates a DcimRackReservationsCreateCreated with default headers values
func NewDcimRackReservationsCreateCreated() *DcimRackReservationsCreateCreated {
	return &DcimRackReservationsCreateCreated{}
}

/* DcimRackReservationsCreateCreated describes a response with status code 201, with default header values.

DcimRackReservationsCreateCreated dcim rack reservations create created
*/
type DcimRackReservationsCreateCreated struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-reservations/][%d] dcimRackReservationsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimRackReservationsCreateCreated) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
