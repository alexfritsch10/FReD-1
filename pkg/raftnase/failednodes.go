package raftnase

import (
	"fmt"
	"math"
	"strings"

	"git.tu-berlin.de/mcc-fred/fred/pkg/fred"
	"github.com/rs/zerolog/log"
)

// ReportFailedNode saves that a node has missed an update to a keygroup that it will get another time
func (n *NameService) ReportFailedNode(nodeID fred.NodeID, kg fred.KeygroupName, id string) error {

}

// RequestNodeStatus request a list of items that the node has missed while it was offline
func (n *NameService) RequestNodeStatus(nodeID fred.NodeID) (kgs []fred.Item) {

}

// GetNodeWithBiggerExpiry if this node has to get an item because it has missed it, it has to get it from a node with a bigger expiry
// if there is no node with a bigger expiry then it returns the node with the highest expiry
func (n *NameService) GetNodeWithBiggerExpiry(kg fred.KeygroupName) (nodeID fred.NodeID, addr string) {


}
