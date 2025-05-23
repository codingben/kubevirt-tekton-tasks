package testoptions

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/util/homedir"
)

var deployNamespace string
var storageClass string
var kubeConfigPath string
var debug string
var skipCreateVMFromManifestTests string
var skipExecuteInVMTests string
var skipGenerateSSHKeysTests string
var skipDiskUploaderTests string

type TestOptions struct {
	DeployNamespace               string
	StorageClass                  string
	KubeConfigPath                string
	Debug                         bool
	SkipCreateVMFromManifestTests bool
	SkipExecuteInVMTests          bool
	SkipGenerateSSHKeysTests      bool
	SkipDiskUploaderTests         bool

	CommonTemplatesVersion string
}

func init() {
	flag.StringVar(&deployNamespace, "deploy-namespace", "", "Namespace where to deploy the tasks and taskrun")
	flag.StringVar(&storageClass, "storage-class", "", "Storage class to be used for creating test DVs/PVCs")
	flag.StringVar(&kubeConfigPath, "kubeconfig-path", "", "Path to the kubeconfig")
	flag.StringVar(&debug, "debug", "", "Debug keeps all the resources alive after the tests complete. One of: true|false")
	flag.StringVar(&skipCreateVMFromManifestTests, "skip-create-vm-from-manifests-tests", "", "Skip create vm from manifests test suite. One of: true|false")
	flag.StringVar(&skipExecuteInVMTests, "skip-execute-in-vm-tests", "", "Skip execute in vm test suite. One of: true|false")
	flag.StringVar(&skipGenerateSSHKeysTests, "skip-generate-ssh-keys-tests", "", "Skip generate ssh keys suite. One of: true|false")
	flag.StringVar(&skipDiskUploaderTests, "skip-disk-uploader-tests", "", "Skip disk uploader suite. One of: true|false")
}

func InitTestOptions(testOptions *TestOptions) error {
	flag.Parse()

	if deployNamespace == "" {
		return errors.New("--deploy-namespace must be specified")
	}

	if kubeConfigPath != "" {
		testOptions.KubeConfigPath = kubeConfigPath
	} else {
		kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
		if file, err := os.Stat(kubeConfigPath); err == nil && file.Mode().IsRegular() {
			testOptions.KubeConfigPath = kubeConfigPath
		}
	}

	testOptions.DeployNamespace = deployNamespace
	testOptions.StorageClass = storageClass
	testOptions.Debug = strings.ToLower(debug) == "true"

	testOptions.SkipCreateVMFromManifestTests = strings.ToLower(skipCreateVMFromManifestTests) == "true"
	testOptions.SkipExecuteInVMTests = strings.ToLower(skipExecuteInVMTests) == "true"
	testOptions.SkipGenerateSSHKeysTests = strings.ToLower(skipGenerateSSHKeysTests) == "true"
	testOptions.SkipDiskUploaderTests = strings.ToLower(skipDiskUploaderTests) == "true"

	return nil
}

func (f *TestOptions) GetDeployNamespace() string {
	return f.DeployNamespace
}
