package main

import (
	"fmt"
	pointer2 "learn-go/pointer2"
	structPoint "learn-go/struct"
)

func main() {
	var i int
	// アドレス演算子
	// 任意の型からそのポインタ型を生成することができる
	p := &i
	fmt.Printf("%T\n", p)
	// ポインタ型の変数から値を参照するには、演算子 * をポインタ型の変数の前に置く
	// これにより、デリファレンスをすることできる
	i = 5
	fmt.Println(*p)
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)

	k := &[3]int{1, 2, 3}
	pow(k)
	fmt.Println(k)

	pointer2.PointerPrint(&i)
	strPoint := structPoint.StructPointerReference()
	fmt.Println(strPoint.Point)
}

func inc(p *int) {
	/* pをでリファレンスして +1 増分 */
	*p++
}

func pow(k *[3]int) {
	i := 0
	for i < 3 {
		(*k)[i] = (*k)[i] * (*k)[i]
		i++
	}
}
