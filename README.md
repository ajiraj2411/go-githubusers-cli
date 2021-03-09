# go-githubusers-cli
Golang CLI app to query GitHub users.

Searching GitHub users using https://api.github.com/users/{username}

It's mainly for learing Implementing CLI

We could use the flag standard library in Go, but, with trial-and-error and a little bit of Googling, you will discover that the standard flag library does not support the long-flag syntax (via double-dash). It only supports single dashes.

So we are going to use already implemented GO library github.com/ogier/pflag


Let's see what and all needed for creating CLI
    pflag (go get github.com/ogier/pflag)
    color (go get github.com/fatih/color)

Task
    Creating Flag user in init Function
    Call the GetUSer funtion for geeting GitHub users data in a user struct format
    Parse the details and print the values (UserName, Name, Email, Bio) If have any