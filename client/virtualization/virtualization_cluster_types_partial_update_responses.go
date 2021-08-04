package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterTypesPartialUpdateReader is a Reader for the VirtualizationClusterTypesPartialUpdate structure.
type VirtualizationClusterTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterTypesPartialUpdateOK creates a VirtualizationClusterTypesPartialUpdateOK with default headers values
func NewVirtualizationClusterTypesPartialUpdateOK() *VirtualizationClusterTypesPartialUpdateOK {
	return &VirtualizationClusterTypesPartialUpdateOK{}
}

/* VirtualizationClusterTypesPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesPartialUpdateOK virtualization cluster types partial update o k
*/
type VirtualizationClusterTypesPartialUpdateOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterTypesPartialUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
