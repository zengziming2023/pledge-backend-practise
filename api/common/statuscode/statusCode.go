package statuscode

const (
	// LangZh language
	LangZh   = 111
	LangEn   = 112
	LangZhTw = 113

	// CommonSuccess common
	CommonSuccess      = 0
	CommonErrServerErr = 1000
	ParameterEmptyErr  = 1001

	TokenErr = 1102 //token error

	// PNameEmpty muti-sign
	PNameEmpty   = 1201 //p_name empty
	ChainIdEmpty = 1202 //chain id empty
	ChainIdErr   = 1203 //chain id error

	NameOrPasswordErr = 1303 //name or password error

)
