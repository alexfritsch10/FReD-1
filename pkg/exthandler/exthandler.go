package exthandler

import (
	"github.com/rs/zerolog/log"

	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/data"
	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/errors"
	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/keygroup"
	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/replication"
)

type handler struct {
	dataService data.Service
	replService replication.Service
}

// New creates a new handler for external request (dataService.e. from clients).
func New(i data.Service, r replication.Service) Handler {
	return &handler{
		dataService: i,
		replService: r,
	}
}

// HandleCreateKeygroup handles requests to the CreateKeygroup endpoint of the client interface.
func (h *handler) HandleCreateKeygroup(k keygroup.Keygroup) error {
	if err := h.dataService.CreateKeygroup(k.Name); err != nil {
		log.Err(err).Msg("Exthandler cannot create keygroup with data service")
		return err
	}

	if err := h.replService.CreateKeygroup(k); err != nil {
		log.Err(err).Msg("Exthandler cannot create keygroup with replication service")
		return err
	}

	return nil
}

// HandleDeleteKeygroup handles requests to the DeleteKeygroup endpoint of the client interface.
func (h *handler) HandleDeleteKeygroup(k keygroup.Keygroup) error {
	if err := h.dataService.DeleteKeygroup(k.Name); err != nil {
		log.Err(err).Msg("Exthandler cannot delete keygroup with data service")
		return err
	}

	if err := h.replService.RelayDeleteKeygroup(keygroup.Keygroup{
		Name: k.Name,
	}); err != nil {
		log.Err(err).Msg("Exthandler cannot delete keygroup with replication service")
		return err
	}

	return nil
}

// HandleRead handles requests to the Read endpoint of the client interface.
func (h *handler) HandleRead(i data.Item) (data.Item, error) {
	result, err := h.dataService.Read(i.Keygroup, i.ID)
	if err != nil {
		return i, errors.New(errors.StatusNotFound, "exthandler: item does not exist")
	}
	// Store result in passed object to return a data.Item and not only the result string
	i.Data = result
	return i, nil
}

// HandleUpdate handles requests to the Update endpoint of the client interface.
func (h *handler) HandleUpdate(i data.Item) error {
	if err := h.dataService.Update(i); err != nil {
		log.Err(err).Msg("Exthandler cannot relay update with data service")
		return err
	}

	if err := h.replService.RelayUpdate(i); err != nil {
		log.Err(err).Msg("Exthandler cannot relay update with replication service")
		return err
	}

	return nil
}

// HandleDelete handles requests to the Delete endpoint of the client interface.
func (h *handler) HandleDelete(i data.Item) error {
	if err := h.dataService.Delete(i.Keygroup, i.ID); err != nil {
		log.Err(err).Msg("Exthandler cannot delete data item with data service")
		return err
	}

	if err := h.replService.RelayDelete(i); err != nil {
		log.Err(err).Msg("Exthandler cannot delete data item with data service")
		return err
	}

	return nil
}

// HandleAddReplica handles requests to the AddKeygroupReplica endpoint of the client interface.
func (h *handler) HandleAddReplica(k keygroup.Keygroup, n replication.Node) error {
	i, err := h.dataService.ReadAll(k.Name)

	if err != nil {
		return err
	}

	if err := h.replService.AddReplica(k, n, i, true); err != nil {
		log.Err(err).Msg("Exthandler cannot add a new keygroup replica")
		return err
	}

	return nil
}

// HandleGetKeygroupReplica handles requests to the GetKeygroupReplica endpoint of the client interface.
func (h *handler) HandleGetKeygroupReplica(k keygroup.Keygroup) ([]replication.Node, error) {
	return h.replService.GetReplica(k)
}

// HandleRemoveReplica handles requests to the RemoveKeygroupReplica endpoint of the client interface.
func (h *handler) HandleRemoveReplica(k keygroup.Keygroup, n replication.Node) error {
	if err := h.replService.RemoveReplica(k, n, true); err != nil {
		return err
	}

	return nil
}

// HandleAddNode handles requests to the AddReplica endpoint of the client interface.
func (h *handler) HandleAddNode(n []replication.Node) error {
	// TODO remove? IDK
	panic("Is it still necessary to add a node? The node can do this on startup by itself")
	// e := make([]string, len(n))
	// ec := 0
	//
	// for _, node := range n {
	// 	if err := h.replService.AddNode(node, true); err != nil {
	// 		log.Err(err).Msgf("Exthandler can not add a new replica node. (node=%#v)", node)
	// 		e[ec] = fmt.Sprintf("%v", err)
	// 		ec++
	// 	}
	// }
	//
	// if ec > 0 {
	// 	return errors.New(errors.StatusInternalError, fmt.Sprintf("exthandler: %v", e))
	// }

	return nil
}

// HandleGetReplica handles requests to the GetAllReplica endpoint of the client interface.
func (h *handler) HandleGetReplica(n replication.Node) (replication.Node, error) {
	return h.replService.GetNode(n)
}

// HandleGetAllReplica handles requests to the GetAllReplica endpoint of the client interface.
func (h *handler) HandleGetAllReplica() ([]replication.Node, error) {
	return h.replService.GetNodes()
}

// HandleRemoveNode handles requests to the RemoveReplica endpoint of the client interface.
func (h *handler) HandleRemoveNode(n replication.Node) error {
	// TODO why would this be called and what to do now.
	//return h.replService.RemoveNode(n, true)
	return nil
}
