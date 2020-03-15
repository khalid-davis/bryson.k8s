package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//使用kubeconfig连接
func K8sClientGoDemo() {
	var clientset *kubernetes.Clientset
	k8sConfig := flag.String("k8sconfig", "./kubeconfig", "kubernetes kubeconfig file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sConfig)
	if err != nil {
		log.Println(err)
		return
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connect k8s success")
	//获取pod
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pods.Items[0].Name)
}

//为请求添加上请求头Authentication: "abcdefg"
type headerAdder struct {
	headers map[string][]string
}
func main() {
	fmt.Println("tes")
}
