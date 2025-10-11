# Changes

Code differences compared to source project demokratos.

## internal/pkg/generate/ast_parse.go (+1 -0)

```diff
+ package generate
```

## internal/pkg/generate/ast_parse_test.go (+84 -0)

```diff
+ package generate_test
+ 
+ import (
+ 	"path/filepath"
+ 	"testing"
+ 
+ 	"github.com/orzkratos/astkratos"
+ 	"github.com/orzkratos/demokratos/demo2kratos"
+ 	"github.com/stretchr/testify/require"
+ 	"github.com/yyle88/must"
+ 	"github.com/yyle88/neatjson/neatjsons"
+ 	"github.com/yyle88/osexistpath/osmustexist"
+ 	"github.com/yyle88/zaplog"
+ )
+ 
+ var projectPath string
+ 
+ func TestMain(m *testing.M) {
+ 	projectPath = demo2kratos.SourceRoot()
+ 	must.Nice(projectPath)
+ 	osmustexist.ROOT(projectPath)
+ 	zaplog.SUG.Debugln(projectPath)
+ 
+ 	m.Run()
+ }
+ 
+ func TestListGrpcClients(t *testing.T) {
+ 	definitions := astkratos.ListGrpcClients(filepath.Join(projectPath, "api"))
+ 	t.Log(neatjsons.S(definitions))
+ 
+ 	require.Len(t, definitions, 1)
+ 	require.Equal(t, "GreeterClient", definitions[0].Name)
+ 	require.Equal(t, "v1", definitions[0].Package)
+ }
+ 
+ func TestListGrpcServers(t *testing.T) {
+ 	definitions := astkratos.ListGrpcServers(filepath.Join(projectPath, "api"))
+ 	t.Log(neatjsons.S(definitions))
+ 
+ 	require.Len(t, definitions, 1)
+ 	require.Equal(t, "GreeterServer", definitions[0].Name)
+ 	require.Equal(t, "v1", definitions[0].Package)
+ }
+ 
+ func TestListGrpcServices(t *testing.T) {
+ 	definitions := astkratos.ListGrpcServices(filepath.Join(projectPath, "api"))
+ 	t.Log(neatjsons.S(definitions))
+ 
+ 	require.Len(t, definitions, 1)
+ 	require.Equal(t, "Greeter", definitions[0].Name)
+ 	require.Equal(t, "v1", definitions[0].Package)
+ }
+ 
+ func TestListGrpcUnimplementedServers(t *testing.T) {
+ 	definitions := astkratos.ListGrpcUnimplementedServers(filepath.Join(projectPath, "api"))
+ 	t.Log(neatjsons.S(definitions))
+ 
+ 	require.Len(t, definitions, 1)
+ 	require.Equal(t, "UnimplementedGreeterServer", definitions[0].Name)
+ 	require.Equal(t, "v1", definitions[0].Package)
+ }
+ 
+ func TestListStructsMap(t *testing.T) {
+ 	structsMap := astkratos.ListStructsMap(filepath.Join(projectPath, "api/helloworld/v1/greeter.pb.go"))
+ 	t.Log(len(structsMap))
+ 
+ 	for name, definition := range structsMap {
+ 		t.Log(name)
+ 
+ 		t.Log(definition.Name, definition.StructCode)
+ 	}
+ 
+ 	require.Len(t, structsMap, 2)
+ 	struct0, ok := structsMap["HelloRequest"]
+ 	require.True(t, ok)
+ 	require.NotNil(t, struct0)
+ 	require.Equal(t, "HelloRequest", struct0.Name)
+ 
+ 	require.Len(t, structsMap, 2)
+ 	struct1, ok := structsMap["HelloReply"]
+ 	require.True(t, ok)
+ 	require.NotNil(t, struct1)
+ 	require.Equal(t, "HelloReply", struct1.Name)
+ }
```

