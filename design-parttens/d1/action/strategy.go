package action

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardStrategy 信用卡支付策略
type CreditCardStrategy struct {
	name    string
	cardNum string
	cvv     string
	expiry  string
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用信用卡支付 %.2f 元，卡号后四位: %s", amount, c.cardNum[len(c.cardNum)-4:])
}

// AlipayStrategy 支付宝支付策略
type AlipayStrategy struct {
	account string
}

func (a *AlipayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝支付 %.2f 元，账户: %s", amount, a.account)
}

// WechatPayStrategy 微信支付策略
type WechatPayStrategy struct {
	openID string
}

func (w *WechatPayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用微信支付 %.2f 元，OpenID: %s", amount, w.openID)
}

// PaymentContext 支付上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) string {
	if p.strategy == nil {
		return "未设置支付策略"
	}
	return p.strategy.Pay(amount)
}

// PaymentType 支付类型枚举
type PaymentType int

const (
	CreditCard PaymentType = iota
	Alipay
	WechatPay
)

// PaymentStrategyFactory 策略工厂
func PaymentStrategyFactory(paymentType PaymentType, params map[string]string) PaymentStrategy {
	switch paymentType {
	case CreditCard:
		return &CreditCardStrategy{
			name:    params["name"],
			cardNum: params["cardNum"],
			cvv:     params["cvv"],
			expiry:  params["expiry"],
		}
	case Alipay:
		return &AlipayStrategy{account: params["account"]}
	case WechatPay:
		return &WechatPayStrategy{openID: params["openID"]}
	default:
		return nil
	}
}

// func main() {
// 	context := &PaymentContext{}

// 	// 使用工厂创建策略
// 	creditCardParams := map[string]string{
// 		"name":    "李四",
// 		"cardNum": "6543210987654321",
// 		"cvv":     "456",
// 		"expiry":  "06/24",
// 	}
// 	context.SetStrategy(PaymentStrategyFactory(CreditCard, creditCardParams))
// 	fmt.Println(context.ExecutePayment(300.25))
// }
