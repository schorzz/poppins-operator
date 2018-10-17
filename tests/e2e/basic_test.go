package e2e
import (
	"github.com/operator-framework/operator-sdk/pkg/test/e2eutil"
	"testing"
	"time"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
)

func TestCreatePoppins(t *testing.T) {
	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup()

	err := ctx.InitializeClusterResources(&framework.CleanupOptions{TestContext: ctx, Timeout: time.Minute, RetryInterval: time.Second*5})
	if err != nil {
		t.Fatalf("failed to initialize cluster resources: %v", err)
	}
	namespace, err := ctx.GetNamespace()
	if err != nil {
		t.Fatal(err)
	}
	// get global framework variables
	f := framework.Global
	// wait for memcached-operator to be ready
	err = e2eutil.WaitForDeployment(t, f.KubeClient, namespace, "poppins-operator", 1, time.Second*5, time.Second*30)
	if err != nil {
		t.Fatal(err)
	}
}