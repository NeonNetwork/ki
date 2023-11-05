package ki

type Controller interface {
	Window() Window
	SetWindow(Window)

	Compute() error
	Render() error
}
