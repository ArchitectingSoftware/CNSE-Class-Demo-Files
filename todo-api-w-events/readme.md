## ToDo API Demo - With Events

This is an extension of the `todo` API that also illustrates adding some basic eventing.  It includes code to create a channel for events, and also a goroutine to process the events asynchronously.  The documentation is basically the same as the `todo` API.  

### Changes over the base todo API

This version of the `todo` API includes the following over the base version:

1. The added endpoints include `/event/true` and `/event/false` to dynamically enable and disable eventing.

2. Demonstration of goroutines to handle events asynchronously. 
3. Demonstration of using a golang context to manage an asynrounous goroutine
4. Demonstration of filtering events using golang channels 