package controllers

import (
	"bryson.k8s/kubernetes-tester/k8s"
	"github.com/astaxie/beego"
	"log"
)

type RBACController struct {
	beego.Controller
}

var k8sClient k8s.Client

var selfNamespace string
var otherNamespace string
var kubeconfigPath string

func init() {
	selfNamespace = "k8s-tester"
	otherNamespace = "greenstock"
	//kubeconfigPath = "E:\\workspace\\learning\\go-learning\\Golang_Workspace\\src\\bryson.k8s\\kubernetes-tester\\kubeconfig"
	kubeconfigPath = "E:\\workspace\\learning\\go-learning\\Golang_Workspace\\src\\bryson.k8s\\kubernetes-tester\\k8s-tester-config"
}

// @router /inclusterselfnamespace [get]
func (controller *RBACController) RBACInClusterSelfNamespace() {
	k8sClient, err := k8s.NewNativeK8sClientConfigByPath(kubeconfigPath)
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	podList, err := k8sClient.Pods(selfNamespace)
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	log.Println("pod size: ", len(podList.Items))
	controller.Ctx.Output.SetStatus(200)
	controller.Data["json"] = podList.Items[0].Name
	controller.ServeJSON()
	return
}

// @route /inclusterothernamespace [get]
func (controller *RBACController) RBACInClusterOtherNamespace() {
	k8sClient, err := k8s.NewNativeK8sClientConfigByPath(kubeconfigPath)
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	podList, err := k8sClient.Pods(otherNamespace)
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	log.Println("pod size: ", len(podList.Items))
	controller.Ctx.Output.SetStatus(200)
	controller.Data["json"] = podList.Items[0].Name
	controller.ServeJSON()
	return
}


// @route /inclusternode [get]
func (controller *RBACController) RBACInClusterNode() {
	k8sClient, err := k8s.NewNativeK8sClientConfigByPath(kubeconfigPath)
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	nodeList, err := k8sClient.Nodes()
	if err != nil {
		log.Println("err: ", err)
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	log.Println("Node size: ", len(nodeList.Items))
	controller.Ctx.Output.SetStatus(200)
	controller.Data["json"] = nodeList.Items[0].Name
	controller.ServeJSON()
	return
}
