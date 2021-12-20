package main

func getMatrix(base, vector Vector) Matrix {
	matrix := Matrix{}

	if abs(base.X) == abs(vector.X) {
		matrix[0][0] = div(base.X, vector.X)
	}
	if abs(base.X) == abs(vector.Y) {
		matrix[0][1] = div(base.X, vector.Y)
	}
	if abs(base.X) == abs(vector.Z) {
		matrix[0][2] = div(base.X, vector.Z)
	}

	if abs(base.Y) == abs(vector.X) {
		matrix[1][0] = div(base.Y, vector.X)
	}
	if abs(base.Y) == abs(vector.Y) {
		matrix[1][1] = div(base.Y, vector.Y)
	}
	if abs(base.Y) == abs(vector.Z) {
		matrix[1][2] = div(base.Y, vector.Z)
	}

	if abs(base.Z) == abs(vector.X) {
		matrix[2][0] = div(base.Z, vector.X)
	}
	if abs(base.Z) == abs(vector.Y) {
		matrix[2][1] = div(base.Z, vector.Y)
	}
	if abs(base.Z) == abs(vector.Z) {
		matrix[2][2] = div(base.Z, vector.Z)
	}

	return matrix
}

func div(a, b int) int {
	if b == 0 {
		return 1
	}

	return a / b
}
