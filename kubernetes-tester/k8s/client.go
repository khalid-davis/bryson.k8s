package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type Client interface {
	Pods(namespace string) (*v1.PodList, error)
}

//默认客户端
type NativeK8sClient struct {
	*kubernetes.Clientset
	config *rest.Config
	url string
}

func NewK8sClient() (Client, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	return &NativeK8sClient{
		Clientset: clientset,
		config:    config,
		url:       "",
	}, nil
}

func (client *NativeK8sClient) Pods(namespace string) (*v1.PodList, error) {
	return client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}


