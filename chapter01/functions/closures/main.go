package main

func main() {
    makeAdder := func(m int) func(int) int {
        return func(n int) int {
            return m + n
        }
    }

    addFive := makeAdder(5)
    addTen := makeAdder(10)

    println(addFive(6))
    println(addTen(6))
    println(addFive(10))
}