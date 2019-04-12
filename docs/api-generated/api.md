# Protocol Documentation
<a name="top"/>

## Table of Contents

- [api.proto](#api.proto)
    - [AWSCredentials](#cmaaws.AWSCredentials)
    - [ClusterDetailItem](#cmaaws.ClusterDetailItem)
    - [ClusterItem](#cmaaws.ClusterItem)
    - [CreateClusterAWSSpec](#cmaaws.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSDataCenter](#cmaaws.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cmaaws.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cmaaws.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaws.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaws.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaws.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaws.DeleteClusterReply)
    - [GetClusterListMsg](#cmaaws.GetClusterListMsg)
    - [GetClusterListReply](#cmaaws.GetClusterListReply)
    - [GetClusterMsg](#cmaaws.GetClusterMsg)
    - [GetClusterReply](#cmaaws.GetClusterReply)
    - [GetVersionMsg](#cmaaws.GetVersionMsg)
    - [GetVersionReply](#cmaaws.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaws.GetVersionReply.VersionInformation)
  
    - [ClusterStatus](#cmaaws.ClusterStatus)
  
  
    - [Cluster](#cmaaws.Cluster)
  

- [api.proto](#api.proto)
    - [AWSCredentials](#cmaaws.AWSCredentials)
    - [ClusterDetailItem](#cmaaws.ClusterDetailItem)
    - [ClusterItem](#cmaaws.ClusterItem)
    - [CreateClusterAWSSpec](#cmaaws.CreateClusterAWSSpec)
    - [CreateClusterAWSSpec.AWSDataCenter](#cmaaws.CreateClusterAWSSpec.AWSDataCenter)
    - [CreateClusterAWSSpec.AWSInstanceGroup](#cmaaws.CreateClusterAWSSpec.AWSInstanceGroup)
    - [CreateClusterAWSSpec.AWSPreconfiguredItems](#cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems)
    - [CreateClusterMsg](#cmaaws.CreateClusterMsg)
    - [CreateClusterProviderSpec](#cmaaws.CreateClusterProviderSpec)
    - [CreateClusterReply](#cmaaws.CreateClusterReply)
    - [DeleteClusterMsg](#cmaaws.DeleteClusterMsg)
    - [DeleteClusterReply](#cmaaws.DeleteClusterReply)
    - [GetClusterListMsg](#cmaaws.GetClusterListMsg)
    - [GetClusterListReply](#cmaaws.GetClusterListReply)
    - [GetClusterMsg](#cmaaws.GetClusterMsg)
    - [GetClusterReply](#cmaaws.GetClusterReply)
    - [GetVersionMsg](#cmaaws.GetVersionMsg)
    - [GetVersionReply](#cmaaws.GetVersionReply)
    - [GetVersionReply.VersionInformation](#cmaaws.GetVersionReply.VersionInformation)
  
    - [ClusterStatus](#cmaaws.ClusterStatus)
  
  
    - [Cluster](#cmaaws.Cluster)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaws.AWSCredentials"/>

### AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |
| region | [string](#string) |  | The Region for API access |






<a name="cmaaws.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cmaaws.ClusterStatus) |  | The status of the cluster |






<a name="cmaaws.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cmaaws.ClusterStatus) |  | The status of the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cmaaws.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cmaaws.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






<a name="cmaaws.CreateClusterAWSSpec.AWSDataCenter"/>

### CreateClusterAWSSpec.AWSDataCenter
Which Data Center


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | Which region (us-east-1, etc.) |
| availability_zones | [string](#string) | repeated | Which availability zones (us-east-1b, us-east-2c, us-west-2d, etc.) |






<a name="cmaaws.CreateClusterAWSSpec.AWSInstanceGroup"/>

### CreateClusterAWSSpec.AWSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Instance type (m5.large, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |
| name | [string](#string) |  | Name of instance group (optional auto-generated if not set) |
| desired_quantity | [int32](#int32) |  | Number of worker node instances desired (defaults to zero) |
| ssh_access | [bool](#bool) |  | SSH Access to open ssh access on the nodes |






<a name="cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems"/>

### CreateClusterAWSSpec.AWSPreconfiguredItems
For when some things are already created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpc_id | [string](#string) |  | The VPC id, blank for for &#34;create one for you&#34;, filled if you are BYO VPC |
| security_group_id | [string](#string) |  | Security group |
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn) |






<a name="cmaaws.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaws.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaws.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cmaaws.CreateClusterAWSSpec) |  | The AWS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaws.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaws.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaws.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to delete the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaws.GetClusterListMsg"/>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to query for the cluster |
| region | [string](#string) |  | Datacenter region (all regions checked if not set) |






<a name="cmaaws.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaws.ClusterItem) | repeated | List of clusters |






<a name="cmaaws.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to query for the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaws.ClusterDetailItem) |  |  |






<a name="cmaaws.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cmaaws.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaws.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaws.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 


<a name="cmaaws.ClusterStatus"/>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |


 

 


<a name="cmaaws.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaws.CreateClusterMsg) | [CreateClusterReply](#cmaaws.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaws.GetClusterMsg) | [GetClusterReply](#cmaaws.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaws.DeleteClusterMsg) | [DeleteClusterReply](#cmaaws.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaws.GetClusterListMsg) | [GetClusterListReply](#cmaaws.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaws.GetVersionMsg) | [GetVersionReply](#cmaaws.GetVersionMsg) | Will return version information about api server |

 



<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="cmaaws.AWSCredentials"/>

### AWSCredentials
The credentials to use for creating the cluster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret_key_id | [string](#string) |  | The SecretKeyId for API Access |
| secret_access_key | [string](#string) |  | The SecretAccessKey for API access |
| region | [string](#string) |  | The Region for API access |






<a name="cmaaws.ClusterDetailItem"/>

### ClusterDetailItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| kubeconfig | [string](#string) |  | What is the kubeconfig to connect to the cluster |
| status | [ClusterStatus](#cmaaws.ClusterStatus) |  | The status of the cluster |






<a name="cmaaws.ClusterItem"/>

### ClusterItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | ID of the cluster |
| name | [string](#string) |  | Name of the cluster |
| status_message | [string](#string) |  | Additional information about the status of the cluster |
| status | [ClusterStatus](#cmaaws.ClusterStatus) |  | The status of the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.CreateClusterAWSSpec"/>

### CreateClusterAWSSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data_center | [CreateClusterAWSSpec.AWSDataCenter](#cmaaws.CreateClusterAWSSpec.AWSDataCenter) |  | The AWS Data Center |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to build the cluster |
| resources | [CreateClusterAWSSpec.AWSPreconfiguredItems](#cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems) |  | BYO items |
| instance_groups | [CreateClusterAWSSpec.AWSInstanceGroup](#cmaaws.CreateClusterAWSSpec.AWSInstanceGroup) | repeated | Instance groups |






<a name="cmaaws.CreateClusterAWSSpec.AWSDataCenter"/>

### CreateClusterAWSSpec.AWSDataCenter
Which Data Center


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| region | [string](#string) |  | Which region (us-east-1, etc.) |
| availability_zones | [string](#string) | repeated | Which availability zones (us-east-1b, us-east-2c, us-west-2d, etc.) |






<a name="cmaaws.CreateClusterAWSSpec.AWSInstanceGroup"/>

### CreateClusterAWSSpec.AWSInstanceGroup
Instance groups define a type and number of instances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Instance type (m5.large, etc.) |
| min_quantity | [int32](#int32) |  | Minimum number of instances (defaults to zero) |
| max_quantity | [int32](#int32) |  | Maximum number of instances (defaults to zero) |
| name | [string](#string) |  | Name of instance group (optional auto-generated if not set) |
| desired_quantity | [int32](#int32) |  | Number of worker node instances desired (defaults to zero) |
| ssh_access | [bool](#bool) |  | SSH Access to open ssh access on the nodes |






<a name="cmaaws.CreateClusterAWSSpec.AWSPreconfiguredItems"/>

### CreateClusterAWSSpec.AWSPreconfiguredItems
For when some things are already created


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vpc_id | [string](#string) |  | The VPC id, blank for for &#34;create one for you&#34;, filled if you are BYO VPC |
| security_group_id | [string](#string) |  | Security group |
| iam_role_arn | [string](#string) |  | The IAM role for the cluster (arn) |






<a name="cmaaws.CreateClusterMsg"/>

### CreateClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be provisioned |
| provider | [CreateClusterProviderSpec](#cmaaws.CreateClusterProviderSpec) |  | The provider specification |






<a name="cmaaws.CreateClusterProviderSpec"/>

### CreateClusterProviderSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the provider - currently this is aws or maas |
| k8s_version | [string](#string) |  | The version of Kubernetes |
| aws | [CreateClusterAWSSpec](#cmaaws.CreateClusterAWSSpec) |  | The AWS specification |
| high_availability | [bool](#bool) |  | Whether or not the cluster is HA |
| network_fabric | [string](#string) |  | The fabric to be used |






<a name="cmaaws.CreateClusterReply"/>

### CreateClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Whether or not the cluster was provisioned by this request |
| cluster | [ClusterItem](#cmaaws.ClusterItem) |  | The details of the cluster request response |






<a name="cmaaws.DeleteClusterMsg"/>

### DeleteClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | What is the cluster&#39;s name to destroy |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to delete the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.DeleteClusterReply"/>

### DeleteClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Could the cluster be destroyed |
| status | [string](#string) |  | Status of the request |






<a name="cmaaws.GetClusterListMsg"/>

### GetClusterListMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to query for the cluster |
| region | [string](#string) |  | Datacenter region (all regions checked if not set) |






<a name="cmaaws.GetClusterListReply"/>

### GetClusterListReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| clusters | [ClusterItem](#cmaaws.ClusterItem) | repeated | List of clusters |






<a name="cmaaws.GetClusterMsg"/>

### GetClusterMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the cluster to be looked up |
| credentials | [AWSCredentials](#cmaaws.AWSCredentials) |  | Credentials to query for the cluster |
| region | [string](#string) |  | Datacenter region |






<a name="cmaaws.GetClusterReply"/>

### GetClusterReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Is the cluster in the system |
| cluster | [ClusterDetailItem](#cmaaws.ClusterDetailItem) |  |  |






<a name="cmaaws.GetVersionMsg"/>

### GetVersionMsg
Get version of API Server






<a name="cmaaws.GetVersionReply"/>

### GetVersionReply
Reply for version request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | If operation was OK |
| version_information | [GetVersionReply.VersionInformation](#cmaaws.GetVersionReply.VersionInformation) |  | Version Information |






<a name="cmaaws.GetVersionReply.VersionInformation"/>

### GetVersionReply.VersionInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| git_version | [string](#string) |  | The tag on the git repository |
| git_commit | [string](#string) |  | The hash of the git commit |
| git_tree_state | [string](#string) |  | Whether or not the tree was clean when built |
| build_date | [string](#string) |  | Date of build |
| go_version | [string](#string) |  | Version of go used to compile |
| compiler | [string](#string) |  | Compiler used |
| platform | [string](#string) |  | Platform it was compiled for / running on |





 


<a name="cmaaws.ClusterStatus"/>

### ClusterStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | Not set |
| PROVISIONING | 1 | The PROVISIONING state indicates the cluster is being created. |
| RUNNING | 2 | The RUNNING state indicates the cluster has been created and is fully usable. |
| RECONCILING | 3 | The RECONCILING state indicates that some work is actively being done on the cluster, such as upgrading the master or node software. |
| STOPPING | 4 | The STOPPING state indicates the cluster is being deleted |
| ERROR | 5 | The ERROR state indicates the cluster may be unusable |
| DEGRADED | 6 | The DEGRADED state indicates the cluster requires user action to restore full functionality |


 

 


<a name="cmaaws.Cluster"/>

### Cluster


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCluster | [CreateClusterMsg](#cmaaws.CreateClusterMsg) | [CreateClusterReply](#cmaaws.CreateClusterMsg) | Will provision a cluster |
| GetCluster | [GetClusterMsg](#cmaaws.GetClusterMsg) | [GetClusterReply](#cmaaws.GetClusterMsg) | Will retrieve the status of a cluster and its kubeconfig for connectivity |
| DeleteCluster | [DeleteClusterMsg](#cmaaws.DeleteClusterMsg) | [DeleteClusterReply](#cmaaws.DeleteClusterMsg) | Will delete a cluster |
| GetClusterList | [GetClusterListMsg](#cmaaws.GetClusterListMsg) | [GetClusterListReply](#cmaaws.GetClusterListMsg) | Will retrieve a list of clusters |
| GetVersionInformation | [GetVersionMsg](#cmaaws.GetVersionMsg) | [GetVersionReply](#cmaaws.GetVersionMsg) | Will return version information about api server |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

