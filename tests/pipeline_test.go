package test

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformInfrastructure(t *testing.T) {
	t.Log("🔧 Initializing and applying Terraform...")
	init := exec.Command("terraform", "init")
	apply := exec.Command("terraform", "apply", "-auto-approve")

	initOutput, err := init.CombinedOutput()
	require.NoError(t, err, string(initOutput))

	applyOutput, err := apply.CombinedOutput()
	require.NoError(t, err, string(applyOutput))

	t.Log("✅ Terraform applied successfully")
}

func TestKubernetesDeployment(t *testing.T) {
	t.Log("🚀 Verifying Kubernetes deployment...")

	options := k8s.NewKubectlOptions("", "", "default")

	// Wait for deployment to be available
	k8s.WaitUntilDeploymentAvailable(t, options, "chat-app", 10, 10*time.Second)

	t.Log("✅ Deployment is available")

	// Optionally, test the service
	service := k8s.GetService(t, options, "chat-service")
	url := fmt.Sprintf("http://%s", service.Spec.ClusterIP)

	t.Logf("🌐 Testing service endpoint: %s", url)

	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode, "Service should return HTTP 200")
}

func TestFrontendStaticCheck(t *testing.T) {
	t.Log("🔍 Running ESLint...")

	cmd := exec.Command("npx", "eslint", "src/")
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("❌ ESLint errors:\n%s", string(out))
	} else {
		t.Log("✅ ESLint passed with no errors")
	}
}

func TestDockerfileBuild(t *testing.T) {
	t.Log("🐳 Testing Dockerfile build...")

	cmd := exec.Command("docker", "build", "-t", "chat-app-test", ".")
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("❌ Docker build failed:\n%s", string(out))
	} else {
		t.Log("✅ Docker image built successfully")
	}
}
