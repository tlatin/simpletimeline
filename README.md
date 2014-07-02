A simple timeline service. Publish an event to the service with an author and content. Use search to return a sorted list of events from multiple authors.

On check-in:
// Make sure GOPATH is set.
% goapp test github.com/tlatin/simpletimeline/...
run gofmt

TODO:
	Add testing for controllers
	Add test for search Controller
	Migrate Search logic to timeline file.
	Create JSON interface for search
	Add CORS headers to search interface
	Add bootstrap & AJAX to appengine site