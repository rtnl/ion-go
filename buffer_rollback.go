package ion

func (b *Buffer) WithRollback(fn func(*Buffer) error) (err error) {
	snapshotCurrRead := b.inner.body.curr_r
	snapshotCurrWrite := b.inner.body.curr_w

	err = fn(b)
	if err != nil {
		b.inner.body.curr_r = snapshotCurrRead
		b.inner.body.curr_w = snapshotCurrWrite

		return
	}

	return
}
