package util

import (
	"regexp"
	"strconv"

	apiv1 "github.com/openshift/api/config/v1"
	clientv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	managedv1alpha1 "github.com/openshift/rbac-permissions-operator/pkg/apis/managed/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	log      = logf.Log.WithName("safelist")
	daLogger = log.WithValues("SafeList", "functions")
)

// ClusterOperatorClient client for managing the ClusterOperator CR
type ClusterOperatorClient struct {
	client              clientv1.ClusterOperatorInterface
	clusterOperatorName string
	namespace           string
	version             string
}

// PopulateCrPermissionClusterRoleNames to see if clusterRoleName exists in permission
// returns list of ClusterRoleNames in permissions that do not exist
func PopulateCrPermissionClusterRoleNames(subjectPermission *managedv1alpha1.SubjectPermission, clusterRoleList *v1.ClusterRoleList) []string {
	//permission ClusterRoleName
	permissions := subjectPermission.Spec.Permissions

	var permissionClusterRoleNames []string
	var found bool

	for _, i := range permissions {
		for _, a := range clusterRoleList.Items {
			if i.ClusterRoleName == a.Name {
				found = true
			}
		}
		if !found {
			permissionClusterRoleNames = append(permissionClusterRoleNames, i.ClusterRoleName)
		}
	}

	// create a map of all unique elements
	encountered := map[string]bool{}
	for v := range permissionClusterRoleNames {
		encountered[permissionClusterRoleNames[v]] = true
	}

	// place all keys from map into slice
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}

	return result
}

// GenerateSafeList by 1st checking allow regex then check denied regex
func GenerateSafeList(allowedRegex string, deniedRegex string, nsList *corev1.NamespaceList) []string {
	safeList := allowedNamespacesList(allowedRegex, nsList)

	updatedSafeList := safeListAfterDeniedRegex(deniedRegex, safeList)

	return updatedSafeList

}

// allowedNamespacesList 1st pass - allowedRegex
func allowedNamespacesList(allowedRegex string, nsList *corev1.NamespaceList) []string {
	var matches []string

	// for every namespace on the cluster
	// check that against the allowedRegex in Permission
	for _, namespace := range nsList.Items {
		rp := regexp.MustCompile(allowedRegex)

		// if namespace on cluster matches with regex, append them to slice
		found := rp.MatchString(namespace.Name)
		if found {
			matches = append(matches, namespace.Name)
		}
	}

	return matches
}

// safeListAfterDeniedRegex 2nd pass - deniedRegex
func safeListAfterDeniedRegex(namespacesDeniedRegex string, safeList []string) []string {
	if namespacesDeniedRegex == "" {
		return safeList
	}
	var updatedSafeList []string

	// for every namespace on SafeList
	// check that against deniedRegex
	for _, namespace := range safeList {
		rp := regexp.MustCompile(namespacesDeniedRegex)

		found := rp.MatchString(namespace)
		// if it does not match then append
		if !found {
			updatedSafeList = append(updatedSafeList, namespace)
		}
	}
	return updatedSafeList

}

// NewRoleBindingForClusterRole creates and returns valid RoleBinding
func NewRoleBindingForClusterRole(clusterRoleName, subjectName, subjectKind, namespace string) *v1.RoleBinding {
	return &v1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterRoleName + "-" + subjectName,
			Namespace: namespace,
		},
		Subjects: []v1.Subject{
			{
				Kind: subjectKind,
				Name: subjectName,
			},
		},
		RoleRef: v1.RoleRef{
			Kind: "ClusterRole",
			Name: clusterRoleName,
		},
	}
}

// UpdateCondition of SubjectPermission
func UpdateCondition(conditions []managedv1alpha1.Condition, message string, clusterRoleNames []string, status bool, state managedv1alpha1.SubjectPermissionState, conditionType managedv1alpha1.SubjectPermissionType) []managedv1alpha1.Condition {
	now := metav1.Now()

	existingCondition := FindRbacCondition(conditions, conditionType)

	// create a map of all unique elements in clusterRoleNames slice
	encountered := map[string]bool{}
	for v := range clusterRoleNames {
		encountered[clusterRoleNames[v]] = true
	}

	// place all keys from map into result slice
	// this prevents the duplication of clusterRoleNames
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}

	if existingCondition == nil {
		conditions = append(
			conditions, managedv1alpha1.Condition{
				LastTransitionTime: now,
				ClusterRoleNames:   result,
				Message:            message,
				Status:             status,
				State:              state,
				Type:               conditionType,
			},
		)
	} else {
		if existingCondition.Status != status {
			existingCondition.LastTransitionTime = now
		}
		existingCondition.Message = message
		existingCondition.ClusterRoleNames = result
		existingCondition.Status = status
		existingCondition.State = state
	}

	return conditions
}

// FindRbacCondition finds in the condition that has the specified condition type in the given list
// if none exists, then returns nil
func FindRbacCondition(conditions []managedv1alpha1.Condition, conditionType managedv1alpha1.SubjectPermissionType) *managedv1alpha1.Condition {
	for i, condition := range conditions {
		if condition.Type == conditionType {
			return &conditions[i]
		}
	}
	return nil
}

// RoleBindingExists checks if a rolebinding exists in the cluster already
func RoleBindingExists(roleBinding *v1.RoleBinding, rbList *v1.RoleBindingList) bool {
	for _, rb := range rbList.Items {
		if roleBinding.Name == rb.Name {
			return true
		}
	}
	return false
}

// NewClusterOperatorClient returns a status reporter instance
func NewClusterOperatorClient(client clientv1.ClusterOperatorInterface, name, namespace, version string) *ClusterOperatorClient {
	return &ClusterOperatorClient{
		client:              client,
		clusterOperatorName: name,
		namespace:           namespace,
		version:             version,
	}
}

// newClusterOperator will create the cluster operator cr
func newClusterOperator(clusterOperatorName string, clusterOperatorNamespace string) *apiv1.ClusterOperator {
	co := &apiv1.ClusterOperator{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "config.openshift.io/v1",
			Kind:       "ClusterOperator",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterOperatorName,
			Namespace: clusterOperatorNamespace,
		},
		Spec:   apiv1.ClusterOperatorSpec{},
		Status: apiv1.ClusterOperatorStatus{},
	}
	return co
}

// GetOrCreateClusterOperator will return the cluster operator instance if it exists, error otherwise
func (c *ClusterOperatorClient) GetOrCreateClusterOperator() (*apiv1.ClusterOperator, error) {
	// Search for the clusterOperator object in this namespace
	instance, err := c.client.Get(c.clusterOperatorName, metav1.GetOptions{})
	if err != nil {
		// ClusterOperator CR does not exists, create it
		instance, err = c.client.Create(newClusterOperator(c.clusterOperatorName, c.namespace))
		if err != nil {
			return instance, err
		}
		return instance, managedv1alpha1.ErrReconcileRequeue
	}
	return instance, nil
}

// ClusterOperatorUpdateConditions will update the conditions of the cluster operator
func (c *ClusterOperatorClient) ClusterOperatorUpdateConditions(condition managedv1alpha1.Condition) error {
	now := metav1.Now()

	// Get clusterOperator cr
	co, err := c.GetOrCreateClusterOperator()
	if err != nil {
		// Failed to retrieve operator, returning error
		return err
	}

	// Update clusterOperatorConditions object
	ocConditions, err := castOperatorCondition(condition, now)
	co.Status.Conditions = append(co.Status.Conditions, *ocConditions)
	return nil
}

// castOperatorCondition will cast subjectPermission conditions into a clusterOperator condition
func castOperatorCondition(condition managedv1alpha1.Condition, time metav1.Time) (*apiv1.ClusterOperatorStatusCondition, error) {
	// Cast condition into
	co := &apiv1.ClusterOperatorStatusCondition{
		Status:  castClusterOperatorConditionStatus(condition.Status),
		Message: condition.Message,
	}
	return co, nil
}

func castClusterOperatorConditionStatus(status bool) apiv1.ConditionStatus {
	str := strconv.FormatBool(status)
	return apiv1.ConditionStatus(str)
}
