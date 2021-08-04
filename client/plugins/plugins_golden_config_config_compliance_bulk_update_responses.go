package plugins

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsGoldenConfigConfigComplianceBulkUpdateReader is a Reader for the PluginsGoldenConfigConfigComplianceBulkUpdate structure.
type PluginsGoldenConfigConfigComplianceBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsGoldenConfigConfigComplianceBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsGoldenConfigConfigComplianceBulkUpdateOK creates a PluginsGoldenConfigConfigComplianceBulkUpdateOK with default headers values
func NewPluginsGoldenConfigConfigComplianceBulkUpdateOK() *PluginsGoldenConfigConfigComplianceBulkUpdateOK {
	return &PluginsGoldenConfigConfigComplianceBulkUpdateOK{}
}

/* PluginsGoldenConfigConfigComplianceBulkUpdateOK describes a response with status code 200, with default header values.

PluginsGoldenConfigConfigComplianceBulkUpdateOK plugins golden config config compliance bulk update o k
*/
type PluginsGoldenConfigConfigComplianceBulkUpdateOK struct {
	Payload *models.ConfigCompliance
}

func (o *PluginsGoldenConfigConfigComplianceBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/golden-config/config-compliance/][%d] pluginsGoldenConfigConfigComplianceBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *PluginsGoldenConfigConfigComplianceBulkUpdateOK) GetPayload() *models.ConfigCompliance {
	return o.Payload
}

func (o *PluginsGoldenConfigConfigComplianceBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigCompliance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
