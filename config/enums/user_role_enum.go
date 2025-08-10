package enums

import (
	"encoding/json"
	"fmt"
)

type UserRoleEnum int

const (
	Admin UserRoleEnum = iota
	User
)

func (u UserRoleEnum) String() string {
	return [...]string{"Admin", "User"}[u]
}

func (u *UserRoleEnum) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }

    switch s {
    case "Admin":
        *u = Admin
    case "User":
        *u = User
    default:
        return fmt.Errorf("invalid role: %s", s)
    }
    return nil
}