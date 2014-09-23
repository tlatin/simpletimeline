A simple timeline service. Publish an event to the service with an author and content. Use search to return a sorted list of events from multiple authors.

On check-in:
// Make sure GOPATH is set.
// From project directory
% goapp test ./...
run gofmt

TODO:
	dynamically change template path for goapp test ./...
	Create JSON interface for search
	Add testing for controllers
	Add test for search Controller
	Migrate Search logic to timeline file.
	Add CORS headers to search interface
	Add bootstrap & AJAX to appengine site