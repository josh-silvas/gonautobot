package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasWebhooksBulkUpdateReader is a Reader for the ExtrasWebhooksBulkUpdate structure.
type ExtrasWebhooksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksBulkUpdateOK creates a ExtrasWebhooksBulkUpdateOK with default headers values
func NewExtrasWebhooksBulkUpdateOK() *ExtrasWebhooksBulkUpdateOK {
	return &ExtrasWebhooksBulkUpdateOK{}
}

/* ExtrasWebhooksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkUpdateOK extras webhooks bulk update o k
*/
type ExtrasWebhooksBulkUpdateOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/][%d] extrasWebhooksBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksBulkUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
