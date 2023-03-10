package entrefine

import (
	"entgo.io/contrib/entgql"
	_ "entgo.io/contrib/entgql"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/diazoxide/entrefine/common"
	"github.com/vektah/gqlparser/v2/ast"
)

// EntgqlExtensionOptionsWrapper Wrap options of EntGQL
// Sometimes entrefine using some custom GraphQL queries and mutations
// Wrapper can add custom configs/hooks
func EntgqlExtensionOptionsWrapper(opts ...entgql.ExtensionOption) []entgql.ExtensionOption {
	plural := gen.Funcs["plural"].(func(string) string)

	return append(
		opts,
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(
			func(graph *gen.Graph, s *ast.Schema) error {
				for _, n := range graph.Nodes {

					name := common.LcFirst(plural(n.Name))
					f := s.Types["Query"].Fields.ForName(name)
					if f == nil {
						return fmt.Errorf("missing query field %q", name)
					}
					f.Arguments = append(f.Arguments, &ast.ArgumentDefinition{
						Description: "Search query",
						Name:        "q",
						Type:        ast.NamedType("String", nil),
					})

				}
				return nil

			},
		),
	)
}
