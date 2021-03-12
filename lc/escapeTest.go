package main

func main(){
	/*
	a := 0
	s := make([]*int,1,5)
	s[0] = &a
	 */
	test()
	testMake()
	s := new([]int) //返回的是指针
	testNew(s)
}

func testNew(p *[]int){
	println(p)
}

func test() (*[]int){
	return &[]int{}
}
func testMake()(*[]int){
	s := make([]int,0,5)
	return &s
}