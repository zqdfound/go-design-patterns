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
	sa "design-parttens/d1/action"
	si "design-parttens/d1/single"
	st "design-parttens/d1/structure"
	"net/http"
	"time"

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
func main3() {
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

// template
func main4() {
	s1 := sa.TemplateWithLog("setUp", func() {
		fmt.Println("设置环境")
		time.Sleep(500 * time.Millisecond)
	})
	s2 := sa.TemplateWithLog("process", func() {
		fmt.Println("执行处理")
		time.Sleep(500 * time.Millisecond)
	})

	s3 := sa.TemplateWithLog("cleanUp", func() {
		fmt.Println("清理资源")
		time.Sleep(500 * time.Millisecond)
	})

	sa.Template(s1, s2, s3)

}

// strategy
func main5() {
	context := &sa.PaymentContext{}

	// 使用工厂创建策略
	creditCardParams := map[string]string{
		"name":    "李四",
		"cardNum": "6543210987654321",
		"cvv":     "456",
		"expiry":  "06/24",
	}
	context.SetStrategy(sa.PaymentStrategyFactory(sa.CreditCard, creditCardParams))
	fmt.Println(context.ExecutePayment(300.25))
}

// observer
func main() {
	appStock := sa.NewStock("AppStock")

	investor1 := sa.NewInvestor(1, 100.0)
	investor2 := sa.NewInvestor(2, 200.0)
	investor3 := sa.NewInvestor(3, 300.0)

	appStock.RegisterObserver(investor1)
	appStock.RegisterObserver(investor2)
	appStock.RegisterObserver(investor3)

	// appStock.UpdatePrice(100.0)
	// appStock.UpdatePrice(200.0)
	// appStock.UpdatePrice(300.0)
	appStock.UpdatePrice(400.0)

}
