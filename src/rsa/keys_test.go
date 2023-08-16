package rsa

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPKCS1PrivateKey(t *testing.T) {
	type args struct {
		privateKey []byte
	}
	tests := []struct {
		name string
		args args

		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "LoadPKCS1PrivateKey-err",
			args: args{
				privateKey: []byte("test"),
			},
			wantErr: assert.Error,
		},
		{
			name: "LoadPKCS1PrivateKey-OK",
			args: args{
				privateKey: []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAoB1eb1FolIkXconqXdc3J12W23KHWluZ/0LC6wVmWBRDztkY
+UmeH7vrKj9XpjCGur3Sieuk6rJKwdzaMzQshlpnnyRr6XyTxYBPL96WjjolWBCh
jclSkX6hg/YTljB2/h4TDp6iVIzk8j9bw1j5X4yEw1NdbIThgQEbxaAHho5eclqP
3YQGWYbx7Hxa54hsClsMZNgJKH4lTX3wuAjtEzeyYnSSl90zof0WFbnBoGU6XDkm
O5L5Aq8oi74L8D5yTAzU6eS3VLDIf3yYNfTqPYywKWPEIYaV4nfV/2v2NbqOR1yS
wv1y5iqIXNgA3K/83MzGFFGwrbTUQGkC7oc/5wIDAQABAoIBACZWYeItj+jg2mhm
dWN6wI4AbrqktZwBCuJ/zcQQSu0UDRheCwjFg8L9b9VxzT7Rp7DW+q5jad14S0YX
53cTKwYQZ3dHdT82wPstOciwd0Qe0ApCESyt48NKmsnKBe6UNCmsccuyWBP+mGMY
oJQpnawvLSrHrzropkYJSekL5EFnPGD1LIf+I+WY9fb4YTt6Hx4IHIZWrYd+Mxg2
hMQsMo0m6p3agJijYIfHntWqB0Ykzj72UtWSF5q+y2MlrG0dQKBvGVTkGGzjhxy7
nUm/1NCJPtxQf5+AYiFvzydjmg/5xOEIO5Ys90HrvdmWEPg9+Q/gg931cG7vhD1U
++rwEoECgYEAz/mXt/GP8fMgy8MlXT0Jqd9WqZ903RsNO8ygW1dyWS+ESs2YHN0s
RmaghFEPzax07nq8pkLCWzHKLukXFONx4gNrx8keNeu61E/GVRpxprUzFWH1xbHZ
T5I1U7l+Iwnr2F30vupGBJ23y0uXQirrm73+Z9c885uTlyu7ZurWMLcCgYEAxRaG
dQofm3wM0gwhzcUyCkYLt9M8FfXjVEnUCkXAjuhwyCXoS+oAld4bqYAhfKlqMEf+
l2aPiXFBRyDGbOwkwvXS0B0ZCHln9NWt4FpFW+60DzxhEv3vavUyHrSO++XbB7KB
yhWGue7eJ2nOkfTNbb/2kfJGr46CSboEcezq2lECgYEAiqyfmJxu4eK+5H8rw972
OZndHF9huijWiyAncKB/c652ZLZwDhb/9bVqpK+0fOzYT0fx0F7FX690ZBEyPdBm
2FB1ppKZHVUgj37d/VMToxhBhql6CqdLAn64JiqSS+TKqMwFbOOjamKL8fdmVU1v
KrrmDvF7B+id4ffcDoTZ60MCgYBYGq62yXTBvB27FGNUKlApWZDJd2uH4ajjODHK
+c2P1Qb94jxLG2txk53IExhlMxLeTIDaS6Xk6jUlR9iMPrBcWyoHkMptCGDZiWA2
SARziW0C1poKtGv/42apZUv5/ZIBieINZbwZiFfVRK5sfwQKiOL/8U96EXna3YY5
K5D9UQKBgQC139cOp2NJKrHIumPbbPs4trm2NL76t6IXOW2HewS8Ibg0o1KnfXBk
S5ZpKRFwpXwXIF5WxBCDVHxVdVLqTbLqzfb16qY08Shg1ERy/uUi09D9qXMV9Z5X
/b2GNpeUBs+1T6JcFhlj7KXj2cviWqy4iXhE42XlNrpg+hkZviBemA==
-----END RSA PRIVATE KEY-----`),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadPKCS1PrivateKey(tt.args.privateKey)
			if !tt.wantErr(t, err, fmt.Sprintf("LoadPKCS1PrivateKey(%v)", tt.args.privateKey)) {
				return
			}

		})
	}
}

func TestSignPKCS1v15WithSha256(t *testing.T) {
	type args struct {
		priv      string
		plaintext []byte
	}
	tests := []struct {
		name string
		args args

		wantErr assert.ErrorAssertionFunc
	}{
		//{
		//	name: "SignPKCS1v15WithSha256-err",
		//	args: args{
		//		priv:      "test",
		//		plaintext: nil,
		//	},
		//	wantErr: assert.Error,
		//},
		{
			name: "SignPKCS1v15WithSha256-OK",
			args: args{
				priv: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAoB1eb1FolIkXconqXdc3J12W23KHWluZ/0LC6wVmWBRDztkY
+UmeH7vrKj9XpjCGur3Sieuk6rJKwdzaMzQshlpnnyRr6XyTxYBPL96WjjolWBCh
jclSkX6hg/YTljB2/h4TDp6iVIzk8j9bw1j5X4yEw1NdbIThgQEbxaAHho5eclqP
3YQGWYbx7Hxa54hsClsMZNgJKH4lTX3wuAjtEzeyYnSSl90zof0WFbnBoGU6XDkm
O5L5Aq8oi74L8D5yTAzU6eS3VLDIf3yYNfTqPYywKWPEIYaV4nfV/2v2NbqOR1yS
wv1y5iqIXNgA3K/83MzGFFGwrbTUQGkC7oc/5wIDAQABAoIBACZWYeItj+jg2mhm
dWN6wI4AbrqktZwBCuJ/zcQQSu0UDRheCwjFg8L9b9VxzT7Rp7DW+q5jad14S0YX
53cTKwYQZ3dHdT82wPstOciwd0Qe0ApCESyt48NKmsnKBe6UNCmsccuyWBP+mGMY
oJQpnawvLSrHrzropkYJSekL5EFnPGD1LIf+I+WY9fb4YTt6Hx4IHIZWrYd+Mxg2
hMQsMo0m6p3agJijYIfHntWqB0Ykzj72UtWSF5q+y2MlrG0dQKBvGVTkGGzjhxy7
nUm/1NCJPtxQf5+AYiFvzydjmg/5xOEIO5Ys90HrvdmWEPg9+Q/gg931cG7vhD1U
++rwEoECgYEAz/mXt/GP8fMgy8MlXT0Jqd9WqZ903RsNO8ygW1dyWS+ESs2YHN0s
RmaghFEPzax07nq8pkLCWzHKLukXFONx4gNrx8keNeu61E/GVRpxprUzFWH1xbHZ
T5I1U7l+Iwnr2F30vupGBJ23y0uXQirrm73+Z9c885uTlyu7ZurWMLcCgYEAxRaG
dQofm3wM0gwhzcUyCkYLt9M8FfXjVEnUCkXAjuhwyCXoS+oAld4bqYAhfKlqMEf+
l2aPiXFBRyDGbOwkwvXS0B0ZCHln9NWt4FpFW+60DzxhEv3vavUyHrSO++XbB7KB
yhWGue7eJ2nOkfTNbb/2kfJGr46CSboEcezq2lECgYEAiqyfmJxu4eK+5H8rw972
OZndHF9huijWiyAncKB/c652ZLZwDhb/9bVqpK+0fOzYT0fx0F7FX690ZBEyPdBm
2FB1ppKZHVUgj37d/VMToxhBhql6CqdLAn64JiqSS+TKqMwFbOOjamKL8fdmVU1v
KrrmDvF7B+id4ffcDoTZ60MCgYBYGq62yXTBvB27FGNUKlApWZDJd2uH4ajjODHK
+c2P1Qb94jxLG2txk53IExhlMxLeTIDaS6Xk6jUlR9iMPrBcWyoHkMptCGDZiWA2
SARziW0C1poKtGv/42apZUv5/ZIBieINZbwZiFfVRK5sfwQKiOL/8U96EXna3YY5
K5D9UQKBgQC139cOp2NJKrHIumPbbPs4trm2NL76t6IXOW2HewS8Ibg0o1KnfXBk
S5ZpKRFwpXwXIF5WxBCDVHxVdVLqTbLqzfb16qY08Shg1ERy/uUi09D9qXMV9Z5X
/b2GNpeUBs+1T6JcFhlj7KXj2cviWqy4iXhE42XlNrpg+hkZviBemA==
-----END RSA PRIVATE KEY-----`,
				plaintext: []byte("test"),
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priv, _ := LoadPKCS1PrivateKey([]byte(tt.args.priv))

			sign, err := SignPKCS1v15WithSha256(priv, tt.args.plaintext)
			if !tt.wantErr(t, err, fmt.Sprintf("SignPKCS1v15WithSha256(%v, %v)", tt.args.priv, tt.args.plaintext)) {
				return
			}
			fmt.Println("_________")
			fmt.Println(sign)
		})
	}
}

func TestSHA256(t *testing.T) {
	strArr := []string{
		"test",
		"test1",
		"test2",
		"test3",
		"test4",
	}

	for _, s2 := range strArr {
		s := sha256.New()
		s.Write([]byte(s2))
		fmt.Println("SHA256 ", fmt.Sprintf("%x", s.Sum(nil)))

		s3 := sha1.New()
		s3.Write([]byte(s2))
		fmt.Println("SHA1 ", fmt.Sprintf("%x", s3.Sum(nil)))
		fmt.Println("_________")
	}

}
