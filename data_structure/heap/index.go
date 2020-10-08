package heap

func Parent(idx int) int {
	return (idx - 1) / 2
}

func Left(idx int) int {
	return 2*idx + 1
}

func Right(idx int) int {
	return Left(idx) + 1
}
