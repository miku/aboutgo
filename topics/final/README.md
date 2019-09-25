# Final project

## Project 1 - Write a small Guestbook application

* write a server
* write two handlers: one for posting a message, one for displaying all messages

Example:

    /        list all messages
    /post    receive posts, through POST request (test with curl)

No need for a database, the service is ephemeral. It uses a fixed length
in-memory queue to store the posts. The oldest post should disappear
automatically.

## Project 2 - Random Image Server

The task is to write a web application to generate pretty random pictures.

* write a helper function to generate random image data
* write a command, that runs a server
* attach a handler to the server
* this handler should parse a query parameter (e.g. "seed") which is then passed to the random image generator
* connect the image output with the http writer to send back a response

## Project 3 - Random text

Random text generator.

* write a server
* write a handler
* write a utility function, that takes a list of URLs and fetches them
* for each URL, parse the page into tokens (try to detect words)
* from each URL select a random token
* concatenate the random tokens and send back the result
