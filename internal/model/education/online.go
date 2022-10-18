package education

import "fmt"

type Online struct {
	id    uint64
	title string
}

func NewEducationOnline(title string) Online {
	return Online{title: title}
}

func (o *Online) String() string {
	return fmt.Sprintf("Course: %s", o.title)
}

func (o *Online) SetID(id uint64) {
	o.id = id
}
