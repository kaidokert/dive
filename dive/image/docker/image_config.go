package docker

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type imageConfig struct {
	History []imageHistoryEntry `json:"history"`
	RootFs  rootFs              `json:"rootfs"`
}

type rootFs struct {
	Type    string   `json:"type"`
	DiffIds []string `json:"diff_ids"`
}

func newDockerImageConfig(configBytes []byte) imageConfig {
	var imageConfig imageConfig
	err := json.Unmarshal(configBytes, &imageConfig)
	if err != nil {
		logrus.Panic(err)
	}

	layerIdx := 0
	for idx := range imageConfig.History {
		if imageConfig.History[idx].EmptyLayer {
			imageConfig.History[idx].ID = "<missing>"
		} else {
			imageConfig.History[idx].ID = imageConfig.RootFs.DiffIds[layerIdx]
			layerIdx++
		}
	}

	return imageConfig
}
