package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasWebhooksUpdateReader is a Reader for the ExtrasWebhooksUpdate structure.
type ExtrasWebhooksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksUpdateOK creates a ExtrasWebhooksUpdateOK with default headers values
func NewExtrasWebhooksUpdateOK() *ExtrasWebhooksUpdateOK {
	return &ExtrasWebhooksUpdateOK{}
}

/* ExtrasWebhooksUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksUpdateOK extras webhooks update o k
*/
type ExtrasWebhooksUpdateOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/{id}/][%d] extrasWebhooksUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
