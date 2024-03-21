package kubernetes

import (
	"context"
	"fmt"
	"path/filepath"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset *kubernetes.Clientset

func InitClient() {
	config, err := rest.InClusterConfig()
	if err != nil {
		zap.L().Info("In-cluster configuration failed, trying out of cluster")
		config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
		if err != nil {
			zap.L().Sugar().Panicf("Kubernetes client configuration failed %s", err)
		}
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		zap.L().Sugar().Panicf("Kubernetes client initialization failed %s", err)
	}
	version, err := clientset.ServerVersion()
	if err != nil {
		zap.L().Sugar().Panicf("Kubernetes client initialization failed while retrieving cluster version %s", err)
	}
	zap.L().Sugar().Infof("Kubernetes client initialized, %s", version.String())
}

func ListPods() {
	pods, err := clientset.CoreV1().Pods("user-f2wbnp").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("Failed retrieving pods %s", err)
	} else {
		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
	}
}
