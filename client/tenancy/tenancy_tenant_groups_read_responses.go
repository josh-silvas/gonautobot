package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsReadReader is a Reader for the TenancyTenantGroupsRead structure.
type TenancyTenantGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsReadOK creates a TenancyTenantGroupsReadOK with default headers values
func NewTenancyTenantGroupsReadOK() *TenancyTenantGroupsReadOK {
	return &TenancyTenantGroupsReadOK{}
}

/* TenancyTenantGroupsReadOK describes a response with status code 200, with default header values.

TenancyTenantGroupsReadOK tenancy tenant groups read o k
*/
type TenancyTenantGroupsReadOK struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsReadOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantGroupsReadOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
