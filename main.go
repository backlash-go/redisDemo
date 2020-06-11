package main

import (
	"fmt"
	"redisDemo/conf"
	"redisDemo/method"
)

func main() {
	//initial redis
	conf.InitialRedis()
	//hmset
	hashK1 := []interface{}{"hk1", "name", "xxb", "age", "18", "sex", "nan", "yanzhi", "100"}
	if err := method.SetHashValues(hashK1); err != nil {
		fmt.Println(err)
	}
	if err := method.SetKeyTtl("hk1", 20); err != nil {
		fmt.Println(err)
	}
	//hmget
	data, _ := method.GetHashValues([]interface{}{"hk1", "name", "sex"})
	fmt.Println(data) //[xixb nang]

    //set and get string key and expires
	method.SetEx("sk1","oh yeah",60)
	s1 , _ := method.GetStringKey("sk1")
	fmt.Println(s1)




}
