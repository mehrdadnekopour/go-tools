package mypes

import (
	"errors"
	"fmt"

	// "github.com/satori/go.uuid"
	"github.com/pborman/uuid"
)

// GUID ...
type GUID struct {
	uuid.UUID
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice is handled by UnmarshalBinary, while
// a longer byte slice or a string is handled by UnmarshalText.
func (u *GUID) Scan(src interface{}) (err error) {
	err = u.UUID.Scan(src)

	if err != nil {
		fmt.Println("Error UUID Scan")
		return
	}

	// copy(u[:], uid[:])

	u.UUID[0], u.UUID[1], u.UUID[2], u.UUID[3] = u.UUID[3], u.UUID[2], u.UUID[1], u.UUID[0]
	u.UUID[4], u.UUID[5] = u.UUID[5], u.UUID[4]
	u.UUID[6], u.UUID[7] = u.UUID[7], u.UUID[6]

	return
}

// NewUUID ...
func NewUUID() (guid *GUID) {
	id := uuid.New()
	guid.ParseID(id)
	return
}

// ParseID ...
func (u *GUID) ParseID(id string) error {
	// return id
	if id == "" {
		return errors.New("id is empty")
	}
	uid := uuid.Parse(id)
	u, err := FromString(uid.String())

	return err

}

// NilID ,,,
func NilID() GUID {
	return GUID{}
}

// GetString ...
func GetString(id GUID) (string, bool) {

	valid := true
	s := id.String()
	// s := id.String()

	// if id == "" {
	// 	valid = false
	// }

	return s, valid
}

// CheckMUID ...
func CheckMUID(id string) (guid *GUID, valid bool) {
	valid = true

	// sid := id.String()

	uid := uuid.Parse(id)
	suuid := uid.String()
	if suuid == "" {
		valid = false
	}

	guid, err := FromString(suuid)
	if err != nil {
		valid = false
	}

	return guid, valid
}

// CheckUUID ...
func CheckUUID(id GUID) bool {
	// valid := true
	// if err != nil {
	// 	valid = false
	// }
	sid := id.String()
	_, valid := CheckMUID(sid)

	return valid
}

// FromString returns UUID parsed from string input.
// Input is expected in a form accepted by UnmarshalText.
func FromString(input string) (u *GUID, err error) {
	err = u.UnmarshalText([]byte(input))
	return
}

// GUID representation compliant with specification
// described in RFC 4122.
// type GUID uuid.UUID

// var uid uuid.UUID

// // Scan implements the sql.Scanner interface.
// // A 16-byte slice is handled by UnmarshalBinary, while
// // a longer byte slice or a string is handled by UnmarshalText.
// func (u *GUID) Scan(src interface{}) (err error) {
// 	err = uid.Scan(src)

// 	if err != nil {
// 		fmt.Println("Error UUID Scan")
// 		return
// 	}

// 	copy(u[:], uid[:])

// 	u[0], u[1], u[2], u[3] = u[3], u[2], u[1], u[0]
// 	u[4], u[5] = u[5], u[4]
// 	u[6], u[7] = u[7], u[6]

// 	return
// }

// // Returns canonical string representation of UUID:
// // xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
// func (u GUID) String() string {
// 	return uid.String()
// }
