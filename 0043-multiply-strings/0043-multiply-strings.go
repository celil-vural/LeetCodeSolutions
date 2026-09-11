const (
	ASCII_ZERO = 48
)
func multiply(num1 string, num2 string) string {
	/*
	       xyz
	       klm
	   x-------
	           (m.x+carry),(m.y+carry),(m.z)
	       (l.x+carry),(l.y+carry)(l.z),0
	   (k.x+carry),(k.y+carry),(k.z),0,0
	*/
	num1Array := []rune(num1)
	num2Array := []rune(num2)
	result := make([]rune, len(num1Array)+len(num2Array))
	for j := len(num2Array) - 1; j >= 0; j-- {
		carry := 0
		for i := len(num1Array) - 1; i >= 0; i-- {
			number1, number2, number3 := int(num1Array[i]-ASCII_ZERO), int(num2Array[j]-ASCII_ZERO), int(result[i+j+1]-ASCII_ZERO)
			if number3 < 0 {
				number3 = 0
			}
			product := number1*number2 + carry + number3
			carry = product / 10
			result[i+j+1] = rune(product%10 + ASCII_ZERO)
		}
		result[j] = rune(carry + ASCII_ZERO)
	}
	res:= string(result)
	for {
		if len(res) > 1 && res[0] == '0' {
			res = res[1:]
		} else {
			break
		}
	}
	return res
}
