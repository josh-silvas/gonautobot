package users

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// UsersPermissionsReadReader is a Reader for the UsersPermissionsRead structure.
type UsersPermissionsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersPermissionsReadOK creates a UsersPermissionsReadOK with default headers values
func NewUsersPermissionsReadOK() *UsersPermissionsReadOK {
	return &UsersPermissionsReadOK{}
}

/* UsersPermissionsReadOK describes a response with status code 200, with default header values.

UsersPermissionsReadOK users permissions read o k
*/
type UsersPermissionsReadOK struct {
	Payload *models.ObjectPermission
}

func (o *UsersPermissionsReadOK) Error() string {
	return fmt.Sprintf("[GET /users/permissions/{id}/][%d] usersPermissionsReadOK  %+v", 200, o.Payload)
}
func (o *UsersPermissionsReadOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
