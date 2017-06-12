var q = 'tasks'

var url = process.env.CLOUDAMQP_URL || 'amqp://localhost';
var open = require('amqplib').connect(url)

open
.then(function(conn) {
  return conn.createChannel()
})
.then(function(ch) {
  return ch
    .assertQueue(q)
    .then(function(ok) {
      ch.consume(q, function(msg) {
        if (msg !== null) {
          console.log('received:', msg.content.toString())
          ch.ack(msg)
        }
      })

      return ch.sendToQueue(q, new Buffer('something to do'))
    })
})
.catch(console.warn)
