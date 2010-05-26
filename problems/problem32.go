//Problem 32
//Find the sum of all numbers that can be written as pandigital products.

package projecteuler

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
)

func pandigital(a,b,c string)bool{
	digits := strings.Split(a+b+c,"",0)
	sort.SortStrings(digits)
	str := strings.Join(digits,"")
	if str == "123456789"{
		return true
	}
	return false
}

func Euler32()string{
	results := make(map[int]int)
	for a:=0;a<5000;a++{
		for b:=a;b<5000;b++{//This is a guess
			stra := strconv.Itoa(a)
			strb := strconv.Itoa(b)
			strab := strconv.Itoa(a*b)
			if len(stra) +len(strb) +len(strab) >9{
				break
			}else if len(stra) +len(strb) +len(strab) <9{
				continue
			}else{
				if pandigital(stra,strb,strab){
					results[a*b]=1
				}
			}
		}
	}
	//Avoiding duplicates
	sum:=0
	for prod := range results{
		sum+=prod
	}
	return fmt.Sprint(sum)
}
