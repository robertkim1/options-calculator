package blackScholes

import "math"

func CDF(x float64) float64 {
	return 0.5 * math.Erfc(-(x) / (math.Sqrt2))
}

func D1(spotPrice, strikePrice, rfRate, volatility, time float64) float64 {
	return (math.Log(spotPrice/strikePrice) + (rfRate + (math.Pow(volatility,2) / 2)) * time) / (volatility * math.Sqrt(time))
}

func D2(d1, volatility, time float64) float64 {
	return d1 - volatility * math.Sqrt(time)
}

func CallPrice(d1 float64, spotPrice float64, d2 float64, strikePrice float64, rfRate float64, time float64) float64 {
	return CDF(d1) * spotPrice - CDF(d2) * strikePrice * math.Exp(-rfRate * time)
}

func PutPrice(d1 float64, spotPrice float64, d2 float64, strikePrice float64, rfRate float64, time float64) float64 {
	return strikePrice * math.Exp(-rfRate * time) * CDF(-d2) - spotPrice * CDF(-d1)
}