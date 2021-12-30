package cmd

import (
	"os"

	"go-demo/pkg/calico"
	"go-demo/pkg/k8sclient"
	"go-demo/router"

	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		klog.Errorln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hff example",
	Short: "only for test...",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		klog.V(3).Info("service start...")
		initializaton()
		test()
		router.InitRouter()
	},
}

// add init func
func initializaton() {
	k8sclient.InitClientSet()
	calico.InitGlobalNetworkPolicy()
}

// add test func
func test() {
	//k8sclient.Test()
}
