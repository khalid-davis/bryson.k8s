package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

type Client interface {
	Pods(namespace string) (*v1.PodList, error)
	Nodes() (*v1.NodeList, error)
}

//默认客户端
type NativeK8sClient struct {
	*kubernetes.Clientset
	config *rest.Config
	url string
}

func NewNativeK8sClient() (Client, error) {
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

func NewNativeK8sClientConfigByPath(kubeconfigPath string) (Client, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
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

func (client *NativeK8sClient) Nodes() (*v1.NodeList, error) {
	return client.CoreV1().Nodes().List(metav1.ListOptions{})
}

