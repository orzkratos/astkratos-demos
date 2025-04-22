package astkratos_demos

import (
	"math/rand/v2"
	"path/filepath"
	"testing"

	"github.com/orzkratos/astkratos"
	"github.com/orzkratos/demokratos/demo1kratos"
	"github.com/orzkratos/demokratos/demo2kratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
)

var projectPath string

func TestMain(m *testing.M) {
	const demo1 = "demo1"
	const demo2 = "demo2"
	choices := []string{demo1, demo2}
	switch choices[rand.IntN(len(choices))] {
	case demo1:
		projectPath = osmustexist.ROOT(demo1kratos.SourceRoot())
	case demo2:
		projectPath = osmustexist.ROOT(demo2kratos.SourceRoot())
	}
	zaplog.SUG.Debugln(must.Nice(projectPath))

	m.Run()
}

func TestListGrpcClients(t *testing.T) {
	definitions := astkratos.ListGrpcClients(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "GreeterClient", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

func TestListGrpcServers(t *testing.T) {
	definitions := astkratos.ListGrpcServers(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "GreeterServer", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

func TestListGrpcServices(t *testing.T) {
	definitions := astkratos.ListGrpcServices(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "Greeter", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

func TestListGrpcUnimplementedServers(t *testing.T) {
	definitions := astkratos.ListGrpcUnimplementedServers(filepath.Join(projectPath, "api"))
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "UnimplementedGreeterServer", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

func TestListStructsMap(t *testing.T) {
	structsMap := astkratos.ListStructsMap(filepath.Join(projectPath, "api/helloworld/v1/greeter.pb.go"))
	t.Log(len(structsMap))

	for name, definition := range structsMap {
		t.Log(name)

		t.Log(definition.Name, definition.StructCode)
	}

	require.Len(t, structsMap, 2)
	struct0, ok := structsMap["HelloRequest"]
	require.True(t, ok)
	require.NotNil(t, struct0)
	require.Equal(t, "HelloRequest", struct0.Name)

	require.Len(t, structsMap, 2)
	struct1, ok := structsMap["HelloReply"]
	require.True(t, ok)
	require.NotNil(t, struct1)
	require.Equal(t, "HelloReply", struct1.Name)
}
