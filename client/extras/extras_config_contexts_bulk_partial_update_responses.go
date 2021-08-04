package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextsBulkPartialUpdateReader is a Reader for the ExtrasConfigContextsBulkPartialUpdate structure.
type ExtrasConfigContextsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsBulkPartialUpdateOK creates a ExtrasConfigContextsBulkPartialUpdateOK with default headers values
func NewExtrasConfigContextsBulkPartialUpdateOK() *ExtrasConfigContextsBulkPartialUpdateOK {
	return &ExtrasConfigContextsBulkPartialUpdateOK{}
}

/* ExtrasConfigContextsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkPartialUpdateOK extras config contexts bulk partial update o k
*/
type ExtrasConfigContextsBulkPartialUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-contexts/][%d] extrasConfigContextsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsBulkPartialUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
