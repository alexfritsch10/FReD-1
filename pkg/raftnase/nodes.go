package raftnase

import (
	"fmt"
	"strings"

	"git.tu-berlin.de/mcc-fred/fred/pkg/fred"
	"github.com/go-errors/errors"
	"github.com/rs/zerolog/log"
)

// GetNodeAddress returns the ip and port of a node
func (n *NameService) GetNodeAddress(nodeID fred.NodeID) (addr string, err error) {

}

// GetAllNodes returns all nodes that are stored in the NaSe in the way they can be reached by other nodes
func (n *NameService) GetAllNodes() (nodes []fred.Node, err error) {

}

// GetAllNodesExternal returns all nodes with the port that the exthandler is running on
func (n *NameService) GetAllNodesExternal() (nodes []fred.Node, err error) {

}

func (n *NameService) getAllNodesBySuffix(suffix string) (nodes []fred.Node, err error) {

}
