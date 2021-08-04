package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasJobResultsUpdateReader is a Reader for the ExtrasJobResultsUpdate structure.
type ExtrasJobResultsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJobResultsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsUpdateOK creates a ExtrasJobResultsUpdateOK with default headers values
func NewExtrasJobResultsUpdateOK() *ExtrasJobResultsUpdateOK {
	return &ExtrasJobResultsUpdateOK{}
}

/* ExtrasJobResultsUpdateOK describes a response with status code 200, with default header values.

ExtrasJobResultsUpdateOK extras job results update o k
*/
type ExtrasJobResultsUpdateOK struct {
	Payload *models.JobResult
}

func (o *ExtrasJobResultsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/job-results/{id}/][%d] extrasJobResultsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasJobResultsUpdateOK) GetPayload() *models.JobResult {
	return o.Payload
}

func (o *ExtrasJobResultsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
