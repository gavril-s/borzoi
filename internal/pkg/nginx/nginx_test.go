package nginx_test

import (
	"strings"
	"testing"

	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/gavril-s/borzoi/internal/pkg/nginx"
	"github.com/stretchr/testify/require"
)

func TestBuildNginxConfig(t *testing.T) {
	t.Parallel()

	expected := `upstream service0 {
	server 127.0.0.1:0;
	server 127.0.0.1:1;
}

upstream service1 {
	server 127.0.0.1:2;
	server 127.0.0.1:3;
}

server {
    listen 5000;
    server_name url www.url;

    location / {
        proxy_pass http://service0;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

server {
	listen 5001;
    server_name url www.url;

	location / {
    	proxy_pass http://service1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
`
	expected = strings.ReplaceAll(expected, "    ", "\t") // 4 spaces

	deploy := domain.Deploy{
		URL: "url",
		Services: []domain.Service{
			{
				Config: domain.ServiceConfig{
					Name:         "service0",
					ExternalPort: 5000,
				},
				Nodes: []domain.Node{
					{
						Port: 0,
					},
					{
						Port: 1,
					},
				},
			},
			{
				Config: domain.ServiceConfig{
					Name:         "service1",
					ExternalPort: 5001,
				},
				Nodes: []domain.Node{
					{
						Port: 2,
					},
					{
						Port: 3,
					},
				},
			},
		},
	}
	actual := nginx.BuildNginxConfig(deploy)
	require.Equal(t, expected, actual)
}
