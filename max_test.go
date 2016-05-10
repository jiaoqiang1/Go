package main

import "testing"

func main(t *testing.T) {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max(a, b)

   t.Printf( "最大值是 : %d\n", ret )
}
