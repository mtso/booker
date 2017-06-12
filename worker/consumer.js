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
      return ch.consume(q, function(msg) {
        if (msg !== null) {
          console.log('consumed:', msg.content.toString())
          ch.ack(msg)

          setTimeout(function() {
            ch.sendToQueue(q, new Buffer('return something done'))
          }, 1000)
        }
      })
    })
})
.catch(console.warn)
