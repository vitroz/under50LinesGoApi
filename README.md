
One page simple CRUD Api in GoLANG

If http router library cannot be found check this article for instrunctions on how to donwload 3rd party libraries
https://wilsonmar.github.io/golang/

How to use:

Run the code:

  go run simpleapi.go

Open a new terminal window, and test it by using curl:


  PUT(CREATE/UPDATE)
  
  curl -X PUT localhost:8083/entry/first/hello
  
  curl -X PUT localhost:8083/entry/second/hi
  
  GET
  
  curl localhost:8083/list
  
  curl localhost:8083/entry/first
  
  curl localhost:8083/entry/second

  DELETE
  
  curl -X DELETE localhost:8083/entry/first


  No database! So the entries only last until the server is alive :)
