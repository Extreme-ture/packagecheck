package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	clientSet "packagecheck/pkg/generated/clientset/versioned"
)

func main(){
	//创建连接kube-apiserver配置文件
	clientConfig, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	//初始化所有客户端连接
	damClient, err := clientSet.NewForConfig(clientConfig)
	if err != nil {
		panic(err)
	}
	get, err := damClient.PackagecheckV1().PackageChecks("default").Get(context.Background(), "hello", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n",get)
}