package add

import (
	"bytes"
	"strings"
	"text/template"
)

const protoTemplate = `
syntax = "proto3";

package {{.Package}};

import "google/api/annotations.proto";
import "validate/validate.proto";

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";

service {{.Service}} {
	rpc Create{{.Service}} (Create{{.Service}}Request) returns (Create{{.Service}}Reply) {
		option(google.api.http) = {
			post:"{{.BaseUrl}}/add"
			body:"*"
		};
	};
	rpc Update{{.Service}} (Update{{.Service}}Request) returns (Update{{.Service}}Reply) {
		option(google.api.http) = {
			post:"{{.BaseUrl}}/{id}/update"
			body:"*"
		};
	};
	rpc Delete{{.Service}} (Delete{{.Service}}Request) returns (Delete{{.Service}}Reply) {
		option(google.api.http) = {
			post:"{{.BaseUrl}}/{id}/delete"
			body:"*"
		};
	};
	rpc Get{{.Service}} (Get{{.Service}}Request) returns (Get{{.Service}}Reply) {

	};
	rpc List{{.Service}} (List{{.Service}}Request) returns (List{{.Service}}Reply) {
		option(google.api.http) = {
			get:"{{.BaseUrl}}/list"
		};
	};
}

message Create{{.Service}}Request {}
message Create{{.Service}}Reply {
	int64  id =1  ;
}

message Update{{.Service}}Request {
	int64  id = 1 ;
}
message Update{{.Service}}Reply {}

message Delete{{.Service}}Request {
	int64  id = 1 ;
}
message Delete{{.Service}}Reply {}

message Get{{.Service}}Request {
	int64  id = 1 ;
}
message Get{{.Service}}Reply {}

message List{{.Service}}Request {
	int64  page_size = 50;
	int64  page_number = 51;
}
message List{{.Service}}Reply {
	message  Item {
		int64  id =1  ;
	}
	repeated  Item  array = 1;
	int32  count = 2 ;
}
`

func (p *Proto) execute() ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("proto").Parse(strings.TrimSpace(protoTemplate))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
