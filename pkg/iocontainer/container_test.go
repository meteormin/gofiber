package iocontainer

import (
	"reflect"
	"testing"
)

type AnyType struct {
	Name string
}

func TestContainerStruct_Bind(t *testing.T) {
	type fields struct {
		instances map[reflect.Type]interface{}
	}
	type args struct {
		keyType  interface{}
		resolver interface{}
	}

	var bindAnyType AnyType

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Bind",
			fields: fields{
				instances: make(map[reflect.Type]interface{}),
			},
			args: args{
				keyType: &bindAnyType,
				resolver: func() AnyType {
					return AnyType{
						Name: "AnyType",
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ContainerStruct{
				instances: tt.fields.instances,
			}
			w.Bind(tt.args.keyType, tt.args.resolver)
		})
	}
}

func TestContainerStruct_Instances(t *testing.T) {
	type fields struct {
		instances map[reflect.Type]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[reflect.Type]interface{}
	}{
		{
			name: "Instances",
			fields: fields{
				instances: make(map[reflect.Type]interface{}),
			},
			want: make(map[reflect.Type]interface{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ContainerStruct{
				instances: tt.fields.instances,
			}
			if got := w.Instances(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Instances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainerStruct_Resolve(t *testing.T) {
	type fields struct {
		instances map[reflect.Type]interface{}
	}
	type args struct {
		keyType interface{}
	}

	var bindAnyType *AnyType
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Resolve",
			fields: fields{
				instances: make(map[reflect.Type]interface{}),
			},
			args: args{
				keyType: &bindAnyType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ContainerStruct{
				instances: tt.fields.instances,
			}
			w.Bind(tt.args.keyType, func() *AnyType {
				return &AnyType{
					Name: "AnyType",
				}
			})

			if err := w.Resolve(tt.args.keyType); (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(bindAnyType.Name)
		})
	}
}

func TestContainerStruct_Singleton(t *testing.T) {
	type fields struct {
		instances map[reflect.Type]interface{}
	}
	type args struct {
		instance interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Singleton",
			fields: fields{
				instances: make(map[reflect.Type]interface{}),
			},
			args: args{
				instance: func() AnyType {
					return AnyType{
						Name: "AnyType",
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ContainerStruct{
				instances: tt.fields.instances,
			}
			w.Singleton(tt.args.instance)
		})
	}
}

func TestContainerStruct_call(t *testing.T) {
	type fields struct {
		instances map[reflect.Type]interface{}
	}
	type args struct {
		callable interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "call",
			fields: fields{
				instances: make(map[reflect.Type]interface{}),
			},
			args: args{
				callable: AnyType{},
			},
			want: AnyType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ContainerStruct{
				instances: tt.fields.instances,
			}
			if got := w.call(tt.args.callable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("call() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContainer(t *testing.T) {
	tests := []struct {
		name string
		want Container
	}{
		{
			name: "NewContainer",
			want: &ContainerStruct{
				instances: make(map[reflect.Type]interface{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContainer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
