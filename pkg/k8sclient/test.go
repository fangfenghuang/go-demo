package k8sclient

import (
	"gopkg.in/ffmt.v1"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func Test() *v1.ControllerRevision {
	crs, err := K8sClientSet.AppsV1().ControllerRevisions("hff").Get("hff-sts-58745bcccf", metav1.GetOptions{})
	if err != nil {
		klog.Errorln(err)
	} else {
		ffmt.Puts(crs)
	}

	return crs
}
