package sdl

import "math"

// [Rect] is a rectangle, with the origin at the upper left (using integers).
//
// [Rect]: https://wiki.libsdl.org/SDL3/SDL_Rect
type Rect struct {
	X, Y, W, H int32
}

// [FRect] is a rectangle, with the origin at the upper left (using floating point values).
//
// [FRect]: https://wiki.libsdl.org/SDL3/SDL_FRect
type FRect struct {
	X, Y, W, H float32
}

// [FPoint] is a structure specifying the structure that defines a point (using floating point values)s.
//
// [FPoint]: https://wiki.libsdl.org/SDL3/SDL_FPoint
type FPoint struct {
	X, Y float32
}

// [Point] is a structure specifying the structure that defines a point (using integers)s.
//
// [Point]: https://wiki.libsdl.org/SDL3/SDL_Point
type Point struct {
	X, Y int32
}

// [RectToFRect] converts an [Rect] to [FRect].
//
// [RectToFRect]: https://wiki.libsdl.org/SDL3/SDL_RectToFRect
func RectToFRect(rect Rect) FRect {
	return FRect{float32(rect.X), float32(rect.Y), float32(rect.W), float32(rect.H)}
}

// [PointInRect] determines whether a point resides inside a rectangle.
//
// [PointInRect]: https://wiki.libsdl.org/SDL3/SDL_PointInRect
func PointInRect(p Point, r Rect) bool {
	return (p.X >= r.X) && (p.X < (r.X + r.W)) && (p.Y >= r.Y) && (p.Y < (r.Y + r.H))
}

// [PointInRectFloat] determines whether a point resides inside a floating point rectangle.
//
// [PointInRectFloat]: https://wiki.libsdl.org/SDL3/SDL_PointInRectFloat
func PointInRectFloat(p FPoint, r FRect) bool {
	return (p.X >= r.X) && (p.X < (r.X + r.W)) && (p.Y >= r.Y) && (p.Y < (r.Y + r.H))
}

// [RectEmpty] determines whether a rectangle has no area.
//
// [RectEmpty]: https://wiki.libsdl.org/SDL3/SDL_RectEmpty
func RectEmpty(r Rect) bool {
	return r.W <= 0 || r.H <= 0
}

// [RectEmptyFloat] determines whether a floating point rectangle takes no space.
//
// [RectEmptyFloat]: https://wiki.libsdl.org/SDL3/SDL_RectEmptyFloat
func RectEmptyFloat(r FRect) bool {
	return r.W <= 0 || r.H <= 0
}

// [GetRectAndLineIntersection] calculates the intersection of a rectangle and line segment.
//
// [GetRectAndLineIntersection]: https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersection
func GetRectAndLineIntersection(rect Rect, x1, y1, x2, y2 *int32) bool {
	return sdlGetRectAndLineIntersection(&rect, x1, y1, x2, y2)
}

// [GetRectAndLineIntersectionFloat] calculates the intersection of a rectangle and line segment with float precision.
//
// [GetRectAndLineIntersectionFloat]: https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat
func GetRectAndLineIntersectionFloat(rect FRect, x1, y1, x2, y2 *float32) bool {
	return sdlGetRectAndLineIntersectionFloat(&rect, x1, y1, x2, y2)
}

// [GetRectEnclosingPoints] calculates a minimal rectangle enclosing a set of points.
//
// [GetRectEnclosingPoints]: https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints
func GetRectEnclosingPoints(points []Point, clip *Rect) (Rect, bool) {
	var result Rect
	var pointsPtr *Point
	if len(points) > 0 {
		pointsPtr = &points[0]
	}
	ret := sdlGetRectEnclosingPoints(pointsPtr, int32(len(points)), clip, &result)
	return result, ret
}

// [GetRectEnclosingPointsFloat] calculates a minimal rectangle enclosing a set of points with float precision.
//
// [GetRectEnclosingPointsFloat]: https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPointsFloat
func GetRectEnclosingPointsFloat(points []FPoint, clip *FRect) (FRect, bool) {
	var result FRect
	var pointsPtr *FPoint
	if len(points) > 0 {
		pointsPtr = &points[0]
	}
	ret := sdlGetRectEnclosingPointsFloat(pointsPtr, int32(len(points)), clip, &result)
	return result, ret
}

// [GetRectIntersection] calculates the intersection of two rectangles.
//
// [GetRectIntersection]: https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection
func GetRectIntersection(a, b Rect) (Rect, bool) {
	var result Rect
	ret := sdlGetRectIntersection(&a, &b, &result)
	return result, ret
}

// [GetRectIntersectionFloat] calculates the intersection of two rectangles with float precision.
//
// [GetRectIntersectionFloat]: https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat
func GetRectIntersectionFloat(a, b FRect) (FRect, bool) {
	var result FRect
	ret := sdlGetRectIntersectionFloat(&a, &b, &result)
	return result, ret
}

// [GetRectUnion] calculates the union of two rectangles.
//
// [GetRectUnion]: https://wiki.libsdl.org/SDL3/SDL_GetRectUnion
func GetRectUnion(a, b Rect) (Rect, bool) {
	var result Rect
	ret := sdlGetRectUnion(&a, &b, &result)
	return result, ret
}

// [GetRectUnionFloat] calculates the union of two rectangles.
//
// [GetRectUnionFloat]: https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat
func GetRectUnionFloat(a, b FRect) (FRect, bool) {
	var result FRect
	ret := sdlGetRectUnionFloat(&a, &b, &result)
	return result, ret
}

// [HasRectIntersection] determines whether two rectangles intersect.
//
// [HasRectIntersection]: https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection
func HasRectIntersection(a, b Rect) bool {
	return sdlHasRectIntersection(&a, &b)
}

// [HasRectIntersectionFloat] determines whether two rectangles intersect with float precision.
//
// [HasRectIntersectionFloat]: https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat
func HasRectIntersectionFloat(a, b FRect) bool {
	return sdlHasRectIntersectionFloat(&a, &b)
}

// [RectsEqualEpsilon] determines whether two floating point rectangles are equal, within some given epsilon.
//
// [RectsEqualEpsilon]: https://wiki.libsdl.org/SDL3/SDL_RectsEqualEpsilon
func RectsEqualEpsilon(a, b FRect, epsilon float32) bool {
	return (float32(math.Abs(float64(a.X-b.X))) <= epsilon) &&
		(float32(math.Abs(float64(a.Y-b.Y))) <= epsilon) &&
		(float32(math.Abs(float64(a.W-b.W))) <= epsilon) &&
		(float32(math.Abs(float64(a.H-b.H))) <= epsilon)
}

// [RectsEqualFloat] determines whether two floating point rectangles are equal, within a default epsilon.
//
// [RectsEqualFloat]: https://wiki.libsdl.org/SDL3/SDL_RectsEqualFloat
func RectsEqualFloat(a, b FRect) bool {
	return RectsEqualEpsilon(a, b, FltEpsilon)
}

// [RectsEqual] determines whether two rectangles are equal.
//
// [RectsEqual]: https://wiki.libsdl.org/SDL3/SDL_RectsEqual
func RectsEqual(a, b Rect) bool {
	return a == b
}
