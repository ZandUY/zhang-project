package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test      bool  `ini:"test"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`

	RedisConfig `ini:"redis"`
}

func lodIni(fileName string, data interface{}) (err error) {
	//0.参数校验
	//0.1 传进来的data参数必须是指针类型（因为在函数中对其赋值）
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err := errors.New("data param should be a pointer")
		return err
	}

	//0.2 传进来的data参数必须是结构体指针类型（因为要对其成员赋值）

	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data param should be a struct pointer")
		return err
	}
	//1.读文件得到字节类型数据
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Println(lineSlice)
	//2.一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//2.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2如果是[开头表示就是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("invalid section at line %d:%s", idx+1, line)
				return err
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			//根据字符串sectionName去data对应的嵌套结构体中找到对应的字段
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					// fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//2.3 如果不是节就是键=值
			//2.3.1 以等号分隔键和值
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err

			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			//2.3.2根据structName去data中找到对应的嵌套结构体取出来

			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sValue.Kind() != reflect.Struct {
				err = fmt.Errorf("field %s should be a struct", structName)
				return err
			}

			// fmt.Printf("sValue Type:%s  sType:%s\n",sValue.Type(),sType)

			var fieldName string
			var fileType reflect.StructField
			//2.3.3遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fileType = field
				if field.Tag.Get("ini") == key {

					fieldName = field.Name
					break
				}
			}

			//2.3.4如果key=tag，给这个字段赋值
			//2.3.4.1 根据fieldName去取这个字段
			if len(fieldName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// fileObj := sValue.FieldByName(structName)
			fmt.Println(333, fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return err
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return err
				}
				fileObj.SetBool(valueBool)
			}
			//2.3.4.2 根据字段的类型去转换value
		}

	}
	return nil
}
func main() {
	cfg,err:=ini.Load("config.ini")
	if err!=nil{
		fmt.Println("load ini failed", err)
		return
	}
	fmt.Println(cfg.Section("redis").Key("host").String())
	// fmt.Println(cfg)
	// err := lodIni("./config.ini", &cfg)
	// if err != nil {
	// 	fmt.Println("load ini failed", err)
	// 	return
	// }
	// fmt.Println(cfg)
	// fmt.Println(cfg.Address, cfg.Port, cfg.Username, cfg.Password)

	

}
