# GlobalWebIndex Engineering Challenge

## Introduction

This challenge is designed to give you the opportunity to demonstrate your abilities as a software engineer and specifically your knowledge of the Go language.

On the surface the challenge is trivial to solve, however you should choose to add features or capabilities which you feel demonstrate your skills and knowledge the best. For example, you could choose to optimise for performance and concurrency, you could choose to add a robust security layer or ensure your application is highly available. Or all of these.

Of course, usually we would choose to solve any given requirement with the simplest possible solution, however that is not the spirit of this challenge.

## Challenge

Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

A working server application with functional API is required, along with a clear readme.md. Useful and passing tests would be also be viewed favourably

It is appreciated, though not required, if a Dockerfile is included.

## Submission

Just create a fork from the current repo and send it to us!

Good luck, potential colleague!


# Code Documentation
## Run WebApp
Currently the application runs on localhost:8000 using the config.env file you can update amd adjust this endpoint if this specific port is un available or IP needs to be changed to your servers IP address.

## API calls
You can call the API from 3 distict URLs 
"/" "GET" request                       ->  This returns all the currently available Data
"/add" "POST" request                   ->  This can be used to add new assets of any type in a specific format
"/update/{assetName}" "POST" request    ->  This can be used to update existing assets of any type to any other type in a specific format (including the name of the asset even though it works as a unique id)
"/delete/{assetName}" "GET" request     ->  This can be used to delete an existing asset by Name

# API Body format
An example of the API raw body can be found in the post_format.json file this can be specificalyly used on the "/add" call

## TEST
A short test data has already been inplemented in frontend/frontend.go call test1, this data can be expanded on with more calls

## Submission

With the short time I had, This is my solution to the problem

I hope you like what you see

My best Regards and Thank you
