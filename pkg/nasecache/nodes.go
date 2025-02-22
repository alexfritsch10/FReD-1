package nasecache

import (
	"git.tu-berlin.de/mcc-fred/fred/pkg/fred"
	"github.com/rs/zerolog/log"
)

// manage information about another node

// GetNodeAddress returns the ip and port of a node
func (n *NameServiceCache) GetNodeAddress(nodeID fred.NodeID) (addr string, err error) {
	key := "GetNodeAddress" + sep + string(nodeID)

	// try to get from cache
	respCache, err := n.cache.Get(key)
	if err != nil {
		log.Debug().Msgf("NaSe: Key %s is not in cache", key)

		resp, err := n.regularNase.GetNodeAddress(nodeID)

		if err != nil {
			return "", err
		}

		// put in cache
		err = n.cache.Set(key, []byte(resp))

		if err != nil {
			log.Err(err).Msg("Could not cache NaSe response")
		}

		return resp, nil
	}

	return string(respCache), nil
}

// GetAllNodes returns all nodes that are stored in the NaSe
func (n *NameServiceCache) GetAllNodes() (nodes []fred.Node, err error) {
	// TODO: implement caching scheme for GetAllNodes
	//key := "GetAllNodes"

	return n.regularNase.GetAllNodes()
}

// GetAllNodesExternal returns all nodes by their externally reachable address
func (n *NameServiceCache) GetAllNodesExternal() (nodes []fred.Node, err error) {
	return n.regularNase.GetAllNodesExternal()
}
