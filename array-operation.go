package main

import "fmt"

func main() {
    var arr [10]int  // 声明了一个int类型的数组
    arr[0] = 42      // 数组下标是从0开始的
    arr[1] = 13      // 赋值操作
    fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
    fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0

    a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

    b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

    c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

    // 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
    doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

    // 上面的声明可以简化，直接忽略内部的类型
    easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

    slice := []byte {'a', 'b', 'c', 'd'}

    // 声明一个含有10个元素元素类型为byte的数组
    var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

    // 声明两个含有byte的slice
    var a, b []byte

    // a指向数组的第3个元素开始，并到第五个元素结束，
    a = ar[2:5]
    //现在a含有的元素: ar[2]、ar[3]和ar[4]

    // b是数组ar的另一个slice
    b = ar[3:5]
    // b的元素是：ar[3]和ar[4]

    // 声明一个数组
    var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
    // 声明两个slice
    var aSlice, bSlice []byte

    // 演示一些简便操作
    aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
    aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
    aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

    // 从slice中获取slice
    aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
    bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
    bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
    bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
    bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

    Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
    Slice_a := Array_a[2:5]
}
