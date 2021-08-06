package main

import (
    "fmt"
    "strconv"
)

/**
思路：
1.求最大值：n位，则每位都是9，再转化为数值类型
2.从i=1开始递增，直到i=最大值，把期间的数值写入切片
*/
func printNumbers(n int) []int {
    if n == 0 {
        return nil
    }
    var maxStr string
    for i := 1; i <= n; i++ {
        maxStr += "9"
    }
    fmt.Println(maxStr)
    max, _ := strconv.Atoi(maxStr)
    var ret []int
    for i := 1; i <= max; i++ {
        ret = append(ret, i)
    }
    return ret
}

func main() {
    nums := printNumbers(2)
    fmt.Println(nums)
}
