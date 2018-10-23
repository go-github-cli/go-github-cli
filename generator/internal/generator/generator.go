package generator

import (
	"fmt"
	"github.com/WillAbides/go-github-cli/generator/internal"
	"github.com/WillAbides/go-github-cli/generator/internal/configparser"
	"github.com/WillAbides/go-github-cli/generator/internal/packagewriter"
	"github.com/WillAbides/go-github-cli/generator/internal/routeparser"
	"github.com/fatih/camelcase"
	"github.com/fatih/structtag"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"reflect"
	"sort"
	"strings"
)

type (
	// Svc represents a group of api endpoints such as Issues, Organizations or Git
	Svc struct {
		Name     string
		Commands []cmd
	}

	// cmd represents a cli command that will be generated
	cmd struct {
		Name     string
		ArgNames []string
		Route    *routeparser.Route
	}

	// pkg represents the go package that will be created for a svc
	pkg struct {
		PackageName   string
		Imports       []string
		CmdHelpers    []*structTmplHelper
		OptionStructs []optionsStruct
	}

	runMethodArgHelper struct {
		Name  string
		IsPtr bool
	}

	runMethodHelper struct {
		StructName string
		HasElement bool
		SvcName    string
		FuncName   string
		Args       []runMethodArgHelper
	}

	structTmplHelper struct {
		Name           string
		Fields         []structField
		RunMethod      *runMethodHelper
		OptionsStructs []optionsStruct
	}

	structField struct {
		Name string
		Type string
		Tags *structtag.Tags
	}

	optionsStruct struct {
		StructName string
		MainStruct structTmplHelper
		ToFunc     toFuncTmplHelper
	}

	valSetter struct {
		TargetIsPtr bool
		Name        string
		FlagName    string
	}

	toFuncTmplHelper struct {
		ReceiverName         string
		TargetName           string
		TargetType           string
		ValSetters           []valSetter
		IncludePointerHelper bool
	}
)

var (
	pkgImports = []string{
		"context",
		"encoding/json",
		"github.com/alecthomas/kong",
		"github.com/google/go-github/github",
		"golang.org/x/oauth2",
		"time",
	}
)

//BuildSvcs builds services
func BuildSvcs(routesPath, configFile string) ([]Svc, error) {
	rt, err := routeparser.ParseRoutesFile(routesPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parsing routes at %q", routesPath)
	}
	hConfig, err := configparser.ParseConfigFile(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parsing config file at %q", configFile)
	}

	var svcs []Svc

	for svcName, hsvc := range hConfig.Service {
		svcRoutesName := hsvc.RouteName
		if svcRoutesName == "" {
			svcRoutesName = flagName(svcName)
		}
		var cmdNames []string
		for cmdName := range hsvc.Command {
			cmdNames = append(cmdNames, cmdName)
		}
		sort.Strings(cmdNames)
		cmds := make([]cmd, len(cmdNames))
		for i, cmdName := range cmdNames {
			hcmd := hsvc.Command[cmdName]
			routesName := hcmd.RoutesName
			if routesName == "" {
				routesName = flagName(cmdName)
			}
			routes := rt[svcRoutesName]
			route := routes.FindByIDName(routesName)
			cmds[i] = newCmd(cmdName, route, hcmd.ArgNames...)
		}
		svcs = append(svcs, Svc{
			Name:     svcName,
			Commands: cmds,
		})
	}
	return svcs, nil
}

//WriteOutput writes services output to the given path
func WriteOutput(outputPath string, svcs []Svc) error {
	for _, s := range svcs {
		p, err := s.pkg()
		if err != nil {
			return errors.Wrap(err, "")
		}
		err = packagewriter.WritePackageFiles(outputPath, p.PackageName, p)
		if err != nil {
			return errors.Wrap(err, "")
		}
	}
	return nil
}

func newCmd(name string, rt *routeparser.Route, argNames ...string) cmd {
	if len(argNames) == 0 {
		for _, p := range rt.Params {
			if p.Location == "url" {
				argNames = append(argNames, internal.ToArgName(p.Name))
			}
		}
	}

	return cmd{
		Name:     name,
		ArgNames: argNames,
		Route:    rt,
	}
}

func (c *cmd) tags() (*structtag.Tags, error) {
	tags := &structtag.Tags{}
	err := tags.Set(newTag("cmd", ""))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	err = tags.Set(&structtag.Tag{Key: "help", Name: c.Route.Name})
	return tags, errors.Wrap(err, "")
}

func newStructField(name, ftype string, tgs *structtag.Tags) *structField {
	return &structField{
		Name: name,
		Type: ftype,
		Tags: tgs,
	}
}

func newTag(key, name string, options ...string) *structtag.Tag {
	return &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	}
}

func newTags(tags ...*structtag.Tag) (*structtag.Tags, error) {
	tgs := &structtag.Tags{}
	for _, tag := range tags {
		err := tgs.Set(tag)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
	}
	return tgs, nil
}

func newSvcCmd(svcName string, cmds []cmd) (*structTmplHelper, error) {
	var fields []structField
	for _, cmd := range cmds {
		tgs, err := cmd.tags()
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		fields = append(fields, structField{
			Name: cmd.Name,
			Type: svcName + cmd.Name + "Cmd",
			Tags: tgs,
		})
	}
	return &structTmplHelper{
		Name:   svcName + "Cmd",
		Fields: fields,
	}, nil
}

func (s *Svc) getStructField() (reflect.StructField, bool) {
	clientType := reflect.TypeOf(github.Client{})
	return clientType.FieldByName(s.Name)
}

func (s *Svc) pkg() (*pkg, error) {
	field, ok := s.getStructField()
	if !ok {
		return nil, errors.New("can't find structField")
	}

	sc, err := newSvcCmd(s.Name, s.Commands)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cmdHelpers := []*structTmplHelper{
		sc,
	}

	for _, c := range s.Commands {
		method, ok := field.Type.MethodByName(c.Name)
		if !ok {
			return nil, errors.New("can't find method")
		}
		cmdHelper, err := buildCommandStruct(s.Name, method.Name, method, c)
		if err != nil {
			return nil, errors.Wrap(err, "creating cmdHelper")
		}
		cmdHelpers = append(cmdHelpers, cmdHelper)
	}

	return &pkg{
		PackageName: strings.ToLower(s.Name + "svc"),
		Imports:     pkgImports,
		CmdHelpers:  cmdHelpers,
	}, nil
}

func flagName(fieldName string) string {
	s := strings.ToLower(strings.Join(camelcase.Split(fieldName), "-"))
	return strings.Replace(s, "_", "-", -1)
}

func funcInTypes(funcType reflect.Type, offset int) []reflect.Type {
	var types []reflect.Type
	for i := offset; i < funcType.NumIn(); i++ {
		types = append(types, funcType.In(i))
	}
	return types
}

func commonStructFields() ([]structField, error) {
	var structFields []structField
	for _, sf := range []struct {
		Name string
		Type string
		Tags []*structtag.Tag
	}{
		{
			Name: "Token",
			Type: "string",
			Tags: []*structtag.Tag{
				newTag("env", "GITHUB_TOKEN"),
				newTag("required", ""),
			},
		},
		{
			Name: "APIBaseURL",
			Type: "string",
			Tags: []*structtag.Tag{
				newTag("env", "GITHUB_API_BASE_URL"),
				newTag("default", "https://api.github.com"),
			},
		},
	} {
		tgs, err := newTags(sf.Tags...)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		structFields = append(structFields, structField{
			Name: sf.Name,
			Type: sf.Type,
			Tags: tgs,
		})
	}
	return structFields, nil
}

func buildCommandStruct(svcName, funcName string, apiFunc reflect.Method, c cmd) (*structTmplHelper, error) {
	structName := svcName + funcName + "Cmd"
	argNames := c.ArgNames
	for i, argName := range argNames {
		argNames[i] = internal.ToArgName(argName)
	}
	fullArgNames := argNames

	structFields, err := commonStructFields()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var oss []optionsStruct

	inTypes := funcInTypes(apiFunc.Type, 2)
	for _, inType := range inTypes {
		switch inType.Kind() {
		case reflect.Ptr:
			if inType.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("only pointers to structs are allowed")
			}
			oStruct, err := generateOptionsStruct(structName, inType.Elem(), c.Route)
			if err != nil {
				return nil, errors.Wrap(err, "")
			}
			oss = append(oss, *oStruct)
			field := structField{
				Type: oStruct.StructName,
			}
			structFields = append(structFields, field)
		case reflect.String, reflect.Int, reflect.Bool:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			param := c.Route.ArgParam(argName)
			field := newStructField(argName, inType.Kind().String(), &structtag.Tags{})
			if param != nil && param.Required {
				err := field.Tags.Set(newTag("required", ""))
				if err != nil {
					return nil, errors.Wrap(err, "")
				}
			}
			structFields = append(structFields, *field)
		case reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			param := c.Route.ArgParam(argName)
			field := newStructField(argName, inType.String(), &structtag.Tags{})
			if param.Required {
				err := field.Tags.Set(newTag("required", ""))
				if err != nil {
					return nil, errors.Wrap(err, "")
				}
			}
			structFields = append(structFields, *field)
		default:
			return nil, fmt.Errorf(`buildCommandStruct: unsupported type: "%s"`, inType.Kind().String())
		}
	}
	runMethod, err := generateRunMethod(svcName, funcName, apiFunc, fullArgNames...)
	if err != nil {
		return nil, err
	}
	return &structTmplHelper{
		Name:           structName,
		Fields:         structFields,
		RunMethod:      runMethod,
		OptionsStructs: oss,
	}, nil
}

func generateRunMethod(svcName, funcName string, apiFunc reflect.Method, argNames ...string) (*runMethodHelper, error) {
	apiFuncType := apiFunc.Type
	numOut := apiFuncType.NumOut()
	structName := svcName + funcName + "Cmd"

	runStruct := &runMethodHelper{
		StructName: structName,
		FuncName:   funcName,
		SvcName:    svcName,
	}

	switch numOut {
	case 3:
		runStruct.HasElement = true
	case 2:
	default:
		return nil, fmt.Errorf("we only take funcs that return 2 or 3 args, not %d", numOut)
	}

	inTypes := funcInTypes(apiFuncType, 2)
	for _, inType := range inTypes {
		switch inType.Kind() {
		case reflect.Ptr:
			if inType.Elem().Kind() != reflect.Struct {
				return nil, fmt.Errorf("only pointers to structs are allowed")
			}
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{Name: inType.Elem().Name(), IsPtr: true})
		case reflect.String, reflect.Int, reflect.Bool:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{Name: argName})
		case reflect.Slice:
			if len(argNames) < 1 {
				return nil, fmt.Errorf("not enough argNames")
			}
			argName := argNames[0]
			argNames = argNames[1:]
			runStruct.Args = append(runStruct.Args, runMethodArgHelper{
				Name: argName,
			})
		default:
			return nil, fmt.Errorf("generateRunMethod: unsupported type: %v", inType.Kind())
		}
	}

	return runStruct, nil
}

func optionsStructName(methodName string, requestType reflect.Type) string {
	return internal.Unexport(methodName, requestType.Name()+"Flags")
}

func typeToFields(inType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < inType.NumField(); i++ {
		field := inType.Field(i)
		if field.Anonymous {
			ft := field.Type
			if ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			fields = append(fields, typeToFields(ft)...)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}

func generateOptionsStruct(cmdName string, requestType reflect.Type, route *routeparser.Route) (*optionsStruct, error) {
	structName := optionsStructName(cmdName, requestType)

	fields := typeToFields(requestType)
	structFields, err := getStructFields(fields, route)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	var keeperFields []reflect.StructField
	for _, field := range fields {
		for _, sf := range structFields {
			if sf.Name == field.Name {
				keeperFields = append(keeperFields, field)
			}
		}
	}
	return &optionsStruct{
		StructName: structName,
		MainStruct: structTmplHelper{
			Name:   structName,
			Fields: structFields,
		},
		ToFunc: generateToRequestFunc(keeperFields, structName, requestType),
	}, err
}

func getStructFields(fields []reflect.StructField, route *routeparser.Route) ([]structField, error) {
	var structFields []structField
	for _, field := range fields {
		if field.Type.Kind() == reflect.Ptr {
			field.Type = field.Type.Elem()
		}
		tags, err := newTags(newTag("name", flagName(field.Name)))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		if route != nil {
			param := route.FieldParam(field)
			if param == nil {
				continue
			}
			if param.Location != "body" {
				continue
			}
			if param != nil {
				if param.Required {
					err := tags.Set(newTag("required", ""))
					if err != nil {
						return nil, errors.Wrap(err, "")
					}
				}

				if param.Description != "" {
					err := tags.Set(newTag("help", param.Description))
					if err != nil {
						return nil, errors.Wrap(err, "")
					}
				}
			}
		}
		sField := newStructField(field.Name, field.Type.String(), tags)
		structFields = append(structFields, *sField)

	}
	return structFields, nil
}

func fieldFlagName(f reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	var name string

	if jsonTag == "" {
		name = strings.Join(camelcase.Split(f.Name), "-")
	} else {
		name = strings.Split(jsonTag, ",")[0]
	}
	return strings.ToLower(strings.Replace(name, "_", "-", -1))
}

func generateToRequestFunc(fields []reflect.StructField, structName string, targetType reflect.Type) toFuncTmplHelper {
	inclPtrHelper := false
	for _, v := range fields {
		if v.Type.Kind() == reflect.Ptr {
			inclPtrHelper = true
			break
		}
	}

	var vss []valSetter
	for _, field := range fields {
		vss = append(vss, valSetter{
			Name:        field.Name,
			FlagName:    fieldFlagName(field),
			TargetIsPtr: field.Type.Kind() == reflect.Ptr,
		})
	}

	return toFuncTmplHelper{
		ReceiverName:         structName,
		TargetName:           targetType.Name(),
		TargetType:           targetType.String(),
		IncludePointerHelper: inclPtrHelper,
		ValSetters:           vss,
	}
}
