package controllers

import (
	fluentbitv1alpha2 "github.com/SeBBBe/fluent-operator/v3/apis/fluentbit/v1alpha2"
	fluentdv1alpha1 "github.com/SeBBBe/fluent-operator/v3/apis/fluentd/v1alpha1"
)

var (
	fluentbitOwnerKey = ".fluentbit.metadata.controller"
	collectorOwnerKey = ".collector.metadata.controller"
	fluentdOwnerKey   = ".fluentd.metadata.controller"
	fluentbitApiGVStr = fluentbitv1alpha2.SchemeGroupVersion.String()
	fluentdApiGVStr   = fluentdv1alpha1.SchemeGroupVersion.String()
	fluentdAgentMode  = "agent"
)
