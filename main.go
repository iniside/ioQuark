package main

import (
	"fmt"

	"encoding/json"
	//"github.com/gorilla/mux"
	//"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
)

type PodData struct {
	PodName string
}

type PodsData []PodData

func main() {

	//router := mux.NewRouter().StrictSlash(false)
	//router.HandleFunc("/", Index)
	// for {

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// 	// Examples for error handling:
	// 	// - Use helper functions like e.g. errors.IsNotFound()
	// 	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// 	_, err = clientSet.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
	// 	if errors.IsNotFound(err) {
	// 		fmt.Printf("Pod not found\n")
	// 	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
	// 	} else if err != nil {
	// 		panic(err.Error())
	// 	} else {
	// 		fmt.Printf("Found pod\n")
	// 	}

	// 	time.Sleep(10 * time.Second)
	// }

	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (pd *PodData) SetName(name string) {
	pd.PodName = name
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientSet.CoreV1().Pods("kube-system").List(v1.ListOptions{})
	podJson := make(PodsData, len(pods.Items))

	for _, element := range pods.Items {
		pt := PodData{}
		pt.SetName(element.GetName())
		podJson = append(podJson, pt)
	}

	json.NewEncoder(w).Encode(podJson)
}
