package store

import (
	"context"
	"time"
)

func (s *Store) warmUpStore(amount int, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	metas, err := s.db.GetRandomQuestionMeta(ctx, int64(amount))
	if err == nil && len(metas) == 0 {
		return nil
	} else if err != nil {
		return err
	}

	if err = s.cache.AddCachedIDs(ctx, metas, timeout); err != nil {
		return err
	}

	return nil
}
