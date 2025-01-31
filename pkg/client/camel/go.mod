module github.com/aanogueira/camel-k/pkg/client/camel

go 1.13

require (
	github.com/aanogueira/camel-k/pkg/apis/camel v0.0.0
	k8s.io/api v0.19.8
	k8s.io/apimachinery v0.19.8
	k8s.io/client-go v0.19.8
	k8s.io/code-generator v0.19.8 // indirect
)

// Local modules
replace github.com/aanogueira/camel-k/pkg/apis/camel => ../../apis/camel
