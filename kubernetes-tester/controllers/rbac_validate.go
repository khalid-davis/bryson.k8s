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

func init() {
	selfNamespace = "k8s-tester"
	otherNamespace = "greenstock"
}

// @router /inclusterselfnamespace [get]
func (controller *RBACController) RBACInClusterSelfNamespace() {
	k8sClient, err := k8s.NewK8sClient()
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
	controller.Data["json"] = podList.Items
	controller.ServeJSON()
	return
}

// @route /inclusterothernamespace [get]
func (controller *RBACController) RBACInClusterOtherNamespace() {
	k8sClient, err := k8s.NewK8sClient()
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
	controller.Data["json"] = podList.Items
	controller.ServeJSON()
	return
}
