Imagine you are building a web application that sends notification to users. When a user performs a specific action, such as submitting a form or completing a purchase, the application needs to send an email notification to the user. The process of sending an email can be time-consuming due to network latency or external email service delays

we want to make sure that users don't experience delays in the web application's response time due to email sending.

To ensure a smooth user experience, we'll use asynchronous email sending. Instead of sending the email synchronously, which would block the main application, we'll use Goroutine to send the email asynchronously in the background. This way, the main application can respond quickly to the user's request, and email sending will happen independently in the background


<!-- you may found the following stackoverflow post helpful -->
https://stackoverflow.com/questions/70656333/how-to-run-2-goroutines-in-a-specific-order