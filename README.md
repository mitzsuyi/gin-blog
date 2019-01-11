# gin-blog

This is a sample [`gb`](http://getgb.io/) project.  Implements the following backend api

GET /articles (optional ?page=x to navigate pages; returns pagination struct)

POST /articles

GET /articles/:id/content

GET /articles/:id/comments

POST /articles/:id/comments

POST /comments/:id/comments


### Parameters for posting look like:

Article:
{
    nickname,
    title,
    content
}

Comment:
{
    nickname,
    content
}

Pagination struct looks like

{
   
    total_record,
    total_page,
    records,
    offset,
    limit,
    page,
    prev_page,
    next_page
}

## Get the code

To fetch this project

     git clone https://github.com/mitzsuyi/gin-blog

# Building the code

To build this project 

     cd gin-blog
     gb vendor restore (to instal dependencies one time)
     gb build

If you don't have `gb` installed, [follow these instructions](http://getgb.io/docs/install/).

# Run the program

To run the program, build it, then run the binary `bin/api`; will start server on port 8080; can change port with flag -p 3000, ie, bin/api -p 3000;

# Requirements

Uses mysql database

# Config

There is a json file src/api/config.json with mysql credentials add your own mysql url
