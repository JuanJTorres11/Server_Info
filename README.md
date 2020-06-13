
# Server Info

Web application that consist of a backend built with go that manage two API endpoints and a frontend built with vue.js  to display the information of these endpoints in a user friendly way.


# How to deploy

Here is the proccess to deploy the app

## Create the database

 1. Go [here](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-linux.html) to get instructions to install cockroachdb 
 2. Follow [this](https://www.cockroachlabs.com/docs/stable/cockroach-start-single-node.html#start-a-single-node-cluster) instructions to create a secure single node cluster.
 3. In the console enter `cockroach sql --certs-dir=certs` to start the SQL shell.
4. Copy and paste the content of [db.sql](https://github.com/JuanJTorres11/Server_Info/blob/master/db.sql)
5. Close the SQL Shell with `\q`
6. Generate a certificate for the user *user_servers* with `cockroach cert create-client user_servers --certs-dir=certs --ca-key=my-safe-directory/ca.key`

## Compile the backend

You can compile the backend with the go console tool `go build ./...` then you can execute the binary like any other program.

## Start the frontend

1. Enter the web directory
2. Install the dependencies with `npm install`
3. Initiate the server with `npm run serve` and it will start in the port 8000

# API Endpoints
The endpoints of the API are:

## servers_info/*domain*
Receives a domain  like amazon.com and return a JSON with the following information:

    {
    “servers”: [
    {
    “address”: “server1”,
    “ssl_grade”: “B”,
    “country”: “US”,
    “owner”: “Amazon.com, Inc.”
    },
    {
    “address”: “server2”,
    “ssl_grade”: “A+”,
    “country”: “US”
    “owner”: “Amazon.com, Inc.”
    },
    {
    “address”: “server3”,
    “ssl_grade”: “A”,
    “country”: “US”
    “owner”: “Amazon.com, Inc.”
    }
    ],
    “servers_changed”: true,
    “ssl_grade”: “B”,
    “previous_ssl_grade”: “A+”,
    “logo”: “ https://server.com/icon.png ”,
    “title”: “Title of the page”,
    “is_down”: false
    }

- **servers:** contains an array of the servers associated with the domain, each object  of the array contains:
	- **address:** the IP or the host of the server  
	- **ssl_grade:** SSL grade graded by SSLabs  
	- **country:** The country of the server as it appears when using the command  whois ip
	- **owner:** the organization that owns the IP, as it appears when using the  whois ip command  
- **servers_changed:** it is true if the servers changed, regarding a  hour or more before  
- **ssl_grade:** lowest grade of all servers  
- **previous_ssl_grade:** the grade that had an hour or more before  
- **logo:** the domain logo taken from the `<head>` of the HTML  
- **title:** the title of the page taken from the `<head>` of the HTML  
- **is_down:** true if the server is down and cannot be contacted

## servers

List the servers that have been previously consulted,  even if the browser is restarted.  
For example, if you have query amazon.com and google.com the endpoint should return both results:

    {
    “items”: [
    {google.com info},
    {amazon.com info}
    ]
    }

