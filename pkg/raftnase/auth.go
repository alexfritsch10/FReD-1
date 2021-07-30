package raftnase

import (
	"fmt"
	"strings"

	"git.tu-berlin.de/mcc-fred/fred/pkg/fred"
)

// RevokeUserPermissions removes user's permission to perform method on kg by deleting the key in etcd.
func (n *NameService) RevokeUserPermissions(user string, method fred.Method, kg fred.KeygroupName) error {

}

// AddUserPermissions adds user's permission to perform method on kg by adding the key to etcd.
func (n *NameService) AddUserPermissions(user string, method fred.Method, kg fred.KeygroupName) error {

}

// GetUserPermissions returns a set of all of the user's permissions on kg from etcd.
func (n *NameService) GetUserPermissions(user string, kg fred.KeygroupName) (map[fred.Method]struct{}, error) {

}
