package bigint

import (
	"strconv"
	"strings"
)

type Bigint struct {
	value string
}

func NewInt(num string) (Bigint, error) { return Bigint{value: num}, nil }

func (z *Bigint) Set(num string) error {
	z.value = num
	return nil
}
func Add(a, b Bigint) Bigint {
	var IsNegativeResult bool
	n1 := IsPositive(a)
	n2 := IsPositive(b)

	if n1 && !n2 { return Sub(a, MakePositive(&b)) }
	else if !n1 && n2 { return Sub(b, MakePositive(&a)) }
	else if !n1 && !n2 { IsNegativeResult = true }

	a1, _ := NewInt(a.value)
	b1, _ := NewInt(b.value)

	s1 := strings.Split(Abs(a1).value, "")
	s2 := strings.Split(Abs(b1).value, "")

	var memory string
	var bigger int
	var change int
	var length int
	var result []string

	if len(s1) > len(s2) {
		length = len(s1)
		bigger = 1
		change = len(s1) - len(s2)
	} else {
		bigger = 2
		length = len(s2)
		change = len(s2) - len(s1)
	}
	if bigger == 1 {
		for i := 0; i < change; i++ { s2 = append([]string{"0"}, s2...) }
	} else {
		for i := 0; i < change; i++ { s1 = append([]string{"0"}, s1...) }
	}

	for i := length - 1; i >= 0; i-- {
		numA, _ := strconv.ParseInt(s1[i], 10, 64)
		numB, _ := strconv.ParseInt(s2[i], 10, 64)

		var results []string
		res := numA + numB
		if memory != "" {
			m, _ := strconv.ParseInt(memory, 10, 64)
			res += m
			memory = ""
		}
		if res/10 >= 1 {
			results = strings.Split(strconv.Itoa(int(res)), "")
			memory = results[0]
			result = append([]string{results[1]}, result...)
		} else {
			r := strconv.Itoa(int(res))
			result = append([]string{r}, result...)
		}
	}
	if memory != "" { result = append([]string{memory}, result...) }
	trimmed := TrimString(strings.Join(result, ""))
	if IsNegativeResult { trimmed = "-" + trimmed }
	return Bigint{value: trimmed}
}
func Sub(a, b Bigint) Bigint {
	n1 := IsPositive(a)
	n2 := IsPositive(b)

	if !n1 && !n2 { return Add(MakePositive(&b), a) }
	else if n1 && !n2 { return Add(a, MakePositive(&b)) }
	else if !n1 && n2 { return Add(a, MakeNegative(&b)) }
	
	if !IsBigger(a, b) {
		x := Sub(b, a)
		z := strings.Split(x.value, "")
		z = append([]string{"-"}, z...)
		return Bigint{value: strings.Join(z, "")}
	}

	s1 := strings.Split(a.value, "")
	s2 := strings.Split(b.value, "")

	var memory int
	var bigger int
	var change int
	var length int
	var result []string

	if len(s1) > len(s2) {
		length = len(s1)
		bigger = 1
		change = len(s1) - len(s2)
	} else {
		bigger = 2
		length = len(s2)
		change = len(s2) - len(s1)
	}
	if bigger == 1 {
		for i := 0; i < change; i++ { s2 = append([]string{"0"}, s2...) }
	} else {
		for i := 0; i < change; i++ { s1 = append([]string{"0"}, s1...) }
	}

	for i := length - 1; i >= 0; i-- {
		numA, _ := strconv.ParseInt(s1[i], 10, 64)
		numB, _ := strconv.ParseInt(s2[i], 10, 64)
		var res int
		if memory != 0 {
			numA -= int64(memory)
			memory = 0
		}
		if numA >= numB { res = int(numA - numB)}
		else {
			res = 10 + int(numA-numB)
			memory = 1
		}
		result = append([]string{strconv.Itoa(res)}, result...)
	}
	trimmed := TrimString(strings.Join(result, ""))
	return Bigint{value: trimmed}
}
func Abs(x Bigint) Bigint {
	if strings.HasPrefix(x.value, "-") { return Bigint{value: strings.Replace(x.value, "-", "", 1)} }
	else { return x }
}

// func Multiply(a, b Bigint) Bigint {
// 	var IsResultPositive bool
// 	n1 := IsPositive(a)
// 	n2 := IsPositive(b)

// 	if (n1 && n2) || (!n1 && !n2) {
// 		IsResultPositive = true
// 	} else {
// 		IsResultPositive = false
// 	}

// 	a1, _ := NewInt(a.value)
// 	b1, _ := NewInt(b.value)

// 	s1 := strings.Split(Abs(a1).value, "")
// 	s2 := strings.Split(Abs(b1).value, "")

// 	var memory string
// 	var bigger int
// 	var change int
// 	var length int
// 	var result []string
// 	if len(s1) > len(s2) {
// 		length = len(s1)
// 		bigger = 1
// 		change = len(s1) - len(s2)
// 	} else {
// 		bigger = 2
// 		length = len(s2)
// 		change = len(s2) - len(s1)
// 	}
// 	if bigger == 1 {
// 		for i := 0; i < change; i++ {
// 			s2 = append([]string{"0"}, s2...)
// 		}
// 	} else {
// 		for i := 0; i < change; i++ {
// 			s1 = append([]string{"0"}, s1...)
// 		}
// 	}

// 	for i := length - 1; i >= 0; i-- {
// 		numA, _ := strconv.ParseInt(s1[i], 10, 64)
// 		numB, _ := strconv.ParseInt(s2[i], 10, 64)
// 		var results []string
// 		res := numA * numB

// 		if memory != "" {
// 			m, _ := strconv.ParseInt(memory, 10, 64)
// 			res += m
// 			memory = ""
// 		}
// 		if res >= 10 {
// 			results = strings.Split(strconv.Itoa(int(res)), "")
// 			memory = results[0]
// 			result = append(result, results[1])
// 		} else {
// 			r := strconv.Itoa(int(res))
// 			result = append([]string{r}, result...)
// 		}
// 		if i-1 < 0 {
// 			result = append([]string{memory}, result...)

// 		}
// 	}
// 	if !IsResultPositive {
// 		result = append([]string{"-"}, result...)
// 		return Bigint{value: strings.Join(result, "")}
// 	} else {
// 		return Bigint{value: strings.Join(result, "")}
// 	}
// }

// func Mod(a, b Bigint) Bigint {
// }

func IsPositive(x Bigint) bool {
	if strings.HasPrefix(x.value, "-") { return false }
	else { return true }
}
func TrimString(x string) string {
	for strings.HasPrefix(x, "0") && len(x) > 1 { x = strings.Replace(x, "0", "", 1) }
	return x
}
func MakePositive(x *Bigint) Bigint {
	if strings.HasPrefix(x.value, "-") { x.value = strings.Replace(x.value, "-", "", 1) }
	return Bigint{ value: x.value }
}
func MakeNegative(x *Bigint) Bigint {
	if !strings.HasPrefix(x.value, "-") {
		s := strings.Split(x.value, "")
		s = append([]string{"-"}, s...)
		res := strings.Join(s, "")
		return Bigint{value: res}

	} else { return Bigint{value: x.value} }
}
func IsBigger(a, b Bigint) bool {

	x1 := MakePositive(&Bigint{value: a.value})
	x2 := MakePositive(&Bigint{value: b.value})

	if len(x1.value) > len(x2.value) { return true }
	else if len(x1.value) == len(x2.value) {

		s1 := strings.Split(x1.value, "")
		s2 := strings.Split(x2.value, "")

		for i := 0; i < len(s1); i++ {
			num1, _ := strconv.ParseInt(s1[i], 10, 64)
			num2, _ := strconv.ParseInt(s2[i], 10, 64)
			if num1 > num2 { return true }
		}
		return false
	} else { return false }
}
