package activity

type Activity struct {
	id       uint
	event_id uint
	user_id  uint
	approved bool
	paid     bool
	check_in bool
}
