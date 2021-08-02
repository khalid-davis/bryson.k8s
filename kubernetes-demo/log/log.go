package main

import (
	"flag"
	"k8s.io/klog/v2"
)

func main()  {
	klog.InitFlags(nil)
	flag.Set("v", "4") // 如果没有这个的话，下面的V(4)那句是打印不出来的
	flag.Parse()

	klog.Infof("test info")
	klog.V(4).Infof("test v4 info")
}
