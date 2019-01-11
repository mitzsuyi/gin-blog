# gin-blog

This is a sample [`gb`](http://getgb.io/) project.  Implements the following backend api

GET /articles
POST /articles
GET /articles/:id/content
GET /articles/:id/comments
POST /articles/:id/comments
POST /comments/:id/comments

Parameters for posting look like:

article:
{
    nickname
    title
    content
}

comment:
{
    nickname
    content
}

## Get the code

To fetch this project

     git clone https://github.com/mitzsuyi/gin-blog

# Building the code

To build this project 

     cd gin-blog
     gb build

If you don't have `gb` installed, [follow these instructions](http://getgb.io/docs/install/).

# Run the program

To run the program, build it, then run the binary `bin/api`; will start server on port 8080; can change port with flag -p 3000, ie, bin/api -p 3000;
