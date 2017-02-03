// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tColumbo struct {}
var Columbo tColumbo


func (_ tColumbo) Search(
		s interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("Columbo.Search", args).Url
}

func (_ tColumbo) Confirm(
		s interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("Columbo.Confirm", args).Url
}

func (_ tColumbo) Result(
		s interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "s", s)
	return revel.MainRouter.Reverse("Columbo.Result", args).Url
}


type tConnectdb struct {}
var Connectdb tConnectdb


func (_ tConnectdb) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Connectdb.Index", args).Url
}


type tConnectdb2 struct {}
var Connectdb2 tConnectdb2


func (_ tConnectdb2) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Connectdb2.Index", args).Url
}


type tHello struct {}
var Hello tHello


func (_ tHello) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Hello.Index", args).Url
}

func (_ tHello) Greet(
		greeting string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "greeting", greeting)
	return revel.MainRouter.Reverse("Hello.Greet", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


