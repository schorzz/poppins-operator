package e2e
import (
	"testing"
	"time"

	//api "github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreatePoppins(t *testing.T) {
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()

	err := ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: time.Minute, RetryInterval: time.Second*30})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}

	//poppinsList := &api.PoppinsList{
	//	TypeMeta: metav1.TypeMeta{
	//		Kind:       "poppins",
	//		APIVersion: "schorzz.poppins.com/v1alpha1",
	//	},
	//}
	//err := framework.AddToFrameworkScheme(api.AddToScheme, poppinsList)
	//if err != nil {
	//	t.Fatalf("could not add scheme to framework scheme: %v", err)
	//}

	//vaultCR, tlsConfig, rootToken := e2eutil.SetupUnsealedVaultCluster(t, f.KubeClient, f.DynamicClient, ctx.Namespace)
	//ctx.AddFinalizerFn(func() error {
	//	if err := e2eutil.DeleteCluster(t, f.DynamicClient, vaultCR); err != nil {
	//		return fmt.Errorf("failed to delete vault cluster: %v", err)
	//	}
	//	return nil
	//})
	//vClient, keyPath, secretData, podName := e2eutil.WriteSecretData(t, vaultCR, f.KubeClient, tlsConfig, rootToken, ctx.Namespace)
	//e2eutil.VerifySecretData(t, vClient, secretData, keyPath, podName)
}