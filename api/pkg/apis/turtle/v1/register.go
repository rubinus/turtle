package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "turtle"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = SchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addKnownTypes, addConversionFuncs)
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&DataCenter{},
		&DataCenterList{},
		&Cluster{},
		&ClusterList{},
		&Node{},
		&NodeList{},
		&Namespace{},
		&NamespaceList{},
		&Deployment{},
		&DeploymentList{},
		&StatefulSet{},
		&StatefulSetList{},
		&PVC{},
		&PVCList{},
		&PV{},
		&PVList{},
		&Event{},
		&EventList{},
		&Pod{},
		&PodList{},
		&ReplicaSet{},
		&ReplicaSetList{},
		&Service{},
		&ServiceList{},
		&Job{},
		&JobList{},
		&CronJob{},
		&CronJobList{},
		&DaemonSet{},
		&DaemonSetList{},
		&HPA{},
		&HPAList{},
		&Ingress{},
		&IngressList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func addConversionFuncs(scheme *runtime.Scheme) (err error) {
	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("DataCenter"),
		dcScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Cluster"),
		clusterScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Zone"),
		zoneScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Deployment"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("StatefulSet"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("DaemonSet"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Job"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("CronJob"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("ReplicaSet"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Namespace"),
		clusterScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Node"),
		nodeFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("PV"),
		clusterScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("PVC"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Service"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Ingress"),
		namespaceScopedFieldLabelConversionFunc)
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Pod"),
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.name",
				"metadata.dcID",
				"metadata.zoneID",
				"metadata.clusterID",
				"metadata.objName",
				"spec.nodeName",
				"metadata.objNamespace",
				"status.healthy":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		})
	if err != nil {
		return
	}

	err = scheme.AddFieldLabelConversionFunc(SchemeGroupVersion.WithKind("Event"),
		func(label, value string) (string, string, error) {
			switch label {
			case "metadata.name",
				"metadata.dcID",
				"metadata.zoneID",
				"metadata.objName",
				"metadata.clusterID",
				"spec.involvedObject.kind",
				"spec.involvedObject.name",
				"spec.involvedObject.namespace":
				return label, value, nil
			default:
				return "", "", fmt.Errorf("field label not supported: %s", label)
			}
		})
	if err != nil {
		return
	}

	return
}

var (
	namespaceScopedFieldLabelConversionFunc = func(label, value string) (string, string, error) {
		switch label {
		case "metadata.name",
			"metadata.dcID",
			"metadata.zoneID",
			"metadata.clusterID",
			"metadata.objName",
			"metadata.objNamespace",
			"status.stateless.healthy":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label not supported: %s", label)
		}
	}

	clusterScopedFieldLabelConversionFunc = func(label, value string) (string, string, error) {
		switch label {
		case "metadata.name",
			"metadata.dcID",
			"metadata.zoneID",
			"metadata.objName",
			"metadata.clusterID":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label not supported: %s", label)
		}
	}

	dcScopedFieldLabelConversionFunc = func(label, value string) (string, string, error) {
		switch label {
		case "metadata.name",
			"metadata.objName",
			"metadata.dcID":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label not supported: %s", label)
		}
	}

	zoneScopedFieldLabelConversionFunc = func(label, value string) (string, string, error) {
		switch label {
		case "metadata.name",
			"metadata.objName",
			"metadata.zoneID",
			"metadata.dcID":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label not supported: %s", label)
		}
	}

	nodeFieldLabelConversionFunc = func(label, value string) (string, string, error) {
		switch label {
		case "metadata.name",
			"metadata.dcID",
			"metadata.zoneID",
			"metadata.objName",
			"metadata.clusterID",
			"spec.unschedulable":
			return label, value, nil
		default:
			return "", "", fmt.Errorf("field label not supported: %s", label)
		}
	}
)
