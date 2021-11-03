package sdk

// AccountType - brokerage or IIS
type AccountType string

const (
	// AccountTinkoff - brokerage account
	AccountTinkoff AccountType = "Tinkoff"
	// AccountTinkoffIIS - IIS
	AccountTinkoffIIS AccountType = "TinkoffIis"
)

// Account - account info
type Account struct {
	Type AccountType `json:"brokerAccountType"`
	ID   string      `json:"brokerAccountId"`
}

// OperationType - all allowed exchange operations
type OperationType string

const (
	BUY                             OperationType = "Buy"
	SELL                            OperationType = "Sell"
	OperationTypeBrokerCommission   OperationType = "BrokerCommission"
	OperationTypeExchangeCommission OperationType = "ExchangeCommission"
	OperationTypeServiceCommission  OperationType = "ServiceCommission"
	OperationTypeMarginCommission   OperationType = "MarginCommission"
	OperationTypeOtherCommission    OperationType = "OtherCommission"
	OperationTypePayIn              OperationType = "PayIn"
	OperationTypePayOut             OperationType = "PayOut"
	OperationTypeTax                OperationType = "Tax"
	OperationTypeTaxLucre           OperationType = "TaxLucre"
	OperationTypeTaxDividend        OperationType = "TaxDividend"
	OperationTypeTaxCoupon          OperationType = "TaxCoupon"
	OperationTypeTaxBack            OperationType = "TaxBack"
	OperationTypeRepayment          OperationType = "Repayment"
	OperationTypePartRepayment      OperationType = "PartRepayment"
	OperationTypeCoupon             OperationType = "Coupon"
	OperationTypeDividend           OperationType = "Dividend"
	OperationTypeSecurityIn         OperationType = "SecurityIn"
	OperationTypeSecurityOut        OperationType = "SecurityOut"
	OperationTypeBuyCard            OperationType = "BuyCard"
)

// OrderStatus - operation result
type OrderStatus string

const (
	OrderStatusNew            OrderStatus = "New"
	OrderStatusPartiallyFill  OrderStatus = "PartiallyFill"
	OrderStatusFill           OrderStatus = "Fill"
	OrderStatusCancelled      OrderStatus = "Cancelled"
	OrderStatusReplaced       OrderStatus = "Replaced"
	OrderStatusPendingCancel  OrderStatus = "PendingCancel"
	OrderStatusRejected       OrderStatus = "Rejected"
	OrderStatusPendingReplace OrderStatus = "PendingReplace"
	OrderStatusPendingNew     OrderStatus = "PendingNew"
)

// OrderType - limited or just market buying
type OrderType string

const (
	// OrderTypeLimit - limited
	OrderTypeLimit OrderType = "Limit"
	// OrderTypeMarket - marked
	OrderTypeMarket OrderType = "Market"
)

// Order
type Order struct {
	ID            string        `json:"orderId"`
	FIGI          string        `json:"figi"`
	Operation     OperationType `json:"operation"`
	Status        OrderStatus   `json:"status"`
	RequestedLots int           `json:"requestedLots"`
	ExecutedLots  int           `json:"executedLots"`
	Type          OrderType     `json:"type"`
	Price         float64       `json:"price"`
}

// DefaultAccount - in system account is "Tinkoff" by default
const DefaultAccount = ""
