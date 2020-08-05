package v1alpha1

import "errors"

// ErrReconcileRequeue indicates that we need a reconcile with requeue true
var ErrReconcileRequeue = errors.New("ErrReconcileRequeue")

// ClusterOperatorName is the string for the cluster operator name
var ClusterOperatorName = "co-rbac-perissions-operator"

// ClusterOperatorNamespace is the string for the cluster operator namespace
var ClusterOperatorNamespace = "rbac-perissions-operator"
