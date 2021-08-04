package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasWebhooksReadReader is a Reader for the ExtrasWebhooksRead structure.
type ExtrasWebhooksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksReadOK creates a ExtrasWebhooksReadOK with default headers values
func NewExtrasWebhooksReadOK() *ExtrasWebhooksReadOK {
	return &ExtrasWebhooksReadOK{}
}

/* ExtrasWebhooksReadOK describes a response with status code 200, with default header values.

ExtrasWebhooksReadOK extras webhooks read o k
*/
type ExtrasWebhooksReadOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/webhooks/{id}/][%d] extrasWebhooksReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksReadOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
