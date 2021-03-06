package Server

import (
	"errors"
)
 
type Exchange struct {
	_userManager UserManager
	//_tradingSessionManager TradingSessionManager 
	_envConfig EnvConfig
	_tradingSession TradingSession
}

func (exchange *Exchange) AddOrder(order Order) (bool, error) {
	//Exchange负责对用户的校验，以及把order分配给对应的Session，现在只有一个Session，所以只做校验
	if exchange._userManager.Check(order.User()) == true {
		return (exchange._tradingSession.AddOrder(order))
	} else {
		return false,errors.New("Invalid User.")
	}
}

func (exchange *Exchange) CancelOrder(order Order) (bool, error) {
	if exchange._userManager.Check(order.User()) == true {
		return exchange._tradingSession.CancelOrder(order)
	} else {
		return false, errors.New("Invalid User.");
	}
}

func (exchange *Exchange) QueryOrderBook(productId int) OrderBook {
	return exchange._tradingSession.QueryOrderBook(productId)
}

func (exchange *Exchange) Init() {
	exchange._userManager = CreateUserManager()
	exchange._envConfig = CreateEnvConfig()
	
	tradingSessionType := exchange._envConfig.TradingSessionType()
	executionSystemType := exchange._envConfig.ExecutionSystemType()
	
	tradingSession := CreateTradingSession(tradingSessionType, executionSystemType)
	//tradingSession.SetExecutionSystemType(executionSystemType)
	
	//定时启动或者挂起或者停止session
	tradingSession.Start()
	//tradingSession.Suspend()
	//tradingSession.Stop()
}

func CreateExchange() *Exchange {
	exchange := new(Exchange)
	exchange.Init()
	return exchange
} 