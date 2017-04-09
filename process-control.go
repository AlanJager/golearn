package main

import "fmt"

func myFunc() {
    i := 0
Here:   //这行的第一个词，以冒号结束作为标签
    println(i)
    i++
    goto Here   //跳转到Here去
}

func main(){
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else {
        fmt.Println("x is less than 10")
    }

    // 计算获取值x,然后根据x返回的大小，判断是否大于10。
    if x := computedValue(); x > 10 {
        fmt.Println("x is greater than 10")
    } else {
        fmt.Println("x is less than 10")
    }

    if integer == 3 {
        fmt.Println("The integer is equal to 3")
    } else if integer < 3 {
        fmt.Println("The integer is less than 3")
    } else {
        fmt.Println("The integer is greater than 3")
    }

    sum := 0;
    for index:=0; index < 10 ; index++ {
        sum += index
    
    }
    fmt.Println("sum is equal to ", sum)


}
