package cb

type ReqType string

const (
	HttpGet  ReqType = "http.get"
	HttpPost ReqType = "http.post"
	Grpc     ReqType = "grpc"
)

type ParamType string

const (
	JSON ParamType = "json"
	XML  ParamType = "xml"
	TEXT ParamType = "text"
	FORM ParamType = "form"
)

type Result struct {
	Code   int    //0成功，非0失败
	Reason string //原因
	Data   interface{}
}

type FallBack struct {
	F func() Result
}
type Request struct {
	ReqType     ReqType           //请求类型,http_get
	ServiceAddr string            //服务地址
	ParamType   ParamType         //参数类型
	Param       interface{}       //参数
	Metadata    map[string]string //元数据或请求头
	Fallback    FallBack          //降级方法
}
