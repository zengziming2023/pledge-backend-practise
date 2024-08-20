package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"pledge-backend-practise/db"
	"regexp"
	"strings"
	"sync"
)

var oneOfValsCache = map[string][]string{}
var oneOfValsCacheRWLock = sync.RWMutex{}
var splitParamsRegexString = `'[^']*'|\S+`
var splitParamsRegex = regexp.MustCompile(splitParamsRegexString)

// IsPassword 判断是否为合法密码
func IsPassword(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9!@#￥%^&*]{6,20}$`, fl.Field().Interface().(string)); isOk {
			return true
		}
	}
	return false
}

func IsPhoneNumber(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^1[0-9]{10}$`, fl.Field().Interface().(string)); isOk {
			return true
		}
	}
	return false
}

func IsEmail(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, fl.Field().Interface().(string)); isOk {
			return true
		}
	}
	return false
}

func CheckUserNicknameLength(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^.{1,20}$`, fl.Field().Interface().(string)); isOk {
			return true
		}
	}
	return false
}

func CheckUserAccount(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[A-Za-z][A-Za-z0-9]{5,19}$`, fl.Field().Interface().(string)); isOk {
			return true
		}
	}
	return false
}

type dataStruct struct {
	DataCount int //这个结构体用来保存查询到的记录条数
}

func OnlyOne(fl validator.FieldLevel) bool {
	vals := parseOnlyOneParam(fl.Param())
	if len(vals) <= 0 {
		panic("OnlyOne parameter err")
	}

	tableName := vals[0]
	fieldName := vals[1]

	//  check db
	var data dataStruct
	sqlStr := fmt.Sprintf("`%s=?`", fieldName)
	db.MySql.Table(tableName).Select("COUNT(*)").Where(sqlStr, fl.Field().Interface()).Scan(&data.DataCount)

	if data.DataCount > 0 {
		return false
	}
	// 没触发false就说明没有重复记录，返回true
	return true

}

func parseOnlyOneParam(s string) []string {
	oneOfValsCacheRWLock.RLock()
	vals, ok := oneOfValsCache[s]
	oneOfValsCacheRWLock.RUnlock()

	if !ok {
		oneOfValsCacheRWLock.Lock()
		vals = splitParamsRegex.FindAllString(s, -1)
		for index, v := range vals {
			vals[index] = strings.ReplaceAll(v, "'", "")
		}
		oneOfValsCache[s] = vals
		oneOfValsCacheRWLock.Unlock()
	}
	return vals
}
