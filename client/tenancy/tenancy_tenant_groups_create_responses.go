package tenancy

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// TenancyTenantGroupsCreateReader is a Reader for the TenancyTenantGroupsCreate structure.
type TenancyTenantGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsCreateCreated creates a TenancyTenantGroupsCreateCreated with default headers values
func NewTenancyTenantGroupsCreateCreated() *TenancyTenantGroupsCreateCreated {
	return &TenancyTenantGroupsCreateCreated{}
}

/* TenancyTenantGroupsCreateCreated describes a response with status code 201, with default header values.

TenancyTenantGroupsCreateCreated tenancy tenant groups create created
*/
type TenancyTenantGroupsCreateCreated struct {
	Payload *models.TenantGroup
}

func (o *TenancyTenantGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenant-groups/][%d] tenancyTenantGroupsCreateCreated  %+v", 201, o.Payload)
}
func (o *TenancyTenantGroupsCreateCreated) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
