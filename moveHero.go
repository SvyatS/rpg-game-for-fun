package main

func (h *Hero) runleft() {
	h.currentEvent = eventEnum.Run
	h.rightSide = false
	h.position[0] = h.position[0] - (h.moveSpeed / 10)
}
func (h *Hero) walkLeft() {
	h.currentEvent = eventEnum.Walk
	h.rightSide = false
	h.position[0] = h.position[0] - (h.moveSpeed / 15)
}
func (h *Hero) step() {
	h.currentEvent = eventEnum.Idle
	return
}
func (h *Hero) walkRight() {
	h.currentEvent = eventEnum.Walk
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 15)
}
func (h *Hero) runRight() {
	h.currentEvent = eventEnum.Run
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 10)
}
func (h *Hero) jumpRunLeft() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = false
	h.position[0] = h.position[0] - (h.moveSpeed / 10)
	var addedY = -(h.currentAnimationFrame * h.currentAnimationFrame) + (h.currentAnimationFrame * (h.events[h.currentEvent].frameCounts - 1))

	h.position[1] = h.prevPosition[1] - addedY
}
func (h *Hero) jumpWalkLeft() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = false
	h.position[0] = h.position[0] - (h.moveSpeed / 15)
	var addedY = -(h.currentAnimationFrame * h.currentAnimationFrame) + (h.currentAnimationFrame * (h.events[h.currentEvent].frameCounts - 1))

	h.position[1] = h.prevPosition[1] - addedY
}
func (h *Hero) jumpUp() {
	h.currentEvent = eventEnum.Jump
	var addedY = -(h.currentAnimationFrame * h.currentAnimationFrame) + (h.currentAnimationFrame * (h.events[h.currentEvent].frameCounts - 1))

	h.position[1] = h.prevPosition[1] - addedY

	// h.position[1] = h.currentAnimationFrame
	// if h.currentAnimationFrame < h.events[h.currentEvent].frameCounts/2 {
	// 	h.position[1] = h.position[1] - (h.moveSpeed / 45)
	// } else {
	// 	h.position[1] = h.position[1] + (h.moveSpeed / 45)
	// }
}
func (h *Hero) jumpWalkRight() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 15)
	var addedY = -(h.currentAnimationFrame * h.currentAnimationFrame) + (h.currentAnimationFrame * (h.events[h.currentEvent].frameCounts - 1))
	h.position[1] = h.prevPosition[1] - addedY
}
func (h *Hero) jumpRunRight() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 10)
	var addedY = -(h.currentAnimationFrame * h.currentAnimationFrame) + (h.currentAnimationFrame * (h.events[h.currentEvent].frameCounts - 1))
	h.position[1] = h.prevPosition[1] - addedY
}
