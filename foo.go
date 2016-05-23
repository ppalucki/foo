package main

import (
	"fmt"

	"github.com/ppalucki/bar/subbar"
	"k8s.io/kubernetes/pkg/api/resource"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {

	println(resource.Scale(1))
	println(client.ConfigMapResourceName)
	fmt.Println(subbar.Subbar)
}
