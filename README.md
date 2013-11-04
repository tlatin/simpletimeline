A simple timeline service. Publish an event to the service with an author and content. Use search to return a sorted list of events from multiple authors.

To Test:
% ~/go_appengine/goapp test ~/simpletimeline/timeline/*

TODO:
	Run golint
	Add testing for controllers
	Create JSON interface for search
	Add CORS headers to search interface
	Add bootstrap & AJAX to appengine site