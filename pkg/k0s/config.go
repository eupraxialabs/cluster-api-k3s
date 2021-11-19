package k0s

import (
	bootstrapv1 "github.com/zawachte-msft/cluster-api-k0s/bootstrap/api/v1alpha3"
)

const DefaultK0sConfigLocation = "/etc/rancher/k0s/config.yaml"

type K0sServerConfig struct {
	DisableCloudController    bool     `json:"disable-cloud-controller,omitempty"`
	KubeAPIServerArgs         []string `json:"kube-apiserver-arg,omitempty"`
	KubeControllerManagerArgs []string `json:"kube-controller-manager-arg,omitempty"`
	TLSSan                    []string `json:"tls-san,omitempty"`
	BindAddress               string   `json:"bind-address,omitempty"`
	HttpsListenPort           string   `json:"https-listen-port,omitempty"`
	AdvertiseAddress          string   `json:"advertise-address,omitempty"`
	AdvertisePort             string   `json:"advertise-port,omitempty"`
	ClusterCidr               string   `json:"cluster-cidr,omitempty"`
	ServiceCidr               string   `json:"service-cidr,omitempty"`
	ClusterDNS                string   `json:"cluster-dns,omitempty"`
	ClusterDomain             string   `json:"cluster-domain,omitempty"`
	DisableComponents         []string `json:"disable,omitempty"`
	ClusterInit               bool     `json:"cluster-init,omitempty"`
	K0sAgentConfig            `json:",inline"`
}

type K0sAgentConfig struct {
	Token           string   `json:"token,omitempty"`
	Server          string   `json:"server,omitempty"`
	KubeletArgs     []string `json:"kubelet-arg,omitempty"`
	NodeLabels      []string `json:"node-labels,omitempty"`
	NodeTaints      []string `json:"node-taints,omitempty"`
	PrivateRegistry string   `json:"private-registry,omitempty"`
	KubeProxyArgs   []string `json:"kube-proxy-arg,omitempty"`
}

func GenerateInitControlPlaneConfig(controlPlaneEndpoint string, token string, serverConfig bootstrapv1.KZerosServerConfig, agentConfig bootstrapv1.KZerosAgentConfig) K0sServerConfig {
	k0sServerConfig := K0sServerConfig{
		DisableCloudController:    true,
		ClusterInit:               true,
		KubeAPIServerArgs:         append(serverConfig.KubeAPIServerArgs, "anonymous-auth=true"),
		TLSSan:                    append(serverConfig.TLSSan, controlPlaneEndpoint),
		KubeControllerManagerArgs: append(serverConfig.KubeControllerManagerArgs, "cloud-provider=external"),
		BindAddress:               serverConfig.BindAddress,
		HttpsListenPort:           serverConfig.HttpsListenPort,
		AdvertiseAddress:          serverConfig.AdvertiseAddress,
		AdvertisePort:             serverConfig.AdvertisePort,
		ClusterCidr:               serverConfig.ClusterCidr,
		ServiceCidr:               serverConfig.ServiceCidr,
		ClusterDNS:                serverConfig.ClusterDNS,
		ClusterDomain:             serverConfig.ClusterDomain,
		DisableComponents:         serverConfig.DisableComponents,
	}

	k0sServerConfig.K0sAgentConfig = K0sAgentConfig{
		Token:           token,
		KubeletArgs:     append(agentConfig.KubeletArgs, "cloud-provider=external"),
		NodeLabels:      agentConfig.NodeLabels,
		NodeTaints:      agentConfig.NodeTaints,
		PrivateRegistry: agentConfig.PrivateRegistry,
		KubeProxyArgs:   agentConfig.KubeProxyArgs,
	}

	return k0sServerConfig
}

func GenerateJoinControlPlaneConfig(serverUrl string, token string, controlplaneendpoint string, serverConfig bootstrapv1.KZerosServerConfig, agentConfig bootstrapv1.KZerosAgentConfig) K0sServerConfig {

	k0sServerConfig := K0sServerConfig{
		DisableCloudController:    true,
		KubeAPIServerArgs:         append(serverConfig.KubeAPIServerArgs, "anonymous-auth=true"),
		TLSSan:                    append(serverConfig.TLSSan, controlplaneendpoint),
		KubeControllerManagerArgs: append(serverConfig.KubeControllerManagerArgs, "cloud-provider=external"),
		BindAddress:               serverConfig.BindAddress,
		HttpsListenPort:           serverConfig.HttpsListenPort,
		AdvertiseAddress:          serverConfig.AdvertiseAddress,
		AdvertisePort:             serverConfig.AdvertisePort,
		ClusterCidr:               serverConfig.ClusterCidr,
		ServiceCidr:               serverConfig.ServiceCidr,
		ClusterDNS:                serverConfig.ClusterDNS,
		ClusterDomain:             serverConfig.ClusterDomain,
		DisableComponents:         serverConfig.DisableComponents,
	}

	k0sServerConfig.K0sAgentConfig = K0sAgentConfig{
		Token:           token,
		Server:          serverUrl,
		KubeletArgs:     append(agentConfig.KubeletArgs, "cloud-provider=external"),
		NodeLabels:      agentConfig.NodeLabels,
		NodeTaints:      agentConfig.NodeTaints,
		PrivateRegistry: agentConfig.PrivateRegistry,
		KubeProxyArgs:   agentConfig.KubeProxyArgs,
	}

	return k0sServerConfig
}

func GenerateWorkerConfig(serverUrl string, token string, agentConfig bootstrapv1.KZerosAgentConfig) K0sAgentConfig {
	return K0sAgentConfig{
		Server:          serverUrl,
		Token:           token,
		KubeletArgs:     append(agentConfig.KubeletArgs, "cloud-provider=external"),
		NodeLabels:      agentConfig.NodeLabels,
		NodeTaints:      agentConfig.NodeTaints,
		PrivateRegistry: agentConfig.PrivateRegistry,
		KubeProxyArgs:   agentConfig.KubeProxyArgs,
	}
}
