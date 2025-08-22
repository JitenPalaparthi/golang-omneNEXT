// Package runtimeid provides a single function to discover identity when running
// inside Kubernetes (or any container).
//
// go get k8s.io/client-go@v0.31.0
package runtimeid

import (
	"bufio"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Identity struct {
	// Basic (no API needed)
	Hostname          string // usually pod name in Kubernetes
	PodName           string // from env/hostname
	Namespace         string // from env/SA namespace file
	PodUID            string // from env (Downward API)
	PodIP             string // from env (Downward API)
	NodeName          string // from env (Downward API)
	ContainerIDCGroup string // parsed from /proc/self/cgroup

	// From Kubernetes API (if queried)
	PodUIDFromAPI     string
	PodIPFromAPI      string
	NodeFromAPI       string
	ContainerIDs      []string // e.g. "containerd://<id>"
	SelfContainerName string   // best-effort match of current process container

	// Non-fatal notes if something couldn't be fetched
	Notes []string
}

// Detect gathers identity. If queryK8sAPI is true, it will try in-cluster
// client-go to enrich details; otherwise it only uses env/hostname/cgroups.
func Detect(ctx context.Context, queryK8sAPI bool) (Identity, error) {
	var id Identity

	// ----- Basics (env/hostname/files) -----
	if hn, _ := os.Hostname(); hn != "" {
		id.Hostname = hn
	}
	id.PodName = firstNonEmpty(os.Getenv("POD_NAME"), id.Hostname)
	id.Namespace = firstNonEmpty(os.Getenv("POD_NAMESPACE"), readFileTrim("/var/run/secrets/kubernetes.io/serviceaccount/namespace"))
	id.PodUID = os.Getenv("POD_UID")
	id.PodIP = os.Getenv("POD_IP")
	id.NodeName = os.Getenv("NODE_NAME")
	id.ContainerIDCGroup = detectContainerIDFromCgroup()

	// If not querying API, we're done.
	if !queryK8sAPI {
		return id, nil
	}

	// ----- Kubernetes API (best-effort) -----
	cfg, err := rest.InClusterConfig()
	if err != nil {
		id.Notes = append(id.Notes, "in-cluster config not available: "+err.Error())
		return id, nil // keep basics; not a fatal error
	}
	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		id.Notes = append(id.Notes, "client-go init failed: "+err.Error())
		return id, nil
	}

	if id.Namespace == "" || id.PodName == "" {
		id.Notes = append(id.Notes, "missing POD_NAMESPACE/POD_NAME (add Downward API envs)")
		return id, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	pod, err := cs.CoreV1().Pods(id.Namespace).Get(ctx, id.PodName, metav1.GetOptions{})
	if err != nil {
		id.Notes = append(id.Notes, "pods.get failed: "+err.Error())
		return id, nil
	}

	id.PodUIDFromAPI = string(pod.UID)
	id.PodIPFromAPI = pod.Status.PodIP
	id.NodeFromAPI = pod.Spec.NodeName

	// Collect runtime container IDs and try to match this process
	var selfSuffix = id.ContainerIDCGroup
	for _, st := range pod.Status.ContainerStatuses {
		id.ContainerIDs = append(id.ContainerIDs, st.ContainerID)
		if selfSuffix != "" && strings.HasSuffix(st.ContainerID, selfSuffix) {
			id.SelfContainerName = st.Name
		}
	}

	return id, nil
}

// ----- helpers -----

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func readFileTrim(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

// Parses /proc/self/cgroup for a 64-hex container ID (works across docker/containerd/cri-o/podman).
func detectContainerIDFromCgroup() string {
	f, err := os.Open("/proc/self/cgroup")
	if err != nil {
		return ""
	}
	defer f.Close()

	re64 := regexp.MustCompile(`[a-f0-9]{64}`)
	reScoped := regexp.MustCompile(`(?:docker|containerd|cri-containerd|crio|libpod|podman)[-:/]([a-f0-9]{64})(?:\.scope)?`)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if m := reScoped.FindStringSubmatch(line); len(m) == 2 {
			return m[1]
		}
		if m := re64.FindString(line); m != "" {
			return m
		}
	}
	return ""
}

// Optional convenience: require K8s envs present.
func RequireDownwardAPIEnv() error {
	if os.Getenv("POD_NAME") == "" || os.Getenv("POD_NAMESPACE") == "" {
		return errors.New("POD_NAME/POD_NAMESPACE not set; add Downward API envs to your Pod spec")
	}
	return nil
}
