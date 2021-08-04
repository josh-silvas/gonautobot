package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasJobResultsReadReader is a Reader for the ExtrasJobResultsRead structure.
type ExtrasJobResultsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobResultsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsReadOK creates a ExtrasJobResultsReadOK with default headers values
func NewExtrasJobResultsReadOK() *ExtrasJobResultsReadOK {
	return &ExtrasJobResultsReadOK{}
}

/* ExtrasJobResultsReadOK describes a response with status code 200, with default header values.

ExtrasJobResultsReadOK extras job results read o k
*/
type ExtrasJobResultsReadOK struct {
	Payload *models.JobResult
}

func (o *ExtrasJobResultsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/job-results/{id}/][%d] extrasJobResultsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasJobResultsReadOK) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *ExtrasJobResultsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
