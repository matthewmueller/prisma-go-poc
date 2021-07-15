package dmmf

type State struct {
	Params *Params `json:"params"`
}

func (s *State) Conditions() (conds []*Condition) {
	prismas := s.Params.Dmmf.Schema.InputObjectTypes.Prisma
	for _, prisma := range prismas {
		conds = append(conds, &Condition{prisma})
	}
	return conds
}

type Condition struct {
	prisma *Prisma
}

func (c *Condition) Field(name string) *VField {
	for _, field := range c.prisma.Fields {
		if name == field.Name {
			return &VField{c.prisma, field}
		}
	}
	return nil
}

func (c *Condition) Name() string {
	return c.prisma.Name
}

type VField struct {
	prisma *Prisma
	field  *Field
}

func (vf *VField) Name() string {
	return vf.field.Name
}

type Params struct {
	Dmmf *DMMF `json:"dmmf"`
}

type DMMF struct {
	Datamodel *Datamodel `json:"datamodel"`
	Schema    *Schema    `json:"schema"`
}

type Datamodel struct {
}

type Schema struct {
	InputObjectTypes *InputObjectTypes `json:"inputObjectTypes"`
}

type InputObjectTypes struct {
	Prisma []*Prisma `json:"prisma"`
}

type Prisma struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	Name       string       `json:"name"`
	IsRequired bool         `json:"isRequired"`
	IsNullable bool         `json:"isNullable"`
	InputTypes []*InputType `json:"inputTypes"`
}

// func (f *Field) Model()

type InputType struct {
	Type      string `json:"type"`
	Namespace string `json:"namespace"`
	Location  string `json:"location"`
	IsList    bool   `json:"isList"`
}

// type A struct {
// 	Params struct {
// 		Datamodel   string `json:"datamodel"`
// 		Datasources []struct {
// 			Name           string   `json:"name"`
// 			Provider       []string `json:"provider"`
// 			ActiveProvider string   `json:"activeProvider"`
// 			URL            struct {
// 				FromEnvVar string      `json:"fromEnvVar"`
// 				Value      interface{} `json:"value"`
// 			} `json:"url"`
// 		} `json:"datasources"`
// 		Generator struct {
// 			Name     string `json:"name"`
// 			Provider struct {
// 				FromEnvVar interface{} `json:"fromEnvVar"`
// 				Value      string      `json:"value"`
// 			} `json:"provider"`
// 			Output struct {
// 				Value      string      `json:"value"`
// 				FromEnvVar interface{} `json:"fromEnvVar"`
// 			} `json:"output"`
// 			Config struct {
// 				Package           string `json:"package"`
// 				DisableGoBinaries string `json:"disableGoBinaries"`
// 			} `json:"config"`
// 			BinaryTargets   []interface{} `json:"binaryTargets"`
// 			PreviewFeatures []interface{} `json:"previewFeatures"`
// 			IsCustomOutput  bool          `json:"isCustomOutput"`
// 		} `json:"generator"`
// 		Dmmf struct {
// 			Datamodel struct {
// 				Enums  []interface{} `json:"enums"`
// 				Models []struct {
// 					Name       string      `json:"name"`
// 					IsEmbedded bool        `json:"isEmbedded"`
// 					DbName     interface{} `json:"dbName"`
// 					Fields     []struct {
// 						Name            string `json:"name"`
// 						Kind            string `json:"kind"`
// 						IsList          bool   `json:"isList"`
// 						IsRequired      bool   `json:"isRequired"`
// 						IsUnique        bool   `json:"isUnique"`
// 						IsID            bool   `json:"isId"`
// 						IsReadOnly      bool   `json:"isReadOnly"`
// 						Type            string `json:"type"`
// 						HasDefaultValue bool   `json:"hasDefaultValue"`
// 						Default         struct {
// 							Name string        `json:"name"`
// 							Args []interface{} `json:"args"`
// 						} `json:"default,omitempty"`
// 						IsGenerated bool `json:"isGenerated"`
// 						IsUpdatedAt bool `json:"isUpdatedAt"`
// 					} `json:"fields"`
// 					IsGenerated   bool          `json:"isGenerated"`
// 					IDFields      []interface{} `json:"idFields"`
// 					UniqueFields  []interface{} `json:"uniqueFields"`
// 					UniqueIndexes []interface{} `json:"uniqueIndexes"`
// 				} `json:"models"`
// 			} `json:"datamodel"`
// 			Schema struct {
// 				InputObjectTypes struct {
// 					Prisma []struct {
// 						Name        string `json:"name"`
// 						Constraints struct {
// 							MaxNumFields interface{} `json:"maxNumFields"`
// 							MinNumFields interface{} `json:"minNumFields"`
// 						} `json:"constraints"`
// 						Fields []struct {
// 							Name       string `json:"name"`
// 							IsRequired bool   `json:"isRequired"`
// 							IsNullable bool   `json:"isNullable"`
// 							InputTypes []struct {
// 								Type      string `json:"type"`
// 								Namespace string `json:"namespace"`
// 								Location  string `json:"location"`
// 								IsList    bool   `json:"isList"`
// 							} `json:"inputTypes"`
// 						} `json:"fields"`
// 					} `json:"prisma"`
// 				} `json:"inputObjectTypes"`
// 				OutputObjectTypes struct {
// 					Prisma []struct {
// 						Name   string `json:"name"`
// 						Fields []struct {
// 							Name string `json:"name"`
// 							Args []struct {
// 								Name       string `json:"name"`
// 								IsRequired bool   `json:"isRequired"`
// 								IsNullable bool   `json:"isNullable"`
// 								InputTypes []struct {
// 									Type      string `json:"type"`
// 									Namespace string `json:"namespace"`
// 									Location  string `json:"location"`
// 									IsList    bool   `json:"isList"`
// 								} `json:"inputTypes"`
// 							} `json:"args"`
// 							IsNullable bool `json:"isNullable"`
// 							OutputType struct {
// 								Type      string `json:"type"`
// 								Namespace string `json:"namespace"`
// 								Location  string `json:"location"`
// 								IsList    bool   `json:"isList"`
// 							} `json:"outputType"`
// 						} `json:"fields"`
// 					} `json:"prisma"`
// 					Model []struct {
// 						Name   string `json:"name"`
// 						Fields []struct {
// 							Name       string        `json:"name"`
// 							Args       []interface{} `json:"args"`
// 							IsNullable bool          `json:"isNullable"`
// 							OutputType struct {
// 								Type     string `json:"type"`
// 								Location string `json:"location"`
// 								IsList   bool   `json:"isList"`
// 							} `json:"outputType"`
// 						} `json:"fields"`
// 					} `json:"model"`
// 				} `json:"outputObjectTypes"`
// 				EnumTypes struct {
// 					Prisma []struct {
// 						Name   string   `json:"name"`
// 						Values []string `json:"values"`
// 					} `json:"prisma"`
// 				} `json:"enumTypes"`
// 			} `json:"schema"`
// 			Mappings struct {
// 				ModelOperations []struct {
// 					Model      string `json:"model"`
// 					Aggregate  string `json:"aggregate"`
// 					CreateOne  string `json:"createOne"`
// 					DeleteMany string `json:"deleteMany"`
// 					DeleteOne  string `json:"deleteOne"`
// 					FindFirst  string `json:"findFirst"`
// 					FindMany   string `json:"findMany"`
// 					FindUnique string `json:"findUnique"`
// 					GroupBy    string `json:"groupBy"`
// 					UpdateMany string `json:"updateMany"`
// 					UpdateOne  string `json:"updateOne"`
// 					UpsertOne  string `json:"upsertOne"`
// 				} `json:"modelOperations"`
// 				OtherOperations struct {
// 					Read  []interface{} `json:"read"`
// 					Write []string      `json:"write"`
// 				} `json:"otherOperations"`
// 			} `json:"mappings"`
// 		} `json:"dmmf"`
// 		OtherGenerators []interface{} `json:"otherGenerators"`
// 		SchemaPath      string        `json:"schemaPath"`
// 		Version         string        `json:"version"`
// 	} `json:"params"`
// 	ID int `json:"id"`
// }
