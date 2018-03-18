package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	// othernat "github.com/docker/docker/vendor/github.com/docker/go-connections/nat"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func docker() {
	ctx := context.Background()
	// client.NewClient
	// client.New
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	// cli, err := client.NewClientWithOpts(client.WithVersion("1.35"))
	// if err != nil {
	// 	panic(err)
	// }

	imageName := "stymphalian/ikuaki-lobby:latest"

	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image: imageName,
			// ExposedPorts: nat.PortSet{
			// 	nat.Port(fmt.Sprintf("%d", freeport.GetPortOrDie())): {},
			// },
		},
		nil,
		// &container.HostConfig{
		// 	Binds: []string{fmt.Sprintf("%d:8083", freeport.GetPortOrDie())},
		// },
		nil,
		// &network.NetworkingConfig{},
		"ikuaki-lobby")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func kube() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod\n")
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	kube()
}
