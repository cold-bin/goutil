// @author cold bin
// @date 2022/11/11

package conv

import (
	"reflect"
	"testing"
)

func TestDecodeBase64(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantDst []byte
		wantErr bool
	}{
		{
			name:    "good case 1",
			args:    args{src: []byte("aGVsbG8gd29ybGQ=")},
			wantDst: []byte("hello world"),
			wantErr: false,
		},
		{
			name:    "good case 2",
			args:    args{src: []byte("Mm56deS9oOWlvSDjgIIuJCBzczEyQA==")},
			wantDst: []byte("2nzu你好 。.$ ss12@"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := DecodeBase64(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("DecodeBase64() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestDecodeBase64Url(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantDst []byte
		wantErr bool
	}{
		{
			name:    "good case 1",
			args:    args{src: []byte("aHR0cHM6Ly9nb3JtLmlvL3poX0NOL2RvY3MvcHJlbG9hZC5odG1sIyVFNSVCOCVBNiVFNiU5RCVBMSVFNCVCQiVCNiVFNyU5QSU4NCVFOSVBMiU4NCVFNSU4QSVBMCVFOCVCRCVCRA==")},
			wantDst: []byte("https://gorm.io/zh_CN/docs/preload.html#%E5%B8%A6%E6%9D%A1%E4%BB%B6%E7%9A%84%E9%A2%84%E5%8A%A0%E8%BD%BD"),
			wantErr: false,
		},
		{
			name:    "good case 2",
			args:    args{src: []byte("aHR0cHM6Ly9jb2xkLWJpbi5naXRodWIuaW8vcG9zdC8lRTUlQjklQjYlRTUlOEYlOTElRTUlQUUlODklRTUlODUlQTglRTQlQjklOEIlRTUlOEUlOUYlRTUlQUQlOTAlRTYlOTMlOEQlRTQlQkQlOUMvIyVlNiVhZiU5NCVlOCViZSU4MyVlNSViOSViNiVlNCViYSVhNCVlNiU4ZCVhMg==")},
			wantDst: []byte("https://cold-bin.github.io/post/%E5%B9%B6%E5%8F%91%E5%AE%89%E5%85%A8%E4%B9%8B%E5%8E%9F%E5%AD%90%E6%93%8D%E4%BD%9C/#%e6%af%94%e8%be%83%e5%b9%b6%e4%ba%a4%e6%8d%a2"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := DecodeBase64Url(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64Url() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("DecodeBase64Url() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestEncodeBase64(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantDst []byte
	}{
		{
			name:    "good case 1",
			args:    args{src: []byte("萨的asd a.a 、？/ @ sa*")},
			wantDst: []byte("6JCo55qEYXNkIGEuYSDjgIHvvJ8vIEAgc2Eq"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := EncodeBase64(tt.args.src); !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("EncodeBase64() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestEncodeBase64Url(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantDst []byte
	}{
		{
			name:    "good case 1",
			args:    args{src: []byte("https://cold-bin.github.io/post/%E5%B9%B6%E5%8F%91%E5%AE%89%E5%85%A8%E4%B9%8B%E5%8E%9F%E5%AD%90%E6%93%8D%E4%BD%9C/#%e6%af%94%e8%be%83%e5%b9%b6%e4%ba%a4%e6%8d%a2")},
			wantDst: []byte("aHR0cHM6Ly9jb2xkLWJpbi5naXRodWIuaW8vcG9zdC8lRTUlQjklQjYlRTUlOEYlOTElRTUlQUUlODklRTUlODUlQTglRTQlQjklOEIlRTUlOEUlOUYlRTUlQUQlOTAlRTYlOTMlOEQlRTQlQkQlOUMvIyVlNiVhZiU5NCVlOCViZSU4MyVlNSViOSViNiVlNCViYSVhNCVlNiU4ZCVhMg=="),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := EncodeBase64Url(tt.args.src); !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("EncodeBase64Url() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
