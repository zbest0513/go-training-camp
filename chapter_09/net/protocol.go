package net

type GoimProtocol struct {
	PackageLength int32  //包总长度 4字节
	HeaderLength  int16  //包头长度 2字节
	Ver           Ver    //请求版本 2字节
	Opt           Opt    //请求指令 4字节
	Seq           int32  //请求序列号 4字节
	Body          []byte //包体内容，可能根据版本和指令去做不同的解析策略
}

type Opt int32
type Ver int16

const (
	OptMsg = 1 //文字消息
	OptImg = 2 //base64图片
)

const (
	Ver1 = 1 // 版本1
)

func NewGoimProtocol() *GoimProtocol {
	return &GoimProtocol{
		HeaderLength: 12, //当前协议包头长度固定，包头2字节，版本2字节，指令4字节，序列号4字节
	}
}

func (p *GoimProtocol) toBytes() []byte {
	bytes1 := Int32ToBytes(p.PackageLength)
	bytes2 := Int16ToBytes(p.HeaderLength)
	bytes3 := Int16ToBytes(p.Ver)
	bytes4 := Int32ToBytes(p.Opt)
	bytes5 := Int32ToBytes(p.Seq)
	bytes6 := Int32ToBytes(p.Seq)
	return BytesCombine(bytes1, bytes2, bytes3, bytes4, bytes5, bytes6, p.Body)
}

func (p *GoimProtocol) Encode(ver Ver, opt Opt, seq int32, body interface{}) []byte {
	//TODO 根据ver 和 opt 构建一个body的编码器
	// 这里简单都看成string 转 byte数组,所以body直接断言成string
	str := body.(string)
	strBytes := Str2Bytes(str)
	//包体长度
	bodyLength := len(strBytes)
	p.Ver = ver
	p.Seq = seq
	p.Body = strBytes
	//PackageLength 本身长度
	pl := 4
	p.PackageLength = int32(p.HeaderLength) + int32(bodyLength) + int32(pl)
	return p.toBytes()
}

func (p *GoimProtocol) DecodePackageLength(data []byte) int32 {
	return BytesToInt32(data)
}

func (p *GoimProtocol) Decode(data []byte) interface{} {
	p.HeaderLength = BytesToInt16(data[0:2])
	p.Ver = Ver(BytesToInt16(data[2:4]))
	p.Opt = Opt(BytesToInt32(data[4:8]))
	//TODO 可能根据解析出来的version 和opt body有不同的解析策略，这里简单只转字符串
	p.Seq = BytesToInt32(data[8:12])
	p.Body = data[12:]
	return Bytes2Str(p.Body)
}
