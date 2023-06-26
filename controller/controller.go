package controller

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"test2/contract"
	"test2/util"
	"time"
)

//合约地址
const contractAddr = "0xD1969F3e73Ea331026B4363361B55fD3f0B2E5a9"
//私链的账户
const decAccount = "0x9bBCc9a79A2503Eb4936e5BD31C94125C0ea25Ee"
//私链的密钥文件在本机的文件位置
const  ksPath = "D:\\BlockChain\\node1-3\\node1\\nodedata\\keystore"
//私链的账户的密码
const passwd ="123456789asd"

//初始化每天的额度
func init()  {
	reset(10,3,3)
}

//申请代币
func Withdraw(a string,ip string) (string,error) {
	//连接geth
	cil,err :=ethclient.Dial("http://localhost:8545")
	if err!=nil{
		log.Fatal("Failed to conn to geth !!")
	}
	//延时关闭geth
	defer cil.Close()

	//将合约进行实例
	test,err:= contract.NewTest(common.HexToAddress(contractAddr),cil)
	if err!=nil{
		log.Fatal(err)
	}

	//获取密钥文件
	ks,err :=util.GetFileContent(decAccount,ksPath)
	if err!=nil{
		log.Fatal(err)
	}

	//进行签名
	opts ,err := bind.NewTransactor(strings.NewReader(ks),passwd)
	if err!=nil{
		log.Fatal(err)
	}

	//执行合约Withdraw的方法
	_ , err= test.Withdraw(opts,common.HexToAddress(a),ip)
	if err!=nil{
		return "",err
	}

	return "获取成功",nil
}

//向合约地址付钱
func SetQuantity() (string,error)  {
	//连接geth
	cil,err:=ethclient.Dial("http://localhost:8545")
	if err!=nil{
		log.Fatal("Failed to conn to geth !!")
	}
	//延时关闭geth
	defer cil.Close()

	//将合约进行实例
	test,err:= contract.NewTest(common.HexToAddress(contractAddr),cil)
	if err!=nil{
		log.Fatal(err)
	}

	//获取密钥文件
	ks,err :=util.GetFileContent(decAccount,ksPath)
	if err!=nil{
		log.Fatal(err)
	}

	//进行签名
	opts ,err := bind.NewTransactor(strings.NewReader(ks),passwd)
	if err!=nil{
		log.Fatal(err)
	}

	//执行合约函数
	opts.Value = big.NewInt(1000000000000000000)
	_ ,err = test.Payme(opts)
	if err!=nil{
		return "",err
	}

	return "设置成功",nil

}

//重置函数
func reset(_num int64,_num1 int64,_num3 int64) (string,error)  {

	cil,err:=ethclient.Dial("http://localhost:8545")
	if err!=nil{
		log.Fatal("Failed to conn to geth !!")
	}
	defer cil.Close()

	test,err:= contract.NewTest(common.HexToAddress(contractAddr),cil)
	if err!=nil{
		log.Fatal(err)
	}

	ks,err :=util.GetFileContent(decAccount,ksPath)
	if err!=nil{
		log.Fatal(err)
	}

	opts ,err := bind.NewTransactor(strings.NewReader(ks),passwd)
	if err!=nil{
		log.Fatal(err)
	}

	_ ,err = test.Reset(opts,big.NewInt(_num),big.NewInt(_num1),big.NewInt(_num3))
	if err!=nil{
		return "",err
	}

	return "设置成功",nil

}

//每隔一天，进行重置
func BoottimeTimingSettlement() {
	reset(10,3,3)
	fmt.Println("计时已开始")
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		reset(10,3,3)
	}
}
