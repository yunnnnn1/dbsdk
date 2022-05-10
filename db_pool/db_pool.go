package db_pool

type GoPool struct {
	MaxLimit int

	tokenChan chan struct{}
}

type GoPoolOption func(*GoPool)

func WithMaxLimit(max int) GoPoolOption  {
	return func(gp *GoPool) {
		gp.MaxLimit = max
		gp.tokenChan = make(chan struct{},gp.MaxLimit)

		for i := 0; i < gp.MaxLimit ; i++{
			gp.tokenChan <- struct{}{}
		}
	}
}

func NewGoPool(options ...GoPoolOption) *GoPool  {
	p := &GoPool{}
	for _,o := range options {
		o(p)
	}
	return  p
}

func (gp *GoPool) Submit(fn func())  {
	token := <- gp.tokenChan
	go func() {
		fn()
		gp.tokenChan <- token
	}()
}

func (gp * GoPool) Wait()  {
	for i:= 0; i< gp.MaxLimit ; i++ {
		<- gp.tokenChan
	}

	close(gp.tokenChan)
}

func (gp *GoPool) size() int   {
	return len(gp.tokenChan)
}
