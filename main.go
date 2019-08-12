package main

import (
	"flag"
	//"fmt"
	"os"
	"path/filepath"
	"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
    log "github.com/sirupsen/logrus"
)

func main() {
    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{})

    // Output to stdout instead of the default stderr
    log.SetOutput(os.Stdout)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}


    // Loop forever
    for {
        // State Cluster
        log.WithFields(log.Fields{
            "Cluster":   config.Host,
        }).Info()
    	// Loop through all pods
    	for _, v := range pods.Items {
    		if v.Status.Phase == "Pending" {
    			if v.Status.ContainerStatuses == nil {
                    for _, pv := range v.Status.Conditions {
    				    //fmt.Printf("Pending: Namespace: %s, Name: %s, Message: %s, Reason: %s\n", v.ObjectMeta.Namespace, v.ObjectMeta.Name, pv.Message, pv.Reason)
                        log.WithFields(log.Fields{
                            "Status":       "Pending",
                            "Namespace":    v.ObjectMeta.Namespace,
                            "Name":         v.ObjectMeta.Name,
                            "Reason":        pv.Reason,
                        }).Warn(pv.Message)
                    }
    			} else {
    				for _, pv := range v.Status.ContainerStatuses {
    					if pv.State.Terminated != nil {
    						//fmt.Printf("Terminating: Namespace: %s, Name: %s, NodeName: %s\n", v.ObjectMeta.Namespace, v.ObjectMeta.Name, v.Spec.NodeName)
                            log.WithFields(log.Fields{
                                "Status":       "Terminating",
                                "Namespace":    v.ObjectMeta.Namespace,
                                "Name":         v.ObjectMeta.Name,
                                "Node":         v.Spec.NodeName,
                            }).Warn()
    					}
    				}
    			}
    		}
    	}
        // Sleep for 60 seconds, start loop over
        time.Sleep(60 * time.Second)
    }
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
