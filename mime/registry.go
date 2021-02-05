package mime

import "fmt"

type Registry string

const (
	Application Registry = "application"
	Audio       Registry = "audio"
	Font        Registry = "font"
	Example     Registry = "example"
	Image       Registry = "image"
	Message     Registry = "message"
	Model       Registry = "model"
	Multipart   Registry = "multipart"
	Text        Registry = "text"
	Video       Registry = "video"
)

var (
	registries = map[string]Registry{
		"application": Application,
		"audio":       Audio,
		"font":        Font,
		"example":     Example,
		"image":       Image,
		"message":     Message,
		"model":       Model,
		"multipart":   Multipart,
		"text":        Text,
		"video":       Video,
	}
)

func RegistryOf(s string) (*Registry, error) {
	r, ok := registries[s]
	if !ok {
		return nil, fmt.Errorf("unknown MIME-Type registry: %s", s)
	}

	return &r, nil
}
