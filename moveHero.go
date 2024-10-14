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
	h.position[0] = h.position[0] - (h.moveSpeed / 10 / h.events[h.currentEvent].frameCounts)
	if h.currentAnimationFrame <= h.events[h.currentEvent].frameCounts/2 {
		h.position[1] = h.position[1] - (h.moveSpeed / 10 / h.events[h.currentEvent].frameCounts)
	} else {
		h.position[1] = h.position[1] + (h.moveSpeed / 10 / h.events[h.currentEvent].frameCounts)
	}
}
func (h *Hero) jumpWalkLeft() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = false
	h.position[0] = h.position[0] - (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	if h.currentAnimationFrame <= h.events[h.currentEvent].frameCounts/2 {
		h.position[1] = h.position[1] - (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	} else {
		h.position[1] = h.position[1] + (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	}
}
func (h *Hero) jumpUp() {
	h.currentEvent = eventEnum.Jump
	if h.currentAnimationFrame <= h.events[h.currentEvent].frameCounts/2 {
		h.position[1] = h.position[1] - (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	} else {
		h.position[1] = h.position[1] + (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	}
}
func (h *Hero) jumpWalkRight() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 15)
	if h.currentAnimationFrame <= h.events[h.currentEvent].frameCounts/2 {
		h.position[1] = h.position[1] - (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	} else {
		h.position[1] = h.position[1] + (h.moveSpeed / 15 / h.events[h.currentEvent].frameCounts)
	}
}
func (h *Hero) jumpRunRight() {
	h.currentEvent = eventEnum.Jump
	h.rightSide = true
	h.position[0] = h.position[0] + (h.moveSpeed / 10)
	if h.currentAnimationFrame <= h.events[h.currentEvent].frameCounts/2 {
		h.position[1] = h.position[1] - (h.moveSpeed / 10 / h.events[h.currentEvent].frameCounts)
	} else {
		h.position[1] = h.position[1] + (h.moveSpeed / 10 / h.events[h.currentEvent].frameCounts)
	}
}
