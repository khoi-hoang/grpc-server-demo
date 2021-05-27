package stream_operation

type StreamOperation interface {
	Feed(int64)

	// Returns an EOF if there is nothing else to fetch
	Fetch() (int64, error)
}
