package raftnase

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-errors/errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// getPrefix gets every key that starts(!) with the specified string
// the keys are sorted ascending by key for easier debugging
func (n *NameService) getPrefix(prefix string) (kv []*mvccpb.KeyValue, err error) {

}

// getExact gets the exact key
func (n *NameService) getExact(key string) (kv []*mvccpb.KeyValue, err error) {

}

func (n *NameService) getKeygroupStatus(kg string) (string, error) {

}

func (n *NameService) getKeygroupMutable(kg string) (string, error) {

}

func (n *NameService) getKeygroupExpiry(kg string, id string) (int, error) {

}

// put puts the value into etcd.
func (n *NameService) put(key, value string) (err error) {

}

// delete removes the value from etcd.
func (n *NameService) delete(key string) (err error) {

}

// addOwnKgNodeEntry adds the entry for this node with a status.
func (n *NameService) addOwnKgNodeEntry(kg string, status string) error {

}

// addOtherKgNodeEntry adds the entry for a remote node with a status.
func (n *NameService) addOtherKgNodeEntry(node string, kg string, status string) error {

}

// addKgStatusEntry adds the entry for a (new!) keygroup with a status.
func (n *NameService) addKgStatusEntry(kg string, status string) error {

}

// addKgMutableEntry adds the ismutable entry for a keygroup with a status.
func (n *NameService) addKgMutableEntry(kg string, mutable bool) error {

}

// addKgExpiryEntry adds the expiry entry for a keygroup with a status.
func (n *NameService) addKgExpiryEntry(kg string, id string, expiry int) error {

}

// fmtKgNode turns a keygroup name into the key that this node will save its state in
// Currently: kg|[keygroup]|node|[NodeID]
func (n *NameService) fmtKgNode(kg string) string {

}

func getNodeNameFromKgNodeString(kgNode string) string {

}
