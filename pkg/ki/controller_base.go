package ki

type ControllerBase struct {
	window Window
}

func (controller *ControllerBase) Window() Window {
	return controller.window
}

func (controller *ControllerBase) SetWindow(value Window) {
	controller.window = value

	return
}

func (controller *ControllerBase) Init() *ControllerBase {
	return controller
}
