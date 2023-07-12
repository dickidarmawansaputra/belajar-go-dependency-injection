//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)

	// kenapa return nil, karna dalam block kode program ini akan diganti oleh google wire
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)

	// kenapa return nil, karna dalam block kode program ini akan diganti oleh google wire
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

// implementasi provider set untuk grouping provider
func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// implementasi binding interface
// error: "no provider found for github.com/dickidarmawansaputra/belajar-go-dependency-injection/simple.SayHello"
// solusi dengan binding interface
// contoh salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

var HelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

// contoh benar dengan binding interface
func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

// implementasi struct provider
func InitializedFooBar() *FooBar {
	// bisa cara ini
	// wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	// atau cara ini jika tidak mau menyebutkan satu-satu
	// tanda * jika ingin semuanya di inject
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

// implementasi binding value
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	// tanda * jika ingin semuanya di inject
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

// implementasi interface value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// implementasi struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

// implementasi cleanup func
// returnnya tambahkan func() (closure)
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
