package kubernetes

import (
	"context"
	"fmt"
	"path/filepath"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
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

func ListPods(namespace string) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("Failed retrieving pods %s", err)
	} else {
		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
	}
}

func GetEvents(namespace string) watch.Interface {
	events, _ := clientset.EventsV1().Events(namespace).Watch(context.TODO(), metav1.ListOptions{})
	return events
}

func InitNamespace(namespace string) {
	namespaceToCreate := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   namespace,
			Labels: map[string]string{"onyxia_owner": "todo"},
		},
	}
	clientset.CoreV1().Namespaces().Create(context.TODO(), &namespaceToCreate, metav1.CreateOptions{})
}

func GetOnyxiaResourceQuota(namespace string) (*corev1.ResourceQuota) {

	resourceQuotas, _ := clientset.CoreV1().ResourceQuotas(namespace).Get(context.TODO(), "onyxia-quota", metav1.GetOptions{})
	return resourceQuotas

}

func PostOnyxiaResourceQuotas(namespace string, quota (*corev1.ResourceQuota)) {

	clientset.CoreV1().ResourceQuotas(namespace).Update(context.TODO(), quota, metav1.UpdateOptions{})
}
