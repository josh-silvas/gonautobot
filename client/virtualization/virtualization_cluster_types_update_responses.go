package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClusterTypesUpdateReader is a Reader for the VirtualizationClusterTypesUpdate structure.
type VirtualizationClusterTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterTypesUpdateOK creates a VirtualizationClusterTypesUpdateOK with default headers values
func NewVirtualizationClusterTypesUpdateOK() *VirtualizationClusterTypesUpdateOK {
	return &VirtualizationClusterTypesUpdateOK{}
}

/* VirtualizationClusterTypesUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterTypesUpdateOK virtualization cluster types update o k
*/
type VirtualizationClusterTypesUpdateOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClusterTypesUpdateOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
