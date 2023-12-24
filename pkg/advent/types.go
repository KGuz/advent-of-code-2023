package advent

import (
	"cmp"
	"math"
)

type number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

type point struct{ i, j int }

func (p point) add(q point) point { return point{p.i + q.i, p.j + q.j} }
func (p point) sub(q point) point { return point{p.i - q.i, p.j - q.j} }
func (p point) mul(s int) point   { return point{p.i * s, p.j * s} }
func (p point) div(s int) point   { return point{p.i / s, p.j / s} }

func (p point) swap() point { return point{p.j, p.i} }

func (p point) l1dist(q point) int { return abs(p.i-q.i) + abs(p.j-q.j) }
func (p point) l2dist(q point) float64 {
	di, dj := float64(p.i-q.i), float64(p.j-q.j)
	return math.Sqrt(di*di + dj*dj)
}

func (p point) cmp(q point) int {
	ord := cmp.Compare(p.i, q.i)
	if ord == 0 {
		return cmp.Compare(p.j, q.j)
	} else {
		return ord
	}
}

type state struct{ pos, dir point }

var N = point{-1, 0}
var W = point{0, -1}
var E = point{0, 1}
var S = point{1, 0}

var NW = point{-1, -1}
var NE = point{-1, 1}
var SW = point{1, -1}
var SE = point{1, 1}

func directions() []point { return []point{NW, N, NE, W, E, SW, S, SE} }
func orthogonal() []point { return []point{N, W, E, S} }
func oblique() []point    { return []point{NW, NE, SW, SE} }

type point3 struct{ x, y, z int }
type point3f struct{ x, y, z float64 }

func (p point3) add(q point3) point3 { return point3{p.x + q.x, p.y + q.y, p.z + q.z} }
func (p point3) sub(q point3) point3 { return point3{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point3) mul(s int) point3    { return point3{p.x * s, p.y * s, p.z * s} }
func (p point3) div(s int) point3    { return point3{p.x / s, p.y / s, p.z / s} }

func (p point3) l1dist(q point3) int { return abs(p.x-q.x) + abs(p.y-q.y) + abs(p.z-q.z) }
func (p point3) l2dist(q point3) float64 {
	dx, dy, dz := float64(p.x-q.x), float64(p.y-q.y), float64(p.z-q.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (p point3f) add(q point3f) point3f { return point3f{p.x + q.x, p.y + q.y, p.z + q.z} }
func (p point3f) sub(q point3f) point3f { return point3f{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point3f) mul(s float64) point3f { return point3f{p.x * s, p.y * s, p.z * s} }
func (p point3f) div(s float64) point3f { return point3f{p.x / s, p.y / s, p.z / s} }

func (p point3) cmp(q point3) int {
	ord := cmp.Compare(p.x, q.x)
	if ord == 0 {
		ord = cmp.Compare(p.y, q.y)
		if ord == 0 {
			return cmp.Compare(p.z, q.z)
		} else {
			return ord
		}
	} else {
		return ord
	}
}
