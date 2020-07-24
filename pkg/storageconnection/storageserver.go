package storageconnection

import (
	"context"

	"github.com/rs/zerolog/log"

	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/fred"
)

// Server implements the DatabaseServer interface
type Server struct {
	store fred.Store
}

// NewStorageServer creates a new Server to serve GRPC Requests. It answers according to the data/Storage Interface
func NewStorageServer(store *fred.Store) *Server {
	log.Debug().Msgf("Setting up new Server with store %#v", *store)
	return &Server{store: *store}
}

// Update calls specific method of the storage interface
func (s *Server) Update(ctx context.Context, item *Item) (*Response, error) {
	log.Debug().Msgf("GRPCServer: Update in=%#v", item)

	err := s.store.Update(item.Keygroup, item.Id, item.Val)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while updating item %#v", item)
		return &Response{Success: false}, err
	}
	return &Response{Success: true}, nil
}

// Delete calls specific method of the storage interface
func (s Server) Delete(ctx context.Context, key *Key) (*Response, error) {
	log.Debug().Msgf("GRPCServer: Delete in=%#v", key)
	err := s.store.Delete(key.Keygroup, key.Id)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while deleting item %#v", key)
		return &Response{Success: false, Message: "Server has encountered an error while deleting an item"}, err
	}
	return &Response{Success: true}, nil
}

// Read calls specific method of the storage interface
func (s Server) Read(ctx context.Context, key *Key) (*Val, error) {
	log.Debug().Msgf("GRPCServer: Read in=%#v", key)
	res, err := s.store.Read(key.Keygroup, key.Id)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while reading item %#v", key)
		return &Val{}, err
	}
	return &Val{Val: res}, nil
}

// ReadAll calls specific method of the storage interface
func (s Server) ReadAll(kg *Keygroup, server Database_ReadAllServer) error {
	// Steam: call server.send for every item, return if none left.
	log.Debug().Msgf("GRPCServer: ReadAll in=%#v", kg)
	res, err := s.store.ReadAll(kg.Keygroup)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while reading whole keygroup %#v", kg)
		return err
	}
	for id, elem := range res {
		err := server.Send(&Item{
			Id:  id,
			Val: elem,
		})
		if err != nil {
			return err
		}
	}
	// Return nil == successful transfer
	return nil
}

// IDs calls specific method of the storage interface
func (s Server) IDs(kg *Keygroup, server Database_IDsServer) error {
	log.Debug().Msgf("GRPCServer: IDs in=%#v", kg)
	res, err := s.store.IDs(kg.Keygroup)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while reading IDs %#v", kg)
		return err
	}
	for _, elem := range res {
		err := server.Send(&Key{Id: elem})
		if err != nil {
			return err
		}
	}
	// Return nil == successful transfer
	return nil
}

// Exists calls specific method of the storage interface
func (s Server) Exists(ctx context.Context, key *Key) (*Response, error) {
	log.Debug().Msgf("GRPCServer: Exists in=%#v", key)
	exists := s.store.Exists(key.Keygroup, key.Id)
	return &Response{Success: exists}, nil
}

// CreateKeygroup calls specific method of the storage interface
func (s Server) CreateKeygroup(ctx context.Context, kg *Keygroup) (*Response, error) {
	log.Debug().Msgf("GRPCServer: CreateKeygroup in=%#v", kg)
	err := s.store.CreateKeygroup(kg.Keygroup)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while creating keygroup %#v", kg)
		return &Response{Success: false}, err
	}
	return &Response{Success: true}, nil
}

// DeleteKeygroup calls specific method of the storage interface
func (s Server) DeleteKeygroup(ctx context.Context, kg *Keygroup) (*Response, error) {
	log.Debug().Msgf("GRPCServer: DeleteKeygroup in=%#v", kg)
	err := s.store.DeleteKeygroup(kg.Keygroup)
	if err != nil {
		log.Err(err).Msgf("GRPCServer has encountered an error while deleting keygroup %#v", kg)
		return &Response{Success: false}, err
	}
	return &Response{Success: true}, nil
}
