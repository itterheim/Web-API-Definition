package main

// App is parsed definition file or its section
type App struct {
	Name      string
	Options   AppOptions
	Sections  []App
	Endpoints []Endpoint
}

// AppOptions is configuration of App
type AppOptions struct {
	RoutePrefix string
	Port        string
}

// Endpoint defines server behavior
type Endpoint struct {
	Method  string
	Route   string
	Name    string
	Options EndpointOptions
}

// EndpointOptions defines route configuration
type EndpointOptions struct {
	Response    string
	ContentType string
}
