package stub

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha.Poppins:
		logrus.Infof("event {}", o)
		//err := sdk.Create(newbusyBoxPod(o))
		//if err != nil && !errors.IsAlreadyExists(err) {
		//	logrus.Errorf("failed to create busybox pod : %v", err)
		//	return err
		//}
	//case *v1alpha.PoppinsList:
		//
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha.Poppins) *corev1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box13",
			//Name:      cr.Labels["name"],
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha.SchemeGroupVersion.Group,
					Version: v1alpha.SchemeGroupVersion.Version,
					Kind:    "Poppins",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "docker.io/busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
