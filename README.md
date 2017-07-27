# [Trello Trends](http://simsalabim.github.io/trello-trends)

Connect to Trello and observe trends of your cards based on assigned labels. Awesome graphs and all in just one file! 

## Development

Due to the nature of Trello's `client.js` authentication if the file isn't served by a web server you will see the 
following error message:
 `trello.com/1/token/approve:3 Failed to execute 'postMessage' on 'DOMWindow': The target origin provided ('file://') 
 does not match the recipient window's origin ('null').`
 
 That's why you need to launch this tiny web server while on localhost: 
 `go run dev_server.go`.
 Default port is `8000` but you can change it to whatever you like like this: `go run dev_server.go -p 3000`.
 All it does is serve [`index.html`](http://0.0.0.0:8000). Now you can connect to your Trello account and use the app 
 on your machine.
 
 ### TODO
 
 The project was written as a proof of concept in a :shipit:-mode and, surely, a lot of things can be improved.
 From my experience, you never get time or enough motivation to fix these things after the fact so such lists just hang
 in the infinity of suspense forever. But I'm gonna leave it here:
 
 - [x] Configure Github Pages to work on `master` branch so that every change is reflected on the
  [demo site](http://simsalabim.github.io/trello-trends).
 - [ ] Let the user know when waiting on external calls to complete (show spinners).
 - [ ] Error handling.
 - [ ] Persist settings between page refreshes.
 - [ ] Refactor JavaScript code to make it more modular and easy to read.
 - [ ] Integrate a router, use React/Vue.js.
 
 ## License
 
 MIT
 
