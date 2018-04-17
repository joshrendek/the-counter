package main

import (
	"flag"
	"os"
	"path/filepath"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeclient := NewKubeClient()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		pods, _ := kubeclient.CoreV1().Pods(currentNamespace()).List(metav1.ListOptions{})
		c.JSON(200, gin.H{
			"count": len(pods.Items),
		})
	})

	r.Run()
}

func currentNamespace() string {
	data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading service account namespace")
	}
	return string(data)
}

func NewKubeClient() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	if os.Getenv("PRODUCTION") != "" {
		*kubeconfig = ""
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal().Err(err)
	}

	// we're running a cluster now
	if config == nil {
		// final try, use in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatal().Err(err)
		}
		return clientset
	}
	// developing locally here
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Info().Err(err)
	}

	return clientset
}
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
