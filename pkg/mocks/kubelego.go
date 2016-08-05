package mocks

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/golang/mock/gomock"
	"k8s.io/kubernetes/pkg/util/intstr"
)

func DummyKubeLego(c *gomock.Controller) *MockKubeLego {
	logrus.SetLevel(logrus.DebugLevel)
	log := logrus.WithField("context", "test-mock")

	kl := NewMockKubeLego(c)
	kl.EXPECT().Log().AnyTimes().Return(log)
	kl.EXPECT().LegoHTTPPort().AnyTimes().Return(intstr.FromInt(8080))
	kl.EXPECT().LegoServiceName().AnyTimes().Return("kube-lego")
	kl.EXPECT().LegoServiceNameGce().AnyTimes().Return("kube-lego-gce")
	kl.EXPECT().Log().AnyTimes().Return(log)
	kl.EXPECT().Version().AnyTimes().Return("mocked-version")
	kl.EXPECT().AcmeUser().AnyTimes().Return(nil, errors.New("I am only mocked"))
	kl.EXPECT().LegoURL().AnyTimes().Return("https://acme-staging.api.letsencrypt.org/directory")
	kl.EXPECT().LegoEmail().AnyTimes().Return("kube-lego-e2e@example.com")

	return kl
}
