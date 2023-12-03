package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"reflect"
	"testing"
)

var reader = rand.Reader
var privateKey *rsa.PrivateKey

func init() {
	var err error
	privateKey, err = rsa.GenerateKey(reader, 4096)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGeneratorStruct_Generate(t *testing.T) {
	type fields struct {
		PrivateKey *rsa.PrivateKey
		PublicKey  crypto.PublicKey
		Exp        int
	}
	type args struct {
		claims     jwt.Claims
		privateKey *rsa.PrivateKey
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "Generate",
			fields: fields{
				PrivateKey: privateKey,
				PublicKey:  privateKey.PublicKey,
				Exp:        0,
			},
			args: args{
				claims:     jwt.MapClaims{"test": "test"},
				privateKey: privateKey,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeneratorStruct{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				Exp:        tt.fields.Exp,
			}
			got, err := g.Generate(tt.args.claims, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratorStruct_GetExp(t *testing.T) {
	type fields struct {
		PrivateKey *rsa.PrivateKey
		PublicKey  crypto.PublicKey
		Exp        int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "GetExp",
			fields: fields{
				PrivateKey: privateKey,
				PublicKey:  privateKey.PublicKey,
				Exp:        3600,
			},
			want: 3600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeneratorStruct{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				Exp:        tt.fields.Exp,
			}
			if got := g.GetExp(); got != tt.want {
				t.Errorf("GetExp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratorStruct_GetPrivateKey(t *testing.T) {
	type fields struct {
		PrivateKey *rsa.PrivateKey
		PublicKey  crypto.PublicKey
		Exp        int
	}
	tests := []struct {
		name   string
		fields fields
		want   *rsa.PrivateKey
	}{
		{
			name: "GetPrivateKey",
			fields: fields{
				PrivateKey: privateKey,
				PublicKey:  privateKey.PublicKey,
				Exp:        3600,
			},
			want: privateKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GeneratorStruct{
				PrivateKey: tt.fields.PrivateKey,
				PublicKey:  tt.fields.PublicKey,
				Exp:        tt.fields.Exp,
			}
			if got := g.GetPrivateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGenerator(t *testing.T) {
	type args struct {
		priv *rsa.PrivateKey
		pub  crypto.PublicKey
		exp  int
	}
	tests := []struct {
		name string
		args args
		want Generator
	}{
		{
			name: "NewGenerator",
			args: args{
				priv: privateKey,
				pub:  privateKey.PublicKey,
				exp:  3600,
			},
			want: &GeneratorStruct{
				PrivateKey: privateKey,
				PublicKey:  privateKey.PublicKey,
				Exp:        3600,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenerator(tt.args.priv, tt.args.pub, tt.args.exp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
