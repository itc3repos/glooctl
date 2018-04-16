package client

import (
	"os"
	"path/filepath"
	"time"

	"log"

	secret "github.com/solo-io/gloo-secret"
	secretcrd "github.com/solo-io/gloo-secret/crd"
	secretfile "github.com/solo-io/gloo-secret/file"
	storage "github.com/solo-io/gloo/pkg/storage"
	"github.com/solo-io/gloo/pkg/storage/crd"
	"github.com/solo-io/gloo/pkg/storage/file"
	"k8s.io/client-go/tools/clientcmd"
)

type StorageOptions struct {
	GlooConfigDir string
	SecretDir     string
	KubeConfig    string
	Namespace     string
	SyncPeriod    int
}

func StorageClient(opts *StorageOptions) (storage.Interface, error) {
	syncPeriod := time.Duration(opts.SyncPeriod) * time.Second
	if opts.GlooConfigDir != "" {
		log.Printf("Using file-based storage for gloo. Gloo must be configured to use file storage with config dir %v", opts.GlooConfigDir)
		return file.NewStorage(opts.GlooConfigDir, syncPeriod)
	}

	kubeConfig := opts.KubeConfig
	if kubeConfig == "" && HomeDir() != "" {
		kubeConfig = filepath.Join(HomeDir(), ".kube", "config")
	}
	kubeClient, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	return crd.NewStorage(kubeClient, opts.Namespace, syncPeriod)
}

func SecretClient(opts *StorageOptions) (secret.SecretInterface, error) {
	secretDir := opts.SecretDir
	if secretDir != "" {
		return secretfile.NewClient(secretDir)
	}

	kubeConfig := opts.KubeConfig
	if kubeConfig == "" && HomeDir() != "" {
		kubeConfig = filepath.Join(HomeDir(), ".kube", "config")
	}
	kubeClient, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	return secretcrd.NewClient(kubeClient, opts.Namespace)
}

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}