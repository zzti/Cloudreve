package util

import (
	"math/rand"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	letterRunes := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// ContainsUint 返回list中是否包含
func ContainsUint(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// IsInExtensionList 返回文件的扩展名是否在给定的列表范围内
func IsInExtensionList(extList []string, fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	// 无扩展名时
	if len(ext) == 0 {
		return false
	}
	// 扩展名匹配,忽略大小写
	if CaseInSensitiveContainsString(extList, ext[1:], false) {
		return true
	}

	return false
}

// ContainsString 返回list中是否包含
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CaseInSensitiveContainsString 返回list中是否包含,考虑大小写
func CaseInSensitiveContainsString(s []string, e string, caseSensitive bool) bool {
	for _, a := range s {
		if caseSensitive {
			// if a == e {
			// 字符串比较使用更高效的方法strings.Compare
			if strings.Compare(a, e) == 0 {
				return true
			}
		} else {
			// 字符串比较使用更高效的方法strings.EqualFold,该方法不区分大小写
			if strings.EqualFold(a, e) {
				return true
			}
		}
	}
	return false
}

// Replace 根据替换表执行批量替换
func Replace(table map[string]string, s string) string {
	for key, value := range table {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

// BuildRegexp 构建用于SQL查询用的多条件正则
func BuildRegexp(search []string, prefix, suffix, condition string) string {
	var res string
	for key, value := range search {
		res += prefix + regexp.QuoteMeta(value) + suffix
		if key < len(search)-1 {
			res += condition
		}
	}
	return res
}

// BuildConcat 根据数据库类型构建字符串连接表达式
func BuildConcat(str1, str2 string, DBType string) string {
	switch DBType {
	case "mysql":
		return "CONCAT(" + str1 + "," + str2 + ")"
	default:
		return str1 + "||" + str2
	}
}

// SliceIntersect 求两个切片交集
func SliceIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// SliceDifference 求两个切片差集
func SliceDifference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := SliceIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
