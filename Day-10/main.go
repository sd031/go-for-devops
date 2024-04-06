package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var namespace string

	var rootCmd = &cobra.Command{
		Use:   "kubectl-myplugin",
		Short: "List pods in a namespace",
		Run: func(cmd *cobra.Command, args []string) {
			configFlags := genericclioptions.NewConfigFlags(false)

			config, err := configFlags.ToRESTConfig()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error getting Kubernetes config: %v\n", err)
				os.Exit(1)
			}

			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating Kubernetes client: %v\n", err)
				os.Exit(1)
			}

			pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				fmt.Fprintf(os.Stderr, "error listing pods: %v\n", err)
				os.Exit(1)
			}

			for _, pod := range pods.Items {
				fmt.Println(pod.Name)
			}
		},
	}

	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "Namespace to list pods in")
	rootCmd.Execute()
}
