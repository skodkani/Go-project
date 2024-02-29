module github.com/skodkani/Go-project/hello-appl/hellogo

go 1.21.3

replace github.com/skodkani/Go-project/hello-appl/mystrings v0.0.0 => ../mystrings

require (
    github.com/skodkani/Go-project/hello-appl/mystrings v0.0.0
)