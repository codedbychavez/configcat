package controllers

type ControllerManager struct {
	PaintController PaintControllerManager
}

func NewControllerManager() ControllerManager {
	paintController := PaintController{}

	return ControllerManager{
		PaintController: paintController,
	}
}