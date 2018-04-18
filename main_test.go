package main

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"os"

	"io/ioutil"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func Test_Healthz(t *testing.T) {
	kubeclient := fake.NewSimpleClientset()
	router := routerSetup(kubeclient)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func Test_Count(t *testing.T) {
	kubeclient := fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind: "pod",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
		},
	},
	)
	router := routerSetup(kubeclient)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"count":1}`, w.Body.String())
}

func Test_currentNamespace(t *testing.T) {
	assert.Equal(t, "default", currentNamespace())
	os.Setenv("GIN_MODE", "release")
	if err := os.MkdirAll("/var/run/secrets/kubernetes.io/serviceaccount/", 0644); err != nil {
		log.Fatal().Err(err)
	}
	ioutil.WriteFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace", []byte("test"), 0644)
	assert.Equal(t, "test", currentNamespace())
}

func Test_homeDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "home dir", want: "/root"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := homeDir(); got != tt.want {
				t.Errorf("homeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
