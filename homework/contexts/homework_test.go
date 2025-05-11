package main

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Group struct {
	ctx      context.Context
	cancel   context.CancelFunc
	errCh    chan error
	wg       *sync.WaitGroup
	jobLimit chan struct{}
}

func NewErrGroup(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{
		ctx:      ctx,
		cancel:   cancel,
		errCh:    make(chan error, 1),
		wg:       &sync.WaitGroup{},
		jobLimit: make(chan struct{}),
	}, ctx
}

func (g *Group) SetLimit(limit int) {
	g.jobLimit = make(chan struct{}, limit)
}

func (g *Group) acquire() {
	jobLimit := int64(cap(g.jobLimit))
	if atomic.LoadInt64(&jobLimit) != 0 {
		g.jobLimit <- struct{}{}
	}
}

func (g *Group) release() {
	jobNum := int64(len(g.jobLimit))
	if atomic.LoadInt64(&jobNum) != 0 {
		<-g.jobLimit
	}
}

func (g *Group) Go(action func() error) {
	g.acquire()

	g.wg.Add(1)

	go func() {
		defer g.wg.Done()
		defer g.release()

		select {
		case <-g.ctx.Done():
			select {
			case g.errCh <- errors.New("ctx done"):
				g.cancel()
			default:
			}
		default:
			if err := action(); err != nil {
				select {
				case g.errCh <- err:
					g.cancel()
				default:
				}
			}
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	close(g.errCh)
	close(g.jobLimit)

	return <-g.errCh
}

func TestErrGroupWithoutError(t *testing.T) {
	var counter atomic.Int32
	group, _ := NewErrGroup(context.Background())

	for i := 0; i < 5; i++ {
		group.Go(func() error {
			time.Sleep(time.Second)
			counter.Add(1)
			return nil
		})
	}

	err := group.Wait()
	assert.Equal(t, int32(5), counter.Load())
	assert.NoError(t, err)
}

func TestErrGroupWithoutErrorWithLimit(t *testing.T) {
	var counter atomic.Int32
	group, _ := NewErrGroup(context.Background())

	group.SetLimit(1)
	timeStart := time.Now()

	for i := 0; i < 5; i++ {
		group.Go(func() error {
			time.Sleep(time.Second)
			counter.Add(1)
			return nil
		})
	}

	err := group.Wait()
	timeEnd := time.Now()
	assert.True(t, timeEnd.Sub(timeStart) > 5*time.Second)
	assert.Equal(t, int32(5), counter.Load())
	assert.NoError(t, err)
}

func TestErrGroupWithError(t *testing.T) {
	var counter atomic.Int32
	group, ctx := NewErrGroup(context.Background())

	for i := 0; i < 5; i++ {
		group.Go(func() error {
			timer := time.NewTimer(time.Second)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				counter.Add(1)
				return nil
			}
		})
	}

	group.Go(func() error {
		return errors.New("error")
	})

	err := group.Wait()
	assert.Equal(t, int32(0), counter.Load())
	assert.Error(t, err)
}

func TestErrGroupWithErrorWithLimit(t *testing.T) {
	var counter atomic.Int32
	group, ctx := NewErrGroup(context.Background())
	group.SetLimit(1)

	for i := 0; i < 5; i++ {
		group.Go(func() error {
			timer := time.NewTimer(10 * time.Millisecond)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				counter.Add(1)
				return nil
			}
		})
	}

	group.Go(func() error {
		return errors.New("error")
	})

	err := group.Wait()
	assert.Equal(t, int32(5), counter.Load())
	assert.Error(t, err)
}
