package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) { //変数レシーバだと毎回変数のコピーが渡される。呼び出し元の変数は変更されない
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) //ポインタレシーバでは、変数から呼び出せる。これはGo側が暗黙的にポインタに変換して引き渡すため
	ScaleFunc(&v, 10)
	ScaleFunc(v, 10) //ポインタを引数とする関数はちゃんとポインタで渡さないとエラー

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
