package kubetype

import "fmt"

type PodNetwork struct {
	PodIP    string
	Port     string
	Protocol string
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
