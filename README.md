# Mappyr Backend api

Soo there are a not too many things this can do yet, but here it is.

## Routes

	/

Index doesn't really do anything

	/comment/{id}
	
returns a specific comment by it's id, the response will look like this:

`{"id":2,"title":"Great food!","description":"Although crap service","latitude":41.353,"longitude":-71.113,"upvotes":-1,"downvotes":1,"date":"2016-10-11T11:27:20.19479779-04:00","user-id":`

	/comments
	
returns a json list of **all** comments.

	/new
	
the create route accepts the POST method, it only needs a title, description, and the latitude and longitude. Notice the numbers aren't in quotes, here's an example curl:

`curl -H "Content-Type: application/json" -d '{"title":"This place sux0rs", "description":"Yikes","latitude":41.33894, "longitude":-71.666}' http://localhost:8080/new `

	/upvote/{id} 
	//OR
	/downvote/{id}

These accept a GET method, and they update the count of downvotes or upvotes for a comment row. It'll then return the very comment which is being commented (maybe this'll just want to be ignored, but I thought'd be a good idea to return *something*.


## TODO

This list ought to be longer

1. Some sort of authentication, maybe hitching of OAuth so people don't need to create an account.
2. Track the tokens, including 'anonomyous' tokens, to keep from 'voter fraud'.
3. Expand the database to include a downvote and upvote table

### todo routes:

	/login
	
	/token
	
	/delete
