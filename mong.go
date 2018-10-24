/*

@Time : 2018/10/24 15:02 

@Author : zhaoxiaoqiang

@File : mong.go

@Software: GoLand

*/
package main
import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("mgo.Dial error:%v",err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"superWang", "13478808311"},
		&Person{"David", "15040268074"})
	if err != nil {
		fmt.Println("c.insert error:%v",err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "superWang"}).One(&result)
	if err != nil {
		fmt.Println("c.find error:%v",err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}
