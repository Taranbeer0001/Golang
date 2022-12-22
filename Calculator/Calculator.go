package Calculator

func Add(x int, y int) (int, int, int) {
	return x, y, x + y
}
func Sub(x int, y int) (int, int, int) {
	return x, y, x - y
}
func Mul(x int, y int) (int, int, int) {
	return x, y, x * y
}
func Div(x int, y int) (int, int, int) {
	return x, y, x / y
}
func Mod(x int, y int) (int, int, int) {
	return x, y, x % y
}
