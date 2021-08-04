package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationClustersUpdateReader is a Reader for the VirtualizationClustersUpdate structure.
type VirtualizationClustersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersUpdateOK creates a VirtualizationClustersUpdateOK with default headers values
func NewVirtualizationClustersUpdateOK() *VirtualizationClustersUpdateOK {
	return &VirtualizationClustersUpdateOK{}
}

/* VirtualizationClustersUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersUpdateOK virtualization clusters update o k
*/
type VirtualizationClustersUpdateOK struct {
	Payload *models.Cluster
}

func (o *VirtualizationClustersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/{id}/][%d] virtualizationClustersUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationClustersUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
