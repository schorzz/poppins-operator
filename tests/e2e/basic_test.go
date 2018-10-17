package e2e

import (
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
	goctx "context"
)

func TestCreatePoppins(t *testing.T) {
	poppinsList := v1alpha.PoppinsList{
		TypeMeta: metav1.TypeMeta{
			Kind:		"Poppins",
			APIVersion: "schorzz.poppins.com/v1alpha",
		},
	}

	err := framework.AddToFrameworkScheme(v1alpha.AddToScheme,&poppinsList)
	if err != nil {
		t.Fatalf("failed to add custom resource scheme to framework: %v", err)
	}

	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()

	err = ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: time.Minute, RetryInterval: time.Second*5})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatal(err)
	}
	// get global framework variables
	f := framework.Global
	// wait for poppins-operator to be ready
	err = e2eutil.WaitForDeployment(t, f.KubeClient, namespace, "poppins-operator", 1, time.Second*5, time.Second*30)
	if err != nil {
		t.Fatal(err)
	}

	poppins := v1alpha.Poppins{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion: "schorzz.poppins.com/v1alpha",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
			Namespace: "poppins",
		},
		Spec: v1alpha.PoppinsSpec{
			ExpireDate: time.Now(),
		},
	}
	err = f.Client.Create(goctx.TODO(), &poppins, &framework.CleanupOptions{TestContext: ctx, Timeout: time.Second*30, RetryInterval: time.Second*5})
	if err != nil {
		return
	}
}