package logger

import "fmt"

var (
	ErrorBase  = NewError(1000, "Base exception")     // 基础错误
	ErrorDB    = NewError(1001, "Database exception") // 数据库错误
	ErrorCache = NewError(1002, "Cache exception")    // 缓存错误
	ErrorLog   = NewError(1003, "Log exception")      // 日志错误

	ErrorWriteFile = NewError(1008, "Write file exception") // 文件相关
	ErrorReadFile  = NewError(1009, "Read file exception")

	ErrorRouter          = NewError(3001, "Router execution exception") // 路由相关
	ErrorRouterRegister  = NewError(3002, "Router CreateUser exception")
	ErrorAgentStart      = NewError(3003, "Agent Startr exception")
	ErrorMakeMiddleware  = NewError(3004, "Make Middleware exception")
	ErrorHandleError     = NewError(3005, "Error Handle exception")
	ErrorRequestExecutor = NewError(3006, "HTTP Request Executor exception")
	ErrorServerPlugin    = NewError(3007, "ServerPlugin exception")

	ErrorParams               = NewError(4000, "Parameters exception") // 参数相关
	ErrorParamsIncomplete     = NewError(4001, "Incomplete parameters")
	ErrorAuthParamsIncomplete = NewError(4101, "CreateUser/UnRegister body is null")

	ErrorHost          = NewError(5001, "Request address exception") // 网络相关
	ErrorHTTPHandle    = NewError(5002, "HTTP Handle exception")
	ErrorNetForwarding = NewError(5003, "Network forwarding error")
	ErrorMiddleware    = NewError(5004, "Middleware Handle error")

	ErrorKafka             = NewError(6001, "Kafka execution exception") // kafka 相关
	ErrorKafkaConsumer     = NewError(6101, "Kafka consumer execution exception")
	ErrorKafkaProducer     = NewError(6201, "Kafka producer execution exception")
	ErrorKafkaProducerSend = NewError(6202, "Kafka producer send data execution exception")
)
var codes = map[int]string{}

type CustomError struct {
	msg  string
	code int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.code, e.msg)
}

func (e CustomError) Code() int {
	return e.code
}

func (e CustomError) Msg() string {
	return e.msg
}

func NewError(code int, msg string) *CustomError {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &CustomError{code: code, msg: msg}
}
