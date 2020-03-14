
# The World's Most Over Engineered BBQ Thermometer

  

## What is this?

  
If you would like a more readable explanation of what this is then please check our wiki.  https://github.com/ssargent/bbq/wiki  It will do a better job of putting things into plain english than we can do here.

This was the repo for the BBQ Coding Challenge. If you're looking for that, its still around just on an branch. https://github.com/ssargent/bbq/tree/dotnet-core-bbq-app This however is no longer the bbq coding challenge. See I overcooked christmas dinner. My solution? Build a kubernetes based bbq thermometer.

## The Parts

 1. monitor-cli is a golang based monitor program that will pull the data off of the bluetooth thermometer and send it to the api server hosted in kubernetes (go-bbq)
 2. bbq-apiserver is an api server that supports rest apis now and grpc future.  It's currently a monolithic api server that provides multi-tenant support for: 
	 - Managing BBQ Devices (Grills, Smokers etc..) 
	 - Managing BBQ Monitors (Thermometers)
	 - Recording and Querying Thermometer Data 
	 - Managing Accounts
	 - Managing Tenants
	 - Temperature Alerts (future work)
	 - ML Based Fault Detection - understand when a probe has failed, the grill has failed or is open etc.. (future work)
 3. bbq-frontend is a react based web application that lets you fully interact with your bbq thermometer.  This will be a modern and responsive web ui that will let you manage devices, monitors, accounts, alerts.
 4. bbq-ml will be the service that performs all machine learning (future work)
## How do I use it
Eventually this will be delivered as software as a service.  You buy a bluetooth thermometer and have a linux computer to run the monitor software (eventually a raspberry pi that runs headless).  But that's in the future.  Currently it works but its very much developers only.  You're welcome to try it but you'll have to know how to compile and run a golang web and cli application.  I hope to have documentation explaining this soon.
