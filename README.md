# go_tui
This is just a place for me to experiment with bubbletea and creating a tui for a go/python application.

The objective here is to have a tui written in GO that can kick off processes written in python. 

The python processes in this case will be very simple placeholders - the main learning is how to marry the two together. 

The reason for this is that the bubbletea framework is very nice and offers greater flexibility than python. THis is also a good use case for how to use these two languages together. 


So where to begin? Since I don't know anything about the GO language, lets start with a really simple tui app straight out of the examples. Maybe the simple list or something. 


Then we need to figure out how to make GO run a python script. Or a shell script. Or both whynot...

I am going to first do this like a caveman, and then improve iteratively since I write code like a caveman. I am just going to make go do like a sys.execute(script) when I choose an option. Once I can do that I have a starting point to improve. At some point I'd like to do something with the dlls and shared libraries, maybe there is even c code involved, idk. We will see. 



---
In my head here is basically how the thing is going to work. In bubbletea there is this concept of commands and messages. And I think each command will kick off a python/bash process which will return the message.  




