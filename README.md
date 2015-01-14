A simple timeline service. Publish an event to the service with an author and content. Use search to return a sorted list of events from multiple authors.

To develop locally:
Update the SDK from here: https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go
// From project directory
% goapp serve app

On check-in:
// Make sure GOPATH is set.
// From project directory
% goapp test ./...
run gofmt

To deploy:
goapp deploy -oauth app

TODO:
	Add testing for controllers
	Migrate Search logic to timeline file.
	Add CORS headers to search interface
	Add bootstrap & AJAX to appengine site