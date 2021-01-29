package turtle

import (
	appv1 "k8s.io/api/apps/v1"
	autov1 "k8s.io/api/autoscaling/v1"
	jobv1 "k8s.io/api/batch/v1"
	cronjobv1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	ingressv1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectMeta struct {
	metav1.ObjectMeta `json:",inline"`

	// +optional
	DCID string `json:"dcID,omitempty" protobuf:"bytes,1,opt,name=dcID"`

	// +optional
	ClusterID string `json:"clusterID,omitempty" protobuf:"bytes,2,opt,name=clusterID"`

	// +optional
	Namespace string `json:"objNamespace,omitempty" protobuf:"bytes,3,opt,name=objNamespace"`

	// +optional
	ObjectName string `json:"objName,omitempty" protobuf:"bytes,4,opt,name=objName"`

	// +optional
	ObjectDCName string `json:"objDCName,omitempty" protobuf:"bytes,5,opt,name=objDCName"`

	// +optional
	ObjectClusterName string `json:"objClusterName,omitempty" protobuf:"bytes,6,opt,name=objClusterName"`

	// +optional
	ObjectLabels map[string]string `json:"objLabels" protobuf:"bytes,7,rep,name=objLabels"`

	// +optional
	ObjectCreationTimestamp metav1.Time `json:"objCreationTimestamp,omitempty" protobuf:"bytes,8,opt,name=objCreationTimestamp"`

	// +optional
	ZoneID string `json:"zoneID,omitempty" protobuf:"bytes,9,opt,name=zoneID"`

	// +optional
	ObjectZoneName string `json:"objZoneName,omitempty" protobuf:"bytes,10,opt,name=objZoneName"`

}

type HostCredential struct {
	// Set to root is not defined
	// +optional
	User string `json:"user,omitempty" protobuf:"bytes,1,opt,name=user"`

	// +optional
	Password string `json:"password,omitempty" protobuf:"bytes,2,opt,name=password"`

	// +optional
	PrivateKey []byte `json:"privateKey,omitempty" protobuf:"bytes,3,opt,name=privateKey"`

	// +optional
	Passphrase string `json:"passphrase,omitempty" protobuf:"bytes,4,opt,name=passphrase"`
}

type Host struct {
	Hostname    string         `json:"hostname,omitempty" protobuf:"bytes,1,opt,name=hostname"`
	SSHEndpoint string         `json:"sshEndpoint,omitempty" protobuf:"bytes,2,opt,name=sshEndpoint"`
	Credential  HostCredential `json:"-"`
	PublicIP    string         `json:"publicIP,omitempty" protobuf:"bytes,3,opt,name=publicIP"`
	PrimaryNIC  string         `json:"primaryNIC,omitempty" protobuf:"bytes,4,opt,name=primaryNIC"`
	PrimaryIP   string         `json:"primaryIP,omitempty" protobuf:"bytes,5,opt,name=primaryIP"`
	NetMask     string         `json:"netMask,omitempty" protobuf:"bytes,6,opt,name=netMask"`
	NodeCIDR    string         `json:"nodeCIDR,omitempty" protobuf:"bytes,7,opt,name=nodeCIDR"`
	NATPort     uint32         `json:"natPort,omitempty" protobuf:"varint,8,opt,name=natPort"`
	Distro      string         `json:"distro,omitempty" protobuf:"bytes,9,opt,name=distro"`
	DataDisks   []string       `json:"dataDisks,omitempty" protobuf:"bytes,10,rep,name=dataDisks"`
	NICs        []string       `json:"nics,omitempty" protobuf:"bytes,11,rep,name=nics"`
}

type StateCount struct {
	NumHealthy   uint32 `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
	NumUnhealthy uint32 `json:"unhealthy" protobuf:"bytes,2,opt,name=unhealthy"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DataCenter struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            DCSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          DCStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DataCenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []DataCenter `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type HostDNAT struct {
	SrcPort   uint32 `protobuf:"varint,1,opt,name=srcPort"`
	NatPort   uint32 `protobuf:"varint,2,opt,name=natPort"`
	Hostname  string `protobuf:"bytes,3,opt,name=hostname"`
	PrimaryIP string `protobuf:"bytes,4,opt,name=primaryIP"`
}

type GatewayDNAT struct {
	NextNatPort uint32 `protobuf:"varint,1,opt,name=nextNatPort"`

	// Hostname:Port => HostDNAT
	NatMap map[string]HostDNAT `protobuf:"bytes,2,rep,name=natMap"`
}

type ClusterCIDR struct {
	DNSIP       string `protobuf:"bytes,1,opt,name=dnsIP"`
	PodCIDR     string `protobuf:"bytes,2,opt,name=podCIDR"`
	ServiceCIDR string `protobuf:"bytes,3,opt,name=serviceCIDR"`
}

type DCSpec struct {
	Warehouse *Host  `json:"warehouse,omitempty" protobuf:"bytes,1,opt,name=warehouse"`
	Gateways  []Host `json:"gateways,omitempty" protobuf:"bytes,2,rep,name=gateways"`
	Registry  *Host  `json:"registry,omitempty" protobuf:"bytes,3,opt,name=registry"`
	Log       *Host  `json:"log,omitempty" protobuf:"bytes,4,opt,name=log"`
	Metric    *Host  `json:"metric,omitempty" protobuf:"bytes,5,opt,name=metric"`
	Ingresses []Host `json:"ingresses,omitempty" protobuf:"bytes,6,rep,name=ingresses"`

	// cluster gateway(typically master) => CIDR
	ClusterCIDRs map[string]ClusterCIDR `json:"-"`
	HostCIDRs    []string               `json:"-"`
	DNATTable    GatewayDNAT            `json:"-"`
}

type DCPhase string

const (
	DCRunning      DCPhase = "Running"
	DCOffline      DCPhase = "Offline"
	DCInstalling   DCPhase = "Installing"
	DCAborted      DCPhase = "Aborted"
	DCInitializing DCPhase = "Initializing"
	DCUpdating     DCPhase = "Updating"
)

type DCStatus struct {
	Phase   DCPhase    `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=DCPhase"`
	Error   string     `json:"error,omitempty" protobuf:"bytes,2,opt,name=error"`
	Healthy bool       `json:"healthy" protobuf:"varint,3,opt,name=healthy"`
	Cluster StateCount `json:"cluster,omitempty" protobuf:"bytes,4,opt,name=cluster"`
	Node    StateCount `json:"node,omitempty" protobuf:"bytes,5,opt,name=node"`
	Pod     StateCount `json:"pod,omitempty" protobuf:"bytes,6,opt,name=pod"`
	Zone	StateCount `json:"zone,omitempty" protobuf:"bytes,7,opt,name=zone"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Zone struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            ZoneSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          ZoneStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Zone `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ZoneSpec struct {
	VisibleInternet bool       `json:"visibleInternet,omitempty" protobuf:"varint,1,opt,name=visibleInternet"`
}

type ZonePhase string

const (
	ZoneAvailable      DCPhase = "Available"
	ZoneUnAvailable      DCPhase = "UnAvailable"
)

type ZoneStatus struct {
	Phase   ZonePhase    `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=ZonePhase"`
	Error   string     `json:"error,omitempty" protobuf:"bytes,2,opt,name=error"`
	Healthy bool       `json:"healthy" protobuf:"varint,3,opt,name=healthy"`
	Cluster StateCount `json:"cluster,omitempty" protobuf:"bytes,4,opt,name=cluster"`
	Node    StateCount `json:"node,omitempty" protobuf:"bytes,5,opt,name=node"`
	Pod     StateCount `json:"pod,omitempty" protobuf:"bytes,6,opt,name=pod"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            ClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          ClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Cluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ClusterContainerPlatform string

const (
	ContainerPlatformKubernetes = "Kubernetes"
	ContainerPlatformOpenShift  = "OpensShift"
)

type InstallationIntegrity string

const (
	InstallationIntegrityFull          = "Full"
	InstallationIntegrityMetricAndLog  = "Metric-Log"
	InstallationIntegrityAgentOrcaOnly = "AgentOrca"
)

type ClusterSpec struct {
	Masters           []Host                   `json:"masters,omitempty" protobuf:"bytes,1,rep,name=masters"`
	Storage           *Host                    `json:"storage,omitempty" protobuf:"bytes,2,opt,name=storage"`
	MasterSchedulable bool                     `json:"masterSchedulable,omitempty" protobuf:"varint,3,opt,name=masterSchedulable"`
	CIDR              ClusterCIDR              `json:"cidr,omitempty" protobuf:"bytes,4,opt,name=cidr"`
	Platform          ClusterContainerPlatform `json:"platform,omitempty" protobuf:"bytes,5,opt,name=platform,casttype=ClusterContainerPlatform"`
	Integrity         InstallationIntegrity    `json:"integrity,omitempty" protobuf:"bytes,6,opt,name=integrity,casttype=InstallationIntegrity"`
}

type ClusterPhase string

const (
	ClusterRunning      ClusterPhase = "Running"
	ClusterOffline      ClusterPhase = "Offline"
	ClusterInstalling   ClusterPhase = "Installing"
	ClusterAborted      ClusterPhase = "Aborted"
	ClusterInitializing ClusterPhase = "Initializing"
	ClusterUpdating     ClusterPhase = "Updating"
)

type ClusterEndpoint struct {
	IP   string `json:"ip,omitempty" protobuf:"bytes,1,opt,name=ip"`
	Port uint32 `json:"port,omitempty" protobuf:"varint,2,opt,name=port"`
}

type ClusterStatus struct {
	Phase      ClusterPhase    `json:"phase,omitempty" protobuf:"bytes,1,opt,name=phase,casttype=ClusterPhase"`
	Error      string          `json:"error,omitempty" protobuf:"bytes,2,opt,name=error"`
	Endpoint   ClusterEndpoint `json:"endpoint,omitempty" protobuf:"bytes,3,opt,name=endpoint"`
	Healthy    bool            `json:"healthy" protobuf:"varint,4,opt,name=healthy"`
	Load       float64         `json:"load,omitempty" protobuf:"fixed64,5,opt,name=load"`
	Node       StateCount      `json:"node,omitempty" protobuf:"bytes,6,opt,name=node"`
	Pod        StateCount      `json:"pod,omitempty" protobuf:"bytes,7,opt,name=pod"`
	Namespace  StateCount      `json:"namespace,omitempty" protobuf:"bytes,8,opt,name=namespace"`
	PVC        StateCount      `json:"pvc,omitempty" protobuf:"bytes,9,opt,name=pvc"`
	PV         StateCount      `json:"pv,omitempty" protobuf:"bytes,10,opt,name=pv"`
	ReplicaSet StateCount      `json:"replicaset,omitempty" protobuf:"bytes,11,opt,name=replicaset"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Node struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            NodeSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NodeSpec struct {
	corev1.NodeSpec `json:",inline"`
}

type NodeStatus struct {
	corev1.NodeStatus `json:",inline"`
	Healthy           bool       `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
	Load              float64    `json:"load,omitempty" protobuf:"bytes,2,opt,name=load"`
	Pod               StateCount `json:"pod,omitempty" protobuf:"bytes,3,opt,name=pod"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Node `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Namespace struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            NamespaceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          NamespaceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NamespaceSpec struct {
	corev1.NamespaceSpec `json:",inline"`
}

type NamespaceStatus struct {
	corev1.NamespaceStatus `json:",inline"`
	Healthy                bool       `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
	Pod                    StateCount `json:"pod,omitempty" protobuf:"bytes,2,opt,name=pod"`
	PVC                    StateCount `json:"pvc,omitempty" protobuf:"bytes,3,opt,name=pvc"`
	ReplicaSet             StateCount `json:"replicaset,omitempty" protobuf:"bytes,4,opt,name=replicaset"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Namespace `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Deployment struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            DeploymentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          DeploymentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DeploymentSpec struct {
	appv1.DeploymentSpec `json:",inline"`
}

type DeploymentStatus struct {
	appv1.DeploymentStatus `json:",inline"`
	Healthy                bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Deployment `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StatefulSet struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            StatefulSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          StatefulSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type StatefulSetSpec struct {
	appv1.StatefulSetSpec `json:",inline"`
}

type StatefulSetStatus struct {
	appv1.StatefulSetStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StatefulSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []StatefulSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReplicaSet struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            ReplicaSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          ReplicaSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ReplicaSetSpec struct {
	Stateless *DeploymentSpec  `json:"stateless,omitempty" protobuf:"bytes,1,opt,name=stateless"`
	Stateful  *StatefulSetSpec `json:"stateful,omitempty" protobuf:"bytes,2,opt,name=stateful"`
}

type ReplicaSetStatus struct {
	Stateless *DeploymentStatus  `json:"stateless,omitempty" protobuf:"bytes,1,opt,name=stateless"`
	Stateful  *StatefulSetStatus `json:"stateful,omitempty" protobuf:"bytes,2,opt,name=stateful"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ReplicaSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PVC struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            PVCSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          PVCStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PVCSpec struct {
	corev1.PersistentVolumeClaimSpec `json:",inline"`
}

type PVCStatus struct {
	corev1.PersistentVolumeClaimStatus `json:",inline"`
	Healthy                            bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PVCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []PVC `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PV struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            PVSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          PVStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PVSpec struct {
	corev1.PersistentVolumeSpec `json:",inline"`
}

type PVStatus struct {
	corev1.PersistentVolumeStatus `json:",inline"`
	Healthy                       bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PVList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []PV `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Event struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Spec            EventSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type EventSpec struct {
	// The object that this event is about.
	InvolvedObject corev1.ObjectReference `json:"involvedObject" protobuf:"bytes,2,opt,name=involvedObject"`

	// This should be a short, machine understandable string that gives the reason
	// for the transition into the object's current status.
	// TODO: provide exact specification for format.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`

	// A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`

	// The component reporting this event. Should be a short machine understandable string.
	// +optional
	Source corev1.EventSource `json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	// +optional
	FirstTimestamp metav1.Time `json:"firstTimestamp,omitempty" protobuf:"bytes,6,opt,name=firstTimestamp"`

	// The time at which the most recent occurrence of this event was recorded.
	// +optional
	LastTimestamp metav1.Time `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp"`

	// The number of times this event has occurred.
	// +optional
	Count int32 `json:"count,omitempty" protobuf:"varint,8,opt,name=count"`

	// Type of this event (Normal, Warning), new types could be added in the future
	// +optional
	Type string `json:"type,omitempty" protobuf:"bytes,9,opt,name=type"`

	// Time when this Event was first observed.
	// +optional
	EventTime metav1.MicroTime `json:"eventTime,omitempty" protobuf:"bytes,10,opt,name=eventTime"`

	// Data about the Event series this event represents or nil if it's a singleton Event.
	// +optional
	Series *corev1.EventSeries `json:"series,omitempty" protobuf:"bytes,11,opt,name=series"`

	// What action was taken/failed regarding to the Regarding object.
	// +optional
	Action string `json:"action,omitempty" protobuf:"bytes,12,opt,name=action"`

	// Optional secondary object for more complex actions.
	// +optional
	Related *corev1.ObjectReference `json:"related,omitempty" protobuf:"bytes,13,opt,name=related"`

	// ObjectName of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	// +optional
	ReportingController string `json:"reportingComponent" protobuf:"bytes,14,opt,name=reportingComponent"`

	// ID of the controller instance, e.g. `kubelet-xyzf`.
	// +optional
	ReportingInstance string `json:"reportingInstance" protobuf:"bytes,15,opt,name=reportingInstance"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EventList is a list of events.
type EventList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of events
	Items []Event `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Pod struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            PodSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PodSpec struct {
	corev1.PodSpec `json:",inline"`
}

type PodStatus struct {
	corev1.PodStatus     `json:",inline"`
	Healthy              bool              `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
	ContainerTerminalKey map[string][]byte `json:"containerTerminalKey,omitempty" protobuf:"bytes,2,opt,name=containerTerminalKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Pod `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Service struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            ServiceSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          ServiceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ServiceSpec struct {
	corev1.ServiceSpec `json:",inline"`
}

type ServiceStatus struct {
	corev1.ServiceStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Service `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DaemonSet struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            DaemonSetSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          DaemonSetStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type DaemonSetSpec struct {
	appv1.DaemonSetSpec `json:",inline"`
}

type DaemonSetStatus struct {
	appv1.DaemonSetStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []DaemonSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}


// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Job struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            JobSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          JobStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type JobSpec struct {
	jobv1.JobSpec `json:",inline"`
}

type JobStatus struct {
	jobv1.JobStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Job `json:"items" protobuf:"bytes,2,rep,name=items"`
}


// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CronJob struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            CronJobSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          CronJobStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type CronJobSpec struct {
	cronjobv1.CronJobSpec `json:",inline"`
}

type CronJobStatus struct {
	cronjobv1.CronJobStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []CronJob `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HPA struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            HPASpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          HPAStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type HPASpec struct {
	autov1.HorizontalPodAutoscalerSpec `json:",inline"`
}

type HPAStatus struct {
	autov1.HorizontalPodAutoscalerStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HPAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []HPA `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Ingress struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec            IngressSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status          IngressStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type IngressSpec struct {
	ingressv1.IngressSpec `json:",inline"`
}

type IngressStatus struct {
	ingressv1.IngressStatus `json:",inline"`
	Healthy                 bool `json:"healthy" protobuf:"bytes,1,opt,name=healthy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Ingress `json:"items" protobuf:"bytes,2,rep,name=items"`
}
