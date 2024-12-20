package sync

import (
	"context"

	"github.com/prysmaticlabs/prysm/v5/beacon-chain/core/feed"
	statefeed "github.com/prysmaticlabs/prysm/v5/beacon-chain/core/feed/state"
	light_client "github.com/prysmaticlabs/prysm/v5/consensus-types/light-client"
	"google.golang.org/protobuf/proto"
)

func (s *Service) lightClientFinalityUpdateSubscriber(_ context.Context, msg proto.Message) error {
	update, err := light_client.NewWrappedFinalityUpdate(msg)
	if err != nil {
		return err
	}

	log.Info("LC: storing new finality update in p2p subscriber")
	s.lcStore.LastLCFinalityUpdate = update

	s.cfg.stateNotifier.StateFeed().Send(&feed.Event{
		Type: statefeed.LightClientFinalityUpdate,
		Data: update,
	})

	return nil
}

func (s *Service) lightClientOptimisticUpdateSubscriber(_ context.Context, msg proto.Message) error {
	update, err := light_client.NewWrappedOptimisticUpdate(msg)
	if err != nil {
		return err
	}

	log.Info("LC: storing new optimistic update in p2p subscriber")
	s.lcStore.LastLCOptimisticUpdate = update

	s.cfg.stateNotifier.StateFeed().Send(&feed.Event{
		Type: statefeed.LightClientOptimisticUpdate,
		Data: update,
	})

	return nil
}
