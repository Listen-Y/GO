package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main1() {
	// 判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。
	fmt.Println(strings.EqualFold("ABC", "abc")) // ture

	// 判断s是否有前缀字符串prefix。
	fmt.Println(strings.HasPrefix("ABC", "AB")) // true

	// 判断字符串s是否包含子串substr。
	fmt.Println(strings.Contains("ABC", "B")) // true

	// 返回将所有字母都转为对应的小写版本的拷贝。
	fmt.Println(strings.ToLower("ABC")) // "abc"

	// 返回将所有字母都转为对应的大写版本的拷贝。
	fmt.Println(strings.ToUpper("abc")) // "ABC"

	// 返回count个s串联的字符串。
	fmt.Println(strings.Repeat("ABC", 2)) // "ABCABC"

	// 返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Println(strings.Trim(" ABC  ", " ")) // "ABC"

	// 返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
	fmt.Println(strings.TrimSpace("ABC  ")) // "ABC"

	// 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	split := strings.Split("a b c d e", " ")
	fmt.Println(split) // [a b c d e]

	// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	parseBool, err := strconv.ParseBool("F")
	fmt.Println(parseBool, err)

	// 返回字符串表示的整数值，接受正负号。
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	parseInt, err := strconv.ParseInt("1212", 10, 0)
	fmt.Println(parseInt, err)

	// 根据b的值返回"true"或"false"。
	fmt.Println(strconv.FormatBool(true))

	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	fmt.Println(strconv.FormatInt(111, 10))

	// Atoi是ParseInt(s, 10, 0)的简写。
	fmt.Println(strconv.Atoi("121"))

	// Itoa是FormatInt(i, 10) 的简写。
	fmt.Println(strconv.Itoa(1212))
}

func main2() {
	// 打开文件
	file, err := os.Open("D:\\test\\study.txt")
	if err != nil {
		log.Fatal("open file error")
		return
	}

	// 关闭
	defer file.Close()

	fmt.Println("file: ", file) // file:  &{0xc0000d0780} 所以返回的file是一个指针
}

func main3() {
	// 打开文件
	file, err := os.Open("D:\\test\\study.txt")
	if err != nil {
		log.Fatal("open file error")
		return
	}
	// 关闭
	defer file.Close()

	// 创建一个带缓冲区的Reader
	reader := bufio.NewReader(file)
	// 读取内容
	for {
		// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
		//如果ReadString方法在读取到delim之前遇到了错误
		//它会返回在错误之前读取的数据以及该错误（一般是io.EOF）
		readString, err := reader.ReadString('\n')
		fmt.Println("error", err)
		if err == io.EOF {
			fmt.Println(readString)
			break
		}
		fmt.Print(readString)
	}
	fmt.Println("文件读取结束")
}

func main4() {
	file, err := os.Open("D:\\test\\study1.txt")
	if err != nil {
		log.Fatal("open file error")
	}

	// 关闭
	defer file.Close()

	writer := bufio.NewWriter(file)

	data := 'a'
	for i := 0; i < 5; i++ {
		writeString, err := writer.WriteString(fmt.Sprintf("%v\n", data))
		if err != nil {
			log.Fatal("write error")
		}
		fmt.Println("write: ", writeString)
	}
	// 记得刷新缓冲区
	_ = writer.Flush()
	fmt.Println("写入成功")
}

func main5() {
	file, err := os.OpenFile("D:\\test\\study1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := "hello"
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%v\n", data)) //这里使用\n是为了换行, 也支持/r/n进行换行
	}

	// 必须进行
	_ = writer.Flush()
	fmt.Println("写入成功")
}

func main6() {
	file, err := os.OpenFile("D:\\test\\study1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := "byeBye"
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%v\r\n", data)) //这里使用\n是为了换行, 也支持/r/n进行换行
	}

	// 必须进行
	_ = writer.Flush()
	fmt.Println("写入成功")
}

func main7() {
	file, err := os.OpenFile("D:\\test\\study1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	data := "Listen"
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%v\r\n", data)) //这里使用\n是为了换行, 也支持/r/n进行换行
	}

	// 必须进行
	_ = writer.Flush()
	fmt.Println("写入成功")
}

func main() {
	file, err := os.OpenFile("D:\\test\\study1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	// 先读出来
	data := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		data = append(data, readString)
	}
	fmt.Println("读取数据成功: ", data)

	// 写入数据
	writer := bufio.NewWriter(file)
	for i := 0; i < len(data); i++ {
		_, _ = writer.WriteString(fmt.Sprintf("%v", data[i])) //这里使用\n是为了换行, 也支持/r/n进行换行
	}

	// 必须进行
	_ = writer.Flush()
	fmt.Println("写入成功")
}
