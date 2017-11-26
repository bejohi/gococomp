package model

type LbpPixel struct {
	X int
	Y int
}

func (lbpPixel LbpPixel)Equals(otherLbpPixel LbpPixel) bool {
	if lbpPixel.X == otherLbpPixel.X && lbpPixel.Y == otherLbpPixel.Y{
		return true
	}
	return false
}