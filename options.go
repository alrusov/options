/*
Package options implements a support for the command-line options
*/

package options

import (
	"flag"
	"go/types"
)

//----------------------------------------------------------------------------------------------------------------------------//

// OptionDef -- command line option definition
type OptionDef struct {
	Name     string
	Comment  string
	Tp       types.BasicKind
	DefValue interface{}
}

type fullOptionDef struct {
	OptionDef
	pValue interface{}
}

var options = make(map[string]fullOptionDef)

//----------------------------------------------------------------------------------------------------------------------------//

// LoadOptions -- load options from command line
func LoadOptions(list map[string]OptionDef) {
	for name, srcDef := range list {
		def := fullOptionDef{srcDef, nil}

		switch def.Tp {
		case types.String:
			v := new(string)
			def.pValue = v
			flag.StringVar(v, def.Name, def.DefValue.(string), def.Comment)
		case types.Int:
			v := new(int)
			def.pValue = v
			flag.IntVar(v, def.Name, def.DefValue.(int), def.Comment)
		}

		options[name] = def
	}

	flag.Parse()
}

// GetOption -- get option from command line
func GetOption(name string) interface{} {
	var ret interface{}
	def := options[name]

	switch def.Tp {
	case types.String:
		ret = *def.pValue.(*string)
	case types.Int:
		ret = *def.pValue.(*int)
	}

	return ret
}

//----------------------------------------------------------------------------------------------------------------------------//
