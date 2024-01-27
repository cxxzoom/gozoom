package tempconv

func F2C(f F) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func C2F(c Celsius) F {
	return F(c*9/5 + 32)
}
