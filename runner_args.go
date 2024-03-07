package spireav

// Runnable defines the interface for an object that creates FFmpeg command line arguments which
// can be executed by the spireav package.
type Runnable interface {
	RunnableArgs() ([]string, error)
}

// RunnableArgs is a simple type that implements the Runnable interface.
type RunnableArgs []string

func (args RunnableArgs) RunnableArgs() ([]string, error) {
	return args, nil
}
