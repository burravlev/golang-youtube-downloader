package youtube

type Request struct {
	Format    Format
	Filepath  string
	Overwrite bool
	// on downloading callback
	Callback func(int)
}
