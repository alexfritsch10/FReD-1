package raftnase

import (
	"bytes"
	"context"
	"fmt"

	"git.tu-berlin.de/mcc-fred/fred/pkg/fred"
	"github.com/go-errors/errors"
	"github.com/rs/zerolog/log"
	"go.etcd.io/etcd/clientv3"
)

// ExitOtherNodeFromKeygroup deletes the node from the NaSe
func (n *NameService) ExitOtherNodeFromKeygroup(kg fred.KeygroupName, nodeID fred.NodeID) error {

}

// DeleteKeygroup marks the keygroup as "deleted" in the NaSe
func (n *NameService) DeleteKeygroup(kg fred.KeygroupName) error {

}

// ExistsKeygroup checks whether a Keygroup exists by checking whether there are keys with the prefix "kg|[kgname]|
func (n *NameService) ExistsKeygroup(kg fred.KeygroupName) (bool, error) {

}

// IsMutable checks whether a Keygroup is mutable.
func (n *NameService) IsMutable(kg fred.KeygroupName) (bool, error) {

}

// GetExpiry checks the expiration time for items of the keygroup on a replica.
func (n *NameService) GetExpiry(kg fred.KeygroupName) (int, error) {

}

// CreateKeygroup created the keygroup status and joins the keygroup
func (n *NameService) CreateKeygroup(kg fred.KeygroupName, mutable bool, expiry int) error {

}

// GetKeygroupMembers returns all IDs of the Members of a Keygroup by iterating over all saved keys that start with the keygroup name.
// The value of the map is the expiry in seconds.
func (n *NameService) GetKeygroupMembers(kg fred.KeygroupName, excludeSelf bool) (ids map[fred.NodeID]int, err error) {

}

// JoinNodeIntoKeygroup joins the node into an already existing keygroup
func (n *NameService) JoinNodeIntoKeygroup(kg fred.KeygroupName, nodeID fred.NodeID, expiry int) error {

}
