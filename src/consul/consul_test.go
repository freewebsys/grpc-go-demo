package src

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"testing"
	"net"
)

const Id = "1234567890"

func TestRegister(t *testing.T) {

	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = Id + "01"
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"group:user-tomcat-01"}
	registration.Address = "127.0.0.1"

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check
	log.Println("get registration : ", registration)
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

	//设置第二个对象。注册两个tomcat。
	registration.ID = Id + "02"
	registration.Name = "user-tomcat"
	registration.Address = "10.0.2.15"
	registration.Tags = []string{"group:user-tomcat-02"}
	log.Println("get registration : ", registration)
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

	//设置第3个对象。注册两个tomcat。
	registration.ID = Id + "03"
	registration.Name = "user-tomcat"
	registration.Address = "10.0.2.16"
	registration.Tags = []string{"group:user-tomcat-03"}
	check.HTTP = fmt.Sprintf("http://%s:%s%s", registration.Address, "8083", "/check")
	log.Println("get registration : ", registration)
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}

}

func TestDregister(t *testing.T) {

	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	err = client.Agent().ServiceDeregister(Id+ "01")
	if err != nil {
		log.Fatal("register server error : ", err)
	}

}


func TestLookup(t *testing.T){
	hostname := "www.baidu.com"
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Printf("lookup host error: %v\n", err)
	} else {
		fmt.Printf("addrs: %v", addrs)
	}

}