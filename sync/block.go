package sync

import (
	"fmt"
	"time"

	"github.com/DefinitelyNotAGoat/go-explore/mongodb"
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
)

// Sync a structure containing the sync(er) config
type Sync struct {
	refresh time.Duration
	mg      *mongodb.Mongo
	gt      *goTezos.GoTezos
}

// NewSync returns a new sync struct for mongodb
func NewSync(refresh time.Duration, gt *goTezos.GoTezos, mg *mongodb.Mongo) *Sync {
	return &Sync{refresh: refresh, gt: gt, mg: mg}
}

// Sync syncs blocks into the db.
func (s *Sync) Sync(quit chan int, errch chan error) {
	quitBlocks := make(chan int)
	s.syncBlocks(quitBlocks, errch)

	go func() {
		for x := range quit {
			quitBlocks <- x
			return
		}
	}()

}

func (s *Sync) syncBlocks(quit chan int, errch chan error) {
	start := 0
	go func() {
		for {
			err := s.handleBlock(errch, start)
			if err == nil {
				start++
			}
		}
	}()
}

func (s *Sync) handleBlock(errch chan error, level int) error {
	block, err := s.gt.Block.Get(level)
	if err != nil {
		errch <- fmt.Errorf("sync:init:block: could not fetch block %d: %v", level, err)
		return err
	}
	err = s.mg.InsertBlock(block)
	if err != nil {
		errch <- fmt.Errorf("sync:init:block: could not insert block %d: %v", level, err)
		return err
	}
	return nil
}
