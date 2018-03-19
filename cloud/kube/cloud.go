package kube

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/Stymphalian/ikuaki/cloud"

	"github.com/google/uuid"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubeCloud struct {
	config *rest.Config
	client *kubernetes.Clientset
}

type KubeAddr struct {
	id       string
	hostport string
}

func (this *KubeAddr) Id() string       { return this.id }
func (this *KubeAddr) Hostport() string { return this.hostport }

// UTILS
// -----------------------------------------------------------------------------
func int32Ptr(i int32) *int32 { return &i }

// Given a POD extract out the ID and hostport as a cloud.Addr object
func GetAddrFromPod(pod *apiv1.Pod) cloud.Addr {
	return &KubeAddr{pod.GetObjectMeta().GetName(),
		fmt.Sprintf("%s:%d", pod.Status.PodIP,
			pod.Spec.Containers[0].Ports[0].ContainerPort)}
}

// KUBE CLOUD
// -----------------------------------------------------------------------------
func NewKubeCloud() (*KubeCloud, error) {
	// Get the config either from the home-directory, or
	// try to get in because we are running in a cluster.
	var config *rest.Config
	var err error
	if home := homedir.HomeDir(); home != "" {
		kubeConfigFilePath := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)
		if err != nil {
			return nil, err
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	// Create a client to the kube cluster
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeCloud{config, clientset}, nil
}

// Internal helper function for starting a POD on kube
// Will wait until the pod is completely ready and will return to you the
// POD object
func (this *KubeCloud) startPod(image string, args map[string]string) (*apiv1.Pod, error) {
	podClient := this.client.Core().Pods(apiv1.NamespaceDefault)

	// Create the spec
	data, err := ioutil.ReadFile(image)
	if err != nil {
		return nil, err
	}
	spec := &apiv1.Pod{}
	_, _, err = scheme.Codecs.UniversalDeserializer().Decode(data, nil, spec)
	if err != nil {
		return nil, err
	}

	// Modify the name so that it is unique
	newName := spec.ObjectMeta.Name + "-" + uuid.New().String()
	spec.ObjectMeta.Name = newName
	spec.Spec.Containers[0].Name = newName
	for k, v := range args {
		spec.Spec.Containers[0].Args = append(
			spec.Spec.Containers[0].Args, fmt.Sprintf("--%s=%s", k, v))
	}

	// Create the pod
	_, err = podClient.Create(spec)
	if err != nil {
		return nil, err
	}

	// Wait until the pod is actually up and assigned an IP
	watcher, err := podClient.Watch(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for c := range watcher.ResultChan() {
		if c.Type == watch.Modified {
			if c.Object.(*apiv1.Pod).Status.Phase == apiv1.PodRunning {
				watcher.Stop()
			}
		}
	}

	// Get the pod
	return podClient.Get(newName, metav1.GetOptions{})
}

// Start the given image with the args and return a cloud.Addr
func (this *KubeCloud) Start(image string, args map[string]string) (cloud.Addr, error) {
	result, err := this.startPod(image, args)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created Pod %q.\n", result.GetObjectMeta().GetName())

	return &KubeAddr{result.GetObjectMeta().GetName(),
		fmt.Sprintf("%s:%d", result.Status.PodIP,
			result.Spec.Containers[0].Ports[0].ContainerPort)}, nil
}

// Stop the server with the specified ID
func (this *KubeCloud) Stop(id string) error {
	podClient := this.client.Core().Pods(apiv1.NamespaceDefault)
	return podClient.Delete(id, &metav1.DeleteOptions{})
}

// Lists all the pods currently running
func (this *KubeCloud) List() ([]cloud.Addr, error) {
	podClient := this.client.Core().Pods(apiv1.NamespaceDefault)
	podList, err := podClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	addrs := make([]cloud.Addr, 0)
	for _, c := range podList.Items {
		addrs = append(addrs, GetAddrFromPod(&c))
	}
	return addrs, nil
}
