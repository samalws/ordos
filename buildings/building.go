package buildings

import "../players"

type Building interface {
	GetType() string
	GetOwner() *players.Player
	GetRevenue() int
	GetRent() int
	GetClosed() bool
	SetClosed(bool)
	RoutinelyCalledFunc(x, y int)
	GetInternalInt(string) int
	SetInternalInt(string, int)
}

type basicBuilding struct {
	owner  *players.Player
	closed bool
}

func (b *basicBuilding) getOwner() *players.Player  { return b.owner }
func (*basicBuilding) GetRevenue() int              { return 0 }
func (b *basicBuilding) GetClosed() bool            { return b.closed }
func (b *basicBuilding) SetClosed(c bool)           { b.closed = c }
func (*basicBuilding) RoutinelyCalledFunc(int, int) {}
func (*basicBuilding) GetInternalInt(string) int    { return 0 }
func (*basicBuilding) SetInternalInt(string, int)   {}