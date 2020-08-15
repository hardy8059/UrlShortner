# UrlShortner

This repo implements a url shortener.
It reads the urls from a config file in which they are defined in the form:
 - path (the url user enters)
 - url (the url to which we need to redirect)
 
Program takes a user input for the name of configuration file and then starts the server by mapping the urls from that config file.
Currently,  two handlers are implemented. Which handle following file format:
 - Yaml Files
 - Json Files

![alt text](https://github.com/hardy8059/UrlShortner/blob/master/Output%20Images/User%Configuration.PNG?raw=true)

### Packages Explored:
 - Json (https://golang.org/pkg/encoding/json/)
 - Yaml (https://github.com/go-yaml/yaml)
 - HTTP (https://golang.org/pkg/net/http/)
 - Filepath (https://golang.org/pkg/path/filepath/)

Inspiration taken from https://courses.calhoun.io/courses/cor_gophercises
