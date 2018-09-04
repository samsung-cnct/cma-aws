package awsmodels

const NewVPCHeptioCFTemplate = `# Copyright 2017 by the contributors
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.

---
AWSTemplateFormatVersion: '2010-09-09'
Description: 'QS(5042) Kubernetes AWS CloudFormation Template: Create a Kubernetes
  cluster in a new VPC. The master node is an auto-recovering Amazon EC2
  instance. 1-20 additional EC2 instances in an AutoScalingGroup join the
  Kubernetes cluster as nodes. An ELB provides configurable external access
  to the Kubernetes API. The new VPC includes a bastion host to grant
  SSH access to the private subnet for the cluster. This template creates
  two stacks: one for the new VPC and one for the cluster. The stack is
  suitable for development and small single-team clusters. **WARNING** This
  template creates four Amazon EC2 instances with default settings. You will
  be billed for the AWS resources used if you create a stack from this template.
  **SUPPORT** Please visit http://jump.heptio.com/aws-qs-help for support.
  **NEXT STEPS** Please visit http://jump.heptio.com/aws-qs-next.'

# The Metadata tells AWS how to display the parameters during stack creation
Metadata:
  AWS::CloudFormation::Interface:
    ParameterGroups:
    - Label:
        default: Required
      Parameters:
      - AvailabilityZone
      - AdminIngressLocation
      - KeyName
    - Label:
        default: Advanced
      Parameters:
      - VPCCIDR
      - PrivateSubnetCIDR
      - PublicSubnetCIDR
      - ClusterDNSProvider
      - NetworkingProvider
      - K8sNodeCapacity
      - InstanceType
      - DiskSizeGb
      - BastionInstanceType
      - QSS3BucketName
      - QSS3KeyPrefix

    ParameterLabels:
      VPCCIDR:
        default: VPC CIDR Block
      PrivateSubnetCIDR:
        default: Private Subnet CIDR Block
      PublicSubnetCIDR:
        default: Public Subnet CIDR Block
      KeyName:
        default: SSH Key
      AvailabilityZone:
        default: Availability Zone
      AdminIngressLocation:
        default: Admin Ingress Location
      InstanceType:
        default: Instance Type
      DiskSizeGb:
        default: Disk Size (GiB)
      BastionInstanceType:
        default: Instance Type (Bastion Host)
      K8sNodeCapacity:
        default: Node Capacity
      QSS3BucketName:
        default: S3 Bucket
      QSS3KeyPrefix:
        default: S3 Key Prefix
      NetworkingProvider:
        default: Networking Provider
      ClusterDNSProvider:
        default: Cluster DNS Provider

# The Parameters allow the user to pass custom settings to the stack before creation
Parameters:
  VPCCIDR:
    Description: VPC CIDR Block
    Type: String
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.
    Default: '10.0.0.0/16'

  PrivateSubnetCIDR:
    Description: CIDR Block for the Private Subnet, must be a valid subnet of the VPC CIDR and not overlap with PublicSubnetCIDR
    Type: String
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.
    Default: '10.0.0.0/19'

  PublicSubnetCIDR:
    Description: CIDR Block for the Public Subnet, must be a valid subnet of the VPC CIDR and not overlap with PrivateSubnetCIDR
    Type: String
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.
    Default: '10.0.128.0/20'

  KeyName:
    Description: Existing EC2 KeyPair for SSH access.
    Type: AWS::EC2::KeyPair::KeyName
    ConstraintDescription: must be the name of an existing EC2 KeyPair.

  InstanceType:
    Description: EC2 instance type for the cluster.
    Type: String
    Default: m4.large
    AllowedValues:
    - m5.large
    - m5.xlarge
    - m5.2xlarge
    - m5.4xlarge
    - m5.12xlarge
    - m5.24xlarge
    - c5.large
    - c5.xlarge
    - c5.2xlarge
    - c5.4xlarge
    - c5.9xlarge
    - c5.18xlarge
    - r4.large
    - r4.xlarge
    - r4.2xlarge
    - r4.4xlarge
    - r4.8xlarge
    - r4.16xlarge
    - x1.16xlarge
    - x1.32xlarge
    - i3.large
    - i3.xlarge
    - i3.2xlarge
    - i3.4xlarge
    - i3.8xlarge
    - i3.16xlarge
    - p2.xlarge
    - p2.8xlarge
    - p2.16xlarge
    - p3.2xlarge
    - p3.8xlarge
    - p3.16xlarge
    - m4.large
    - m4.xlarge
    - m4.2xlarge
    - m4.4xlarge
    - m4.10xlarge
    - m4.16xlarge
    - c4.large
    - c4.xlarge
    - c4.2xlarge
    - c4.4xlarge
    - c4.8xlarge
    - g3.4xlarge
    - g3.8xlarge
    - g3.16xlarge
    - r3.large
    - r3.xlarge
    - r3.2xlarge
    - r3.4xlarge
    - r3.8xlarge
    ConstraintDescription: must be a valid Current Generation (non-burstable) EC2 instance type.

  # Specifies the size of the root disk for all EC2 instances, including master
  # and nodes.
  DiskSizeGb:
    Description: 'Size of the root disk for the EC2 instances, in GiB.  Default: 40'
    Default: 40
    Type: Number
    MinValue: 8
    MaxValue: 1024

  BastionInstanceType:
    Description: EC2 instance type for the bastion host (used for public SSH access).
    Type: String
    Default: t2.micro
    AllowedValues:
    - t2.nano
    - t2.micro
    - t2.small
    - t2.medium
    - t2.large
    - t2.xlarge
    - t2.2xlarge
    - m5.large
    - m5.xlarge
    - m5.2xlarge
    - m5.4xlarge
    - m5.12xlarge
    - m5.24xlarge
    - c5.large
    - c5.xlarge
    - c5.2xlarge
    - c5.4xlarge
    - c5.9xlarge
    - c5.18xlarge
    - r4.large
    - r4.xlarge
    - r4.2xlarge
    - r4.4xlarge
    - r4.8xlarge
    - r4.16xlarge
    - x1.16xlarge
    - x1.32xlarge
    - i3.large
    - i3.xlarge
    - i3.2xlarge
    - i3.4xlarge
    - i3.8xlarge
    - i3.16xlarge
    - p2.xlarge
    - p2.8xlarge
    - p2.16xlarge
    - p3.2xlarge
    - p3.8xlarge
    - p3.16xlarge
    - m4.large
    - m4.xlarge
    - m4.2xlarge
    - m4.4xlarge
    - m4.10xlarge
    - m4.16xlarge
    - c4.large
    - c4.xlarge
    - c4.2xlarge
    - c4.4xlarge
    - c4.8xlarge
    - g3.4xlarge
    - g3.8xlarge
    - g3.16xlarge
    - r3.large
    - r3.xlarge
    - r3.2xlarge
    - r3.4xlarge
    - r3.8xlarge
    ConstraintDescription: must be a valid Current Generation EC2 instance type.

  AvailabilityZone:
    Description: The Availability Zone for this cluster. Heptio recommends
      that you run one cluster per AZ and use tooling to coordinate across AZs.
    Type: AWS::EC2::AvailabilityZone::Name
    ConstraintDescription: must be the name of an AWS Availability Zone

  AdminIngressLocation:
    Description: CIDR block (IP address range) to allow SSH access to the
      bastion host and HTTPS access to the Kubernetes API. Use 0.0.0.0/0
      to allow access from all locations.
    Type: String
    MinLength: '9'
    MaxLength: '18'
    AllowedPattern: "(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,2})"
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.

  K8sNodeCapacity:
    Default: '2'
    Description: Initial number of Kubernetes nodes (1-20).
    Type: Number
    MinValue: '1'
    MaxValue: '20'
    ConstraintDescription: must be between 1 and 20 EC2 instances.

  # S3 Bucket configuration: allows users to use their own downstream snapshots
  # of the quickstart-aws-vpc and quickstart-linux-bastion templates
  QSS3BucketName:
    AllowedPattern: "^[0-9a-zA-Z]+([0-9a-zA-Z-]*[0-9a-zA-Z])*$"
    ConstraintDescription: Quick Start bucket name can include numbers, lowercase
      letters, uppercase letters, and hyphens (-). It cannot start or end with a hyphen
      (-).

    Default: aws-quickstart
    Description: Only change this if you have set up assets, like your own networking
      configuration, in an S3 bucket. S3 bucket name for the Quick Start assets.
      Quick Start bucket name can include numbers, lowercase letters, uppercase
      letters, and hyphens (-). It cannot start or end with a hyphen (-).
    Type: String

  QSS3KeyPrefix:
    AllowedPattern: ^[0-9a-zA-Z-/]*$
    ConstraintDescription: Quick Start key prefix can include numbers, lowercase
      letters, uppercase letters, hyphens (-), and forward slash (/).
    Default: quickstart-heptio/
    Description: Only change this if you have set up assets in an S3 bucket, as explained
      in the S3 Bucket parameter. S3 key prefix for the Quick Start assets.
      Quick Start key prefix can include numbers, lowercase letters, uppercase
      letters, hyphens (-), and forward slash (/).
    Type: String

  NetworkingProvider:
    AllowedValues:
    - calico
    - weave
    ConstraintDescription: 'Currently supported values are "calico" and "weave"'
    Default: calico
    Description: Choose the networking provider to use for communication between
      pods in the Kubernetes cluster. Supported configurations are calico
      (https://docs.projectcalico.org/v2.6/getting-started/kubernetes/installation/hosted/kubeadm/)
      and weave (https://github.com/weaveworks/weave/blob/master/site/kubernetes/kube-addon.md).
    Type: String

  ClusterDNSProvider:
    AllowedValues:
    - CoreDNS
    - KubeDNS
    ConstraintDescription: 'Currently supported values are "CoreDNS" and "KubeDNS"'
    Default: CoreDNS
    Description: Choose the cluster DNS provider to use for internal cluster DNS. Supported
      configurations are CoreDNS and KubeDNS
    Type: String

Mappings:
  RegionMap:
    ap-northeast-1:
      '64': ami-d39a02b5
    ap-northeast-2:
      '64': ami-67973709
    ap-south-1:
      '64': ami-5d055232
    ap-southeast-1:
      '64': ami-325d2e4e
    ap-southeast-2:
      '64': ami-37df2255
    ca-central-1:
      '64': ami-f0870294
    eu-central-1:
      '64': ami-af79ebc0
    eu-west-1:
      '64': ami-4d46d534
    eu-west-2:
      '64': ami-d7aab2b3
    eu-west-3:
      '64': ami-5e0eb923
    sa-east-1:
      '64': ami-1157157d
    us-east-1:
      '64': ami-41e0b93b
    us-east-2:
      '64': ami-2581aa40
    us-west-1:
      '64': ami-79aeae19
    us-west-2:
      '64': ami-1ee65166
Conditions:
  UsEast1Condition:
    Fn::Equals:
    - !Ref AWS::Region
    - "us-east-1"

Resources:
  # Resources for new VPC
  VPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: !Ref VPCCIDR
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      Tags:
      - Key: Name
        Value: !Ref AWS::StackName

  DHCPOptions:
    Type: AWS::EC2::DHCPOptions
    Properties:
      DomainName:
        # us-east-1 needs .ec2.internal, the rest of the regions get <region>.compute.internal.
        # See http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html
        Fn::If:
        - UsEast1Condition
        - "ec2.internal"
        - !Sub "${AWS::Region}.compute.internal"
      DomainNameServers:
      - AmazonProvidedDNS

  VPCDHCPOptionsAssociation:
    Type: AWS::EC2::VPCDHCPOptionsAssociation
    Properties:
      VpcId: !Ref VPC
      DhcpOptionsId: !Ref DHCPOptions

  InternetGateway:
    Type: AWS::EC2::InternetGateway
    Properties:
      Tags:
      - Key: Network
        Value: Public

  VPCGatewayAttachment:
    Type: AWS::EC2::VPCGatewayAttachment
    Properties:
      VpcId: !Ref VPC
      InternetGatewayId: !Ref InternetGateway

  PrivateSubnet:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref VPC
      CidrBlock: !Ref PrivateSubnetCIDR
      AvailabilityZone: !Ref AvailabilityZone
      Tags:
      - Key: Name
        Value: Private subnet
      - Key: Network
        Value: Private

  PublicSubnet:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref VPC
      CidrBlock: !Ref PublicSubnetCIDR
      AvailabilityZone: !Ref AvailabilityZone
      Tags:
      - Key: Name
        Value: Public subnet
      - Key: Network
        Value: Public
      - Key: KubernetesCluster
        Value: !Ref AWS::StackName
      MapPublicIpOnLaunch: true

  # The NAT IP for the private subnet, as seen from within the public one
  NATEIP:
    DependsOn: VPCGatewayAttachment
    Type: AWS::EC2::EIP
    Properties:
      Domain: vpc

  # The NAT gateway for the private subnet
  NATGateway:
    DependsOn: VPCGatewayAttachment
    Type: AWS::EC2::NatGateway
    Properties:
      AllocationId: !GetAtt NATEIP.AllocationId
      SubnetId: !Ref PublicSubnet

  PrivateSubnetRouteTable:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId: !Ref VPC
      Tags:
      - Key: Name
        Value: Private subnets
      - Key: Network
        Value: Private

  PrivateSubnetRoute:
    DependsOn: VPCGatewayAttachment
    Type: AWS::EC2::Route
    Properties:
      RouteTableId: !Ref PrivateSubnetRouteTable
      DestinationCidrBlock: 0.0.0.0/0
      NatGatewayId: !Ref NATGateway

  PrivateSubnetRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref PrivateSubnet
      RouteTableId: !Ref PrivateSubnetRouteTable

  PublicSubnetRouteTable:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId: !Ref VPC
      Tags:
      - Key: Name
        Value: Public Subnets
      - Key: Network
        Value: Public

  PublicSubnetRoute:
    DependsOn: VPCGatewayAttachment
    Type: AWS::EC2::Route
    Properties:
      RouteTableId: !Ref PublicSubnetRouteTable
      DestinationCidrBlock: 0.0.0.0/0
      GatewayId: !Ref InternetGateway

  PublicSubnetRouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref PublicSubnet
      RouteTableId: !Ref PublicSubnetRouteTable

  # Taken from github.com/aws-quickstart/quickstart-linux-bastion.  We don't
  # call it directly because that quickstart forces 2 bastion hosts and we only
  # want one
  BastionHost:
    Type: AWS::EC2::Instance
    Properties:
      ImageId:
        Fn::FindInMap:
        - RegionMap
        - Ref: AWS::Region
        - '64'
      InstanceType: !Ref BastionInstanceType
      NetworkInterfaces:
      - AssociatePublicIpAddress: true
        DeleteOnTermination: true
        DeviceIndex: 0
        SubnetId: !Ref PublicSubnet
        # This address is chosen because our public subnet begins at 10.0.128.0/20
        PrivateIpAddress: '10.0.128.5'
        GroupSet:
        - Ref: BastionSecurityGroup
      Tags:
      - Key: Name
        Value: bastion-host
      KeyName: !Ref KeyName
      UserData:
        Fn::Base64:
          Fn::Sub: |
            #!/bin/bash

            BASTION_BOOTSTRAP_FILE=bastion_bootstrap.sh
            BASTION_BOOTSTRAP=https://s3.amazonaws.com/aws-quickstart/quickstart-linux-bastion/scripts/bastion_bootstrap.sh

            curl -s -o $BASTION_BOOTSTRAP_FILE $BASTION_BOOTSTRAP
            chmod +x $BASTION_BOOTSTRAP_FILE

            # This gets us far enough in the bastion script to be useful.
            apt-get -y update && apt-get -y install python-pip
            pip install --upgrade pip &> /dev/null

            ./$BASTION_BOOTSTRAP_FILE --banner https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}scripts/banner_message.txt --enable true

  # Open up port 22 for SSH for the bastion host
  BastionSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Enable SSH access via port 22
      VpcId: !Ref VPC
      SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: '22'
        ToPort: '22'
        CidrIp: !Ref AdminIngressLocation

  # Call the cluster template and supply its parameters
  # This creates a second stack that creates the actual Kubernetes cluster
  # within the new VPC
  K8sStack:
    Type: AWS::CloudFormation::Stack
    Properties:
      TemplateURL: !Sub "https://${QSS3BucketName}.s3.amazonaws.com/${QSS3KeyPrefix}templates/kubernetes-cluster.template"
      Parameters:
        VPCID: !Ref VPC
        AvailabilityZone: !Ref AvailabilityZone
        InstanceType: !Ref InstanceType
        DiskSizeGb: !Ref DiskSizeGb
        ClusterSubnetId: !Ref PrivateSubnet
        # Direct SSH access only from the bastion host itself
        SSHLocation: !Sub "${BastionHost.PrivateIp}/32"
        ApiLbLocation: !Ref AdminIngressLocation
        KeyName: !Ref KeyName
        K8sNodeCapacity: !Ref K8sNodeCapacity
        QSS3BucketName: !Ref QSS3BucketName
        QSS3KeyPrefix: !Ref QSS3KeyPrefix
        ClusterAssociation: !Ref AWS::StackName
        NetworkingProvider: !Ref NetworkingProvider
        LoadBalancerSubnetId: !Ref PublicSubnet
        ClusterDNSProvider: !Ref ClusterDNSProvider

Outputs:
  # Outputs from VPC creation
  VPCID:
    Description: ID of the newly-created EC2 VPC.
    Value: !Ref VPC

  BastionHostPublicIp:
    Description: IP Address of the bastion host for the newly-created EC2 VPC.
    Value: !GetAtt BastionHost.PublicIp

  BastionHostPublicDNS:
    Description: Public DNS FQDN of the bastion host for the newly-created EC2 VPC.
    Value: !GetAtt BastionHost.PublicDnsName

  SSHProxyCommand:
    Description: Run locally - SSH command to proxy to the master instance
      through the bastion host, to access port 8080 (command to SSH to the master Kubernetes node).
    Value: !Sub >-
      SSH_KEY="path/to/${KeyName}.pem";
      ssh
      -i $SSH_KEY
      -A -L8080:localhost:8080
      -o ProxyCommand="ssh -i \"${!SSH_KEY}\" ubuntu@${BastionHost.PublicIp} nc %h %p"
      ubuntu@${K8sStack.Outputs.MasterPrivateIp}

  GetKubeConfigCommand:
    Description: Run locally - SCP command to download the Kubernetes configuration
      file for accessing the new cluster via kubectl, a Kubernetes command line tool.
      Creates a "kubeconfig" file in the current directory. Then, you can run
      "export KUBECONFIG=$(pwd)/kubeconfig" to ensure kubectl uses this configuration file.
      About kubectl - https://kubernetes.io/docs/user-guide/prereqs/
    Value: !Sub >-
      SSH_KEY="path/to/${KeyName}.pem";
      scp
      -i $SSH_KEY
      -o ProxyCommand="ssh -i \"${!SSH_KEY}\" ubuntu@${BastionHost.PublicIp} nc %h %p"
      ubuntu@${K8sStack.Outputs.MasterPrivateIp}:~/kubeconfig ./kubeconfig

  # Outputs forwarded from the k8s template
  MasterInstanceId:
    Description: InstanceId of the master EC2 instance.
    Value: !GetAtt K8sStack.Outputs.MasterInstanceId

  MasterPrivateIp:
    Description: Private IP address of the master.
    Value: !GetAtt K8sStack.Outputs.MasterPrivateIp

  NodeGroupInstanceId:
    Description: InstanceId of the newly-created NodeGroup.
    Value: !GetAtt K8sStack.Outputs.NodeGroupInstanceId

  JoinNodes:
    Description: Command to join more nodes to this cluster.
    Value: !GetAtt K8sStack.Outputs.JoinNodes

  NextSteps:
    Description: Verify your cluster and deploy a test application. Instructions -
      http://jump.heptio.com/aws-qs-next
    Value: !GetAtt K8sStack.Outputs.NextSteps
`
