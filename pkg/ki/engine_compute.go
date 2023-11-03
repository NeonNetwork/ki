package ki

func (engine *Engine) Compute() (err error) {
	err = engine.ComputeWindows()
	if err != nil {
		return
	}

	return
}

func (engine *Engine) ComputeWindows() (err error) {
	return
}
