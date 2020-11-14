package main
import (
	"fmt"
	"strconv"
	"strings"
	"time"
)
func main() {
	i:="3.1415926"
	j,_:=strconv.ParseFloat(i,64)
	fmt.Println(j)
	func () {
		var x int64 = 123
		y := strconv.FormatInt(x, 2)
		fmt.Println(y)
	}()
	test()
	testall()
	Demo()
	test1()
	test2()
	timer()
	calc(2,3)
}
func Demo(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("芜湖")
	}()
}
func test(){
	var x ="hongyannb,fenggenb,起飞！！！"
	y:=strings.Index(x,"e")
	fmt.Println(y)
	i:=strings.Split(x ,",")
	fmt.Println(i)
	h:=strings.Count(x,"n")
	if h>5{
		fmt.Println(x)
	}else{
		fmt.Println("打工人，打工魂")
	}
	j:=strings.Replace(x,"起飞","666",1)
	fmt.Println(j)
}
func testall(){
	m:=make(map[string] string)
	m["a"]="a"
	m["b"] = "B"
	fmt.Println(m)
	delete(m, "b")
	fmt.Println(m)
}
func fun(m map[string]string, key, value string) {
	m[key] = value
}
func test1() {
	m := make(map[string]string)
	fun(m, "z", "ZZ")
	fmt.Println(m)
}
func timer(){
	j:=time.Now().Format("2016-01-02 15:04:05")
	fmt.Println(j)
	for {
		fmt.Println("芜湖，起飞")
		timer := time.NewTimer(5*time.Second)
		<-timer.C
	}
}
func test2() {
	arr := make([]int, 2)
	arr[0] = 1;
	arr[1] = 2;
	for i := 3; i < 16; i++ {
		arr = append(arr, i)
		defer func(a []int) {
			fmt.Println(a, &a[0], len(a))
		}(arr)
	}
	fmt.Println(arr, &arr[0], len(arr))
}
func calc(a int,b...int) int{
	sum:=a
	for i:=0;i<len(b);i++{
		sum+=b[i]
	}
	fmt.Println(sum)
	return sum
}