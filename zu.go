package zu

var (
	tick = make(chan struct{})
	tock = make(chan struct{})
)

func Main(app func() error) error {
	var err error
	done := make(chan struct{})
	go func() {
		err = app()
		close(done)
	}()
	<-done
	return err
}

func Next() bool {
	tick <- struct{}{}
	<-tock
	return true
}
