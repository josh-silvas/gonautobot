package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextsBulkUpdateReader is a Reader for the ExtrasConfigContextsBulkUpdate structure.
type ExtrasConfigContextsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsBulkUpdateOK creates a ExtrasConfigContextsBulkUpdateOK with default headers values
func NewExtrasConfigContextsBulkUpdateOK() *ExtrasConfigContextsBulkUpdateOK {
	return &ExtrasConfigContextsBulkUpdateOK{}
}

/* ExtrasConfigContextsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextsBulkUpdateOK extras config contexts bulk update o k
*/
type ExtrasConfigContextsBulkUpdateOK struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/config-contexts/][%d] extrasConfigContextsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextsBulkUpdateOK) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
