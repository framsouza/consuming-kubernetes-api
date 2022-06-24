package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var resources []string

func main() {
	listResources("pods")
	for index := range resources {
		fmt.Println(resources[index])
	}

}

func listResources(rs string) string {
	kubeconfig := flag.String("kubeconfig", "/Users/francismarasouza/.kube/config", "file localtion")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if rs == "pods" {
		pods, _ := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
		for _, pod := range pods.Items {
			resources = append(resources, pod.Name)
		}
	}

	if rs == "nodes" {
		nodes, _ := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		for _, node := range nodes.Items {
			resources = append(resources, node.Name)
		}
	}

	if rs == "deployments" {
		deploy, _ := clientset.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
		for _, deploy := range deploy.Items {
			resources = append(resources, deploy.Name)
		}
	}

	return rs
}
