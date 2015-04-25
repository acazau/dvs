package domain

import (
	"errors"
)

type Backend struct {
	ID       string `json:"Id"`
	Type     string `json:"Type"`
	Settings struct {
		Timeouts struct {
			Read         string `json:"Read"`
			Dial         string `json:"Dial"`
			Tlshandshake string `json:"TLSHandshake"`
		} `json:"Timeouts"`
		Keepalive struct {
			Period              string `json:"Period"`
			Maxidleconnsperhost int    `json:"MaxIdleConnsPerHost"`
		} `json:"KeepAlive"`
	} `json:"Settings"`
}

type IVulcandAPIClientManager interface {
	ListBackends(socketPath string) ([]*Backend, error)
}

type VulcandAPIClientManager struct {
	InjectedVulcandAPIClientManager IVulcandAPIClientManager
}

func (manager *VulcandAPIClientManager) ListBackends(apiUrl string) ([]*Backend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	backends, err := manager.InjectedVulcandAPIClientManager.ListBackends(apiUrl)

	return backends, err
}
