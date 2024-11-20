package pkg

import (
	"math"
)

// Sigmoid 正太分布的平滑曲线
// 得出一个结论，值是0 - 2 之间，当两个值相近的时候，返回1
// 现在返回的是 最大值 0.7 - 0.8 之间， 两个数乘2 最大值才1.5 ，到达2， 需要调整
func SigmoidN(A, B float64) float64 {
	x := (A - B) / A
	// 为了避免数值溢出，我们可以根据 x 的正负来选择计算方式
	if x < 0 {
		return math.Exp(x) / (1 + math.Exp(x)) * 2
	} else {
		return 1 / (1 + math.Exp(-x)) * 2
	}
}

// Sigmoid 可以简单出个公式
// 相等的时候 是1
// 大于1 是助攻点击
// 小于1 是减少点击
func Sigmoid(A, B float64) float64 {
	if B == 0 {
		return 1
	}
	// 保留三位小数
	factor := math.Round(A/B*1000) / 1000
	if factor <= 0 {
		return 1
	}
	switch {
	case factor > 1 && factor <= 1.2:
		return RandFloat(1.0, 1.1, 3)
	case factor > 1.2 && factor <= 1.5:
		return RandFloat(1.1, 1.3, 3)
	case factor > 1.5 && factor <= 1.6:
		return RandFloat(1.3, 1.4, 3)
	case factor > 1.6 && factor <= 1.8:
		return RandFloat(1.4, 1.5, 3)
	case factor > 1.8 && factor <= 2:
		return RandFloat(1.5, 1.6, 3)
	case factor > 2:
		return RandFloat(1.6, 2, 3)
	case factor < 1:
		return factor
	}
	return 1
}

// AB 分组，当分到A组的时候返回true ， 分到B组是false
func AB(id, a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	if a > b {
		return !ab(id, b, a)
	} else {
		return ab(id, a, b)
	}
}

// ab 私有的分组方法，a 值必须小于b 值
func ab(id, a, b int) bool {
	m := Gcd(a, b)
	newA := a / m
	newB := b / m
	// 大余数
	bigCost := newA + newB
	// ab 之间的倍数
	// ab 是不是整倍数
	isInt := (newB % newA) == 0
	flat := newB / newA
	smallCost := 1 + flat
	// 整数倍, 求余即可得到结果
	if isInt {
		return id%bigCost < 1
	}
	// 非整数倍 ，前面的数据都进行比例分配，后面的部分给策略组
	val := id % bigCost
	if val < smallCost*newA {
		return val%smallCost < 1
	}
	// 超过2倍的小余数都是 b分类里
	return false
}

// Gcd 函数计算两个正整数的最大公约数
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
