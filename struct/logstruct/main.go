package logstruct

import "time"

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type GceLog struct {
	InsertId string `json:"insertId"`
	Labels   struct {
		ComputeGoogleapisComRootTriggerId string `json:"compute.googleapis.com/root_trigger_id"`
	} `json:"labels"`
	LogName   string `json:"logName"`
	Operation struct {
		First    bool   `json:"first"`
		Id       string `json:"id"`
		Producer string `json:"producer"`
	} `json:"operation"`
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		AuthorizationInfo []struct {
			Granted            bool   `json:"granted"`
			Permission         string `json:"permission"`
			PermissionType     string `json:"permissionType"`
			Resource           string `json:"resource"`
			ResourceAttributes struct {
				Name    string `json:"name"`
				Service string `json:"service"`
				Type    string `json:"type"`
			} `json:"resourceAttributes"`
		} `json:"authorizationInfo"`
		MethodName string `json:"methodName"`
		Request    struct {
			Type                       string `json:"@type"`
			CanIpForward               bool   `json:"canIpForward"`
			ConfidentialInstanceConfig struct {
				EnableConfidentialCompute bool `json:"enableConfidentialCompute"`
			} `json:"confidentialInstanceConfig"`
			DeletionProtection bool   `json:"deletionProtection"`
			Description        string `json:"description"`
			Disks              []struct {
				AutoDelete       bool   `json:"autoDelete"`
				Boot             bool   `json:"boot"`
				DeviceName       string `json:"deviceName"`
				InitializeParams struct {
					DiskSizeGb  string `json:"diskSizeGb"`
					DiskType    string `json:"diskType"`
					SourceImage string `json:"sourceImage"`
				} `json:"initializeParams"`
				Mode string `json:"mode"`
				Type string `json:"type"`
			} `json:"disks"`
			DisplayDevice struct {
				EnableDisplay bool `json:"enableDisplay"`
			} `json:"displayDevice"`
			KeyRevocationActionType string `json:"keyRevocationActionType"`
			MachineType             string `json:"machineType"`
			Name                    string `json:"name"`
			NetworkInterfaces       []struct {
				AccessConfigs []struct {
					Name        string `json:"name"`
					NetworkTier string `json:"networkTier"`
				} `json:"accessConfigs"`
				StackType  string `json:"stackType"`
				Subnetwork string `json:"subnetwork"`
			} `json:"networkInterfaces"`
			ReservationAffinity struct {
				ConsumeReservationType string `json:"consumeReservationType"`
			} `json:"reservationAffinity"`
			Scheduling struct {
				AutomaticRestart  bool   `json:"automaticRestart"`
				OnHostMaintenance string `json:"onHostMaintenance"`
				ProvisioningModel string `json:"provisioningModel"`
			} `json:"scheduling"`
			ServiceAccounts []struct {
				Email  string   `json:"email"`
				Scopes []string `json:"scopes"`
			} `json:"serviceAccounts"`
			ShieldedInstanceConfig struct {
				EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring"`
				EnableSecureBoot          bool `json:"enableSecureBoot"`
				EnableVtpm                bool `json:"enableVtpm"`
			} `json:"shieldedInstanceConfig"`
		} `json:"request"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			DestinationAttributes   struct {
			} `json:"destinationAttributes"`
			RequestAttributes struct {
				Auth struct {
				} `json:"auth"`
				Reason string    `json:"reason"`
				Time   time.Time `json:"time"`
			} `json:"requestAttributes"`
		} `json:"requestMetadata"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
		ResourceName string `json:"resourceName"`
		Response     struct {
			Type           string    `json:"@type"`
			Id             string    `json:"id"`
			InsertTime     time.Time `json:"insertTime"`
			Name           string    `json:"name"`
			OperationType  string    `json:"operationType"`
			Progress       string    `json:"progress"`
			SelfLink       string    `json:"selfLink"`
			SelfLinkWithId string    `json:"selfLinkWithId"`
			StartTime      time.Time `json:"startTime"`
			Status         string    `json:"status"`
			TargetId       string    `json:"targetId"`
			TargetLink     string    `json:"targetLink"`
			User           string    `json:"user"`
			Zone           string    `json:"zone"`
		} `json:"response"`
		ServiceName string `json:"serviceName"`
	} `json:"protoPayload"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
	Resource         struct {
		Labels struct {
			InstanceId string `json:"instance_id"`
			ProjectId  string `json:"project_id"`
			Zone       string `json:"zone"`
		} `json:"labels"`
		Type string `json:"type"`
	} `json:"resource"`
	Severity  string    `json:"severity"`
	Timestamp time.Time `json:"timestamp"`
}
type DiskLog struct {
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time   time.Time `json:"time"`
				Reason string    `json:"reason"`
				Auth   struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
				Type    string `json:"type"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Name   string `json:"name"`
			SizeGb string `json:"sizeGb"`
			Type   string `json:"type"`
			Labels []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"labels"`
			ResourcePolicys []string `json:"resourcePolicys"`
			Type1           string   `json:"@type"`
		} `json:"request"`
		Response struct {
			Id             string    `json:"id"`
			Name           string    `json:"name"`
			Zone           string    `json:"zone"`
			OperationType  string    `json:"operationType"`
			TargetLink     string    `json:"targetLink"`
			TargetId       string    `json:"targetId"`
			Status         string    `json:"status"`
			User           string    `json:"user"`
			Progress       string    `json:"progress"`
			InsertTime     time.Time `json:"insertTime"`
			StartTime      time.Time `json:"startTime"`
			SelfLink       string    `json:"selfLink"`
			SelfLinkWithId string    `json:"selfLinkWithId"`
			Type           string    `json:"@type"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			Zone      string `json:"zone"`
			ProjectId string `json:"project_id"`
			DiskId    string `json:"disk_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	Labels    struct {
		ComputeGoogleapisComRootTriggerId string `json:"compute.googleapis.com/root_trigger_id"`
	} `json:"labels"`
	LogName   string `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type IpaddressLog struct {
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
		} `json:"requestMetadata"`
		ServiceName  string `json:"serviceName"`
		MethodName   string `json:"methodName"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Type string `json:"@type"`
		} `json:"request"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ReservedAddressId string `json:"reserved_address_id"`
			ProjectId         string `json:"project_id"`
			Location          string `json:"location"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	Labels    struct {
		ComputeGoogleapisComRootTriggerId string `json:"compute.googleapis.com/root_trigger_id"`
	} `json:"labels"`
	LogName   string `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		Last     bool   `json:"last"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type SqlLog struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp          string `json:"callerIp"`
			RequestAttributes struct {
				Time   time.Time `json:"time"`
				Reason string    `json:"reason"`
				Auth   struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
				Type    string `json:"type"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			DatabaseId string `json:"database_id"`
			Region     string `json:"region"`
			ProjectId  string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		Last     bool   `json:"last"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type RedisLog struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp          string `json:"callerIp"`
			RequestAttributes struct {
				Time   time.Time `json:"time"`
				Reason string    `json:"reason"`
				Auth   struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			InstanceId string `json:"instance_id"`
			Instance   struct {
				MemorySizeGb          int    `json:"memory_size_gb"`
				AuthorizedNetwork     string `json:"authorized_network"`
				ConnectMode           string `json:"connect_mode"`
				TransitEncryptionMode string `json:"transit_encryption_mode"`
				PersistenceConfig     struct {
					PersistenceMode string `json:"persistence_mode"`
				} `json:"persistence_config"`
				RedisVersion string `json:"redis_version"`
				Tier         string `json:"tier"`
			} `json:"instance"`
			Type   string `json:"@type"`
			Parent string `json:"parent"`
		} `json:"request"`
		Response struct {
			Type string `json:"@type"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			Method    string `json:"method"`
			ProjectId string `json:"project_id"`
			Service   string `json:"service"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type DatasetlogBg struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Metadata     struct {
			Type            string `json:"@type"`
			DatasetCreation struct {
				Reason  string `json:"reason"`
				Dataset struct {
					UpdateTime  time.Time `json:"updateTime"`
					DatasetName string    `json:"datasetName"`
					CreateTime  time.Time `json:"createTime"`
					Acl         struct {
						Policy struct {
							Bindings []struct {
								Role    string   `json:"role"`
								Members []string `json:"members"`
							} `json:"bindings"`
						} `json:"policy"`
					} `json:"acl"`
				} `json:"dataset"`
			} `json:"datasetCreation"`
		} `json:"metadata"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			DatasetId string `json:"dataset_id"`
			ProjectId string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp        time.Time `json:"timestamp"`
	Severity         string    `json:"severity"`
	LogName          string    `json:"logName"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type TablelogBG struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		ServiceData  struct {
			Type               string `json:"@type"`
			TableInsertRequest struct {
				Resource struct {
					TableName struct {
						ProjectId string `json:"projectId"`
						DatasetId string `json:"datasetId"`
						TableId   string `json:"tableId"`
					} `json:"tableName"`
					Info struct {
					} `json:"info"`
					View struct {
					} `json:"view"`
					SchemaJson string `json:"schemaJson"`
				} `json:"resource"`
			} `json:"tableInsertRequest"`
			TableInsertResponse struct {
				Resource struct {
					TableName struct {
						ProjectId string `json:"projectId"`
						DatasetId string `json:"datasetId"`
						TableId   string `json:"tableId"`
					} `json:"tableName"`
					Info struct {
					} `json:"info"`
					View struct {
					} `json:"view"`
					CreateTime time.Time `json:"createTime"`
					SchemaJson string    `json:"schemaJson"`
					UpdateTime time.Time `json:"updateTime"`
				} `json:"resource"`
			} `json:"tableInsertResponse"`
		} `json:"serviceData"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ProjectId string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp        time.Time `json:"timestamp"`
	Severity         string    `json:"severity"`
	LogName          string    `json:"logName"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type ClusterlogDP struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail   string `json:"principalEmail"`
			PrincipalSubject string `json:"principalSubject"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Cluster struct {
				ProjectId   string `json:"projectId"`
				ClusterName string `json:"clusterName"`
				Config      struct {
					GceClusterConfig struct {
						SubnetworkUri          string `json:"subnetworkUri"`
						InternalIpOnly         bool   `json:"internalIpOnly"`
						ShieldedInstanceConfig struct {
							EnableSecureBoot          bool `json:"enableSecureBoot"`
							EnableVtpm                bool `json:"enableVtpm"`
							EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring"`
						} `json:"shieldedInstanceConfig"`
					} `json:"gceClusterConfig"`
					MasterConfig struct {
						NumInstances   int    `json:"numInstances"`
						MachineTypeUri string `json:"machineTypeUri"`
						DiskConfig     struct {
							BootDiskSizeGb int    `json:"bootDiskSizeGb"`
							BootDiskType   string `json:"bootDiskType"`
						} `json:"diskConfig"`
					} `json:"masterConfig"`
					WorkerConfig struct {
						NumInstances   int    `json:"numInstances"`
						MachineTypeUri string `json:"machineTypeUri"`
						DiskConfig     struct {
							BootDiskSizeGb int    `json:"bootDiskSizeGb"`
							BootDiskType   string `json:"bootDiskType"`
						} `json:"diskConfig"`
					} `json:"workerConfig"`
					SecondaryWorkerConfig struct {
					} `json:"secondaryWorkerConfig"`
					SoftwareConfig struct {
						ImageVersion string `json:"imageVersion"`
					} `json:"softwareConfig"`
					EncryptionConfig struct {
					} `json:"encryptionConfig"`
					SecurityConfig struct {
						KerberosConfig struct {
						} `json:"kerberosConfig"`
					} `json:"securityConfig"`
					LifecycleConfig struct {
					} `json:"lifecycleConfig"`
					AutoscalingConfig struct {
					} `json:"autoscalingConfig"`
					EndpointConfig struct {
					} `json:"endpointConfig"`
				} `json:"config"`
				Status struct {
				} `json:"status"`
				StatusHistory []struct {
				} `json:"statusHistory"`
				Metrics struct {
				} `json:"metrics"`
			} `json:"cluster"`
			Region    string `json:"region"`
			ProjectId string `json:"projectId"`
			Type      string `json:"@type"`
		} `json:"request"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ClusterUuid string `json:"cluster_uuid"`
			ProjectId   string `json:"project_id"`
			Region      string `json:"region"`
			ClusterName string `json:"cluster_name"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type JoblogDP struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail   string `json:"principalEmail"`
			PrincipalSubject string `json:"principalSubject"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Job struct {
				Reference struct {
					JobId string `json:"jobId"`
				} `json:"reference"`
				Placement struct {
					ClusterName string `json:"clusterName"`
				} `json:"placement"`
				HadoopJob struct {
					MainClass string `json:"mainClass"`
				} `json:"hadoopJob"`
			} `json:"job"`
			Region    string `json:"region"`
			ProjectId string `json:"projectId"`
			Type      string `json:"@type"`
		} `json:"request"`
		Response struct {
			Reference struct {
				ProjectId string `json:"projectId"`
				JobId     string `json:"jobId"`
			} `json:"reference"`
			HadoopJob struct {
				MainClass string `json:"mainClass"`
			} `json:"hadoopJob"`
			JobUuid               string `json:"jobUuid"`
			DriverControlFilesUri string `json:"driverControlFilesUri"`
			Placement             struct {
				ClusterName string `json:"clusterName"`
				ClusterUuid string `json:"clusterUuid"`
			} `json:"placement"`
			DriverOutputResourceUri string `json:"driverOutputResourceUri"`
			Status                  struct {
				State          string    `json:"state"`
				StateStartTime time.Time `json:"stateStartTime"`
			} `json:"status"`
			Type string `json:"@type"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ProjectId   string `json:"project_id"`
			Region      string `json:"region"`
			ClusterUuid string `json:"cluster_uuid"`
			ClusterName string `json:"cluster_name"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp        time.Time `json:"timestamp"`
	Severity         string    `json:"severity"`
	LogName          string    `json:"logName"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type Gcslog struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		ServiceData  struct {
			Type        string `json:"@type"`
			PolicyDelta struct {
				BindingDeltas []struct {
					Action string `json:"action"`
					Role   string `json:"role"`
					Member string `json:"member"`
				} `json:"bindingDeltas"`
			} `json:"policyDelta"`
		} `json:"serviceData"`
		Request struct {
			DefaultObjectAcl struct {
				Type     string `json:"@type"`
				Bindings []struct {
					Role    string   `json:"role"`
					Members []string `json:"members"`
				} `json:"bindings"`
			} `json:"defaultObjectAcl"`
		} `json:"request"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			Location   string `json:"location"`
			ProjectId  string `json:"project_id"`
			BucketName string `json:"bucket_name"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp        time.Time `json:"timestamp"`
	Severity         string    `json:"severity"`
	LogName          string    `json:"logName"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type Gkelog struct {
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time   time.Time `json:"time"`
				Reason string    `json:"reason"`
				Auth   struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
				Type    string `json:"type"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Cluster struct {
				NetworkPolicy struct {
				} `json:"networkPolicy"`
				NotificationConfig struct {
				} `json:"notificationConfig"`
				SecurityPostureConfig struct {
				} `json:"securityPostureConfig"`
				AuthenticatorGroupsConfig struct {
				} `json:"authenticatorGroupsConfig"`
				MonitoringConfig struct {
					ComponentConfig struct {
						EnableComponents []string `json:"enableComponents"`
					} `json:"componentConfig"`
					ManagedPrometheusConfig struct {
						Enabled bool `json:"enabled"`
					} `json:"managedPrometheusConfig"`
				} `json:"monitoringConfig"`
				InitialClusterVersion string `json:"initialClusterVersion"`
				Location              string `json:"location"`
				BinaryAuthorization   struct {
					EvaluationMode string `json:"evaluationMode"`
				} `json:"binaryAuthorization"`
				Autoscaling struct {
				} `json:"autoscaling"`
				NetworkConfig struct {
					DatapathProvider string `json:"datapathProvider"`
				} `json:"networkConfig"`
				Network                        string `json:"network"`
				MasterAuthorizedNetworksConfig struct {
				} `json:"masterAuthorizedNetworksConfig"`
				Locations          []string `json:"locations"`
				IpAllocationPolicy struct {
					StackType    string `json:"stackType"`
					UseIpAliases bool   `json:"useIpAliases"`
				} `json:"ipAllocationPolicy"`
				Name      string `json:"name"`
				NodePools []struct {
					NetworkConfig struct {
					} `json:"networkConfig"`
					UpgradeSettings struct {
						MaxSurge int    `json:"maxSurge"`
						Strategy string `json:"strategy"`
					} `json:"upgradeSettings"`
					QueuedProvisioning struct {
					} `json:"queuedProvisioning"`
					Name   string `json:"name"`
					Config struct {
						ImageType              string `json:"imageType"`
						DiskType               string `json:"diskType"`
						ShieldedInstanceConfig struct {
						} `json:"shieldedInstanceConfig"`
						AdvancedMachineFeatures struct {
							EnableNestedVirtualization bool `json:"enableNestedVirtualization"`
						} `json:"advancedMachineFeatures"`
						MachineType string   `json:"machineType"`
						DiskSizeGb  int      `json:"diskSizeGb"`
						OauthScopes []string `json:"oauthScopes"`
						Metadata    struct {
							DisableLegacyEndpoints string `json:"disable-legacy-endpoints"`
						} `json:"metadata"`
					} `json:"config"`
					InitialNodeCount int `json:"initialNodeCount"`
					Autoscaling      struct {
					} `json:"autoscaling"`
					Management struct {
						AutoUpgrade bool `json:"autoUpgrade"`
						AutoRepair  bool `json:"autoRepair"`
					} `json:"management"`
				} `json:"nodePools"`
				ShieldedNodes struct {
				} `json:"shieldedNodes"`
				ReleaseChannel struct {
					Channel string `json:"channel"`
				} `json:"releaseChannel"`
				LoggingConfig struct {
					ComponentConfig struct {
						EnableComponents []string `json:"enableComponents"`
					} `json:"componentConfig"`
				} `json:"loggingConfig"`
				AddonsConfig struct {
					HorizontalPodAutoscaling struct {
					} `json:"horizontalPodAutoscaling"`
					DnsCacheConfig struct {
					} `json:"dnsCacheConfig"`
					GcePersistentDiskCsiDriverConfig struct {
						Enabled bool `json:"enabled"`
					} `json:"gcePersistentDiskCsiDriverConfig"`
					GcsFuseCsiDriverConfig struct {
					} `json:"gcsFuseCsiDriverConfig"`
					HttpLoadBalancing struct {
					} `json:"httpLoadBalancing"`
				} `json:"addonsConfig"`
				Subnetwork         string `json:"subnetwork"`
				DatabaseEncryption struct {
				} `json:"databaseEncryption"`
				DefaultMaxPodsConstraint struct {
				} `json:"defaultMaxPodsConstraint"`
			} `json:"cluster"`
			Parent string `json:"parent"`
			Type   string `json:"@type"`
		} `json:"request"`
		Response struct {
			Name          string    `json:"name"`
			OperationType string    `json:"operationType"`
			Status        string    `json:"status"`
			Type          string    `json:"@type"`
			SelfLink      string    `json:"selfLink"`
			TargetLink    string    `json:"targetLink"`
			StartTime     time.Time `json:"startTime"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
		PolicyViolationInfo struct {
			OrgPolicyViolationInfo struct {
			} `json:"orgPolicyViolationInfo"`
		} `json:"policyViolationInfo"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ClusterName string `json:"cluster_name"`
			Location    string `json:"location"`
			ProjectId   string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type Arlog struct {
	ProtoPayload struct {
		Type               string `json:"@type"`
		AuthenticationInfo struct {
			PrincipalEmail   string `json:"principalEmail"`
			PrincipalSubject string `json:"principalSubject"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Repository struct {
				Format       string `json:"format"`
				DockerConfig struct {
					ImmutableTags bool `json:"immutableTags"`
				} `json:"dockerConfig"`
				CleanupPolicyDryRun bool `json:"cleanupPolicyDryRun"`
			} `json:"repository"`
			RepositoryId string `json:"repositoryId"`
			Type         string `json:"@type"`
			Parent       string `json:"parent"`
		} `json:"request"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			Service   string `json:"service"`
			Method    string `json:"method"`
			ProjectId string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type FilestoreInstanceLog struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail   string `json:"principalEmail"`
			PrincipalSubject string `json:"principalSubject"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			InstanceId string `json:"instance_id"`
			Instance   struct {
				Tier       string `json:"tier"`
				FileShares []struct {
					CapacityGb int    `json:"capacity_gb"`
					Name       string `json:"name"`
				} `json:"file_shares"`
				Networks []struct {
					Network         string   `json:"network"`
					Modes           []string `json:"modes"`
					ReservedIpRange string   `json:"reserved_ip_range"`
				} `json:"networks"`
				Description string `json:"description"`
			} `json:"instance"`
			Parent string `json:"parent"`
			Type   string `json:"@type"`
		} `json:"request"`
		Response struct {
			Type string `json:"@type"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ProjectId string `json:"project_id"`
			Method    string `json:"method"`
			Service   string `json:"service"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
type FilestoreBackupLog struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail   string `json:"principalEmail"`
			PrincipalSubject string `json:"principalSubject"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIp                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
			RequestAttributes       struct {
				Time time.Time `json:"time"`
				Auth struct {
				} `json:"auth"`
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
		} `json:"requestMetadata"`
		ServiceName       string `json:"serviceName"`
		MethodName        string `json:"methodName"`
		AuthorizationInfo []struct {
			Resource           string `json:"resource"`
			Permission         string `json:"permission"`
			Granted            bool   `json:"granted"`
			ResourceAttributes struct {
				Service string `json:"service"`
				Name    string `json:"name"`
			} `json:"resourceAttributes"`
			PermissionType string `json:"permissionType"`
		} `json:"authorizationInfo"`
		ResourceName string `json:"resourceName"`
		Request      struct {
			Backup struct {
				SourceFileShare string `json:"source_file_share"`
				Description     string `json:"description"`
				SourceInstance  string `json:"source_instance"`
			} `json:"backup"`
			BackupId string `json:"backup_id"`
			Type     string `json:"@type"`
			Parent   string `json:"parent"`
		} `json:"request"`
		Response struct {
			Type string `json:"@type"`
		} `json:"response"`
		ResourceLocation struct {
			CurrentLocations []string `json:"currentLocations"`
		} `json:"resourceLocation"`
	} `json:"protoPayload"`
	InsertId string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			Service   string `json:"service"`
			ProjectId string `json:"project_id"`
			Method    string `json:"method"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	LogName   string    `json:"logName"`
	Operation struct {
		Id       string `json:"id"`
		Producer string `json:"producer"`
		First    bool   `json:"first"`
	} `json:"operation"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
