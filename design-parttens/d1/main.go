package main

// import (
// 	fa "design-parttens/d1/factory"
// )

// func main() {
// 	var fruitFactory = &fa.FruitFactory{}
// 	fruitFactory.CreateFruit("apple").Show()
// 	fruitFactory.CreateFruit("peach").Show()

// 	fa.ShowFruit("apple")
// 	fa.ShowFruit("peach")
// }

import (
	si "design-parttens/d1/single"
	st "design-parttens/d1/structure"
	"net/http"

	"fmt"
)

func main1() {
	// for i := 0; i < 10; i++ {
	// 	s := si.GetSingeltoneDelay()
	// 	fmt.Printf("%p\n", s)
	// }

	// for i := 0; i < 10; i++ {
	// 	s := si.GetSingeltoneLock()
	// 	fmt.Printf("%p\n", s)
	// }

	for i := 0; i < 10; i++ {
		s := si.GetSingeltoneOnce()
		fmt.Printf("%p\n", s)
	}

}

// proxy
func main2() {
	//第一次获取
	service := st.NewProxyDataService()
	fmt.Println(service.GetData(1))
	//第二次获取，应该从缓存中获取
	fmt.Println(service.GetData(1))
	//获取其他数据
	fmt.Println(service.GetData(2))

}

// decorator
func main() {
	// 创建基础处理器
	h := &st.BasicHandler{}

	// 添加日志装饰
	loggedHandler := &st.LoggingDecorator{
		Handler: h}

	// 添加认证装饰
	secureHandler := &st.AuthDecorator{Handler: loggedHandler}

	// 注册处理器
	http.Handle("/basic", h)
	http.Handle("/logged", loggedHandler)
	http.Handle("/secure", secureHandler)

	fmt.Println("服务器启动在 :8089")
	http.ListenAndServe(":8089", nil)
}
