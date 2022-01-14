## BACKEND TEST 1
<i>Submitted by: Mitchell Sullivan</i>

### Intro
This is basically my first-ever project in Go. As such, it's quite simple 
and bears little resemblance to the hulking behemoth I'd have wrought in .NET. 
<i>But, hey, isn't that the whole point of Go?</i>

Because I'm a novice in this language, there are some sacrifices, but I 
thought that the trade-off of using Array's technology was worth it. 
This project lacks demonstrations of more advanced practices with which I am familiar in 
other languages, like dependency injection and exhaustive unit testing.
Proper commenting and git commits fell by the wayside while I focused on learning, as 
did a certain degree of polish.

### Directions 
To run the app, just make sure that port 8099 is free on your machine, or change `HOST_PORT` in the `.env` file,
and run `docker-compose up`.

A Postman collection is included in `docs`, as well an OpenApi 3.0 spec.

In Postman, after signing up, logging in, or refreshing, the token will need be retrieved from the 
response body and set as the Bearer token with the Authorization header.


### Design 
<b>Login</b> is facilitated via a JWT that expires in 1 hour. For the sake of a 
minimum viable product, I have <b>NOT</b> implemented a separate refresh token with a 
longer expiry time; the unexpired JWT will need to be used in a refresh request within the hour
in order to obtain a new one, and the old one will be revoked.

<b>Logout and token revocation</b> are facilitated via a blacklist, stored in Redis, with a time-to-live of 1 hour,
since the JWT would become invalid after that time, anyway.

I gave <b>Redis</b> a replica, probably out of insecurity over 
my experience level with Go and the desire to show off more skills. I had originally set up and was developing for a 
cluster, but even the Redis docs state that clustering is a clunky experience within Docker containers, 
so I chose something simpler. 
 
The <b>Gin web framework</b> was chosen due to its popularity and the 
preponderance of examples which use it.

<b>PostgreSQL</b> is the database instead of MS SQL Server due to MSSQL's lack of 
initialization options when using docker (setup scripts basically require 
building a new container), and Postgres supports automatic database creation and an init 
directory for scripts. 

The app uses <b>GORM</b>, an ORM, for data access; there is code for automatic migration to
create the relevant table(s), but, since this exercise requires the inclusion of such DDL, 
the script is located at `docker-vols/pg_init/setup.sql` and is run on database startup, preventing the migration. 
For the primary key of Users, I've used a ULID (generated in-app), instead of a UUID vX, 
for improved indexing and because it seemed more appropriate for use as a JWT claim than a serial integer. 


### Thank you!!



