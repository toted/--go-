package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

//后台用逆波兰算法
//用数组模拟栈
const imaxn int = 10000000
const fmaxn float64 = 1000000

var postfix [100]string
var now int = 0

func stringToNum(str string) (int, float64) {
	if strings.Contains(str, ".") == false {
		i, _ := strconv.ParseInt(str, 10, 64)
		return int(i), fmaxn
	} else {
		i, _ := strconv.ParseFloat(str, 10)
		return imaxn, i
	}
}

func Property(a string) int {
	switch a {
	case "/":
		return 3
	case "ln", "cos", "sin", "g", "p":
		return 2
	case "m", "d":
		return 1
	case "+", "-":
		return 0
	}
	return -1
}

func isOption(a string) bool {
	if a == "sin" || a == "cos" || a == "+" || a == "-" || a == "m" || a == "d" || a == "p" || a == "g" || a == "/" || a == "ln" {
		return true
	}
	return false
}

func infixToPostfix(reply string) {
	//用正则表达式匹配数字和符号
	//m表示✖，d表示÷，g表示开根号，p表示平方
	var stack [100]string
	count := -1
	r := regexp.MustCompile(`\d+\.?\d*|sin|cos|\+|\-|m|d|p|g|\/|ln`)
	str := r.FindAllString(reply, -1)
	for i := 0; i < len(str); i++ {
		if isOption(str[i]) {
			if count == -1 {
				count++
				stack[count] = str[i]
			} else {
				for {
					if count == -1 {
						count++
						stack[count] = str[i]
						break
					}
					if Property(str[i]) > Property(stack[count]) {
						count++
						stack[count] = str[i]
						break
					} else {
						postfix[now] = stack[count]
						now++
						count--
					}
				}
			}
		} else {
			postfix[now] = str[i]
			now++
		}
	}
	for i := count; i >= 0; i-- {
		postfix[now] = stack[i]
		now++
	}
}

func calculate() string {
	var stack [100]string
	count := 0
	flag := 1
	for i := 0; i < now; i++ {
		num11, num12 := stringToNum(stack[count])
		if postfix[i] == "+" || postfix[i] == "-" || postfix[i] == "m" || postfix[i] == "d" || postfix[i] == "/" {
			count--
		}
		num21, num22 := stringToNum(stack[count])
		if postfix[i] == "+" {
			if num11 == imaxn {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num12+num22, 'g', 6, 64)
				} else {
					stack[count] = strconv.FormatFloat(num12+float64(num21), 'g', 6, 64)
				}
			} else {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22+float64(num11), 'g', 6, 64)
				} else {
					stack[count] = strconv.Itoa(num11 + num21)

				}
			}
		} else if postfix[i] == "-" {
			if num11 == imaxn {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22-num12, 'g', 6, 64)
				} else {
					stack[count] = strconv.FormatFloat(float64(num21)-num12, 'g', 6, 64)
				}
			} else {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22-float64(num11), 'g', 6, 64)
				} else {
					stack[count] = strconv.Itoa(num21 - num11)
				}
			}
		} else if postfix[i] == "m" {
			if num11 == imaxn {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num12*num22, 'g', 6, 64)
				} else {
					stack[count] = strconv.FormatFloat(num12*float64(num21), 'g', 6, 64)
				}
			} else {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22*float64(num11), 'g', 6, 64)
				} else {
					stack[count] = strconv.Itoa(num11 * num21)
				}
			}
		} else if postfix[i] == "d" || postfix[i] == "/" {
			if num11 == imaxn {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22/num12, 'g', 6, 64)
				} else {
					stack[count] = strconv.FormatFloat(float64(num21)/num12, 'g', 6, 64)
				}
			} else {
				if num21 == imaxn {
					stack[count] = strconv.FormatFloat(num22/float64(num11), 'g', 6, 64)
				} else {
					stack[count] = strconv.FormatFloat(float64(num21)/float64(num11), 'g', 6, 64)
				}
			}
		} else if postfix[i] == "sin" {
			if num11 == imaxn {
				stack[count] = strconv.FormatFloat(math.Sin(num12), 'g', 6, 64)
			} else {
				stack[count] = strconv.FormatFloat(math.Sin(float64(num11)), 'g', 6, 64)
			}
		} else if postfix[i] == "cos" {
			if num11 == imaxn {
				stack[count] = strconv.FormatFloat(math.Cos(num12), 'g', 6, 64)
			} else {
				stack[count] = strconv.FormatFloat(math.Cos(float64(num11)), 'g', 6, 64)
			}
		} else if postfix[i] == "g" {
			if num11 == imaxn {
				stack[count] = strconv.FormatFloat(math.Sqrt(num12), 'g', 6, 64)
			} else {
				stack[count] = strconv.FormatFloat(math.Sqrt(float64(num11)), 'g', 6, 64)
			}
		} else if postfix[i] == "p" {
			if num11 == imaxn {
				stack[count] = strconv.FormatFloat(num12*num12, 'g', 6, 64)
			} else {
				stack[count] = strconv.Itoa((num11) * num11)
			}
		} else if postfix[i] == "ln" {
			if num11 == imaxn {
				stack[count] = strconv.FormatFloat(math.Log(num12), 'g', 6, 64)
			} else {
				stack[count] = strconv.FormatFloat(math.Log(float64(num11)), 'g', 6, 64)
			}
		} else {
			if flag == 1 {
				flag = 0
			} else {
				count++
			}
			stack[count] = postfix[i]
		}
	}
	return stack[0]
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		now = 0
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("接收失败：", err)
			break
		}
		fmt.Println("接收到来自前端的信息: " + reply)
		infixToPostfix(reply)
		//处理
		fmt.Printf("转为后缀表达式：")
		for i := 0; i < now; i++ {
			fmt.Printf("%s ", postfix[i])
		}
		fmt.Println()
		msg := calculate()
		fmt.Println("发送信息到前端：" + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("发送失败：", err)
			break
		}
	}
}

func main() {
	fmt.Println("后台服务开启...")
	http.Handle("/websocket", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
