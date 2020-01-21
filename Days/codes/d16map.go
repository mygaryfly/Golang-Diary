/*
牛奶 Milk | Le lait
柠檬水 Lemonade | Limonade
咖啡 Coffee | Le café
橙汁 Orange juice | Jus d'orange
酸奶 Yogurt | Yaourt
椰汁 Coconut milk | Lait de coco
*/

// 构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
//先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。

package main

import (
"fmt"
"sort"
)

func main(){
	var fd map[string]string
	fd = map[string]string{"牛奶":"Le lait","柠檬水":"Limonade","橙汁":"Jus d'orange","咖啡":"Le café,","酸奶":"Yaourt","椰汁":"Lait de coco"}
	for k1,_:=range fd{
		fmt.Println(k1)
	}
	
	fmt.Println("-----打印原名和翻译后的名字-----")
	
	for k1,v1:=range fd{
	fmt.Printf("%s 法语名称:%s\n",k1,v1)
	}
		
	fmt.Println("-----按照英文名排序-----")
	ed:=[]string{}
	ed = []string{"Milk","Lemonade","Coffee","Orange juice","Yogurt","Coconut milk"}
	sort.Strings(ed)
	for _,v1:=range ed{
	fmt.Printf("英: %s\n",v1)
	}
}