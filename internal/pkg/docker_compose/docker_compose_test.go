package dockercompose_test

import (
	"strings"
	"testing"

	dockercompose "github.com/gavril-s/borzoi/internal/pkg/docker_compose"
	"github.com/gavril-s/borzoi/internal/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestBuildDockerCompose(t *testing.T) {
	t.Parallel()

	expected := `services:
  service1-0:
    build:
      context: repo/s1
	  dockerfile: ./s1-build/Dockerfile
    environment:
	  - SERVICE1_URL=url:5000
	  - SERVICE2_URL=url:5001
    ports:
      - "0:80"
  service1-1:
    build:
      context: repo/s1
	  dockerfile: ./s1-build/Dockerfile
    environment:
	  - SERVICE1_URL=url:5000
	  - SERVICE2_URL=url:5001
    ports:
      - "1:80"
  service2-0:
    build:
      context: repo/s2
	  dockerfile: ./s2-build/Dockerfile
    environment:
	  - SERVICE1_URL=url:5000
	  - SERVICE2_URL=url:5001
    ports:
      - "2:80"
  service2-1:
    build:
      context: repo/s2
	  dockerfile: ./s2-build/Dockerfile
    environment:
	  - SERVICE1_URL=url:5000
	  - SERVICE2_URL=url:5001
    ports:
      - "3:80"
`
	expected = strings.ReplaceAll(expected, "\t", "    ") // 4 spaces

	deploy := domain.Deploy{
		URL: "url",
		Services: []domain.Service{
			{
				Config: domain.ServiceConfig{
					Name:           "service1",
					UpperName:      "SERVICE1",
					RootPath:       "./s1",
					DockerfilePath: "./s1-build/Dockerfile",
					InternalPort:   80,
					ExternalPort:   5000,
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
					Name:           "service2",
					UpperName:      "SERVICE2",
					RootPath:       "./s2",
					DockerfilePath: "./s2-build/Dockerfile",
					InternalPort:   80,
					ExternalPort:   5001,
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
	actual := dockercompose.BuildDockerCompose(deploy, "repo")
	require.Equal(t, expected, actual)
}
