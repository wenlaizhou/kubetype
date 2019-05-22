package kubetype

import (
	"fmt"
)

type PodNetwork struct {
	PodIP    string
	Port     string
	Protocol string
}

type ContainerResource struct {
	Name    string
	Limit   map[string]string
	Request map[string]string
}

type NodeNetwork struct {
	Hostname    string
	ExternalIP  string
	InternalIP  string
	ExternalDNS string
	InternalDNS string
}

// 获取pod网络信息
func GetPodNetwork(pod Pod) []PodNetwork {
	var res []PodNetwork
	podIp := pod.Status.PodIP
	for _, container := range pod.Spec.Containers {
		if len(container.Ports) > 0 {
			for _, port := range container.Ports {
				network := PodNetwork{
					PodIP:    podIp,
					Port:     fmt.Sprintf("%v", port.ContainerPort),
					Protocol: fmt.Sprintf("%v", port.Protocol),
				}
				res = append(res, network)
			}
		}
	}
	return res
}

// 获取pod资源信息
func GetPodResource(pod Pod) []ContainerResource {
	var res []ContainerResource
	for _, container := range pod.Spec.Containers {
		resource := ContainerResource{
			Name:    container.Name,
			Limit:   map[string]string{},
			Request: map[string]string{},
		}

		for k, v := range container.Resources.Limits {
			resource.Limit[fmt.Sprintf("%v", k)] = v
		}

		for k, v := range container.Resources.Requests {
			resource.Request[fmt.Sprintf("%v", k)] = v
		}
		res = append(res, resource)
	}
	return res
}

// 获取node网络信息
func GetNodeNetwork(node Node) NodeNetwork {
	res := NodeNetwork{}
	for _, addr := range node.Status.Addresses {
		switch addr.Type {

		case NodeHostName:
			res.Hostname = addr.Address
			break

		case NodeInternalIP:
			res.InternalIP = addr.Address
			break

		case NodeExternalIP:
			res.ExternalIP = addr.Address
			break

		case NodeInternalDNS:
			res.InternalDNS = addr.Address
			break

		case NodeExternalDNS:
			res.ExternalDNS = addr.Address
			break

		}
	}
	return res
}
